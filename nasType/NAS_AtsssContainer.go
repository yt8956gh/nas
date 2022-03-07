package nasType

import (
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
	var start int
	aps := []AtsssParameter{}
	for start < int(a.Len) {
		var ap AtsssParameter
		id := a.Buffer[start]
		contentLen := binary.BigEndian.Uint16(a.Buffer[start+1 : start+3])
		start += 3

		switch id {
		case AtsssParameterIdentifierAtsssRule:
			ap = new(AtsssRule)
		case AtsssParameterIdentifierNetworkSteeringfuncInfo:
			ap = new(AtsssNetworkSteeringFuncInfo)
		case AtsssParameterIdentifierMeasurementAssistanceInfo:
			ap = new(AtsssMeasurementAssistanceInfo)
		default:
			return nil, fmt.Errorf("Unknown ATSSS Parameter id: %d", id)
		}

		if err := ap.Decode(a.Buffer[start : start+int(contentLen)]); err != nil {
			return nil, err
		}
		start += int(contentLen)
		aps = append(aps, ap)
	}
	return aps, nil
}
