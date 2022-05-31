package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// TS 24.193 6.1.5.2
// PMF MAC addr hasn't been supported

const (
	AtsssMeasurementAssistanceInfoIpTypeIPv4 uint8 = iota + 1
	AtsssMeasurementAssistanceInfoIpTypeIPv6
	AtsssMeasurementAssistanceInfoIpTypeIPv4v6
)

type AtsssMeasurementAssistanceInfo struct {
	PmfIpAddrType  uint8
	PmfIpAddr      []byte
	Pmf3gppPort    uint16
	PmfNon3gppPort uint16
	AARI           bool
}

func NewAtsssMeasurementAssistanceInfo() *AtsssMeasurementAssistanceInfo {
	return &AtsssMeasurementAssistanceInfo{}
}

func (a *AtsssMeasurementAssistanceInfo) GetIdentifier() AtsssParameterIdentifier {
	return AtsssParameterIdentifierMeasurementAssistanceInfo
}

func (a *AtsssMeasurementAssistanceInfo) SetPmfIPv4Addr(ip net.IP) {
	a.PmfIpAddrType = AtsssMeasurementAssistanceInfoIpTypeIPv4
	a.PmfIpAddr = ip.To4()
}

func (a *AtsssMeasurementAssistanceInfo) GetPmfIPv4Addr() net.IP {
	return a.PmfIpAddr
}

func (a *AtsssMeasurementAssistanceInfo) SetPmf3gppPort(port uint16) {
	a.Pmf3gppPort = port
}

func (a *AtsssMeasurementAssistanceInfo) GetPmf3gppPort() uint16 {
	return a.Pmf3gppPort
}

func (a *AtsssMeasurementAssistanceInfo) SetPmfNon3gppPort(port uint16) {
	a.PmfNon3gppPort = port
}

func (a *AtsssMeasurementAssistanceInfo) GetPmfNon3gppPort() uint16 {
	return a.PmfNon3gppPort
}

func (a *AtsssMeasurementAssistanceInfo) SetAARI(aari bool) {
	a.AARI = aari
}

func (a *AtsssMeasurementAssistanceInfo) GetAARI() uint16 {
	return a.PmfNon3gppPort
}

func (a *AtsssMeasurementAssistanceInfo) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, &a.PmfIpAddrType); err != nil {
		return err
	}
	if a.PmfIpAddrType != AtsssMeasurementAssistanceInfoIpTypeIPv4 {
		return fmt.Errorf("Only support IPv4 in AtsssMeasurementAssistanceInfo")
	}
	if err := binary.Read(buffer, binary.BigEndian, a.PmfIpAddr[:4]); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Pmf3gppPort); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PmfNon3gppPort); err != nil {
		return err
	}
	var aariByte byte
	if err := binary.Read(buffer, binary.BigEndian, &aariByte); err != nil {
		return err
	}
	a.AARI = (aariByte == 1)

	return nil
}

func (a *AtsssMeasurementAssistanceInfo) Encode() ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, &a.PmfIpAddrType); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PmfIpAddr); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Pmf3gppPort); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.PmfNon3gppPort); err != nil {
		return nil, err
	}
	var arriByte byte
	if a.AARI {
		arriByte = 1
	}
	if err := binary.Write(buffer, binary.BigEndian, &arriByte); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
