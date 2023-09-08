package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/CosmosContracts/juno/v17/x/feepay/types"
)

// NewQueryCmd returns the cli query commands for this module
func NewQueryCmd() *cobra.Command {
	feepayQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	feepayQueryCmd.AddCommand(
		NewQueryFeePayContract(),
		NewQueryFeePayContracts(),
	)

	return feepayQueryCmd
}

// Query all fee pay contracts
func NewQueryFeePayContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract [address]",
		Short: "Query a FeePay contract by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]

			req := &types.QueryFeePayContract{
				ContractAddress: address,
			}

			res, err := queryClient.FeePayContract(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// Query all fee pay contracts
func NewQueryFeePayContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contracts",
		Short: "Query all FeePay contracts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryFeePayContracts{
				Pagination: pageReq,
			}

			res, err := queryClient.FeePayContracts(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "all-contracts")
	return cmd
}