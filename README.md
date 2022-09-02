

# Introduction

A Non-Cryptocurrency Public Chain is a transformed public chain framework based on an existing public chain. Gas Credit transfers are not permitted between standard wallets. There will be no cryptocurrency incentives for mining or participating in consensus.

## 1. About Spartan-II Chain (Powered by NC Cosmos)

This document is a guide to install, configure and run a full node in the Non-Crypto Polygon Edge (NC-Cosmos) public blockchain.

NC-Polygon Edge networks have two identifiers, a network ID and a chain ID. Although they often have the same value, they have different uses. Peer-to-peer communication between nodes uses the network ID, while the transaction signature process uses the chain ID.

EVM module:  Network ID = Chain ID = 9003

Native module:  Network ID = Chain Id = starmint

## 2. Hardware Requirement

Before joining the mainnet/testnet, it should be noted that running a node can be quite resource-intensive.

It's recommended that you run Spartan-II Chain（Powered by NC Cosmos) nodes on Linux Server with the following minimum requirement.

#### Minimum Requirement

- 2 CPU
- Memory: 6GB
- Disk: 256GB SSD
- OS: Ubuntu 16.04 LTS +
- Bandwidth: 20Mbps
- Allow all incoming connections on TCP port 26656 and 26657

#### Recommended Requirement

- 4 CPU
- Memory: 12GB
- Disk: 512GB SSD
- OS: Ubuntu 18.04 LTS +
- Bandwidth: 20Mbps
- Allow all incoming connections on TCP port 26656 and 26657

## 3. How to Install a Full Node

### 3.1 Prerequisites:

::: tip
**Go 1.18** is recommended for building and installing the Spartan-Cosmos software.
:::

Install `go` by following the [official docs](https://go.dev/doc/install).

Remember to set your `$GOPATH`, `$GOBIN`, and `$PATH` environment variables, for example:

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bashrc
source ~/.bashrc
echo "export GOBIN=$GOPATH/bin" >> ~/.bashrc
source ~/.bashrc
echo "export PATH=$PATH:$GOBIN" >> ~/.bashrc
source ~/.bashrc
```

Verify that `go` has been installed successfully.

```bash
go version
```

### 3.2 Installation

After setting up `go` correctly, you should be able to compile and run `spartan`.

Make sure that your server can access Google because our project depends on some libraries provided by Google. (If you don't have access to google.com, you can also try to add a proxy: `export GOPROXY=https://goproxy.io`)

```bash
git clone https://github.com/BSN-Spartan/NC-Cosmos.git
cd NC-Cosmos
git checkout <version>
make install
```

Have you set up the Environment Variables correctly, you will get no error during `spartan` installation.

Now check your `spartan` version.

```bash
spartan version
```




## 4. Run the Full Node

### 4.1 Initialize Node

Select a Chain ID and initialize the node. By default, the `testnet` command creates the `~/.spartan` directory with subfolders `config` and `data`. In the following steps, we assume you use this default path.

Note that you should specify `-v` as 1, otherwise it will create multiple node config files, which is useful for local testnet.

```shell
# initialize node configurations
spartan testnet -v 1 --chain-id=<spartan-mainnet>
```

### 4.2 Genesis & Seeds
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

### 4.3 Start Mainnet

:::tip NOTE
By now can **Run the Node as a Full Node** with this last step to join the mainnet. No extra steps are required.
:::

The final step is to start the node. The node will start producing blocks.

```shell
spartan start
```


## 5. Generate the Node Signature

When joining the Spartan Network as a VDC, the VDC Owner will be rewarded a certain amount of NTT Incentives based on the quantity of the registered node. To achieve this, the VDC Owner should firstly provide the signature of the VDC node to verify the node's ownership.

####     Node Installed by Commands:

Execute the following command in the node's data directory after the node is started.

```
spartan node sign-info --home spartan
```

* `--home` is the data directory of the node.  you should specify this directory to store the data file of the node.

#### Node Installed by Docker

Execute below command:

  ```
docker exec bsnspartan/nc-cosmos:0.45.1 spartan node sign-info --home spartan 
  
  ```
#### Node Signature

After executing the above commands，you will get the following information. Please submit it to the Spartan Governance System when registering the node .

```
{
      "node_id":"3b1e874ae98c36cdd5da32f9623fa0a4586b4445",
      "pub_key":"{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"3ggBZ5e7lJCoGsRE3CdNESSDv1fU8mbGDK4k8oNzJZQ=\"}",
      "signature":"Za1mHGzMXdYFs6RAmIr9Nge7P5Zw6k1m/eOW9iMzP9qhnaBMQEZ+LMoYxy+BqmZuJVxJALs2vKXcMXt5eHLtBQ=="
  }
```



## 6.Resources

To find out more about Spartan-II Chain (Powered by NC Cosmos), visit here
