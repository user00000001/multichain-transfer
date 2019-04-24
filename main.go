package main

import (
	"os"

	"github.com/ontio/multichain-transfer/config"
	"github.com/ontio/multichain-transfer/core"
	"github.com/ontio/multichain-transfer/utils"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/urfave/cli"
)

func main() {
	Run(NewClient(), os.Args...)
}

func Run(client *core.Client, args ...string) {
	app := cli.NewApp()
	app.Usage = "cli for multi chain transfer tools"
	app.Commands = []cli.Command{
		{
			Name:   "lock",
			Usage:  "lock ong",
			Action: client.Lock,
		},
	}
	app.Flags = []cli.Flag{
		utils.ConfigFlag,
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
