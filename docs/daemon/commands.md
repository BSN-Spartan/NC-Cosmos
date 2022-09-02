---
order: 2
---

# Commands

## Introduction

Spartan-Cosmos Daemon Commands allow you to init, start and reset a node, or generate a genesis file, etc.

You can get familiar with these commands by creating a [Local Testnet](local-testnet.md).

## Usage

```bash
spartan [command]
```

## Available Commands

| Name                  | Description                                                                                                                |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| add-genesis-account   | Add a genesis account to genesis.json                                                                                      |
| add-genesis-validator | Generate a genesis tx to create a validator                                                                                |
| config                | Create or query an application CLI configuration file                                                                      |
| export                | Export state to JSON                                                                                                       |
| help                  | Help about any command                                                                                                     |
| init                  | Initialize private validator, p2p, genesis, and application configuration files                                            |
| keys                  | Manage your application's keys                                                                                             |
| query                 | Querying subcommands                                                                                                       |
| snapshot              | Snapshot the latest information and drop the others                                                                        |
| start                 | Run the full node                                                                                                          |
| status                | Query remote node for status                                                                                               |
| tendermint            | Tendermint subcommands                                                                                                     |
| testnet               | Initialize files for a spartan testnet                                                                                     |
| tx                    | Transactions subcommands                                                                                                   |
| unsafe-reset-all      | Resets the blockchain database, removes address book files, and resets data/priv_validator_state.json to the genesis state |
| validate-genesis      | validates the genesis file at the default location or at the location passed as an arg                                     |
| version               | Print the application binary version information                                                                           |

## Global Flags

| Name,shorthand | Default         | Description                          |
| -------------- | --------------- | ------------------------------------ |
| -h, --help     |                 | Help for spartan                     |
| --home         | /$HOME/.spartan | Directory for config and data        |
| --log_format   | plain           | The logging format                   |
| --log_level    | info            | The logging level                    |
| --trace        |                 | Print out full stack trace on errors |
