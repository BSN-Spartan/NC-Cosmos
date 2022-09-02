---
order: 2
---

# Node

## Accounting Node

Accounting Nodes are nodes that do not participate in transaction consensus. In Spartan-Cosmos, such nodes are full nodes.

A full-node is a program that fully validates transactions and blocks of a blockchain. It is distinct from a light-node that only processes block headers and a small subset of transactions. Running a full-node requires more resources than a light-node but is necessary in order to be a validator. In practice, running a full-node only implies running a non-compromised and up-to-date version of the software with low network latency and without downtime.

## Consensus Node

Consensus Nodes are nodes that participate in transaction consensus. In Spartan-Cosmos, such nodes are Validators.

Based on [Cosmos SDK](https://cosmos.network/docs/intro/)and [Tendermint](https://tendermint.com/docs/introduction/what-is-tendermint.html), Spartan-Cosmos relies on a set of validators to secure the network. The role of validators is to run a full-node and participate in consensus by broadcasting votes which contain cryptographic signatures signed by their private keys. Validators commit new blocks to the blockchain and they must also participate in governance by voting on proposals.

In the Non-Crypto Public Chain Network,  Validators, or Consensus Nodes, will not get any block reward for their work. Also, their weight power is decided in their joining process to be  Consensus Nodes.

### Responsibilites

Validators have two main responsibilities:

- **Be able to constantly run a correct version of the software:** Validators need to make sure that their servers are always online and their private keys are not compromised.
- **Actively participate in governance:** Validators are required to vote on every proposal.

Additionally, validators are expected to be active members of the community. They should always be up-to-date with the current state of the ecosystem so that they can easily adapt to any change.
