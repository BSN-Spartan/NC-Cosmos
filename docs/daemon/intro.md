---
order: 1
---

# Introduction

The `spartan` executable program is the entry point for running an Spartan-Cosmos node. All the validator nodes and full nodes need to install the `spartan` and launching the daemon to join the Spartan-Cosmos network. You can also use `spartan` to start your own test network locally.

## Home Directory

The home directory is the working directory of the spartan node. The home directory contains all the configuration information and all the data that the node runs.

In the `spartan` command, you can specify the home directory of the node by using flag `--home`. If you run multiple nodes on the same machine, you need to specify different home directories for them. If the `--home` flag is not specified in the spartan command, the default value `$HOME/.spartan` is used as the home directory.

The `spartan init` command is responsible for initializing the specified `--home` directory and creating the default configuration files. Except the `spartan init` command, the home directory used by any other `spartan` sub-commands must be initialized, otherwise an error will be reported.

The data of the Spartan-Cosmos node is stored in the `data` directory of the home, including blockchain data, application layer data, and index data.

All configuration files are stored in the `<home-dir>/config` directory:

### genesis.json

genesis.json defines the genesis block data, which specifies the system parameters such as chain_id, consensus parameters, initial account token allocation, creation of validators, and parameters for modules. See [genesis-file](../concepts/genesis-file.md) for details.

### node_key.json

node_key.json is used to store the node's key. The node-id queried by `spartan tendermint show-node-id` is derived by the key, which is used to indicate the unique identity of the node. It is used in p2p connection.

### priv_validator.json

pri_validator.json is the [Tendermint Key](../concepts/validator-faq.md#tendermint-key) file that the validator will use to sign Pre-vote/Pre-commit in each round of consensus voting. As the consensus progresses, the tendermint consensus engine will continuously update `last_height`/`last_round`/`last_step` values.

### config.toml

config.toml is the non-consensus configuration of the node. Different nodes can configure themselves according to their own situation. Common modifications are `persistent_peers`, `moniker`, `laddr`.

### app.toml

app.toml provides base configuration, telemetry configuration, API configuration, gRPC configuration and state sync configuration for Spartan-Cosmos.
