package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// TS 24.193 6.1.3
type AtsssRule struct {
	Len                       uint16
	RuleID                    uint8
	RuleOperation             uint8
	Precedence                uint8
	LenTrafficDescriptor      uint16
	TrafficDescriptor         *AtsssTrafficDescriptor
	AccessSelectionDescriptor *AtsssAccessSelectionDescriptor
}

func NewAtsssRule() *AtsssRule {
	a := &AtsssRule{}
	a.TrafficDescriptor = NewAtsssTrafficDescriptor()
	a.AccessSelectionDescriptor = NewAtsssAccessSelectionDescriptor()
	return a
}

func (a *AtsssRule) SetLenTrafficDescriptor(len uint16) {
	a.Len = len
}

func (a *AtsssRule) GetLen() uint16 {
	return a.Len
}

func (a *AtsssRule) GetIdentifier() uint8 {
	return AtsssParameterIdentifierAtsssRule
}

// TODO
func (a *AtsssRule) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
		return err
	}
	if buffer.Len() != int(a.Len) {
		return fmt.Errorf("The length of data doesn't match length field.")
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RuleID); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.RuleOperation); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Precedence); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.LenTrafficDescriptor); err != nil {
		return err
	}

	content := make([]byte, a.LenTrafficDescriptor)
	if err := binary.Read(buffer, binary.BigEndian, content[:]); err != nil {
		return err
	}
	if err := a.TrafficDescriptor.Decode(content); err != nil {
		return err
	}

	content = make([]byte, buffer.Len())
	if err := binary.Read(buffer, binary.BigEndian, content[:]); err != nil {
		return err
	}
	if err := a.AccessSelectionDescriptor.Decode(content); err != nil {
		return err
	}

	return nil
}

func (a *AtsssRule) Encode() ([]byte, error) {
	var result []byte

	return result, nil
}
