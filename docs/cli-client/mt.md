# MT

The MT module provides ERC-1155 capacity and is able to model off-chain assets into unique on-chain assets.

## Available Commands

| Type  | Name                                            | Description                     |
| ----- | ----------------------------------------------- | ------------------------------- |
| Tx    | [burn](#spartan-tx-mt-burn)                     | Burn amounts of an MT           |
| Tx    | [edit](#spartan-tx-mt-edit)                     | Edit the metadata of an MT      |
| Tx    | [mint](#spartan-tx-mt-mint)                     | Issue or mint MT                |
| Tx    | [transfer](#spartan-tx-mt-transfer)             | Transfer an MT to a recipient   |
| Tx    | [transfer-denom](#spartan-tx-mt-transfer-denom) | Transfer a denom to a recipient |
| Query | [balances](#spartan-query-mt-transfer-balances) | Query balances of an owner      |
| Query | [denom](#spartan-query-mt-transfer-denom)       | Query denom by ID               |
| Query | [denoms](#spartan-query-mt-transfer-denoms)     | Query all denoms                |
| Query | [token](#spartan-query-mt-transfer-token)       | Query MT by ID                  |
| Query | [tokens](#spartan-query-mt-transfer-tokens)     | Query all MTs of a denom        |

## spartan tx mt burn

Burn amounts of an MT.

```bash
spartan tx mt burn [denom-id] [mt-id] [amount] [flags]
```

## spartan tx mt edit

Edit the metadata of an MT.

```bash
spartan tx mt edit [denom-id] [mt-id] [flags]
```

## spartan tx mt mint

Issue or mint MT.

```bash
spartan tx mt mint [denom-id] [flags]
```

## spartan tx mt transfer

Transfer an MT to a recipient.

```bash
spartan tx mt transfer [from_key_or_address] [recipient] [denom-id] [mt-id] [amount] [flags]
```

## spartan tx mt transfer-denom

Transfer a denom to a recipient.

```bash
spartan tx mt transfer-denom [from_key_or_address] [recipient] [denom-id] [flags]
```

## spartan query mt balances

Query balances of an owner.

```bash
spartan query mt balances [owner] [denom-id] [flags]
```

## spartan query mt denom

Query denom by ID.

```bash
spartan query mt denom [denom-id] [flags]
```

## spartan query mt denoms

Query all denoms.

```bash
spartan query mt denoms [flags]
```

## spartan query mt token

Query MT by ID.

```bash
spartan query mt token [denom-id] [mt-id] [flags]
```

## spartan query mt tokens

Query all MTs of a denom.

```bash
spartan query mt tokens [denom-id] [flags]
```
