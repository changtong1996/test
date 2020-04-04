package cli

import (
	"fmt"
	"bufio"
	"strconv"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/changtong1996/test/x/test1/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	test1TxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	test1TxCmd.AddCommand(flags.PostCommands(
		// TODO: Add tx based commands
		// GetCmd<Action>(cdc)
	)...)

	return test1TxCmd
}


func GetCmdNewVote(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add-vote [article_id][voteup]",
		Short: "add a vote",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			voteUP, err := strconv.Atoi(args[1])
			msg := types.NewMsgCreateAVote(cliCtx.GetFromAddress(), args[0], voteUP)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// Example:
//
// GetCmd<Action> is the CLI command for doing <Action>
// func GetCmd<Action>(cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "/* Describe your action cmd */",
// 		Short: "/* Provide a short description on the cmd */",
// 		Args:  cobra.ExactArgs(2), // Does your request require arguments
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

// 			msg := types.NewMsg<Action>(/* Action params */)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
// 		},
// 	}
// }
