package utils

import (
	"github.com/ontio/multichain-transfer/config"
	"github.com/urfave/cli"
	"strings"
)

var (
	ConfigFlag = cli.StringFlag{
		Name:  "config",
		Usage: "Genesis block config `<file>`. If doesn't specifies, use config.json as default.",
		Value: config.DEFAULT_CONFIG_FILE_NAME,
	}
	FeeFlag = cli.Uint64Flag{
		Name:  "fee",
		Usage: "relayer fee",
		Value: config.DEFAULT_FEE,
	}
	ChainIDFlag = cli.Uint64Flag{
		Name:  "chain-id",
		Usage: "chain id of destination",
	}
	WalletFileFlag = cli.StringFlag{
		Name:  "wallet,w",
		Value: config.DEFAULT_WALLET_FILE_NAME,
		Usage: "Wallet `<file>`",
	}
	AmountFlag = cli.Uint64Flag{
		Name:  "amount",
		Usage: "amount of asset that need cross chain transfer",
	}
)

func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
