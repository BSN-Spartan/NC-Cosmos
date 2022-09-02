---
order: 4
---

# Join the Spartan-II Chain (Powered by NC-Cosmos) Network



## Chain ID

EVM module:  Network ID = Chain ID = 9003

Native module:  Network ID = Chain Id = starmint

## Install Spartan

Follow the [Installation](installation.md) document to install the Spartan binary.

## Initialize Node

Select a Chain ID and initialize the node. By default, the `testnet` command creates the `~/.spartan` directory with subfolders `config` and `data`. In the following steps, we assume you use this default path.

Note that you should specify `-v` as 1, otherwise it will create multiple node config files, which is useful for local testnet.

```shell
# initialize node configurations
spartan testnet -v 1 --chain-id=<spartan-mainnet>
```

## Genesis & Seeds

In the `~/.spartan/config` directory, the most important files for configuration are `app.toml` and `config.toml`.

In this step, you should download the public `genesis.json` from the official site.

```shell
# download mainnet public config.toml and genesis.json
curl -o ~/.spartan/config/genesis.json <official-genesis-url>
```

To find peers, it's necessary to add healthy seed nodes to `config.toml`. As well, you can download it from the official site.

```shell
curl -o ~/.spartan/config/config.toml <official-config-url>
```

## Start Mainnet

:::tip NOTE
By now can **Run the Node as a Full Node** with this last step to join the mainnet. No extra steps are required. 
:::

The final step is to start the node. The node will start producing blocks.

```shell
spartan start
```





