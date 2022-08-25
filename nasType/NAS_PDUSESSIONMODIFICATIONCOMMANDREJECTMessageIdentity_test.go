package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yt8956gh/nas"
	"github.com/yt8956gh/nas/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationCommandReject, nas.MsgTypePDUSessionModificationCommandReject},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONCOMMANDREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
