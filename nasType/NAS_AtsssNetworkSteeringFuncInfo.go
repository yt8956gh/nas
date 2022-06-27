package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// TS 24.193 6.1.4

const (
	AtsssNetworkSteeringFuncInfoIpTypeIPv4 uint8 = iota + 1
	AtsssNetworkSteeringFuncInfoIpTypeIPv6
	AtsssNetworkSteeringFuncInfoIpTypeIPv4v6
)

type AtsssNetworkSteeringFuncInfo struct {
	Ue3gppIpType       uint8
	Ue3gppIpAddr       []byte
	UeNon3gppIpType    uint8
	UeNon3gppIpAddr    []byte
	MptcpProxyInfoLen  uint8
	MptcpProxyInfoList []MptcpProxyInfo
}

func NewAtsssNetworkSteeringFuncInfo() *AtsssNetworkSteeringFuncInfo {
	return &AtsssNetworkSteeringFuncInfo{}
}

func (a *AtsssNetworkSteeringFuncInfo) GetIdentifier() AtsssParameterIdentifier {
	return AtsssParameterIdentifierNetworkSteeringfuncInfo
}

func (a *AtsssNetworkSteeringFuncInfo) SetUe3gppIPv4Addr(ip net.IP) {
	a.Ue3gppIpType = AtsssNetworkSteeringFuncInfoIpTypeIPv4
	a.Ue3gppIpAddr = ip.To4()
}

func (a *AtsssNetworkSteeringFuncInfo) GetUe3gppIPv4Addr() net.IP {
	return a.Ue3gppIpAddr
}

func (a *AtsssNetworkSteeringFuncInfo) SetUeNon3gppIPv4Addr(ip net.IP) {
	a.UeNon3gppIpType = AtsssNetworkSteeringFuncInfoIpTypeIPv4
	a.UeNon3gppIpAddr = ip.To4()
}

func (a *AtsssNetworkSteeringFuncInfo) GetUeNon3gppIPv4Addr() net.IP {
	return a.UeNon3gppIpAddr
}

func (a *AtsssNetworkSteeringFuncInfo) SetMptcpProxyInfo(len uint8, list []MptcpProxyInfo) {
	a.MptcpProxyInfoLen = len
	a.MptcpProxyInfoList = list
}

func (a *AtsssNetworkSteeringFuncInfo) GetMptcpProxyInfo() []MptcpProxyInfo {
	return a.MptcpProxyInfoList
}

func (a *AtsssNetworkSteeringFuncInfo) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, &a.Ue3gppIpType); err != nil {
		return err
	}
	if a.Ue3gppIpType != AtsssNetworkSteeringFuncInfoIpTypeIPv4 {
		return fmt.Errorf("Only support IPv4 in AtsssNetworkSteeringFuncInfo")
	}
	a.Ue3gppIpAddr = make([]byte, net.IPv4len)
	if err := binary.Read(buffer, binary.BigEndian, a.Ue3gppIpAddr[:]); err != nil {
		return fmt.Errorf("Binary.Read Ue3gppIpAddr Fail: %+v", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.UeNon3gppIpType); err != nil {
		return fmt.Errorf("Binary.Read UeNon3gppIpType Fail: %+v", err)
	}
	if a.UeNon3gppIpType != AtsssNetworkSteeringFuncInfoIpTypeIPv4 {
		return fmt.Errorf("Only support IPv4 in AtsssNetworkSteeringFuncInfo")
	}
	a.UeNon3gppIpAddr = make([]byte, net.IPv4len)
	if err := binary.Read(buffer, binary.BigEndian, a.UeNon3gppIpAddr[:]); err != nil {
		return fmt.Errorf("Binary.Read UeNon3gppIpAddr Fail: %+v", err)
	}

	if err := binary.Read(buffer, binary.BigEndian, &a.MptcpProxyInfoLen); err != nil {
		return fmt.Errorf("Binary.Read MptcpProxyInfoLen Fail: %+v", err)
	}

	if buffer.Len() != int(a.MptcpProxyInfoLen) {
		return fmt.Errorf("MptcpProxyInfoLen doesn't match the exact size.")
	}

	for buffer.Len() > 0 {
		info := MptcpProxyInfo{}
		if err := binary.Read(buffer, binary.BigEndian, &info.IpAddrType); err != nil {
			return fmt.Errorf("Binary.Read IpAddrType Fail: %+v", err)
		}
		if info.IpAddrType != AtsssNetworkSteeringFuncInfoIpTypeIPv4 {
			return fmt.Errorf("Only support IPv4 in AtsssNetworkSteeringFuncInfo")
		}
		info.IpAddr = make([]byte, net.IPv4len)
		if err := binary.Read(buffer, binary.BigEndian, info.IpAddr[:]); err != nil {
			return fmt.Errorf("Binary.Read IpAddr Fail: %+v", err)
		}
		if err := binary.Read(buffer, binary.BigEndian, &info.Port); err != nil {
			return fmt.Errorf("Binary.Read Port Fail: %+v", err)
		}
		if err := binary.Read(buffer, binary.BigEndian, &info.Type); err != nil {
			return fmt.Errorf("Binary.Read Type Fail: %+v", err)
		}
		a.MptcpProxyInfoList = append(a.MptcpProxyInfoList, info)
	}

	return nil
}

func (a *AtsssNetworkSteeringFuncInfo) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, &a.Ue3gppIpType); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Ue3gppIpAddr); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.UeNon3gppIpType); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, a.UeNon3gppIpAddr); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.MptcpProxyInfoLen); err != nil {
		return nil, err
	}

	for _, info := range a.MptcpProxyInfoList {
		if err := binary.Write(buffer, binary.BigEndian, &info.IpAddrType); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, info.IpAddr); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &info.Port); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &info.Type); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

const (
	MptcpProxyTypeTransportConverter uint8 = 1
)

type MptcpProxyInfo struct {
	IpAddrType uint8
	IpAddr     []byte
	Port       uint16
	Type       uint8
}

func NewMptcpProxyInfo(v4 bool, v6 bool, ipv4Addr net.IP, ipv6Addr net.IP, port uint16, proxyType uint8) (uint8, MptcpProxyInfo) {
	len := uint8(4)
	m := MptcpProxyInfo{
		IpAddrType: uint8(0),
		IpAddr:     make([]byte, 0),
		Port:       port,
		Type:       proxyType,
	}

	if v4 {
		len += net.IPv4len
		m.IpAddrType += AtsssNetworkSteeringFuncInfoIpTypeIPv4
		m.IpAddr = append(m.IpAddr, ipv4Addr.To4()...)
	}
	if v6 {
		len += net.IPv6len
		m.IpAddrType += AtsssNetworkSteeringFuncInfoIpTypeIPv6
		m.IpAddr = append(m.IpAddr, ipv6Addr.To16()...)
	}

	return len, m
}
