package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/bianjieai/spartan-cosmos/module/node"
)

type (
	ValidatorCreateProposalJSON struct {
		Deposit string                      `json:"deposit"`
		Content CreateValidatorProposalJSON `json:"content"`
	}

	CreateValidatorProposalJSON struct {
		Title       string `json:"title"`
		Summary     string `json:"summary"`
		Name        string `json:"name"`
		Pubkey      string `json:"pubkey"`
		Power       int64  `json:"power"`
		Description string `json:"description"`
		Operator    string `json:"operator"`
	}

	ValidatorUpdateProposalJSON struct {
		Deposit string                       `json:"deposit"`
		Content node.UpdateValidatorProposal `json:"content"`
	}

	ValidatorRemoveProposalJSON struct {
		Deposit string                       `json:"deposit"`
		Content node.RemoveValidatorProposal `json:"content"`
	}
)

// ParseValidatorCreateJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseValidatorCreateJSON(cdc *codec.LegacyAmino, proposalFile string) (ValidatorCreateProposalJSON, error) {
	proposal := ValidatorCreateProposalJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// ParseValidatorCreateJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseValidatorUpdateJSON(cdc *codec.LegacyAmino, proposalFile string) (ValidatorUpdateProposalJSON, error) {
	proposal := ValidatorUpdateProposalJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// ParseValidatorCreateJSON reads and parses a ParamChangeProposalJSON from
// file.
func ParseValidatorRemoveJSON(cdc *codec.LegacyAmino, proposalFile string) (ValidatorRemoveProposalJSON, error) {
	proposal := ValidatorRemoveProposalJSON{}
	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
