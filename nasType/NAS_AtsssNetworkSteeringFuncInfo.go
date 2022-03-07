package nasType

// TS 24.193 6.1.4
type AtsssNetworkSteeringFuncInfo struct {
	Ue3gppIpType        uint8
	Ue3gppIpAddr        []byte
	UeNon3gppIpType     uint8
	UeNon3gppIpAddr     []byte
	LenOfMptcpProxyInfo uint8
	MptcpProxyInfoList  []MptcpProxyInfo
}

func (a *AtsssNetworkSteeringFuncInfo) GetIdentifier() uint8 {
	return AtsssParameterIdentifierNetworkSteeringfuncInfo
}

func (a *AtsssNetworkSteeringFuncInfo) Decode([]byte) error {
	return nil
}

func (a *AtsssNetworkSteeringFuncInfo) Encode() ([]byte, error) {
	var result []byte

	return result, nil
}

type MptcpProxyInfo struct {
	IpAddrType uint8
	IpAddr     []byte
	Port       uint16
	Type       uint8
}
