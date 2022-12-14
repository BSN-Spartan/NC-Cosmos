# Spartan NC Cosmos Node Installation User Guide

## Introduction

A Non-Cryptocurrency Public Chain is a transformed public chain framework based on an existing public chain. Gas Credit transfers are not permitted between standard wallets. There are no cryptocurrency incentives for mining or participating in consensus. On Spartan Network, there are three Non-Cryptocurrency Public Chains at launch. We except to add more in the foreseeable future.

> As a clear demonstration, all commands in this document are run with root permission. These commands can also be run under normal user permissions, please set the file storage and configure the parameters properly.

## 1. About Spartan-II Chain (Powered by NC Cosmos)

This document is a guide to install, configure and run a full node in the Spartan-II Chain (powered by NC Cosmos) .
Cosmos-based networks have two identifiers, a network ID and a chain ID. Although they often have the same value, they have different uses. Peer-to-peer communication between nodes uses the network ID, while the transaction signature process uses the chain ID.

EVM module:  Network ID = **Chain ID = 9003**

Native module:  Network ID = **Chain ID = starmint**

## 2. Hardware Requirements

Before joining the mainnet/testnet, it should be noted that running a node can be quite resource-intensive.

It's recommended that you run Spartan-II Chain (Powered by NC Cosmos) nodes on Linux Server with the following minimum requirement.

Below is the instruction for Linux system.

#### Minimum Requirements:

- 2 CPU
- Memory: 6GB
- Disk: 256GB SSD
- Bandwidth: 20Mbps
- Allow all incoming connections on TCP port 26656 and 26657

#### Recommended Requirements:

- 4 CPU
- Memory: 12GB
- Disk: 512GB SSD
- Bandwidth: 20Mbps
- Allow all incoming connections on TCP port 26656 and 26657

## 3. Full Node Installation by Commands

In this chapter, we will build a full node by commands. If you prefer to build the node by Docker Images, please go to chapter 4 Full Node Installation by Docker.


### 3.1 Prerequisites

| Software  | Version  |
| ----- | ----- |
| Golang | 1.15+ |
| GCC | latest |
| Git | 1.8.3.1+ |
| tree (optional) | 1.6.0 |

#### 3.1.1 Installing Golang

::: tip
**Go 1.18** is recommended for building and installing the Spartan NC Cosmos software.
:::

Install `go` by the following steps:

Download and untar the installation file

```shell
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

```shell
source /etc/profile
```

Now, check whether `go` is correctly installed:

```shell
go version
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-Cosmos/main/.github/images/1.go_version.jpg)

#### 3.1.2 Installing GCC

Install `gcc` by the following command:

```
yum install gcc
```

Check the version

```
gcc -v
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-Cosmos/main/.github/images/2.%20gcc.jpg)


#### 3.1.3 Installing Git

Install `gcc` by the system command

```
yum install git
```

Check the version

```
git version
```

![git](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/.github/images/git_version.jpg)

### 3.2 Building the Node

Now, you can compile and run the full node of Spartan-II Chain.

Make sure that your server can access Google because the Spartan project depends on some libraries provided by Google. (If not, you can also try to add a proxy in /etc/profile: `export GOPROXY=https://goproxy.io`)

```bash
git clone https://github.com/BSN-Spartan/NC-Cosmos.git
cd NC-Cosmos
git checkout v1.0.2
make install
```

Have you set up the Environment Variables correctly, you will get no error during Spartan Network installation.

Now check the `spartan` version.

```bash
spartan version
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-Cosmos/main/.github/images/3.spartanversion.jpg)


### 3.3 Initializing the Node

Create a directory `/spartan/` for node installation and the sub directory `config/` and `data/` to store the configuration file:

```shell
mkdir -p /spartan/config
mkdir -p /spartan/data
```

Copy [genesis.json](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/genesis.json), [app.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/app.toml) and [config.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/config.toml) from `spartan/` directory to `/spartan/config/` directory.

```shell
cp ./spartan/*  /spartan/config
```

Now the structure of /spartan is like below:

```shell
/spartan
????????? config
??????? ????????? app.toml
??????? ????????? config.toml
??????? ????????? genesis.json
????????? data

2 directories, 3 files
```
You can edit the ports in app.toml and config toml.
Some important ports are shown below:

* `Json RPC`:
  * `EVM module: ` http://<ip_address>:8545
  * `Native module: `  http://<ip_address>:26657
* `WebSocket`:
  * `EVM module:`  ws://<ip_address>:8546
  * `Native module:`  ws://<ip_address>:26657/websocket
* `gRPC`:
  * <ip_address>:9090


### 3.4 Starting Mainnet

:::tip NOTE
By now, you can run the node as a Full Node with this last step to join the mainnet. No extra steps are required.
:::

The final step is to??start the node. The node will then start producing blocks:

```shell
spartan start --home /spartan
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-Cosmos/main/.github/images/5.startnode.jpg)

Or you can start the node in the background via `nohup` command:

```shell
nohup spartan start --home /spartan >./output.log 2>&1 &
```

You can check the process of block synchronization from the log:

```
tail -f output.log
```

## 4. Full Node Installation by Docker

In this chapter, we will build a full node by docker images. If you have built a node by commands, you can skip this chapter and move forward.

### 4.1 Prerequisites

| Software  | Version  |
| ----- | ----- |
| Docker-ce | 18+ |

Run the following command to install Docker:

```shell
wget -qO- https://get.docker.com/ | sh
```

Grant your user permission to execute Docker commands:

```shell
sudo usermod -aG docker your-user
```

Now, check the docker version:

```shell
docker version
```
![](https://raw.githubusercontent.com/BSN-Spartan/NC-Cosmos/main/.github/images/4.1dockerversion.jpg)

Start Docker:

```shell
systemctl start docker
```


### 4.2 Building the Node

Official Docker images are hosted under the hub.docker.com registry. Run the following command to pull them to the server:

```shell
docker pull bsnspartan/nc-cosmos:latest
```


### 4.3 Initializing the Node

Create a directory `/spartan/` for node installation and the sub directory `config/` and `data/` to store the configuration file:

```shell
mkdir -p /spartan/config
mkdir -p /spartan/data
```

Copy [genesis.json](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/genesis.json), [app.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/app.toml) and [config.toml](https://github.com/BSN-Spartan/NC-Cosmos/blob/main/spartan/config.toml) from `spartan/` directory to `/spartan/config/` directory.

```shell
cp ./spartan/*  /spartan/config
```

Now the structure of /spartan is like below:

```shell
/spartan
????????? config
??????? ????????? app.toml
??????? ????????? config.toml
??????? ????????? genesis.json
????????? data

2 directories, 3 files
```

### 4.4 Starting Mainnet

Run command below to start the node:

```shell
docker run -d -p 9090:9090 -p 26656:26656 -p 26657:26657 -p 8545:8545 -p 8546:8546 -v /spartan:/spartan --restart=always --name spartan-nc-cosmos bsnspartan/nc-cosmos:latest spartan start --home /spartan
```

* `Json RPC`:
  * `EVM module: ` http://<ip_address>:8545
  * `Native module: `  http://<ip_address>:26657
* `WebSocket`:
  * `EVM module:`  ws://<ip_address>:8546
  * `Native module:`  ws://<ip_address>:26657/websocket
* `gRPC`:
  * <ip_address>:9090

You can check logs by command below:

```shell
docker logs -f spartan-nc-cosmos
```

## 5. Generating the Node Signature

When joining the Spartan Network as a Data Center, the Data Center Operator will be rewarded a certain amount of NTT Incentives based on the quantity of the registered node. To achieve this, the Data Center Operator should first provide the signature of the full node to verify the node's ownership.

### 5.1 Node Installed by Commands

Execute the following command in the node's data directory after the node is started.

```shell
spartan node sign-info --home /spartan
```

* `--home` is the data directory of the node. You need to specify this directory to store the data file of the node.

### 5.2 Node Installed by Docker

Execute below command:

```shell
docker exec spartan-nc-cosmos spartan node sign-info --home /spartan 
```

### 5.3 Node Signature

The node signature is like below:

```
{
      "node_id":"3b1e874ae98c36cdd5da32f9623fa0a4586b4445",
      "pub_key":"{\"@type\":\"/cosmos.crypto.ed25519.PubKey\",\"key\":\"3ggBZ5e7lJCoGsRE3CdNESSDv1fU8mbGDK4k8oNzJZQ=\"}",
      "signature":"Za1mHGzMXdYFs6RAmIr9Nge7P5Zw6k1m/eOW9iMzP9qhnaBMQEZ+LMoYxy+BqmZuJVxJALs2vKXcMXt5eHLtBQ=="
  }
```

Please submit it to the Data Center System when registering the node. Note that the value of "**pub_key**" filled into the "**Node Address**" blank needs to be unescaped like below:

```
 {"@type":"/cosmos.crypto.ed25519.PubKey","key":"3ggBZ5e7lJCoGsRE3CdNESSDv1fU8mbGDK4k8oNzJZQ="}
```

### 5.4 Node RPC URL

When joining the Spartan Network as a Data Center, the RPC URL is also needed. It is configured in the `config.toml` file, the default configuration is:

* `RPC URL: ` http://<ip_address>:26657

## 6. Deleting the Node

You can use the following command to stop the running node and delete it, and also clear the node data by deleting the data directory.

If the node has been registered in the Data Center, you can back up the `node_key.json` file stored in `/spartan/config` directory, so that you can recover this registered node when needed.

### 6.1 Ceasing the Node started by Commands

Use the following command to stop the running node:
```
pkill -INT spartan
```

### 6.2 Ceasing the Node started by Docker

Use the following command to stop the running container and delete the container and the image file:

```
docker stop spartan-nc-cosmos
docker rm spartan-nc-cosmos
docker rmi bsnspartan/nc-cosmos:latest
```

### 6.3 Deleting Node Data

If you need to completely delete all data of the node, you can use the following command to delete the data directory.

```
rm -rf /spartan
```

## 7.Resources

To find out more information about Spartan-II Chain (Powered by NC Cosmos), please visit [here](https://github.com/BSN-Spartan/NC-Cosmos/tree/main/docs)
