package ante

import (
	"fmt"
	"runtime/debug"

	tmlog "github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	ethermintante "github.com/tharsis/ethermint/app/ante"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	opbkeeper "github.com/bianjieai/iritamod/modules/opb/keeper"
	"github.com/bianjieai/iritamod/modules/perm"

	evmante "github.com/bianjieai/spartan-cosmos/module/evm"
	"github.com/bianjieai/spartan-cosmos/module/gas"
	spartantypes "github.com/bianjieai/spartan-cosmos/types"
)

type HandlerOptions struct {
	PermKeeper      perm.Keeper
	AccountKeeper   authkeeper.AccountKeeper
	BankKeeper      bankkeeper.Keeper
	FeegrantKeeper  authante.FeegrantKeeper
	TokenKeeper     tokenkeeper.Keeper
	OpbKeeper       opbkeeper.Keeper
	SigGasConsumer  ante.SignatureVerificationGasConsumer
	SignModeHandler signing.SignModeHandler

	// evm config
	EvmKeeper          evmante.EVMKeeper
	EvmFeeMarketKeeper evmtypes.FeeMarketKeeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, deducts fees from the first
// signer, and performs other module-specific logic.
func NewAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		//defer Recover(ctx.Logger(), &err)
		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx
					anteHandler = sdk.ChainAnteDecorators(
						ethermintante.NewEthSetUpContextDecorator(options.EvmKeeper), // outermost AnteDecorator. SetUpContext must be called first
						evmante.NewEthValidateBasicDecorator(options.EvmKeeper),
						evmante.NewEthContractCallableDecorator(options.PermKeeper),
						evmante.NewEthSigVerificationDecorator(options.EvmKeeper, options.AccountKeeper, options.SignModeHandler),
						ethermintante.NewCanTransferDecorator(options.EvmKeeper),
						//evmante.NewCanTransferDecorator(options.EvmKeeper, options.OpbKeeper, options.TokenKeeper, options.PermKeeper),
						ethermintante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.BankKeeper, options.EvmKeeper),
						ethermintante.NewEthGasConsumeDecorator(options.EvmKeeper),
						ethermintante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
						ethermintante.NewEthMempoolFeeDecorator(options.EvmKeeper),                  // Check eth effective gas price against minimal-gas-prices
					)

				default:
					return ctx, sdkerrors.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s",
						typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}
		switch tx.(type) {
		case sdk.Tx:
			anteHandler = sdk.ChainAnteDecorators(
				gas.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
				perm.NewAuthDecorator(options.PermKeeper),
				ante.NewMempoolFeeDecorator(),
				ante.NewValidateBasicDecorator(),
				ante.NewValidateMemoDecorator(options.AccountKeeper),
				ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
				ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
				ante.NewValidateSigCountDecorator(options.AccountKeeper),
				ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
				ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
				ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
				ante.NewIncrementSequenceDecorator(options.AccountKeeper),
				ante.NewRejectExtensionOptionsDecorator(),
				ante.NewTxTimeoutHeightDecorator(),
				tokenkeeper.NewValidateTokenFeeDecorator(options.TokenKeeper, options.BankKeeper),
				opbkeeper.NewValidateTokenTransferDecorator(options.OpbKeeper,
					spartantypes.WrapTokenKeeper(options.TokenKeeper),
					options.PermKeeper,
				).DefaultValidateFn(),
			)
		default:
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		return anteHandler(ctx, tx, sim)

	}
}

func Recover(logger tmlog.Logger, err *error) {
	if r := recover(); r != nil {
		*err = sdkerrors.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}
