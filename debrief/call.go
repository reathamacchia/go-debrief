package debrief

import (
	"errors"

	"github.com/Debrief-BC/go-debrief/common"
)

var (
	// RegisterCallAddress doc soon
	RegisterCallAddress = common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff")
)

// CallType doc soon
type CallType int

const (
	// NotCall doc soon
	NotCall CallType = iota
	// RegisterCall doc soon
	RegisterCall
)

// GetCallType doc soon
func GetCallType(address common.Address) CallType {
	switch address {
	case RegisterCallAddress:
		return RegisterCall
	default:
		return NotCall
	}
}

// SerializeRegisterCall doc soon
func SerializeRegisterCall(nickname string, pubKey []byte) ([]byte, error) {

	if len(pubKey) == 65 {
		pubKey = pubKey[1:]
	}

	if len(pubKey) != 64 {
		return nil, errors.New("invalid public key")
	}

	if len(nickname) == 0 || len(nickname) > 16 {
		return nil, errors.New("invalid nickname, should be less than 16")
	}

	return append(pubKey, ([]byte(nickname))...), nil
}

// DeserializeRegisterCall doc soon
func DeserializeRegisterCall(data []byte) (string, []byte, error) {
	if len(data) < 64 {
		return "", nil, errors.New("invalid data")
	}
	return string(data[64:]), data[:64], nil
}
