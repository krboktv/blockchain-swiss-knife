package Knife

import (
	. "github.com/krboktv/blockchain-swiss-knife/bitcoin"
	. "github.com/krboktv/blockchain-swiss-knife/bitcoinGold"
	. "github.com/krboktv/blockchain-swiss-knife/dash"
	. "github.com/krboktv/blockchain-swiss-knife/ethereum"
	. "github.com/krboktv/blockchain-swiss-knife/ethereumClassic"
	. "github.com/krboktv/blockchain-swiss-knife/monero"
	. "github.com/krboktv/blockchain-swiss-knife/ripple"
	. "github.com/krboktv/blockchain-swiss-knife/stellar"
	. "github.com/krboktv/blockchain-swiss-knife/tether"
	. "github.com/krboktv/blockchain-swiss-knife/zcash"
)

type Knife struct {
	Bitcoin         Bitcoin
	BitcoinGold     BitcoinGold
	Dash            Dash
	Ethereum        Ethereum
	EthereumClassic EthereumClassic
	Monero          Monero
	Ripple          Ripple
	Stellar         Stellar
	Tether          Tether
	ZCash           ZCash
}
