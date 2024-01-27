package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/evmos/evmos/v12/x/evm/precompiles/contracts"
)

const (
	EmptyEvmAddress = "0x0000000000000000000000000000000000000000"
	StakingAddress  = "0x0000000000000000000000000000000000001003"
)

type Contract struct {
	Address common.Address
	ABI     abi.ABI
	Bin     []byte
	Code    []byte
}

func (c Contract) CodeHash() common.Hash {
	return crypto.Keccak256Hash(c.Code)
}

var (
	erc1967Proxy = Contract{
		Address: common.Address{},
		ABI:     MustABIJson(contracts.ERC1967ProxyMetaData.ABI),
		Bin:     MustDecodeHex(contracts.ERC1967ProxyMetaData.Bin),
		Code:    []byte{},
	}
)

func GetERC1967Proxy() Contract {
	return erc1967Proxy
}

func MustDecodeHex(str string) []byte {
	bz, err := hexutil.Decode(str)
	if err != nil {
		panic(err)
	}
	return bz
}

func MustABIJson(str string) abi.ABI {
	j, err := abi.JSON(strings.NewReader(str))
	if err != nil {
		panic(err)
	}
	return j
}
