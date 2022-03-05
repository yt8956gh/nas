package nasType

// TS 24.501 9.11.4.22 ATSSS container
// ATSSSContainer
type ATSSSContainer struct {
	Iei    uint8
	Len    uint16
	Buffer []uint8
}

func NewATSSSContainer(iei uint8) (atsssContainer *ATSSSContainer) {
	atsssContainer = &ATSSSContainer{}
	atsssContainer.SetIei(iei)
	return atsssContainer
}

func (a *ATSSSContainer) SetIei(iei uint8) {
	a.Iei = iei
}

func (a *ATSSSContainer) SetLen(len uint16) {
	a.Len = len
	a.Buffer = make([]uint8, a.Len)
}

func (a *ATSSSContainer) GetLen() uint16 {
	return a.Len
}
