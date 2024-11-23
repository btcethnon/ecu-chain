package chainid

import "math/big"

const ChainID = 3687

// 3687 ->main

const GasLimit = 1000000000000

const EnforceTips = true

const Period = 3
const Epoch = 30000

const YearBlocks = 15768000

const StartTime = 1732347524

const Reward = 10e18

const CoinBase = "0x43EA608b6443DED78881079c42f5ab185E8f8D93"

//const CoinBase = "0x0000000000000000000000000000000000999999"

const FeesCoinBase = "0x43EA608b6443DED78881079c42f5ab185E8f8D93"

//const FeesCoinBase = "0x0000000000000000000000000000000000888888"

func CalcReward(number uint64) *big.Int {
	//year := number / YearBlocks
	//
	//blockReward := uint64(Reward)
	//for i := uint64(0); i < year; i++ {
	//	blockReward = blockReward * 9 / 10
	//}
	return new(big.Int).SetUint64(uint64(Reward))
}

const Addr1 = "0xC9f241BDcE437aD1d09C1aF47cDA062Af1178781"
const Pri1 = "e9c97c9426b9fc4cf757fcfa8d3ae4680f65ae4f159cf1f1c73f72271ccc65ee"

const Addr2 = "0xB74801A6e21EC8A311F31Ac9FEA9AAFc1Be8B51E"
const Pri2 = "aa7c2418722923298ecf200944713414f7ab68551f7ca4e07cc0cdc919d79598"
