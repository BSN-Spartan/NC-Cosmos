package keeper

import (
	"context"

	"github.com/bianjieai/iritamod/modules/node/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) CreateValidator(_ context.Context, _ *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to create validator")
}

func (k Keeper) UpdateValidator(_ context.Context, _ *types.MsgUpdateValidator) (*types.MsgUpdateValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to update validator")
}

func (k Keeper) RemoveValidator(_ context.Context, _ *types.MsgRemoveValidator) (*types.MsgRemoveValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to remove validator")
}

func (k Keeper) GrantNode(_ context.Context, _ *types.MsgGrantNode) (*types.MsgGrantNodeResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to grant node")
}

func (k Keeper) RevokeNode(_ context.Context, _ *types.MsgRevokeNode) (*types.MsgRevokeNodeResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to revoke node")
}
