package msp

import (
	"os"
	"reflect"
	"testing"

	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/stretchr/testify/assert"
)

func TestNoopMSP(t *testing.T) {
	noopmsp := NewNoopMsp()

	id, err := noopmsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := noopmsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("foo")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}
}

func TestMSPSetupBad(t *testing.T) {
	_, err := GetLocalMspConfig("barf")
	if err == nil {
		t.Fatalf("Setup should have failed on an invalid config file")
		return
	}
}

func TestGetIdentities(t *testing.T) {
	_, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetDefaultSigningIdentity failed with err %s", err)
		return
	}
}

func TestSerializeIdentities(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded, got err %s", err)
		return
	}

	err = localMsp.Validate(idBack)
	if err != nil {
		t.Fatalf("The identity should be valid, got err %s", err)
		return
	}

	if !reflect.DeepEqual(id.GetPublicVersion(), idBack) {
		t.Fatalf("Identities should be equal (%s) (%s)", id, idBack)
		return
	}
}

func TestSerializeIdentitiesWithWrongMSP(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	sid := &SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = localMsp.DeserializeIdentity(serializedID)
	assert.Error(t, err)
}

func TestSerializeIdentitiesWithMSPManager(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded, got err %s", err)
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded, got err %s", err)
		return
	}

	_, err = mspMgr.DeserializeIdentity(serializedID)
	assert.NoError(t, err)

	sid := &SerializedIdentity{}
	err = proto.Unmarshal(serializedID, sid)
	assert.NoError(t, err)

	sid.Mspid += "BARF"

	serializedID, err = proto.Marshal(sid)
	assert.NoError(t, err)

	_, err = mspMgr.DeserializeIdentity(serializedID)
	assert.Error(t, err)
}

func TestSignAndVerify(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("foo")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}
}

func TestSignAndVerify_longMessage(t *testing.T) {
	id, err := localMsp.GetDefaultSigningIdentity()
	if err != nil {
		t.Fatalf("GetSigningIdentity should have succeeded")
		return
	}

	serializedID, err := id.Serialize()
	if err != nil {
		t.Fatalf("Serialize should have succeeded")
		return
	}

	idBack, err := localMsp.DeserializeIdentity(serializedID)
	if err != nil {
		t.Fatalf("DeserializeIdentity should have succeeded")
		return
	}

	msg := []byte("ABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFGABCDEFG")
	sig, err := id.Sign(msg)
	if err != nil {
		t.Fatalf("Sign should have succeeded")
		return
	}

	err = id.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}

	err = idBack.Verify(msg, sig)
	if err != nil {
		t.Fatalf("The signature should be valid")
		return
	}
}

var conf *msp.MSPConfig
var localMsp MSP
var mspMgr MSPManager

func TestMain(m *testing.M) {
	var err error
	conf, err = GetLocalMspConfig("./sampleconfig/")
	if err != nil {
		fmt.Printf("Setup should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	localMsp, err = NewBccspMsp()
	if err != nil {
		fmt.Printf("Constructor for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	err = localMsp.Setup(conf)
	if err != nil {
		fmt.Printf("Setup for msp should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	mspMgr = NewMSPManager()
	err = mspMgr.Setup([]*msp.MSPConfig{conf})
	if err != nil {
		fmt.Printf("Setup for msp manager should have succeeded, got err %s instead", err)
		os.Exit(-1)
	}

	retVal := m.Run()
	os.Exit(retVal)
}
