package params

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params/types/genesisT"
)

var DefiChainGenesisHash = common.HexToHash("0x49fab206bd26391d036121d239a7d95cec2b3b4747e65b61ffba1c67dae8b1c1")

func DefaultDefiChainGenesisBlock() *genesisT.Genesis {
	return &genesisT.Genesis{
		Config:     DefiChainConfig,
		Nonce:      hexutil.MustDecodeUint64("0x0"),
		ExtraData:  hexutil.MustDecode("0x42"),
		GasLimit:   hexutil.MustDecodeUint64("0x2fefd8"),
		Difficulty: hexutil.MustDecodeBig("0x2000"),
		Timestamp:  1623664673,
		Alloc:      genesisT.DecodePreAlloc(allocDefiChain),
	}
}
