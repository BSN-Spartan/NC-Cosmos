package node

import (
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/bianjieai/iritamod/modules/node/types"
)

const (
	ProposalTypeCreateValidator string = "CreateValidator"
	ProposalTypeUpdateValidator string = "UpdateValidator"
	ProposalTypeRemoveValidator string = "RemoveValidator"

	// DoNotModifyDesc used in flags to indicate that description field should not be updated
	DoNotModifyDesc = "[do-not-modify]"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateValidator)
	govtypes.RegisterProposalType(ProposalTypeUpdateValidator)
	govtypes.RegisterProposalType(ProposalTypeRemoveValidator)
}

// Implements Proposal Interface
var (
	_ govtypes.Content = &CreateValidatorProposal{}
	_ govtypes.Content = &UpdateValidatorProposal{}
	_ govtypes.Content = &RemoveValidatorProposal{}

	_ codectypes.UnpackInterfacesMessage = (*CreateValidatorProposal)(nil)
)

func (sup *CreateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *CreateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *CreateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *CreateValidatorProposal) ProposalType() string   { return ProposalTypeCreateValidator }
func (sup *CreateValidatorProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(sup); err != nil {
		return err
	}

	if len(sup.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator cannot be blank")
	}

	if _, err := sdk.AccAddressFromBech32(sup.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	name := strings.TrimSpace(sup.Name)
	if len(name) == 0 || DoNotModifyDesc == name {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator name cannot be blank")
	}

	if sup.Power <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "power must be positive")
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (sup CreateValidatorProposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(sup.Pubkey, &pubKey)
}
func (sup *UpdateValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *UpdateValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *UpdateValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *UpdateValidatorProposal) ProposalType() string   { return ProposalTypeUpdateValidator }
func (sup *UpdateValidatorProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(sup); err != nil {
		return err
	}

	if len(sup.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "validator cannot be blank")
	}

	if len(sup.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator cannot be blank")
	}

	if _, err := sdk.AccAddressFromBech32(sup.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if sup.Power <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "power must be positive")
	}
	return nil
}

func (sup *RemoveValidatorProposal) GetTitle() string       { return sup.Title }
func (sup *RemoveValidatorProposal) GetDescription() string { return sup.Summary }
func (sup *RemoveValidatorProposal) ProposalRoute() string  { return types.RouterKey }
func (sup *RemoveValidatorProposal) ProposalType() string   { return ProposalTypeRemoveValidator }
func (sup *RemoveValidatorProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(sup); err != nil {
		return err
	}

	if len(sup.Operator) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "operator cannot be blank")
	}

	if _, err := sdk.AccAddressFromBech32(sup.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if len(sup.Id) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator id cannot be blank")
	}
	return nil
}
