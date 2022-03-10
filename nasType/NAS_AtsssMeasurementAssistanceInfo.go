package nasType

// TS 24.193 6.1.5.2
// PMF MAC addr hasn't been supported
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

func (a *AtsssMeasurementAssistanceInfo) GetIdentifier() uint8 {
	return AtsssParameterIdentifierMeasurementAssistanceInfo
}

func (a *AtsssMeasurementAssistanceInfo) Decode([]byte) error {
	return nil
}

func (a *AtsssMeasurementAssistanceInfo) Encode() ([]byte, error) {
	var result []byte

	return result, nil
}
