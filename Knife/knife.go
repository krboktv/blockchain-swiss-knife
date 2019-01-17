package Knife

import (
	. "github.com/krboktv/blockchain-swiss-knife/bitcoin"
	. "github.com/krboktv/blockchain-swiss-knife/bitcoinGold"
	. "github.com/krboktv/blockchain-swiss-knife/dash"
	. "github.com/krboktv/blockchain-swiss-knife/ethereum"
	. "github.com/krboktv/blockchain-swiss-knife/ethereumClassic"
	. "github.com/krboktv/blockchain-swiss-knife/ripple"
)

type Knife struct {
	Bitcoin         Bitcoin
	BitcoinGold     BitcoinGold
	Dash            Dash
	Ethereum        Ethereum
	EthereumClassic EthereumClassic
	Ripple          Ripple
}
