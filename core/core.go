package core

import (
	"fmt"
	"github.com/ontio/multichain-transfer/config"
	"github.com/ontio/multichain-transfer/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common/password"
	"github.com/ontio/ontology/smartcontract/service/native/ong"
	outils "github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/urfave/cli"
)

var OntIDVersion = byte(0)

type Client struct {
	OntSdk *sdk.OntologySdk
}

func (client *Client) Lock(c *cli.Context) {
	err := config.DefConfig.Init(c.String(utils.GetFlagName(utils.ConfigFlag)))
	if err != nil {
		fmt.Println("DefConfig.Init error: ", err)
		return
	}
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
	params := &ong.OngLockParam{
		OngxFee:    c.Uint64(utils.GetFlagName(utils.FeeFlag)),
		ToChainID:  c.Uint64(utils.GetFlagName(utils.ChainIDFlag)),
		Address:    user.Address,
		OngxAmount: c.Uint64(utils.GetFlagName(utils.AmountFlag)),
	}
	method := "ongLock"
	contractAddress := outils.OngContractAddress
	txHash, err := client.OntSdk.Native.InvokeNativeContract(config.DefConfig.ChainID, config.DefConfig.GasPrice,
		config.DefConfig.GasLimit, user, OntIDVersion, contractAddress, method, []interface{}{params})
	if err != nil {
		fmt.Println("invokeNativeContract error :", err)
		return
	}
	fmt.Println("txhash is: ", txHash)
}
