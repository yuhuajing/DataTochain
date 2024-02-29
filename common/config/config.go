package config

import (
	"main/common/ethconn"
)

var (
	EthServer = "https://eth-sepolia.g.alchemy.com/v2/t39LUOfkEMNMz_uxQk2gkwK3kJ1HwDZF"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0xC6F48C4a0EEc939b0841c07AAe1836b179EA7a84"
)

// 浏览器
// https://sepolia.etherscan.io/address/0xC6F48C4a0EEc939b0841c07AAe1836b179EA7a84
