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
			id     uint8
			length uint16
		)

		if err := binary.Read(buffer, binary.BigEndian, &id); err != nil {
			return nil, err
		}

		switch id {
		case AtsssParameterIdentifierAtsssRule:
			ap = NewAtsssRule()
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
