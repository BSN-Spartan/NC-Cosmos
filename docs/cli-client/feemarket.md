# Feemarket

The Feemarket module allows to define a global transaction fee for the network.

## Availale Commands

| Type  | Name                                            | Description                                     |
| ----- | ----------------------------------------------- | ----------------------------------------------- |
| Query | [base-fee](#spartan-query-feemarket-base-fee)   | Get the base fee amount at a given block height |
| Query | [block-gas](#spartan-query-feemarket-block-gas) | Get the block gas used at a given block height  |
| Query | [params](#spartan-query-feemarket-params)       | Get the fee market params                       |

## spartan query feemarket base-fee

Get the base fee amount at a given block height.

```bash
spartan query feemarket base-fee [height] [flags]
```

## spartan query feemarket block-gas

Get the block gas used at a given block height.

```bash
spartan query feemarket block-gas [height] [flags]
```

## spartan query feemarket params

Get the fee market parameter values.

```bash
spartan query feemarket params [flags]
```
