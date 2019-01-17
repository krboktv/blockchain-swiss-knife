package knife

import (
	"github.com/krboktv/blockchain-swiss-knife/bitcoin"
	"github.com/krboktv/blockchain-swiss-knife/bitcoinGold"
)

type Knife struct {
	Bitcoin     bitcoin.Bitcoin
	BitcoinGold bitcoinGold.BitcoinGold
}
