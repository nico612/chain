package client

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// NewEthClientAuthOpt new eth client auth opt
func NewEthClientAuthOpt(chainId int64, privateKey string) (*bind.TransactOpts, error) {
	ecdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(ecdsa, big.NewInt(chainId))
	if err != nil {
		return nil, err
	}
	return auth, nil
}
