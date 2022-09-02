# NFT

The NFT module provides ERC-721 capacity and is able to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain nft.

## Available Commands

| Type  | Name                                             | Description                                        |
| ----- | ------------------------------------------------ | -------------------------------------------------- |
| Tx    | [burn](#spartan-tx-nft-burn)                     | Burn an NFT                                        |
| Tx    | [edit](#spartan-tx-nft-edit)                     | Edit the token data of an NFT                      |
| Tx    | [issue](#spartan-tx-nft-issue)                   | Issue a new denom                                  |
| Tx    | [mint](#spartan-tx-nft-mint)                     | Mint an NFT and set the owner to the recipient     |
| Tx    | [transfer](#spartan-tx-nft-transfer)             | Transfer an NFT to a recipient                     |
| Tx    | [transfer-denom](#spartan-tx-nft-transfer-denom) | Transfer an Denom to a recipient                   |
| Query | [collection](#spartan-query-nft-collection)      | Get all the NFTs from a given collection           |
| Query | [denom](#spartan-query-nft-denom)                | Query the denom by the specified denom id          |
| Query | [denoms](#spartan-query-nft-denoms)              | Query all denominations of all collections of NFTs |
| Query | [token](#spartan-query-nft-token)                | Query a single NFT from a collection               |
| Query | [owner](#spartan-query-nft-owner)                | Get the NFTs owned by an account address           |
| Query | [supply](#spartan-query-nft-supply)              | Get total supply of a collection or owner of NFTs  |

## spartan tx nft burn

Burn an NFT.

```bash
spartan tx nft burn [denom-id] [nft-id] [flags]
```

## spartan tx nft edit

Edit the token data of an NFT.

```bash
spartan tx nft edit [denom-id] [nft-id] [flags]
```

## spartan tx nft issue

Issue a new denom.

```bash
spartan tx nft issue [denom-id] [flags]
```

## spartan tx nft mint

Mint an NFT and set the owner to the recipient.

```bash
spartan tx nft mint [denom-id] [nft-id] [flags]
```

## spartan tx nft transfer

Transfer an NFT to a recipient.

```bash
spartan tx nft transfer [recipient] [denom-id] [nft-id] [flags]
```

## spartan tx nft transfer-denom

Transfer an Denom to a recipient.

```bash
spartan tx nft transfer-denom [recipient] [denom-id] [flags]
```

## spartan query nft collection

Get all the NFTs from a given collection.

```bash
spartan query nft collection [denom-id] [flags]
```

## spartan query nft denom

Query the denom by the specified denom id.

```bash
spartan query nft denom [denom-id] [flags]
```

## spartan query nft denoms

Query all denominations of all collections of NFTs.

```bash
spartan query nft denoms [flags]
```

## spartan query nft token

Query a single NFT from a collection.

```bash
spartan query nft token [denom-id] [nft-id] [flags]
```

## spartan query nft owner

Get the NFTs owned by an account address.

```bash
spartan query nft owner [address] [flags]
```

## spartan query nft supply

Get total supply of a collection or owner of NFTs.

```bash
spartan query nft supply [denom-id] [flags]
```
