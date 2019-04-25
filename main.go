package main

import (
	"os"

	"fmt"
	"github.com/ontio/multichain-transfer/config"
	"github.com/ontio/multichain-transfer/core"
	"github.com/ontio/multichain-transfer/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

func main() {
	err := config.DefConfig.Init(config.DEFAULT_CONFIG_FILE_NAME)
	if err != nil {
		fmt.Println("DefConfig.Init error: ", err)
		return
	}
	Run(NewClient(), os.Args...)
}

func Run(client *core.Client, args ...string) {
	app := cli.NewApp()
	app.Usage = "cli for multi chain transfer tools"
	app.Action = client.Lock
	app.Flags = []cli.Flag{
		//common setting
		utils.FeeFlag,
		utils.ChainIDFlag,
		utils.WalletFileFlag,
		utils.AmountFlag,
	}

	app.Run(args)
}

func NewClient() *core.Client {
	ontSdk := sdk.NewOntologySdk()
	ontSdk.NewRpcClient().SetAddress(config.DefConfig.JsonRpcAddress)
	return &core.Client{
		OntSdk: ontSdk,
	}
}
