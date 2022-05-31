package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// TS 24.501 9.11.4.22 Atsss container
// AtsssContainer
type AtsssContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewAtsssContainer(iei uint8) (atsssContainer *AtsssContainer) {
	atsssContainer = &AtsssContainer{}
	atsssContainer.SetIei(iei)
	return atsssContainer
}

func (a *AtsssContainer) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *AtsssContainer) GetIei() uint8 {
	return a.Iei
}

func (a *AtsssContainer) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

func (a *AtsssContainer) GetLen() uint16 {
	return a.Len
}

func (a *AtsssContainer) GetAtsssParameters() ([]AtsssParameter, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	aps := []AtsssParameter{}
	for buffer.Len() > 0 {
		var (
			ap     AtsssParameter
			id     AtsssParameterIdentifier
			length uint16
		)

		if err := binary.Read(buffer, binary.BigEndian, &id); err != nil {
			return nil, err
		}

		switch id {
		case AtsssParameterIdentifierAtsssRule:
			ap = NewAtsssRules()
		case AtsssParameterIdentifierNetworkSteeringfuncInfo:
			ap = NewAtsssNetworkSteeringFuncInfo()
		case AtsssParameterIdentifierMeasurementAssistanceInfo:
			ap = NewAtsssMeasurementAssistanceInfo()
		default:
			return nil, fmt.Errorf("Unknown ATSSS Parameter id: %d", id)
		}

		if err := binary.Read(buffer, binary.BigEndian, &length); err != nil {
			return nil, err
		}

		content := make([]byte, length)

		if err := binary.Read(buffer, binary.BigEndian, &content); err != nil {
			return nil, err
		}

		if err := ap.Decode(content); err != nil {
			return nil, err
		}
		aps = append(aps, ap)
	}
	return aps, nil
}

func (a *AtsssContainer) SetAtsssParameters(aps []AtsssParameter) error {
	buffer := bytes.NewBuffer(nil)
	length := uint16(0)
	for _, ap := range aps {
		if err := binary.Write(buffer, binary.BigEndian, ap.GetIdentifier()); err != nil {
			return fmt.Errorf("Write ATSSS Parameter ID to buffer Fail: %s", err)
		}

		content, err := ap.Encode()
		length += uint16(len(content) + 3)
		if err != nil {
			return fmt.Errorf("ATSSS Parameter Encode Fail: %s", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, uint16(len(content))); err != nil {
			return fmt.Errorf("Write ATSSS Parameter Length to buffer Fail: %s", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, content); err != nil {
			return fmt.Errorf("Write ATSSS Parameter Content to buffer Fail: %s", err)
		}
	}
	a.SetLen(length)
	a.Buffer = buffer.Bytes()
	return nil
}
