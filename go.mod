module github.com/assetsadapterstore/syscoin-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/bitcoin-adapter v1.5.2
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.6
	github.com/pborman/uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
)

//replace github.com/blocktree/bitcoin-adapter => ../bitcoin-adapter

//replace github.com/blocktree/go-owcdrivers => ../../go-owcdrivers
