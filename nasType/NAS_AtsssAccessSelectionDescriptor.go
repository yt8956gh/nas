package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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

func NewAtsssAccessSelectionDescriptor() *AtsssAccessSelectionDescriptor {
	return &AtsssAccessSelectionDescriptor{}
}

func (a *AtsssAccessSelectionDescriptor) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, a.Len); err != nil {
		return err
	}
	if buffer.Len() != int(a.Len) {
		return fmt.Errorf("The length of data doesn't match length field.")
	}
	if err := binary.Read(buffer, binary.BigEndian, a.SteeringFunc); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, a.SteeringMode); err != nil {
		return err
	}

	if a.SteeringMode == AtsssAccessSelectionDescriptorSteeringModeSmallestDelay {
		if a.Len != AtsssAccessSelectionDescriptorLenSmallestDelay {
			return fmt.Errorf("The length of smallest delay selection descriptor doesn't match spec.")
		}
	} else if a.SteeringMode == AtsssAccessSelectionDescriptorSteeringModeActiveStandby ||
		a.SteeringMode == AtsssAccessSelectionDescriptorSteeringModeLoadBalancing ||
		a.SteeringMode == AtsssAccessSelectionDescriptorSteeringModePriorityBased {
		if err := binary.Read(buffer, binary.BigEndian, a.SteeringModeInfo); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("The steering mode doesn't match anything.")
	}

	return nil
}

func (a *AtsssAccessSelectionDescriptor) Encode() ([]byte, error) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	if err := binary.Write(buffer, binary.BigEndian, &a.Len); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SteeringFunc); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.SteeringMode); err != nil {
		return nil, err
	}

	if a.SteeringMode != AtsssAccessSelectionDescriptorSteeringModeSmallestDelay {
		if err := binary.Write(buffer, binary.BigEndian, &a.SteeringModeInfo); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
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
