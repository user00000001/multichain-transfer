module github.com/ontio/multichain-transfer

go 1.13

require (
	github.com/Workiva/go-datastructures v1.0.50 // indirect
	github.com/alecthomas/log4go v0.0.0-20180109082532-d146e6b86faa
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c // indirect
	github.com/itchyny/base58-go v0.1.0 // indirect
	github.com/ontio/ontology v1.6.1-0.20190530080932-4167ffd8250b
	github.com/ontio/ontology-crypto v1.0.7 // indirect
	github.com/ontio/ontology-eventbus v0.9.1 // indirect
	github.com/ontio/ontology-go-sdk v1.0.4-0.20190327114318-7c510760ac7a
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/urfave/cli v1.22.2
	golang.org/x/crypto v0.0.0-20191219195013-becbf705a915 // indirect
)

// exclude (
// 	github.com/siovanus/multichain-transfer v0.0.0-20190516065132-e0d83e9c1740
// 	github.com/siovanus/ontology v1.6.1-0.20190530080932-4167ffd8250b
// 	github.com/siovanus/ontology-go-sdk v1.0.4-0.20190327114318-7c510760ac7a
// )

replace (
	github.com/ontio/multichain-transfer v0.0.0-20190516065132-e0d83e9c1740 => github.com/siovanus/multichain-transfer v0.0.0-20190516065132-e0d83e9c1740
	github.com/ontio/ontology v1.6.1-0.20190530080932-4167ffd8250b => github.com/siovanus/ontology v1.6.1-0.20190530080932-4167ffd8250b
	github.com/ontio/ontology-go-sdk v1.0.4-0.20190327114318-7c510760ac7a => github.com/siovanus/ontology-go-sdk v1.0.4-0.20190327114318-7c510760ac7a
)
