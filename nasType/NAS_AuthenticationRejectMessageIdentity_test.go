package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yt8956gh/nas/nasMessage"
	"github.com/yt8956gh/nas/nasType"
)

type nasTypeRejectMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeRejectMessageIdentityTable = []nasTypeRejectMessageIdentityData{
	{nasMessage.PDUSessionEstablishmentRejectEAPMessageType, nasMessage.PDUSessionEstablishmentRejectEAPMessageType},
}

func TestNasTypeNewAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetAuthenticationRejectMessageIdentity(t *testing.T) {
	a := nasType.NewAuthenticationRejectMessageIdentity()
	for _, table := range nasTypeRejectMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
