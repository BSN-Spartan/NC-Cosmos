---
order: 4
---

# Account

In the Spartan-Cosmos, an account designates a pair of public key `PubKey` and private key `PrivKey`. The `PubKey` can be derived to generate various `Addresses`, which are used to identify users (among other parties) in the application. `Addresses` are also associated with messages to identify the sender of the message. The `PrivKey` is used to generate digital signatures to prove that an `Address` associated with the `PrivKey` approved of a given message.

By default, the keyring (which holds the public/private key paris to interact with a node) generates a `eth_secp256k1` key. Cosmos `secp256k1` are not used due to compatibility issues with Ethereum transactions.

## Account Types

### Chain Account

The wallet address or contract address of a Non-Crypto Public Chain on the Spartan Network, which can be created by a Spartan Network user.

### Admin Account

The wallet address or contract address of a Non-Crypto Public Chain on the Spartan Network with Gas Credit management authority, which can only be created by the BSN Foundation.

## Key

### HD wallet

HD Wallets, originally specified in Bitcoin's [BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki), are a special kind of wallet that let users derive any number of accounts from a single seed. To understand what that means, let us first define some terminology:

- **Wallet**: Set of accounts obtained from a given seed.
- **Account**: A pair of public key/private key.
- **Private Key**: A private key is a secret piece of information used to sign messages. In the blockchain context, a private key identifies the owner of an account. The private key of a user should never be revealed to others.
- **Public Key**: A public key is a piece of information obtained by applying a one-way mathematical function on a private key. From it, an address can be derived. A private key cannot be found from a public key.
- **Address**: An address is a public string with a human-readable prefix that identifies an account. It is obtained by applying mathematical transformations to a public key.
- **Digital Signature**: A digital signature is a piece of cryptographic information that proves the owner of a given private key approved of a given message without revealing the private key.
- **Seed**: Same as Mnemonic.
- **Mnemonic**: A mnemonic is a sequence of words that is used as seed to derive private keys. The mnemonic is at the core of each wallet. **NEVER LOSE YOUR MNEMONIC. WRITE IT DOWN ON A PIECE OF PAPER AND STORE IT SOMEWHERE SAFE. IF YOU LOSE IT, THERE IS NO WAY TO RETRIEVE IT. IF SOMEONE GAINS ACCESS TO IT, THEY GAIN ACCESS TO ALL THE ASSOCIATED ACCOUNTS.**

At the core of a HD wallet, there is a seed. From this seed, users can deterministically generate accounts. To generate an account from a seed, one-way mathematical transformations are applied. To decide which account to generate, the user specifies a `path`, generally an `integer` (`0`, `1`, `2`, ...).

By specifying `path` to be `0` for example, the Wallet will generate `Private Key 0` from the seed. Then, `Public Key 0` can be generated from `Private Key 0`.  Finally, `Address 0` can be generated from `Public Key 0`. All these steps are one way only, meaning the `Public Key` cannot be found from the `Address`, the `Private Key` cannot be found from the `Public Key`, ...

```bash
     Account 0                         Account 1                         Account 2

+------------------+              +------------------+               +------------------+
|                  |              |                  |               |                  |
|    Address 0     |              |    Address 1     |               |    Address 2     |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|  Public key 0    |              |  Public key 1    |               |  Public key 2    |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|  Private key 0   |              |  Private key 1   |               |  Private key 2   |
|        ^         |              |        ^         |               |        ^         |
+------------------+              +------------------+               +------------------+
         |                                 |                                  |
         |                                 |                                  |
         |                                 |                                  |
         +--------------------------------------------------------------------+
                                           |
                                           |
                                 +---------+---------+
                                 |                   |
                                 |  Mnemonic (Seed)  |
                                 |                   |
                                 +-------------------+
```

The process of derivating accounts from the seed is deterministic. This means that given the same path, the derived private key will always be the same.

The funds stored in an account are controlled by the private key. This private key is generated using a one-way function from the mnemonic. If you lose the private key, you can retrieve it using the mnemonic. However, if you lose the mnemonic, you will lose access to all the derived private keys. Likewise, if someone gains access to your mnemonic, they gain access to all the associated accounts.

### Validator Key

This is a unique key used to sign consensus votes. It is generated when the node is created by `spartan init`. It is stored in the `config/priv_validator.json`

### Account Key

This key is created from `spartan` and used to sign transactions. Application keys are associated with a public key prefixed by `iap` and an address prefixed by `iaa`. Both are derived from account keys generated by `spartan keys add`.

**Note:** A validator's operator key is directly tied to an application key, but uses reserved prefixes solely for this purpose: `iva` and `ivp`.

## Bech32

Bech32 is a new Bitcoin address format proposed by Pieter Wuille and Greg Maxwell. Besides Bitcoin addresses, Bech32 can encode any short binary data. In the Spartan-Cosmos network, keys and addresses may refer to a number of different roles in the network like accounts, validators etc. The Spartan-Cosmos network is designed to use the Bech32 address format to provide robust integrity checks on data. The human readable part(HRP) makes it more efficient to read and the users could see error messages. More details in [bip-0173](https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki)

### Human Readable Part Table

| HRP | Definition                                     |
| --- | ---------------------------------------------- |
| iaa | Spartan-Cosmos Account Address                 |
| iap | Spartan-Cosmos Account Public Key              |
| iva | Spartan-Cosmos Validator's Operator Address    |
| ivp | Spartan-Cosmos Validator's Operator Public Key |
| ica | Tendermint Consensus Address                   |
| icp | Tendermint Consensus Public Key                |

### Encoding

Not all interfaces to Spartan-Cosmos users should be exposed as bech32 interfaces. Many addresses are still in hex or base64 encoded form.

To covert between other binary representation of addresses and keys, it is important to first apply the Amino encoding process before bech32 encoding.
