# NFT

## Overview

The NFT module provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain NFT.

_For NFT commands, refer to [NFT CLI client](../cli-client/nft.md)_

## Concepts

NFT on the chain is identified by `ID`. With the help of the security and non-tamperable features of the blockchain, the ownership of NFT will be clarified. The transaction process of NFT among members will also be publicly recorded to facilitate traceability and dispute settlement.

NFT metadata (`metadata`) can be stored directly on the chain, or the `URI` of its storage source outside the chain can be stored on the chain. NFT metadata is organized according to a specific [JSON Schema](https://JSON-Schema.org/).

NFT need to be issued before creation to declare their abstract properties:

- **Denom:** the globally unique NFT classification name

- **Denom ID:** the globally unique NFT classification identifier of Denom

- **Symbol:** the symbol of the token

- **Mint-restricted:** This field indicates whether there are restrictions on the issuance of NFTs under this classification, true means that only Denom owners can issue NFTs under this classification, and false means anyone can

- **Update-restricted:** This field indicates whether there are restrictions on updating NFTs under this classification, true means that no one under this classification can update the NFT, and false means that only the owner of this NFT can update

- **Metadata Specification:** The JSON Schema that NFT metadata should follow

Each specific NFT is described by the following elements:

- **Denom:** the classification of the nft

- **ID:** The identifier of the NFT, which is unique in this NFT denom; this ID is generated off-chain

- **Metadata:** The structure containing the specific data of the nft

- **Metadata URI:** When metadata is stored off-chain, this URI indicates its storage location
