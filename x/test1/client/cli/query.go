package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/changtong1996/test/x/test1/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group test1 queries under a subcommand
	test1QueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	test1QueryCmd.AddCommand(
		flags.GetCommands(
	// TODO: Add query Cmds
		)...,
	)

	return test1QueryCmd
}

// TODO: Add Query Commands
func GetCmdArticleName(queryRoute string, cdc *codec.Codec) *cobra.Command {
		return &cobra.Command{
			Use:     "article [name]",
			Short:   "article name",
			Args:    cobra.ExactArgs(1),
			RunE:    func(cmd *cobra.Command, agrs []string) error{
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not get  - %s \n", name)
				return nil
			}
			var out types.Articles
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
			},
		}
}
