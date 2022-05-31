package nasType

type AtsssParameterIdentifier uint8

// TS 24.193 6.1.2
const (
	AtsssParameterIdentifierAtsssRule                 AtsssParameterIdentifier = 0x1
	AtsssParameterIdentifierNetworkSteeringfuncInfo   AtsssParameterIdentifier = 0x2
	AtsssParameterIdentifierMeasurementAssistanceInfo AtsssParameterIdentifier = 0x3
)

type AtsssParameter interface {
	GetIdentifier() AtsssParameterIdentifier
	Decode([]byte) error
	Encode() ([]byte, error)
}
