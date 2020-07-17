module github.com/assetsadapterstore/kunpeng-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owaddress v1.1.11
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/openwallet/v2 v2.0.2
	github.com/blocktree/qtum-adapter v1.4.2-0.20200717093751-19ad7771dbe4
)

//replace github.com/blocktree/qtum-adapter => ../qtum-adapter
