package cli

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Team-Kujira/core/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	oracleQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	oracleQueryCmd.AddCommand(
		GetCmdQueryExchangeRates(),
		GetCmdQueryActives(),
		GetCmdQueryParams(),
		GetCmdQueryFeederDelegation(),
		GetCmdQueryMissCounter(),
		GetCmdQueryAggregatePrevote(),
		GetCmdQueryAggregateVote(),
	)

	return oracleQueryCmd
}

// GetCmdQueryExchangeRates implements the query rate command.
func GetCmdQueryExchangeRates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rates [denom]",
		Args:  cobra.RangeArgs(0, 1),
		Short: "Query the current exchange rate of an asset",
		Long: strings.TrimSpace(`
Query the current exchange rate of USD with an asset. 
You can find the current list of active denoms by running

$ kujirad query oracle exchange-rates 

Or, can filter with denom

$ kujirad query oracle exchange-rates KUJI
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 0 {
				res, err := queryClient.ExchangeRates(context.Background(), &types.QueryExchangeRatesRequest{})
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

			denom := args[0]
			res, err := queryClient.ExchangeRate(
				context.Background(),
				&types.QueryExchangeRateRequest{Denom: denom},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryActives implements the query actives command.
func GetCmdQueryActives() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actives",
		Args:  cobra.NoArgs,
		Short: "Query the active list of assets recognized by the oracle",
		Long: strings.TrimSpace(`
Query the active list of assets recognized by the types.

$ kujirad query oracle actives
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Actives(context.Background(), &types.QueryActivesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current Oracle params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryFeederDelegation implements the query feeder delegation command
func GetCmdQueryFeederDelegation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feeder [validator]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the oracle feeder delegate account",
		Long: strings.TrimSpace(`
Query the account the validator's oracle voting right is delegated to.

$ kujirad query oracle feeder kujiravaloper...
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			valString := args[0]
			validator, err := sdk.ValAddressFromBech32(valString)
			if err != nil {
				return err
			}

			res, err := queryClient.FeederDelegation(
				context.Background(),
				&types.QueryFeederDelegationRequest{ValidatorAddr: validator.String()},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryMissCounter implements the query miss counter of the validator command
func GetCmdQueryMissCounter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miss [validator]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the # of the miss count",
		Long: strings.TrimSpace(`
Query the # of vote periods missed in this oracle slash window.

$ kujirad query oracle miss kujiravaloper...
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			valString := args[0]
			validator, err := sdk.ValAddressFromBech32(valString)
			if err != nil {
				return err
			}

			res, err := queryClient.MissCounter(
				context.Background(),
				&types.QueryMissCounterRequest{ValidatorAddr: validator.String()},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAggregatePrevote implements the query aggregate prevote of the validator command
func GetCmdQueryAggregatePrevote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-prevotes [validator]",
		Args:  cobra.RangeArgs(0, 1),
		Short: "Query outstanding oracle aggregate prevotes.",
		Long: strings.TrimSpace(`
Query outstanding oracle aggregate prevotes.

$ kujirad query oracle aggregate-prevotes

Or, can filter with voter address

$ kujirad query oracle aggregate-prevotes kujiravaloper...
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 0 {
				res, err := queryClient.AggregatePrevotes(
					context.Background(),
					&types.QueryAggregatePrevotesRequest{},
				)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

			valString := args[0]
			validator, err := sdk.ValAddressFromBech32(valString)
			if err != nil {
				return err
			}

			res, err := queryClient.AggregatePrevote(
				context.Background(),
				&types.QueryAggregatePrevoteRequest{ValidatorAddr: validator.String()},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAggregateVote implements the query aggregate prevote of the validator command
func GetCmdQueryAggregateVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregate-votes [validator]",
		Args:  cobra.RangeArgs(0, 1),
		Short: "Query outstanding oracle aggregate votes.",
		Long: strings.TrimSpace(`
Query outstanding oracle aggregate vote.

$ kujirad query oracle aggregate-votes 

Or, can filter with voter address

$ kujirad query oracle aggregate-votes kujiravaloper...
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 0 {
				res, err := queryClient.AggregateVotes(
					context.Background(),
					&types.QueryAggregateVotesRequest{},
				)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

			valString := args[0]
			validator, err := sdk.ValAddressFromBech32(valString)
			if err != nil {
				return err
			}

			res, err := queryClient.AggregateVote(
				context.Background(),
				&types.QueryAggregateVoteRequest{ValidatorAddr: validator.String()},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
