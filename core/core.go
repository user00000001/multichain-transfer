package core

import (
	"fmt"

	"github.com/siovanus/multichain-transfer/config"
	"github.com/siovanus/multichain-transfer/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common/password"
	"github.com/ontio/ontology/common"
	outils "github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/urfave/cli"
)

var OntIDVersion = byte(0)

type Client struct {
	OntSdk *sdk.OntologySdk
}

type OngLockParam struct {
	Fee	uint64
	ToChainID uint64
	Address   common.Address
	Amount    uint64
}

func (client *Client) Lock(c *cli.Context) {
	wallet, err := client.OntSdk.OpenWallet(c.String(utils.GetFlagName(utils.WalletFileFlag)))
	if err != nil {
		fmt.Println("client.OntSdk.OpenWallet error: ", err)
		return
	}
	pwd, err := password.GetPassword()
	if err != nil {
		fmt.Println("password.GetPassword error: ", err)
		return
	}
	user, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		fmt.Println("wallet.GetDefaultAccount error: ", err)
		return
	}
	params := &OngLockParam{
		Fee:       c.Uint64(utils.GetFlagName(utils.FeeFlag)),
		ToChainID: c.Uint64(utils.GetFlagName(utils.ChainIDFlag)),
		Address:   user.Address,
		Amount:    c.Uint64(utils.GetFlagName(utils.AmountFlag)),
	}
	// params := &ong.OngLockParam{
	// 	Fee:       c.Uint64(utils.GetFlagName(utils.FeeFlag)),
	// 	ToChainID: c.Uint64(utils.GetFlagName(utils.ChainIDFlag)),
	// 	Address:   user.Address,
	// 	Amount:    c.Uint64(utils.GetFlagName(utils.AmountFlag)),
	// }
	method := "ongLock"
	contractAddress := outils.OngContractAddress
	txHash, err := client.OntSdk.Native.InvokeNativeContract(config.DefConfig.GasPrice,
		config.DefConfig.GasLimit, user, OntIDVersion, contractAddress, method, []interface{}{params})
	// txHash, err := client.OntSdk.Native.InvokeNativeContract(config.DefConfig.ChainID, config.DefConfig.GasPrice,
	// 	config.DefConfig.GasLimit, user, OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		fmt.Println("invokeNativeContract error :", err)
		return
	}
	fmt.Println("txhash is: ", txHash.ToHexString())
}
