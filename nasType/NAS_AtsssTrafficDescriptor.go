package nasType

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
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

const (
	AtsssTrafficDescriptorIp3TupleFieldIPv4Address uint8 = 0b00000001 << iota
	AtsssTrafficDescriptorIp3TupleFieldIPv6Address
	AtsssTrafficDescriptorIp3TupleFieldProtocolIdNextHeader
	AtsssTrafficDescriptorIp3TupleFieldSingleRemotePort
	AtsssTrafficDescriptorIp3TupleFieldRemotePortRange
)

// TS 24.526 5.2.1
// The Set functions of some types have not been supported yet
const (
	AtsssTrafficDescriptorLenMatchAll uint8 = 0
	// AtsssTrafficDescriptorLenOsIdAnsOsAppIS                   uint8 =
	AtsssTrafficDescriptorLenIPv4RemoteAddress                uint8 = 4 + 4
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
	Buffer []uint8
}

func NewAtsssTrafficDescriptor() *AtsssTrafficDescriptor {
	return &AtsssTrafficDescriptor{}
}

func (a *AtsssTrafficDescriptor) SetTypeID(id uint8) {
	a.TypeID = id
}

func (a *AtsssTrafficDescriptor) GetTypeID() uint8 {
	return a.TypeID
}

func (a *AtsssTrafficDescriptor) SetMatchAll() {
	a.TypeID = AtsssTrafficDescriptorTypeIdMatchAll
	a.Buffer = nil
}

func (a *AtsssTrafficDescriptor) SetIPv4RemoteAddress(ip net.IPNet) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdIPv4RemoteAddress
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, ip.IP.To4()); err != nil {
		return err
	}
	if err := binary.Write(buffer, binary.BigEndian, ip.Mask); err != nil {
		return err
	}
	a.Buffer = buffer.Bytes()
	return nil
}

func (a *AtsssTrafficDescriptor) GetIPv4RemoteAddress() (*net.IPNet, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	var ip net.IPNet
	tmp := make([]byte, 4)
	if err := binary.Read(buffer, binary.BigEndian, &tmp); err != nil {
		return nil, err
	}
	ip.IP = tmp
	if err := binary.Read(buffer, binary.BigEndian, &tmp); err != nil {
		return nil, err
	}
	ip.Mask = tmp
	return &ip, nil
}

func (a *AtsssTrafficDescriptor) SetProtocolIdNextHeader(number uint8) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdProtocolIdNextHeader
	a.Buffer = []uint8{number}
	return nil
}

func (a *AtsssTrafficDescriptor) GetProtocolIdNextHeader() uint8 {
	return a.Buffer[0]
}

func (a *AtsssTrafficDescriptor) SetSingleRemotePort(port uint16) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdSingleRemotePort
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, port); err != nil {
		return err
	}
	a.Buffer = buffer.Bytes()
	return nil
}

func (a *AtsssTrafficDescriptor) GetSingleRemotePort() (uint16, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	var port uint16
	if err := binary.Read(buffer, binary.BigEndian, port); err != nil {
		return 0, err
	}
	return port, nil
}

func (a *AtsssTrafficDescriptor) SetRemotePortRange(low uint16, high uint16) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdRemotePortRange
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, low); err != nil {
		return err
	}
	if err := binary.Write(buffer, binary.BigEndian, high); err != nil {
		return err
	}
	a.Buffer = buffer.Bytes()
	return nil
}

func (a *AtsssTrafficDescriptor) GetRemotePortRange() (uint16, uint16, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	var low, high uint16
	if err := binary.Read(buffer, binary.BigEndian, low); err != nil {
		return 0, 0, err
	}
	if err := binary.Read(buffer, binary.BigEndian, high); err != nil {
		return 0, 0, err
	}
	return low, high, nil
}

// If the length of port is 1, which represents SingleRemotePort.
// If it's 2, which represents RemotePortRange.
func (a *AtsssTrafficDescriptor) SetIp3Tuple(ip net.IP, protocolIdNextHeader uint8, ports *[]uint16) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdIp3Tuple
	buffer := bytes.NewBuffer(nil)
	var bitmap byte

	// Only support IPv4
	if err := binary.Write(buffer, binary.BigEndian, ip.To4()); err != nil {
		return err
	}
	bitmap |= AtsssTrafficDescriptorIp3TupleFieldIPv4Address

	if err := binary.Write(buffer, binary.BigEndian, protocolIdNextHeader); err != nil {
		return err
	}
	bitmap |= AtsssTrafficDescriptorIp3TupleFieldProtocolIdNextHeader

	if len(*ports) == 1 {
		bitmap |= AtsssTrafficDescriptorIp3TupleFieldSingleRemotePort
	} else if len(*ports) == 2 {
		bitmap |= AtsssTrafficDescriptorIp3TupleFieldRemotePortRange
	} else {
		return fmt.Errorf("The argument \"ports\" is illegal.")
	}

	if err := binary.Write(buffer, binary.BigEndian, ports); err != nil {
		return err
	}

	a.Buffer = []byte{bitmap}
	a.Buffer = append(a.Buffer, buffer.Bytes()...)
	return nil
}

func (a *AtsssTrafficDescriptor) GetIp3Tuple() (*net.IP, *uint8, *[]uint16, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	var (
		bitmap               byte
		ip                   net.IP
		protocolIdNextHeader byte
		ports                []uint16
	)
	if err := binary.Read(buffer, binary.BigEndian, &bitmap); err != nil {
		return nil, nil, nil, err
	}

	// Only support IPv4
	if bitmap&AtsssTrafficDescriptorIp3TupleFieldIPv4Address != 0 {
		ipTmp := make([]byte, 4)
		if err := binary.Read(buffer, binary.BigEndian, &ipTmp); err != nil {
			return nil, nil, nil, err
		}
		ip = ipTmp
	}

	if bitmap&AtsssTrafficDescriptorIp3TupleFieldProtocolIdNextHeader != 0 {
		if err := binary.Read(buffer, binary.BigEndian, &protocolIdNextHeader); err != nil {
			return nil, nil, nil, err
		}
	}

	if bitmap&AtsssTrafficDescriptorIp3TupleFieldSingleRemotePort != 0 {
		if err := binary.Read(buffer, binary.BigEndian, ports[:1]); err != nil {
			return nil, nil, nil, err
		}
	} else if bitmap&AtsssTrafficDescriptorIp3TupleFieldRemotePortRange != 0 {
		if err := binary.Read(buffer, binary.BigEndian, ports[:2]); err != nil {
			return nil, nil, nil, err
		}
	}
	return &ip, &protocolIdNextHeader, &ports, nil
}

func (a *AtsssTrafficDescriptor) SetSecurityParameterIndex(spi uint32) error {
	a.TypeID = AtsssTrafficDescriptorTypeIdSecurityParameterIndex
	buffer := bytes.NewBuffer(nil)
	if err := binary.Write(buffer, binary.BigEndian, spi); err != nil {
		return err
	}
	a.Buffer = buffer.Bytes()
	return nil
}

func (a *AtsssTrafficDescriptor) GetSecurityParameterIndex() (uint32, error) {
	buffer := bytes.NewBuffer(a.Buffer)
	var spi uint32
	if err := binary.Read(buffer, binary.BigEndian, spi); err != nil {
		return 0, err
	}
	return spi, nil
}

func (a *AtsssTrafficDescriptor) Decode(b []byte) error {
	buffer := bytes.NewBuffer(b)
	if err := binary.Read(buffer, binary.BigEndian, &a.TypeID); err != nil {
		return err
	}
	if err := binary.Read(buffer, binary.BigEndian, a.Buffer); err != nil {
		return err
	}
	return nil
}

func (a *AtsssTrafficDescriptor) Encode() ([]byte, error) {
	var b []byte
	buffer := bytes.NewBuffer(b)
	if err := binary.Write(buffer, binary.BigEndian, &a.TypeID); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
