---
order: 3
---

# Local Testnet

For testing or developing purposes, you may want to set up a local testnet.

It is recommended to use the `spartan testnet` command to setup directly. Currently the `init` subcommand is not supported.

## Set up with `spartan testnet` Command

This is the most convenient way to set up a multiple-node local testnet on your host.

### init testnet

```bash
spartan testnet --v <validator-nums> --chain-id <chain-name> --keyring-backend file --output-dir <dir-name> --node-dir-prefix <node-prefix>
```

Here we specify validator numbers as one. You are asked to settle keyring passphrase.

```bash
spartan testnet --v 1 --chain-id spartan-chain --keyring-backend file --output-dir ./testnet

> Enter keyring passphrase: ********
> Re-enter keyring passphrase: ********
> Successfully initialized 1 node directories
```

The node data directory looks like this. The `node0` dir contains all the config and data files of your validator node.

```bash
├── testnet
    ├── node0
    │   ├── spartan
    │   │   ├── config
    │   │   │   ├── app.toml
    │   │   │   ├── config.toml
    │   │   │   ├── genesis.json
    │   │   │   ├── node_key.json
    │   │   │   └── priv_validator_key.json
    │   │   └── data
    │   │       └── priv_validator_state.json
    │   └── spartancli
    │       ├── keyring-file
    │       │   ├── 6fab5aa7a1e04a5c0575190faf77dd9678de90db.address
    │       │   ├── keyhash
    │       │   └── node0.info
    │       └── key_seed.json
    ├── root_cert.pem
    └── root_key.pem
```

### spartan start

Simply specify the node home directory where you want to start, then run your chain.

```bash
spartan start --home ./testnet/node0/spartan
```

### spartan unsafe-reset-all

This command will reset all your local data and reset the chain to Genesis status.

```bash
spartan unsafe-reset-all --home ./testnet/node0/spartan

> INF Removed existing address book file=testnet/node0/spartan/config/addrbook.json
> INF Removed all blockchain history dir=testnet/node0/spartan/data
> INF Reset private validator file to genesis state keyFile=testnet/node0/spartan/config/priv_validator_key.json stateFile=testnet/node0/spartan/data/priv_validator_state.json
```

## Testnet with Docker

**Requirements:**

- [Install spartan](../get-started/install.md)
- [Install jq](https://stedolan.github.io/jq/download/)
- [Install docker](https://docs.docker.com/engine/installation/)
- [Install docker-compose](https://docs.docker.com/compose/install/)

:::tip NOTE
Currently, the spartan-cosmos docker image has not been prepared yet.
:::
