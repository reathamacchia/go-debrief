package debrief

import (
	"bytes"
	"testing"

	"github.com/Debrief-BC/go-debrief/common/hexutil"
)

func TestSerialize(t *testing.T) {

	nickname := "HelloWorld"
	pubKey := hexutil.MustDecode("0x3df5e2bc52991aa2e6739b7fcd08a6f32327d30cd1d2b236334d3be7e04685e82c8a4e4025e33dce0e7a484be16ec02823a3f45753b261af4fe1add2b7470dd6")

	data, err := SerializeRegisterCall(nickname, pubKey)

	if err != nil {
		t.Error(err)
	}

	nickname1, pubKey1, err := DeserializeRegisterCall(data)

	if err != nil {
		t.Error(err)
	}

	if nickname != nickname1 {
		t.Errorf("got %v, want %v", nickname, nickname1)
	}

	if !bytes.Equal(pubKey, pubKey1) {
		t.Errorf(" got %v, want %v", pubKey, pubKey1)
	}

}
