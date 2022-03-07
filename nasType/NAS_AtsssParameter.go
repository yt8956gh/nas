package nasType

// TS 24.193 6.1.2
const (
	AtsssParameterIdentifierAtsssRule                 uint8 = 0x1
	AtsssParameterIdentifierNetworkSteeringfuncInfo   uint8 = 0x2
	AtsssParameterIdentifierMeasurementAssistanceInfo uint8 = 0x3
)

type AtsssParameter interface {
	GetIdentifier() uint8
	Decode([]byte) error
	Encode() ([]byte, error)
}
