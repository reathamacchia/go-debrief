package ethapi

import (
	"context"
	"errors"
	"math/big"

	"github.com/Debrief-BC/go-debrief/accounts"
	"github.com/Debrief-BC/go-debrief/common"
	"github.com/Debrief-BC/go-debrief/common/hexutil"
	"github.com/Debrief-BC/go-debrief/debrief"
	"github.com/Debrief-BC/go-debrief/rpc"
)

// PublicDebriefAPI  doc soon
type PublicDebriefAPI struct {
	b         Backend
	nonceLock *AddrLocker
}

// NewPublicDebriefAPI creates a new Debrief protocol API.
func NewPublicDebriefAPI(b Backend, nonceLock *AddrLocker) *PublicDebriefAPI {
	return &PublicDebriefAPI{b, nonceLock}
}

// Register   doc soon
func (s *PublicDebriefAPI) Register(ctx context.Context, address common.Address, nickname string, publicKey hexutil.Bytes, blockNrOrHash rpc.BlockNumberOrHash) (common.Hash, error) {

	state, _, err := s.b.StateAndHeaderByNumberOrHash(ctx, blockNrOrHash)
	if state == nil || err != nil {
		return common.Hash{}, err
	}

	u, err := state.GetUser(address)
	if err != nil {
		return common.Hash{}, err
	}
	if u != nil {
		return common.Hash{}, errors.New("Registered")
	}

	data, err := debrief.SerializeRegisterCall(nickname, publicKey)
	if err != nil {
		return common.Hash{}, err
	}

	account := accounts.Account{Address: address}
	wallet, err := s.b.AccountManager().Find(account)
	if err != nil {
		return common.Hash{}, err
	}

	var argsData = hexutil.Bytes(data)
	args := SendTxArgs{
		From:  address,
		To:    &debrief.RegisterCallAddress,
		Value: (*hexutil.Big)(big.NewInt(0)),
		Data:  &argsData,
	}

	if args.Nonce == nil {
		s.nonceLock.LockAddr(args.From)
		defer s.nonceLock.UnlockAddr(args.From)
	}

	if err := args.setDefaults(ctx, s.b); err != nil {
		return common.Hash{}, err
	}

	tx := args.toTransaction()
	signed, err := wallet.SignTx(account, tx, s.b.ChainConfig().ChainID)

	if err != nil {
		return common.Hash{}, err
	}
	return SubmitTransaction(ctx, s.b, signed)
}

// GetUser   doc soon
func (s *PublicDebriefAPI) GetUser(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*debrief.User, error) {
	state, _, err := s.b.StateAndHeaderByNumberOrHash(ctx, blockNrOrHash)
	if state == nil || err != nil {
		return nil, err
	}
	return state.GetUser(address)
}

// GetUsers   doc soon
func (s *PublicDebriefAPI) GetUsers(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]debrief.User, error) {
	state, _, err := s.b.StateAndHeaderByNumberOrHash(ctx, blockNrOrHash)
	if state == nil || err != nil {
		return nil, err
	}
	return state.GetUsers()
}
