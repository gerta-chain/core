package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	gertaapp "github.com/gerta-chain/core/app"
)

func main() {
	rootCmd, _ := NewRootCmd()

	if err := svrcmd.Execute(rootCmd, gertaapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
