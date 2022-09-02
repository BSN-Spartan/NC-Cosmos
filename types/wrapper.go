package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	opbtypes "github.com/bianjieai/iritamod/modules/opb/types"
)

var _ opbtypes.TokenKeeper = TokenKeeper{}

type TokenKeeper struct {
	k tokenkeeper.Keeper
}

func WrapTokenKeeper(k tokenkeeper.Keeper) TokenKeeper {
	return TokenKeeper{k}
}

// GetOwner implements types.TokenKeeper
func (tk TokenKeeper) GetOwner(ctx sdk.Context, denom string) (sdk.AccAddress, error) {
	return tk.k.GetOwner(ctx, denom)
}

// GetToken implements types.TokenKeeper
func (tk TokenKeeper) GetToken(ctx sdk.Context, denom string) (opbtypes.Token, error) {
	token, err := tk.k.GetToken(ctx, denom)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// MintToken implements types.TokenKeeper
func (tk TokenKeeper) MintToken(ctx sdk.Context, symbol string, amount uint64, recipient sdk.AccAddress, owner sdk.AccAddress) error {
	return tk.k.MintToken(ctx, symbol, amount, recipient, owner)
}
