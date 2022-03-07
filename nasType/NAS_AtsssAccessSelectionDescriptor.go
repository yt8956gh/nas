package nasType

// TS 24.193 6.1.3.2
// Length of access selection descriptor
const (
	AtsssAccessSelectionDescriptorLenSmallestDelay = 0x3 + uint8(iota)
	AtsssAccessSelectionDescriptorLenNotSmallestDelay
)

// Steering functionality
const (
	AtsssAccessSelectionDescriptorSteeringFuncUeSupported = uint8(iota)
	AtsssAccessSelectionDescriptorSteeringFuncMPTCP
	AtsssAccessSelectionDescriptorSteeringFuncAtsssLL
)

// Steering mode
const (
	AtsssAccessSelectionDescriptorSteeringModeActiveStandby = uint8(iota)
	AtsssAccessSelectionDescriptorSteeringModeSmallestDelay
	AtsssAccessSelectionDescriptorSteeringModeLoadBalancing
	AtsssAccessSelectionDescriptorSteeringModePriorityBased
)

// Active-standby steering mode information
const (
	AtsssAccessSelectionDescriptorActive3gppAndNoStandby = uint8(iota)
	AtsssAccessSelectionDescriptorActive3gppAndNon3gppStandby
	AtsssAccessSelectionDescriptorActiveNon3gppAndNoStandby
	AtsssAccessSelectionDescriptorActiveNon3gppAnd3gppStandby
)

// Load-balancing steering mode information
const (
	AtsssAccessSelectionDescriptorLoadBalance3gpp100Percent = uint8(iota)
	AtsssAccessSelectionDescriptorLoadBalance3gpp90Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp80Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp70Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp60Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp50Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp40Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp30Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp20Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp10Percent
	AtsssAccessSelectionDescriptorLoadBalance3gpp0Percent
)

//  Priority-based steering mode information
const (
	AtsssAccessSelectionDescriptorLoadbalancePriorityBased3gppHigher = uint8(iota)
	AtsssAccessSelectionDescriptorLoadbalancePriorityBasedNon3gppHigher
)

type AtsssAccessSelectionDescriptor struct {
	Len              uint8
	SteeringFunc     uint8
	SteeringMode     uint8
	SteeringModeInfo uint8
}

func NewAtsssAccessSelectionDescriptor(len, steeringFunc, steeringMode, steeringModeInfo uint8) (
	a *AtsssAccessSelectionDescriptor) {
	a = &AtsssAccessSelectionDescriptor{
		Len:              len,
		SteeringFunc:     steeringFunc,
		SteeringMode:     steeringMode,
		SteeringModeInfo: steeringModeInfo,
	}
	return a
}

func (a *AtsssAccessSelectionDescriptor) SetLen(len uint8) {
	a.Len = len
}

func (a *AtsssAccessSelectionDescriptor) GetLen() uint8 {
	return a.Len
}

func (a *AtsssAccessSelectionDescriptor) SetSteeringFunc(sf uint8) {
	a.SteeringFunc = sf
}

func (a *AtsssAccessSelectionDescriptor) GetSteeringFunc() uint8 {
	return a.SteeringFunc
}

func (a *AtsssAccessSelectionDescriptor) SetSteeringMode(sm uint8) {
	a.SteeringMode = sm
}

func (a *AtsssAccessSelectionDescriptor) GetSteeringMode() uint8 {
	return a.SteeringMode
}

func (a *AtsssAccessSelectionDescriptor) SetSteeringModeInfo(smi uint8) {
	a.SteeringModeInfo = smi
}

func (a *AtsssAccessSelectionDescriptor) GetSteeringModeInfo() uint8 {
	return a.SteeringModeInfo
}
