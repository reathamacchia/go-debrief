package debrief

import (
	"encoding/json"

	"github.com/Debrief-BC/go-debrief/common"
	"github.com/Debrief-BC/go-debrief/common/hexutil"
)

// User doc soon
type User struct {
	Address   common.Address
	Nickname  []byte
	PublicKey []byte
}

// MarshalJSON doc soon
func (u User) MarshalJSON() ([]byte, error) {
	type User struct {
		Address   common.Address
		Nickname  string
		PublicKey string
	}
	var user User

	user.Address = u.Address
	user.Nickname = string(u.Nickname)
	user.PublicKey = hexutil.Encode(u.PublicKey)

	return json.Marshal(&user)
}
