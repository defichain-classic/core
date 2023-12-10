package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params/types/coregeth"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
)

var (
	DefiChainConfig = &coregeth.CoreGethChainConfig{
		NetworkID: 106910,
		ChainID:   big.NewInt(696969),
		Lyra2:     new(ctypes.Lyra2Config),

		EIP2FBlock: big.NewInt(0),
		EIP7FBlock: big.NewInt(0),

		EIP150Block: big.NewInt(0),

		EIP155Block: big.NewInt(0),

		// EIP158 eq
		EIP160FBlock: big.NewInt(0),
		EIP161FBlock: big.NewInt(0),
		EIP170FBlock: big.NewInt(0),

		// Byzantium eq
		EIP100FBlock: big.NewInt(0),
		EIP140FBlock: big.NewInt(0),
		EIP198FBlock: big.NewInt(0),
		EIP211FBlock: big.NewInt(0),
		EIP212FBlock: big.NewInt(0),
		EIP213FBlock: big.NewInt(0),
		EIP214FBlock: big.NewInt(0),
		EIP658FBlock: big.NewInt(0),

		// Constantinople eq, aka Agharta
		EIP145FBlock:  big.NewInt(0),
		EIP1014FBlock: big.NewInt(0),
		EIP1052FBlock: big.NewInt(0),

		// Istanbul eq, aka Phoenix
		// ECIP-1088
		EIP152FBlock:  big.NewInt(0),
		EIP1108FBlock: big.NewInt(0),
		EIP1344FBlock: big.NewInt(0),
		EIP1884FBlock: big.NewInt(0),
		EIP2028FBlock: big.NewInt(0),
		EIP2200FBlock: big.NewInt(0), // RePetersburg (== re-1283)

		ECIP1099FBlock: nil, // Etchash

		DisposalBlock:      big.NewInt(0), // Dispose difficulty bomb
		ECIP1017FBlock:     nil,           // Ethereum Classic's disinflationary monetary policy
		ECIP1017EraRounds:  nil,
		ECIP1010PauseBlock: nil, // No need to delay difficulty bomb, is defused by default
		ECIP1010Length:     nil,
		ECBP1100FBlock:     nil, // ECBP1100 (MESS artificial finality)
		RequireBlockHashes: map[uint64]common.Hash{},

		Lyra2NonceTransitionBlock: big.NewInt(0),
	}
)
