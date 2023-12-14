package lyra_rewards

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/lyra2"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
)

var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

func _getEraUncleBlockReward(minerReward *big.Int) *big.Int {
	return new(big.Int).Div(minerReward, big32)
}

func Lyra2BlockReward(config ctypes.ChainConfigurator, header *types.Header, uncles []*types.Header) (*big.Int, []*big.Int) {
	eraLen := big.NewInt(100000)
	era := lyra2.GetBlockEra(header.Number, eraLen)
	era = era.Add(era, big.NewInt(72))

	minerReward := lyra2.GetBlockWinnerRewardByEra(era)
	uncleReward := _getEraUncleBlockReward(minerReward)

	// Create an array of *big.Int with the same length as uncles
	uncleRewards := make([]*big.Int, len(uncles))

	// Initialize each element of uncleRewards with uncleReward
	for i := range uncleRewards {
		// Deep copy the uncleReward if needed, or simply assign the pointer
		uncleRewards[i] = new(big.Int).Set(uncleReward)
	}

	return minerReward, uncleRewards
}
