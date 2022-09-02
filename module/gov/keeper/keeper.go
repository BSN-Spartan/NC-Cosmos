package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Keeper defines the governance module Keeper
type Keeper struct {
	govkeeper.Keeper
	sk       types.StakingKeeper
	storeKey sdk.StoreKey
}

// NewKeeper returns a governance keeper. It handles:
// - submitting governance proposals
// - depositing funds into proposals, and activating upon sufficient funds being deposited
// - users voting on proposals, with weight proportional to stake in the system
// - and tallying the result of the vote.
//
// CONTRACT: the parameter Subspace must have the param key table already initialized
func NewKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramSpace types.ParamSubspace,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper, sk types.StakingKeeper, rtr types.Router,
) Keeper {

	return Keeper{
		storeKey: key,
		sk:       sk,
		Keeper:   govkeeper.NewKeeper(cdc, key, paramSpace, authKeeper, bankKeeper, sk, rtr),
	}
}

// deleteVote deletes a vote from a given proposalID and voter from the store
func (keeper Keeper) deleteVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(types.VoteKey(proposalID, voterAddr))
}
