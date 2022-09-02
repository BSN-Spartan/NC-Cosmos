# Bank

The Bank module allows you to manage assets in your local accounts.

## Available Commands

| Type  | Name                                                | Description                                      |
| ----- | --------------------------------------------------- | ------------------------------------------------ |
| Tx    | [send](spartan-tx-bank-send)                        | Send funds from one account to another           |
| Query | [balances](spartan-query-bank-balances)             | Query for account balances by address            |
| Query | [denom-metadata](spartan-query-bank-denom-metadata) | Query the client metadata for coin denominations |
| Query | [total](spartan-query-bank-total)                   | Query the total supply of coins of the chain     |

## spartan tx bank send

Send funds from one account to another.

```bash
spartan tx bank send [from_key_or_address] [to_address] [amount] [flags]
```

## spartan query bank balances

Query the total balance of an account or of a specific denomination.

```bash
spartan query bank balances [address] [flags]
```

## spartan query bank denom-metadata

Query the client metadata for all the registered coin denominations.

```bash
spartan query bank denom-metadata [flags]
```

## spartan query bank total

Query total supply of coins that are held by accounts in the chain.

```bash
spartan query bank total [flags]
```
