package nasType

import (
	"bytes"
	"encoding/binary"
)

// TS 24.526 5.2.1
const (
	AtsssTrafficDescriptorTypeIdMatchAll                         uint8 = 0b00000001
	AtsssTrafficDescriptorTypeIdOsIdAnsOsAppIS                   uint8 = 0b00001000
	AtsssTrafficDescriptorTypeIdIPv4RemoteAddress                uint8 = 0b00010000
	AtsssTrafficDescriptorTypeIdIPv6RemoteAddressAndPrefixLength uint8 = 0b00100001
	AtsssTrafficDescriptorTypeIdProtocolIdNextHeader             uint8 = 0b00110000
	AtsssTrafficDescriptorTypeIdSingleRemotePort                 uint8 = 0b01010000
	AtsssTrafficDescriptorTypeIdRemotePortRange                  uint8 = 0b01010001
	AtsssTrafficDescriptorTypeIdIp3Tuple                         uint8 = 0b01010010
	AtsssTrafficDescriptorTypeIdSecurityParameterIndex           uint8 = 0b01100000
	AtsssTrafficDescriptorTypeIdTypeOfServiceAndTrafficClass     uint8 = 0b01110000
	AtsssTrafficDescriptorTypeIdFlowLabel                        uint8 = 0b10000000
	AtsssTrafficDescriptorTypeIdDestinationMacAddress            uint8 = 0b10000001
	AtsssTrafficDescriptorTypeIdIEEE_802_1Q_C_TAG_VID            uint8 = 0b10000011
	AtsssTrafficDescriptorTypeIdIEEE_802_1Q_S_TAG_VID            uint8 = 0b10000100
	AtsssTrafficDescriptorTypeIdIEEE_802_1Q_C_TAG_PCP_DEI        uint8 = 0b10000101
	AtsssTrafficDescriptorTypeIdIEEE_802_1Q_S_TAG_PCP_DEI        uint8 = 0b10000110
	AtsssTrafficDescriptorTypeIdEthertype                        uint8 = 0b10000111
	AtsssTrafficDescriptorTypeIdDNN                              uint8 = 0b10001000
	AtsssTrafficDescriptorTypeIdDestinationFQDN                  uint8 = 0b10010001
	AtsssTrafficDescriptorTypeIdRegularExpression                uint8 = 0b10010010
	AtsssTrafficDescriptorTypeIdOsAppID                          uint8 = 0b10100000
)

// TS 24.526 5.2.1
// Some decoders have not been supported yet
const (
	AtsssTrafficDescriptorLenMatchAll uint8 = 0
	// AtsssTrafficDescriptorLenOsIdAnsOsAppIS                   uint8 =
	AtsssTrafficDescriptorLenIPv4RemoteAddress                uint8 = 4
	AtsssTrafficDescriptorLenIPv6RemoteAddressAndPrefixLength uint8 = 16 + 1
	AtsssTrafficDescriptorLenProtocolIdNextHeader             uint8 = 1
	AtsssTrafficDescriptorLenSingleRemotePort                 uint8 = 2
	AtsssTrafficDescriptorLenRemotePortRange                  uint8 = 2 + 2
	AtsssTrafficDescriptorLenIp3Tuple                         uint8 = 0
	AtsssTrafficDescriptorLenSecurityParameterIndex           uint8 = 4
	AtsssTrafficDescriptorLenTypeOfServiceAndTrafficClass     uint8 = 1 + 1
	// AtsssTrafficDescriptorLenFlowLabel                        uint8 =
	AtsssTrafficDescriptorLenDestinationMacAddress uint8 = 6
	// AtsssTrafficDescriptorLenIEEE_802_1Q_C_TAG_VID            uint8 =
	// AtsssTrafficDescriptorLenIEEE_802_1Q_S_TAG_VID            uint8 =
	// AtsssTrafficDescriptorLenIEEE_802_1Q_C_TAG_PCP_DEI        uint8 =
	// AtsssTrafficDescriptorLenIEEE_802_1Q_S_TAG_PCP_DEI        uint8 =
	AtsssTrafficDescriptorLenEthertype uint8 = 16
	AtsssTrafficDescriptorLenDNN       uint8 = 2
	// AtsssTrafficDescriptorLenDestinationFQDN                  uint8 =
	// AtsssTrafficDescriptorLenRegularExpression                uint8 =
	// AtsssTrafficDescriptorLenOsAppID                          uint8 =
)

type AtsssTrafficDescriptor struct {
	TypeID uint8
	Len    uint8
	Buffer []uint8
}

func NewAtsssTrafficDescriptor() *AtsssTrafficDescriptor {
	return &AtsssTrafficDescriptor{}
}

func (a *AtsssTrafficDescriptor) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, &a.TypeID); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Len); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, a.Buffer[:a.Len]); err != nil {
		return err
	}

	return nil
}
