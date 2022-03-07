package nasType

// TS 24.193 6.1.3
type AtsssRule struct {
	Len                       uint16
	RuleID                    uint8
	RuleOperation             uint8
	Precedence                uint8
	LenTrafficDescriptor      uint16
	TrafficDescriptor         []AtsssTrafficDescriptor
	AccessSelectionDescriptor AtsssAccessSelectionDescriptor
}

func NewAtsssRule() (atsssRule *AtsssRule) {
	atsssRule = &AtsssRule{}
	return atsssRule
}

func (a *AtsssRule) SetLenTrafficDescriptor(len uint16) {
	a.Len = len
	a.TrafficDescriptor = make([]AtsssTrafficDescriptor, a.Len)
}

func (a *AtsssRule) GetLen() uint16 {
	return a.Len
}

func (a *AtsssRule) GetIdentifier() uint8 {
	return AtsssParameterIdentifierAtsssRule
}

// TODO
func (a *AtsssRule) Decode([]byte) error {
	return nil
}

func (a *AtsssRule) Encode() ([]byte, error) {
	var result []byte

	return result, nil
}
