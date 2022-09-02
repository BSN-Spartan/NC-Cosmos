---
order: 4
---

# Join the Mainnet

:::warning
Since the Miannet is not prepared, this document provides a limited reference.
:::

## Mainnet Infos

The table below gives an overview of all Mainnet Chain IDs.

| Chain ID      | Description | Site | Version | Status |
| ------------- | ----------- | ---- | ------- | ------ |
| mainnet-1 |             |      |         |        |
| mainnet-2 |             |      |         |        |

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

## Join the Mainnet as a Validator [Optional]

In Spartan-Cosmos, one node can become a validator only by submitting a `create-validator` proposal. 

```shell
spartan tx gov submit-proposal create-validator <create-validator.json>
```

Here is a template for the  `create-validator.json` file.

```json
{
        "deposit":"1000usnmt",
        "content":{
                "title":"proposal title",
                "summary":"proposal description",
                "name":"node1",
                "pubkey":"{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"FU1a1Yhu+eE0XlZWSwywZur2uItmtpYZGs8TpMK4im0=\"}",
            "power":"100",
            "description":"my node1",
            "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
        }
}

```
