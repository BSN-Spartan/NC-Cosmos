package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/bianjieai/spartan-cosmos/module/node"
	"github.com/bianjieai/spartan-cosmos/module/node/client/utils"
)

// NewCreateValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewCreateValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator to create proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal create-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"name":"node1",
		"pubkey":"{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"FU1a1Yhu+eE0XlZWSwywZur2uItmtpYZGs8TpMK4im0=\"}",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorCreateJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			if len(proposal.Content.Pubkey) == 0 {
				return errors.New("pubkey cannot be blank")
			}

			var pk cryptotypes.PubKey
			if err := clientCtx.Codec.UnmarshalInterfaceJSON([]byte(proposal.Content.Pubkey), &pk); err != nil {
				return err
			}

			var pkAny *codectypes.Any
			if pkAny, err = codectypes.NewAnyWithValue(pk); err != nil {
				return err
			}

			content := &node.CreateValidatorProposal{
				Title:       proposal.Content.Title,
				Summary:     proposal.Content.Summary,
				Name:        proposal.Content.Name,
				Pubkey:      pkAny,
				Power:       proposal.Content.Power,
				Description: proposal.Content.Description,
				Operator:    proposal.Content.Operator,
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// NewUpdateValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewUpdateValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator to update proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal update-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"3B0CB3D0CFD6590F1101543B9608CF590BB4CF9380C86407BA9DE1E7DA12F43E",
		"name":"node1",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorUpdateJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal.Content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// NewRemoveValidatorProposalTxCmd returns a CLI command handler for creating
// a validator create proposal governance transaction.
func NewRemoveValidatorProposalTxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove-validator [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a validator to remove proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Example:
$ %s tx gov submit-proposal remove-validator proposal.json --from mykey

Where proposal.json contains:

{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"3B0CB3D0CFD6590F1101543B9608CF590BB4CF9380C86407BA9DE1E7DA12F43E",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := utils.ParseValidatorRemoveJSON(clientCtx.LegacyAmino, args[0])
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&proposal.Content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
