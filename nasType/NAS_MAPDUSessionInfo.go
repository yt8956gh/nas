package nasType

// TS 24.501 9.11.3.31A MA PDU session information
// IEI Row, sBit, len = [0, 0], 8 , 4
// InfoValue Row, sBit, len = [0, 0], 4 , 4
type MAPDUSessionInfo struct {
	Octet uint8
}

const (
	MAPDUSessionNetworkUpgradeAllowed uint8 = 0x01
)

func NewMAPDUSessionInfo(iei uint8) (mAPDUSessionInfo *MAPDUSessionInfo) {
	mAPDUSessionInfo = &MAPDUSessionInfo{}
	mAPDUSessionInfo.SetIei(iei)
	return mAPDUSessionInfo
}

func (m *MAPDUSessionInfo) SetIei(iei uint8) {
	m.Octet = (m.Octet & 15) + ((iei & 15) << 4)
}

func (m *MAPDUSessionInfo) GetIei() (iei uint8) {
	return m.Octet & GetBitMask(8, 4) >> (4)
}

func (m *MAPDUSessionInfo) SetValue(v uint8) {
	m.Octet = (m.Octet & 240) + (v & 15)
}

func (m *MAPDUSessionInfo) GetValue() (v uint8) {
	return m.Octet & GetBitMask(4, 0)
}
