package params

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params/types/genesisT"
)

var DefiChainGenesisHash = common.HexToHash("0x107687ef26cd54744a0c60d8b1e88b28f1f768caeef8d0459b577cf47e8cccd4")

func DefaultDefiChainGenesisBlock() *genesisT.Genesis {
	return &genesisT.Genesis{
		Config:     DefiChainConfig,
		Nonce:      hexutil.MustDecodeUint64("0x0"),
		ExtraData:  hexutil.MustDecode("0x42"),
		GasLimit:   hexutil.MustDecodeUint64("0x2fefd80"),
		Difficulty: hexutil.MustDecodeBig("0x2000"),
		Timestamp:  1623664673,
		Alloc:      genesisT.DecodePreAlloc(allocDefiChain),
	}
}
