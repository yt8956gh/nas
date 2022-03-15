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

func (a *AtsssNetworkSteeringFuncInfo) GetIdentifier() uint8 {
	return AtsssParameterIdentifierNetworkSteeringfuncInfo
}

func (a *AtsssNetworkSteeringFuncInfo) SetUe3gppIPv4Addr(ip net.IP) {
	a.Ue3gppIpType = AtsssNetworkSteeringFuncInfoIpTypeIPv4
	a.Ue3gppIpAddr = ip
}

func (a *AtsssNetworkSteeringFuncInfo) GetUe3gppIPv4Addr() net.IP {
	return a.Ue3gppIpAddr
}

func (a *AtsssNetworkSteeringFuncInfo) SetUeNon3gppIPv4Addr(ip net.IP) {
	a.UeNon3gppIpType = AtsssNetworkSteeringFuncInfoIpTypeIPv4
	a.UeNon3gppIpAddr = ip
}

func (a *AtsssNetworkSteeringFuncInfo) GetUeNon3gppIPv4Addr() net.IP {
	return a.UeNon3gppIpAddr
}

func (a *AtsssNetworkSteeringFuncInfo) SetMptcpProxyInfo(list []MptcpProxyInfo) {
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
	if err := binary.Read(buffer, binary.BigEndian, a.Ue3gppIpAddr[:4]); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.UeNon3gppIpType); err != nil {
		return err
	}
	if a.UeNon3gppIpType != AtsssNetworkSteeringFuncInfoIpTypeIPv4 {
		return fmt.Errorf("Only support IPv4 in AtsssNetworkSteeringFuncInfo")
	}
	if err := binary.Read(buffer, binary.BigEndian, a.UeNon3gppIpAddr[:4]); err != nil {
		return err
	}

	if err := binary.Read(buffer, binary.BigEndian, &a.MptcpProxyInfoLen); err != nil {
		return err
	}

	if buffer.Len() != int(a.MptcpProxyInfoLen) {
		return fmt.Errorf("MptcpProxyInfoLen doesn't match the exact size.")
	}

	for buffer.Len() > 0 {
		info := MptcpProxyInfo{}
		if err := binary.Read(buffer, binary.BigEndian, &info.IpAddrType); err != nil {
			return err
		}
		if info.IpAddrType != AtsssNetworkSteeringFuncInfoIpTypeIPv4 {
			return fmt.Errorf("Only support IPv4 in AtsssNetworkSteeringFuncInfo")
		}
		if err := binary.Read(buffer, binary.BigEndian, info.IpAddr[:4]); err != nil {
			return err
		}
		if err := binary.Read(buffer, binary.BigEndian, &info.Port); err != nil {
			return err
		}
		if err := binary.Read(buffer, binary.BigEndian, &info.Type); err != nil {
			return err
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
