package cmd

// DONTCOVER

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	tmconfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	evmhd "github.com/tharsis/ethermint/crypto/hd"
	evmosConfig "github.com/tharsis/ethermint/server/config"
	ethermint "github.com/tharsis/ethermint/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	evmfmttypes "github.com/tharsis/ethermint/x/feemarket/types"

	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/bianjieai/iritamod/modules/genutil"
	genutiltypes "github.com/bianjieai/iritamod/modules/genutil/types"
	"github.com/bianjieai/iritamod/modules/node"
	opbtypes "github.com/bianjieai/iritamod/modules/opb/types"
	"github.com/bianjieai/iritamod/modules/perm"

	"github.com/bianjieai/iritamod/utils"
)

const (
	nodeDirPerm         = 0755
	DefaultPointDenom   = "point"
	DefaultPointMinUnit = "upoint"
	NewEvmDenom         = "gas"
	DefaultEvmMinUnit   = "ugas"
)

var PowerReduction = sdk.NewIntFromUint64(1000000000000000000)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCLIHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
)

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(mbm module.BasicManager, genBalIterator banktypes.GenesisBalancesIterator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a spartan testnet",
		Long: "testnet will create \"v\" number of directories and populate each with " +
			"necessary files (private validator, genesis, config, etc.).\n" +
			"Note, strict routability for addresses is turned off in the config file.",
		Example: "spartan testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			outputDir := viper.GetString(flagOutputDir)
			chainID := viper.GetString(flags.FlagChainID)
			minGasPrices := viper.GetString(server.FlagMinGasPrices)
			nodeDirPrefix := viper.GetString(flagNodeDirPrefix)
			nodeDaemonHome := viper.GetString(flagNodeDaemonHome)
			nodeCLIHome := viper.GetString(flagNodeCLIHome)
			startingIPAddress := viper.GetString(flagStartingIPAddress)
			numValidators := viper.GetInt(flagNumValidators)
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)

			return InitTestnet(
				clientCtx, cmd, config, mbm, genBalIterator, outputDir, chainID, minGasPrices,
				nodeDirPrefix, nodeDaemonHome, nodeCLIHome, startingIPAddress, numValidators, algo,
			)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "spartan", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeCLIHome, "spartancli", "Home directory of the node's cli configuration")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1", "Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(evmhd.EthSecp256k1Type), "Key signing algorithm to generate keys for")
	return cmd
}

// Initialize the testnet
func InitTestnet(
	clientCtx client.Context, cmd *cobra.Command,
	config *tmconfig.Config, mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	outputDir, chainID, minGasPrices, nodeDirPrefix,
	nodeDaemonHome, nodeCLIHome, startingIPAddress string,
	numValidators int, algoStr string,
) error {
	if chainID == "" {
		chainID = fmt.Sprintf("chain_%d-1", tmrand.Int63n(9999999999999)+1)
	}

	monikers := make([]string, numValidators)
	nodeIDs := make([]string, numValidators)
	validators := make([]node.Validator, numValidators)
	peers := make([]string, numValidators)
	signingInfos := make([]slashingtypes.SigningInfo, numValidators)

	spartanConfig := evmosConfig.DefaultConfig()
	spartanConfig.MinGasPrices = minGasPrices
	spartanConfig.API.Enable = true
	spartanConfig.Telemetry.Enabled = true
	spartanConfig.Telemetry.PrometheusRetentionTime = 60
	spartanConfig.Telemetry.EnableHostnameLabel = false
	spartanConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", chainID}}

	//nolint:prealloc
	var (
		genAccounts []authtypes.GenesisAccount
		genBalances []banktypes.Balance
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())

	rootKeyPath := filepath.Join(outputDir, "root_key.pem")
	rootCertPath := filepath.Join(outputDir, "root_cert.pem")
	if err := os.MkdirAll(outputDir, nodeDirPerm); err != nil {
		_ = os.RemoveAll(outputDir)
		return err
	}

	utils.GenRootCert(rootKeyPath, rootCertPath, "/C=CN/ST=root/L=root/O=root/OU=root/CN=root")
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		clientDir := filepath.Join(outputDir, nodeDirName, nodeCLIHome)

		config.SetRoot(nodeDir)
		config.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		if err := os.MkdirAll(clientDir, nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		monikers[i] = nodeDirName
		config.Moniker = nodeDirName

		ip, err := getIP(i, startingIPAddress)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		nodeKey, filePv, err := genutil.InitializeNodeValidatorFiles(config)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		nodeIDs[i] = string(nodeKey.ID())
		peers[i] = fmt.Sprintf("%s@%s:26656", nodeIDs[i], ip)

		genFiles = append(genFiles, config.GenesisFile())

		kb, err := keyring.New(
			sdk.KeyringServiceName(),
			viper.GetString(flags.FlagKeyringBackend),
			clientDir,
			inBuf,
			evmhd.EthSecp256k1Option(),
		)
		if err != nil {
			return err
		}

		keyringAlgos, _ := kb.SupportedAlgorithms()
		signAlgo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
		if err != nil {
			return err
		}

		addr, secret, err := server.GenerateSaveCoinKey(kb, nodeDirName, true, signAlgo)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		// save private key seed words
		if err := writeFile(fmt.Sprintf("%v.json", "key_seed"), clientDir, cliPrint); err != nil {
			return err
		}

		accPointTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
		accNativeTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
		accEvmTokens := sdk.TokensFromConsensusPower(5000, PowerReduction)
		coins := sdk.Coins{
			sdk.NewCoin(DefaultPointMinUnit, accPointTokens),
			sdk.NewCoin(tokentypes.GetNativeToken().MinUnit, accNativeTokens),
			sdk.NewCoin(DefaultEvmMinUnit, accEvmTokens),
		}

		genBalances = append(genBalances, banktypes.Balance{Address: addr.String(), Coins: coins.Sort()})

		genAccounts = append(genAccounts, &ethermint.EthAccount{
			BaseAccount: authtypes.NewBaseAccount(addr, nil, 0, 0),
			CodeHash:    common.BytesToHash(evmtypes.EmptyCodeHash).Hex(),
		})

		tmValPubKey, err := filePv.GetPubKey()
		if err != nil {
			return err
		}

		pubkey, err := cryptocodec.FromTmPubKeyInterface(tmValPubKey)
		if err != nil {
			return err
		}

		validators[i] = node.NewValidator(
			tmbytes.HexBytes(tmhash.Sum(pubkey.Bytes())),
			nodeDirName,
			nodeDirName,
			pubkey,
			"",
			100,
			addr,
		)

		consAddr, err := validators[i].GetConsAddr()
		if err != nil {
			return err
		}

		signingInfos[i] = slashingtypes.SigningInfo{
			Address: consAddr.String(),
			ValidatorSigningInfo: slashingtypes.NewValidatorSigningInfo(
				consAddr,
				0,
				0,
				time.Unix(0, 0),
				false,
				0,
			),
		}

		customAppTemplate, customAppConfig := evmosConfig.AppConfig(ethermint.AttoPhoton)
		srvconfig.SetConfigTemplate(customAppTemplate)
		if err := server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig); err != nil {
			return err
		}

		spartanConfigFilePath := filepath.Join(nodeDir, "config/app.toml")
		srvconfig.WriteConfigFile(spartanConfigFilePath, spartanConfig)
	}

	if err := initGenFiles(DefaultEvmMinUnit, clientCtx, mbm, chainID, genAccounts, genBalances, genFiles, validators,
		signingInfos, monikers, nodeIDs, rootCertPath); err != nil {
		return err
	}

	if err := collectGenFiles(
		clientCtx, config, chainID, monikers, nodeIDs, peers, numValidators,
		outputDir, nodeDirPrefix, nodeDaemonHome,
	); err != nil {
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", numValidators)
	return nil
}

func initGenFiles(
	coinDenom string, clientCtx client.Context, mbm module.BasicManager, chainID string,
	genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance,
	genFiles []string, validators []node.Validator, signingInfos []slashingtypes.SigningInfo, monikers []string, nodeIDs []string,
	rootCertPath string,
) error {
	rootCertBz, err := ioutil.ReadFile(rootCertPath)
	if err != nil {
		return fmt.Errorf("failed to read root certificate: %s", err.Error())
	}
	jsonMarshaler := clientCtx.Codec

	appGenState := mbm.DefaultGenesis(jsonMarshaler)

	var nodeGenState node.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[node.ModuleName], &nodeGenState)

	nodeGenState.RootCert = string(rootCertBz)
	nodeGenState.Validators = validators
	appGenState[node.ModuleName] = jsonMarshaler.MustMarshalJSON(&nodeGenState)

	var slashingGenState slashingtypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[slashingtypes.ModuleName], &slashingGenState)
	slashingGenState.SigningInfos = signingInfos
	appGenState[slashingtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&slashingGenState)

	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return err
	}

	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&bankGenState)

	// set the point token in the genesis state
	var tokenGenState tokentypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[tokentypes.ModuleName], &tokenGenState)

	pointToken := tokentypes.Token{
		Symbol:        DefaultPointDenom,
		Name:          "Spartan Network Point Token",
		Scale:         6,
		MinUnit:       DefaultPointMinUnit,
		InitialSupply: 1000000000,
		MaxSupply:     math.MaxUint64,
		Mintable:      true,
		Owner:         genAccounts[0].GetAddress().String(),
	}

	gasToken := tokentypes.Token{
		Symbol:        NewEvmDenom,
		Name:          "Spartan Network Fee Token",
		Scale:         18,
		MinUnit:       DefaultEvmMinUnit,
		InitialSupply: 1000000000,
		MaxSupply:     math.MaxUint64,
		Mintable:      true,
		Owner:         genAccounts[0].GetAddress().String(),
	}
	nativeToken := tokentypes.GetNativeToken()

	tokenGenState.Tokens = append(tokenGenState.Tokens, pointToken)
	tokenGenState.Tokens = append(tokenGenState.Tokens, gasToken)
	tokenGenState.Params.IssueTokenBaseFee = sdk.NewCoin(nativeToken.Symbol, sdk.NewInt(60000))
	appGenState[tokentypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&tokenGenState)

	// modify the native token denoms in the opb genesis
	var opbGenState opbtypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[opbtypes.ModuleName], &opbGenState)

	opbGenState.Params.BaseTokenDenom = nativeToken.MinUnit
	opbGenState.Params.PointTokenDenom = DefaultPointMinUnit
	appGenState[opbtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&opbGenState)

	// modify the constant fee denoms in the crisis genesis
	var crisisGenState crisistypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[crisistypes.ModuleName], &crisisGenState)

	crisisGenState.ConstantFee.Denom = nativeToken.MinUnit
	appGenState[crisistypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&crisisGenState)

	// modify the constant fee denoms in the crisis genesis
	var serviceGenState servicetypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[servicetypes.ModuleName], &serviceGenState)

	serviceGenState.Params.MinDeposit = sdk.NewCoins(sdk.NewCoin(DefaultPointMinUnit, sdk.NewInt(5000)))
	serviceGenState.Params.BaseDenom = DefaultPointMinUnit
	appGenState[servicetypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&serviceGenState)

	var evmGenState evmtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[evmtypes.ModuleName], &evmGenState)

	evmGenState.Params.EvmDenom = coinDenom
	appGenState[evmtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&evmGenState)

	var fMKGenState evmfmttypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[evmfmttypes.ModuleName], &fMKGenState)
	fMKGenState.Params.NoBaseFee = true
	appGenState[evmfmttypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&fMKGenState)

	// add all genesis accounts as root admins
	var permGenState perm.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[perm.ModuleName], &permGenState)
	for _, account := range genAccounts {
		permGenState.RoleAccounts = append(
			permGenState.RoleAccounts,
			perm.RoleAccount{
				Address: account.GetAddress().String(),
				Roles:   []perm.Role{perm.RoleRootAdmin},
			},
		)
	}
	appGenState[perm.ModuleName] = jsonMarshaler.MustMarshalJSON(&permGenState)

	var govGenState govtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[govtypes.ModuleName], &govGenState)
	govGenState.DepositParams.MinDeposit = sdk.NewCoins(
		sdk.NewCoin(nativeToken.MinUnit, sdk.NewInt(100)),
	)
	appGenState[govtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&govGenState)

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < len(validators); i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context,
	config *tmconfig.Config,
	chainID string,
	monikers, nodeIDs, peers []string,
	numValidators int,
	outputDir, nodeDirPrefix, nodeDaemonHome string,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		config.Moniker = nodeDirName

		config.SetRoot(nodeDir)

		genDoc, err := types.GenesisDocFromFile(config.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := GenAppStateFromConfig(clientCtx.Codec,
			clientCtx.TxConfig,
			config,
			*genDoc,
			nodeIDs[i],
			peers,
		)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

// GenAppStateFromConfig gets the genesis app state from the config
func GenAppStateFromConfig(
	cdc codec.JSONCodec,
	txEncodingConfig client.TxEncodingConfig,
	config *cfg.Config,
	genDoc tmtypes.GenesisDoc,
	nodeID string,
	peers []string,
) (appState json.RawMessage, err error) {

	config.P2P.PersistentPeers = filterPeers(nodeID, peers)
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	// create the app state
	appGenesisState, err := genutiltypes.GenesisStateFromGenDoc(genDoc)
	if err != nil {
		return appState, err
	}

	appGenesisState, err = genutil.AddGenTxsInAppGenesisState(cdc, txEncodingConfig.TxJSONEncoder(), appGenesisState, []sdk.Tx{})
	if err != nil {
		return appState, err
	}

	appState, err = json.MarshalIndent(appGenesisState, "", "  ")
	if err != nil {
		return appState, err
	}

	genDoc.AppState = appState
	err = genutil.ExportGenesisFile(&genDoc, config.GenesisFile())

	return appState, err
}

func filterPeers(nodeID string, peers []string) string {
	var dstPeers []string
	for _, p := range peers {
		if !strings.HasPrefix(p, nodeID) {
			dstPeers = append(dstPeers, p)
		}
	}
	if len(dstPeers) == 0 {
		return ""
	}

	if len(dstPeers) == 1 {
		return dstPeers[0]
	}

	sort.Strings(dstPeers)
	return strings.Join(dstPeers, ",")
}

func getIP(i int, startingIPAddr string) (ip string, err error) {
	if len(startingIPAddr) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIPAddr, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	if err := tmos.EnsureDir(writePath, 0700); err != nil {
		return err
	}

	return tmos.WriteFile(file, contents, 0600)
}
