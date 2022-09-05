

# Introduction

A Non-Cryptocurrency Public Chain is a transformed public chain framework based on an existing public chain. Gas Credit transfers are not permitted between standard wallets. There will be no cryptocurrency incentives for mining or participating in consensus.

## 1. About Spartan-II Chain (Powered by NC Cosmos)

This document is a guide to install, configure and run a full node in the Non-Crypto Polygon Edge (NC-Cosmos) public blockchain.

NC-Polygon Edge networks have two identifiers, a network ID and a chain ID. Although they often have the same value, they have different uses. Peer-to-peer communication between nodes uses the network ID, while the transaction signature process uses the chain ID.

EVM module:  Network ID = Chain ID = 9003

Native module:  Network ID = Chain ID = starmint

## 2. Hardware Requirement

Before joining the mainnet/testnet, it should be noted that running a node can be quite resource-intensive.

It's recommended that you run Spartan-II Chain (Powered by NC Cosmos) nodes on Linux Server with the following minimum requirement.

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

Install `go` by the following steps:

Download and untar the installation file

```
wget https://go.dev/dl/go1.18.5.linux-amd64.tar.gz

tar -C /usr/local -zxvf go1.18.5.linux-amd64.tar.gz
```

Modify environment variables, for example in bash

```shell
vim /etc/profile

# insert at the bottom of the file
export PATH=$PATH:/usr/local/go/bin:/opt/gopath/bin
export GOPATH=/opt/gopath
```

Then, make the /etc/profile file take effect after modification

```
source /etc/profile
```

Check the installation result

```
go version
```

### 3.2 Installation

#### 3.2.1 Building from Source
After setting up `go` correctly, you should be able to compile and run `spartan`.

Make sure that your server can access Google because our project depends on some libraries provided by Google. (If you don't have access to google.com, you can also try to add a proxy: `export GOPROXY=https://goproxy.io`)

```bash
git clone https://github.com/BSN-Spartan/NC-Cosmos.git
cd NC-Cosmos
git checkout v0.45.1
make install
```

Have you set up the Environment Variables correctly, you will get no error during `spartan` installation.

Now check your `spartan` version.

```bash
spartan version
```

#### 3.2.2 Using Docker Images


Before installing the node by Docker images, Docker 18 or later version should be installed in your server.

Run the following command to install the Docker image:

```
wget -qO- https://get.docker.com/ | sh
```

Grant your user permission to execute Docker commands:

```
sudo usermod -aG docker your-user
```


Official Docker images are hosted under the hub.docker.com registry. Run the following command to pull them to the server:

```
docker pull bsnspartan/nc-cosmos:0.45.1
```


## 4. Run the Full Node

### 4.1 Initialize Node

Create the directory `/spartan` for node installation and the sub directory `/config` and `/data` to store the configuration file:

```shell
mkdir -p /spartan/config
mkdir -p /spartan/data
```

Download [genesis.json](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/genesis.json), [app.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/app.toml) and [config.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/config.toml) to the `/spartan/config` directory.

Now the structure is like below:

```shell
/spartan/
├── config
│   ├── app.toml
│   ├── config.toml
│   └── genesis.json
└── data
```

### 4.2 Start Mainnet

#### 4.2.1 Binary
:::tip NOTE
By now can **Run the Node as a Full Node** with this last step to join the mainnet. No extra steps are required.
:::

The final step is to start the node. The node will start producing blocks.

```shell
spartan start --home /spartan
```
Or you can execute in the background via `nohup`:

```shell
nohup spartan start --home /spartan >./output.log 2>&1 &
```

#### 4.2.2 Docker

You can also start the node by Docker:

```shell
docker run -d -p 9090:9090 -p 26656:26656 -p 26657:26657 -p 8545:8545 -p 8546:8546 -v /spartan:/spartan --restart=always --name spartan-nc-cosmos bsnspartan/nc-cosmos:0.45.1 spartan start --home /spartan
```

You can check logs by command below:

```shell
docker logs -f spartan-nc-cosmos
```


## 5. Generate the Node Signature

When joining the Spartan Network as a VDC, the VDC Owner will be rewarded a certain amount of NTT Incentives based on the quantity of the registered node. To achieve this, the VDC Owner should firstly provide the signature of the VDC node to verify the node's ownership.

####     Node Installed by Commands:

Execute the following command in the node's data directory after the node is started.

```
spartan node sign-info --home /spartan
```

* `--home` is the data directory of the node.  you should specify this directory to store the data file of the node.

#### Node Installed by Docker

Execute below command:

```
docker exec spartan-nc-cosmos spartan node sign-info --home /spartan 
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

To find out more about Spartan-II Chain (Powered by NC Cosmos), visit [here](https://github.com/BSN-Spartan/NC-Cosmos/tree/main/docs)
