# MT

## Overview

The MT module provides ERC-1155 capacity. The MT module is able to model off-chain assets into unique on-chain assets.

_For MT commands, refer to [MT CLI client](../cli-client/mt.md)_

## Concepts

On-chain assets are identified by`ID` . With the secure and tamper-proof nature of blockchains, the ownership of assets can be confirmed and verified. Transactions of assets between multiple parties can also be publicly recorded, in order to facilitate traceability and dispute resolution. The metadata (`data` ) of assets can be stored directly on the chain, or be used to record the off-chain storage addresses.

Assets need to be issued before they are created to declare their abstract properties.

- **DenomID:** the globally unique asset class identifier, which is generated on the chain.
- **DenomName:** the name of the asset class.

Each specific asset is described by the following elements:

- **DenomID:** the class of the asset.
- **ID:** the identifier of the asset, which is unique within the corresponding asset class. This ID is also generated on the chain.
- **Metadata:** a structure that contains specific data of the assets.
