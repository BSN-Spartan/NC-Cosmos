package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/gov/client/rest"

	"github.com/bianjieai/spartan-cosmos/module/node/client/cli"
)

var (
	EmptyRESTHandlerFn = func(client.Context) rest.ProposalRESTHandler { return rest.ProposalRESTHandler{} }

	CreateValidatorProposalHandler = govclient.NewProposalHandler(cli.NewCreateValidatorProposalTxCmd, EmptyRESTHandlerFn)
	UpdateValidatorProposalHandler = govclient.NewProposalHandler(cli.NewUpdateValidatorProposalTxCmd, EmptyRESTHandlerFn)
	RemoveValidatorProposalHandler = govclient.NewProposalHandler(cli.NewRemoveValidatorProposalTxCmd, EmptyRESTHandlerFn)
)
