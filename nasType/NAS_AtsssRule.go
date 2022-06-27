package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type AtsssRules struct {
	atsssRuleList []*AtsssRule
}

func NewAtsssRules() *AtsssRules {
	as := &AtsssRules{
		atsssRuleList: make([]*AtsssRule, 0),
	}
	return as
}

func (as *AtsssRules) AddAtsssRule(a *AtsssRule) {
	as.atsssRuleList = append(as.atsssRuleList, a)
}

func (as *AtsssRules) GetIdentifier() AtsssParameterIdentifier {
	return AtsssParameterIdentifierAtsssRule
}

func (as *AtsssRules) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)

	for buffer.Len() > 0 {
		a := NewAtsssRule()
		if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
			return fmt.Errorf("binary.Read Len Fail: %+v", err)
		}
		if buffer.Len() < int(a.Len)-2 {
			return fmt.Errorf("The length of ATSSS rules doesn't match length field. %d, %d.", buffer.Len(), int(a.Len))
		}
		if err := binary.Read(buffer, binary.BigEndian, &a.RuleID); err != nil {
			return fmt.Errorf("binary.Read RuleID Fail: %+v", err)
		}
		if err := binary.Read(buffer, binary.BigEndian, &a.RuleOperation); err != nil {
			return fmt.Errorf("binary.Read RuleOperation Fail: %+v", err)
		}
		if err := binary.Read(buffer, binary.BigEndian, &a.Precedence); err != nil {
			return fmt.Errorf("binary.Read Precedence Fail: %+v", err)
		}
		if err := binary.Read(buffer, binary.BigEndian, &a.LenTrafficDescriptor); err != nil {
			return fmt.Errorf("binary.Read LenTrafficDescriptor Fail: %+v", err)
		}
		if a.LenTrafficDescriptor != 0 {
			content := make([]byte, a.LenTrafficDescriptor)
			if err := binary.Read(buffer, binary.BigEndian, content[:]); err != nil {
				return fmt.Errorf("LenTrafficDescriptor Fail: %+v", err)
			}
			if err := a.TrafficDescriptor.Decode(content); err != nil {
				return fmt.Errorf("TrafficDescriptor decode Fail: %+v", err)
			}
		}

		content := make([]byte, a.Len-a.LenTrafficDescriptor-7)
		if err := binary.Read(buffer, binary.BigEndian, content[:]); err != nil {
			return err
		}
		if err := a.AccessSelectionDescriptor.Decode(content); err != nil {
			return err
		}
		as.atsssRuleList = append(as.atsssRuleList, a)
	}
	return nil
}

func (as *AtsssRules) Encode() ([]byte, error) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	for _, a := range as.atsssRuleList {
		if err := binary.Write(buffer, binary.BigEndian, &a.Len); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RuleID); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.RuleOperation); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.Precedence); err != nil {
			return nil, err
		}
		if err := binary.Write(buffer, binary.BigEndian, &a.LenTrafficDescriptor); err != nil {
			return nil, err
		}

		if err := binary.Write(buffer, binary.BigEndian, a.TrafficDescriptor.Buffer); err != nil {
			return nil, err
		}

		content, err := a.AccessSelectionDescriptor.Encode()
		if err != nil {
			return nil, err
		}

		if err := binary.Write(buffer, binary.BigEndian, content); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

type AtsssRuleOperationCode uint8

const (
	OperationCodeCreateOrReplaceAtsssRule = AtsssRuleOperationCode(iota) + 1
	OperationCodeDeleteExistingAtsssRule
)

// TS 24.193 6.1.3
type AtsssRule struct {
	Len                       uint16
	RuleID                    uint8
	RuleOperation             AtsssRuleOperationCode
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

func (a *AtsssRule) SetLen(len uint16) {
	a.Len = len
}

func (a *AtsssRule) GetLen() uint16 {
	return a.Len
}

func (a *AtsssRule) SetRuleID(ri uint8) {
	a.RuleID = ri
}

func (a *AtsssRule) GetRuleID() uint8 {
	return a.RuleID
}

func (a *AtsssRule) SetRuleOperation(ro AtsssRuleOperationCode) {
	a.RuleOperation = ro
}

func (a *AtsssRule) GetRuleOperation() AtsssRuleOperationCode {
	return a.RuleOperation
}

func (a *AtsssRule) SetPrecedence(p uint8) {
	a.Precedence = p
}

func (a *AtsssRule) GetPrecedence() uint8 {
	return a.Precedence
}

func (a *AtsssRule) SetLenTrafficDescriptor(len uint16) {
	a.LenTrafficDescriptor = len
}

func (a *AtsssRule) GetLenTrafficDescriptor() uint16 {
	return a.LenTrafficDescriptor
}

func (a *AtsssRule) SetTrafficDescriptor(td *AtsssTrafficDescriptor) {
	a.TrafficDescriptor = td
}

func (a *AtsssRule) GetTrafficDescriptor() *AtsssTrafficDescriptor {
	return a.TrafficDescriptor
}

func (a *AtsssRule) SetAccessSelectionDescriptor(sd *AtsssAccessSelectionDescriptor) {
	a.AccessSelectionDescriptor = sd
}

func (a *AtsssRule) GetAccessSelectionDescriptor() *AtsssAccessSelectionDescriptor {
	return a.AccessSelectionDescriptor
}
