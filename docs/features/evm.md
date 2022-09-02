# EVM

## Overview

The growth of EVM-based chains (e.g. Ethereum), however, has uncovered several scalability challenges that are often referred to as the [Trilemma of decentralization, security, and scalability](https://vitalik.ca/general/2021/04/07/sharding.html). Developers are frustrated by high gas fees, slow transaction speed & throughput, and chain-specific governance that can only undergo slow change because of its wide range of deployed applications. A solution is required that eliminates these concerns for developers, who build applications within a familiar EVM environment.

The EVM module provides this EVM familiarity on a scalable, high-throughput Proof-of-Stake blockchain. It is built as a  [Cosmos SDK module](https://docs.cosmos.network/master/building-modules/intro.html)which allows for the deployment of smart contracts, interaction with the EVM state machine (state transitions), and the use of EVM tooling.

_For EVM Commands, refer to [EVM CLI client](../cli-client/evm.md)_

_For EVM Endpoints, refer to [EVM JSON-RPC](../endpoints/evm-json-rpc.md)_

## Concepts

### EVM

The Ethereum Virtual Machine (EVM) is a computation engine which can be thought of as one single entity maintained by thousands of connected computers (nodes) running an Ethereum client. As a virtual machine ([VM ](https://en.wikipedia.org/wiki/Virtual_machine)), the EVM is responisble for computing changes to the state deterministically regardless of its environment (hardware and OS). This means that every node has to get the exact same result given an identical starting state and transaction (tx).

The EVM is considered to be the part of the Ethereum protocol that handles the deployment and execution of [smart contracts](https://ethereum.org/en/developers/docs/smart-contracts/). To make a clear distinction:

- The Ethereum protocol describes a blockchain, in which all Ethereum accounts and smart contracts live. It has only one canonical state (a data structure, which keeps all accounts) at any given block in the chain.
- The EVM, however, is the [state machine](https://en.wikipedia.org/wiki/Finite-state_machine)that defines the rules for computing a new valid state from block to block. It is an isolated runtime, which means that code running inside the EVM has no access to network, filesystem, or other processes (not external APIs).

The EVM module implements the EVM as a Cosmos SDK module. It allows users to interact with the EVM by submitting Ethereum txs and executing their containing messages on the given state to evoke a state transition.

### State

The Ethereum state is a data structure, implemented as a [Merkle Patricia Trie](https://en.wikipedia.org/wiki/Merkle_tree), that keeps all accounts on the chain. The EVM makes changes to this data structure resulting in a new state with a different State Root. Ethereum can therefore be seen as a state chain that transitions from one state to another by executing transations in a block using the EVM. A new block of txs can be described through its Block header (parent hash, block number, time stamp, nonce, receipts,...).

### Accounts

There are two types of accounts that can be stored in state at a given address:

- **Externally Owned Account (EOA)**: Has nonce (tx counter) and balance
- **Smart Contract**: Has nonce, balance, (immutable) code hash, storage root (another Merkle Patricia Trie)

Smart contracts are just like regular accounts on the blockchain, which additionally store executable code in an Ethereum-specific binary format, known as **EVM bytecode**. They are typically written in an Ethereum high level language such as Solidity which is compiled down to EVM bytecode and deployed on the blockchain, by submitting a tx using an Ethereum client.

## Functions

### API-related features

Spartan-Cosmos supports all features of `EVM`. For the related `API` documentation, you can refer to: [EVM API](https://eth.wiki/json-rpc/API)

### Export the `ETH` private key of the account

```shell
spartan keys unsafe-export-eth-key [name]
```

### Import the account's `ETH` private key

```shell
spartan keys unsafe-import-eth-key [name] [pk]
```

### spartan tx evm raw

Build cosmos transaction from raw ethereum transaction

```bash
spartan tx evm raw [tx-hex] [flags]
```

### spartan query evm code

Gets code from an account.

```bash
spartan query evm code [tx-hex] [flags]
```

### spartan query evm storage

Gets storage for an account with a given key and height.

```bash
spartan query evm storage [address] [key] [flags]
```

## Hardhat: Deploying a Smart Contract

[Hardhat](https://hardhat.org/) is a flexible development environment for building Ethereum-based smart contracts. It is designed with integrations and extensibility in mind

### Install Dependencies

Before proceeding, you need to install Node.js (we'll use v16.x) and the npm package manager. You can download directly from [Node.js](https://nodejs.org/en/download/) or in your terminal:

```bash
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -

sudo apt install -y nodejs
```

You can verify that everything is installed correctly by querying the version for each package:

```bash
$ node -v
...

$ npm -v
...
```

### Create Hardhat Project

To create a new project, navigate to your project directory and run:

```bash
$ npx hardhat

888    888                      888 888               888
888    888                      888 888               888
888    888                      888 888               888
8888888888  8888b.  888d888 .d88888 88888b.   8888b.  888888
888    888     "88b 888P"  d88" 888 888 "88b     "88b 888
888    888 .d888888 888    888  888 888  888 .d888888 888
888    888 888  888 888    Y88b 888 888  888 888  888 Y88b.
888    888 "Y888888 888     "Y88888 888  888 "Y888888  "Y888

👷 Welcome to Hardhat v2.9.3 👷‍

? What do you want to do? …
  Create a basic sample project
❯ Create an advanced sample project
  Create an advanced sample project that uses TypeScript
  Create an empty hardhat.config.js
  Quit
```

Following the prompts should create a new project structure in your directory. Consult the [Hardhat config page](https://hardhat.org/config/) for a list of configuration options to specify in `hardhat.config.js`. Most importantly, you should set the `defaultNetwork` entry to point to your desired JSON-RPC network:

```javascript
// Local Node
module.exports = {
  defaultNetwork: "local",
  networks: {
    hardhat: {
    },
    local: {
      url: "http://localhost:8545/",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}
```

```javascript
// Testnet
module.exports = {
  defaultNetwork: "testnet",
  networks: {
    hardhat: {
    },
    testnet: {
      url: "https://eth.bd.evmos.dev:8545",
      accounts: [privateKey1, privateKey2, ...]
    }
  },
  ...
}
```

To ensure you are targeting the correct network, you can query for a list of accounts available to you from your default network provider:

```bash
$ npx hardhat accounts
0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
0x70997970C51812dc3A010C7d01b50e0d17dc79C8
0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
0x90F79bf6EB2c4f870365E785982E1f101E93b906
...
```

### Deploying a Smart Contract

You will see that a default smart contract, written in Solidity, has already been provided under `contracts/Greeter.sol`:

```javascript
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract Greeter {
    string private greeting;

    constructor(string memory _greeting) {
        console.log("Deploying a Greeter with greeting:", _greeting);
        greeting = _greeting;
    }

    function greet() public view returns (string memory) {
        return greeting;
    }

    function setGreeting(string memory _greeting) public {
        console.log("Changing greeting from '%s' to '%s'", greeting, _greeting);
        greeting = _greeting;
    }
}
```

This contract allows you to set and query a string `greeting`. Hardhat also provides a script to deploy smart contracts to a target network; this can be invoked via the following command, targeting your default network:

```bash
npx hardhat run scripts/deploy.js
```

Hardhat also lets you manually specify a target network via the `--network <your-network>` flag:

Local Node

```bash
npx hardhat run --network {{ $themeConfig.project.rpc_url_local }} scripts/deploy.js
```

Testnet

```bash
npx hardhat run --network {{ $themeConfig.project.rpc_url_testnet }} scripts/deploy.js
```

Finally, try running a Hardhat test:

```bash
$ npx hardhat test
Compiling 1 file with 0.8.4
Compilation finished successfully


  Greeter
Deploying a Greeter with greeting: Hello, world!
Changing greeting from 'Hello, world!' to 'Hola, mundo!'
    ✓ Should return the new greeting once it's changed (803ms)


  1 passing (805ms)
```
