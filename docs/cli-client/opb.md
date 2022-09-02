# OPB

The OPB module restricts transfers between accounts.

## Available Commands

| Type  | Name                                | Description                                                                           |
| ----- | ----------------------------------- | ------------------------------------------------------------------------------------- |
| Tx    | [mint](#spartan-tx-opb-mint)        | Mint the base native token by the given amount in main unit                           |
| Tx    | [reclaim](#spartan-tx-opb-reclaim)  | Reclaim the native token of the specified denom from the corresponding escrow account |
| Query | [params](#spartan-query-opb-params) | Query the current OPB parameter set                                                   |

## spartan tx opb mint

Mint the base native token by the given amount in main unit.

```bash
spartan tx opb mint [amount] [to] [flags]
```

## spartan tx opb reclaim

Reclaim the native token of the specified denom from the corresponding escrow account.

```bash
spartan tx opb reclaim [denom] [to] [flags]
```

## spartan query opb params

Query the current OPB parameter set.

```bash
spartan query opb params [flags]
```
