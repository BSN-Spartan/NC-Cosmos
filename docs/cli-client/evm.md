# EVM

The EVM module enables to make of an Ethereum smart contract transaction.

## Availale Commands

| Type  | Name                                  | Description                                             |
| ----- | ------------------------------------- | ------------------------------------------------------- |
| Tx    | [raw](#spartan-tx-evm-raw)            | Build cosmos transaction from raw ethereum transaction  |
| Query | [code](#spartan-query-evm-code)       | Gets code from an account                               |
| Query | [storage](#spartan-query-evm-storage) | Gets storage for an account with a given key and height |

## spartan tx evm raw

Build cosmos transaction from raw ethereum transaction

```bash
spartan tx evm raw [tx-hex] [flags]
```

## spartan query evm code

Gets code from an account.

```bash
spartan query evm code [tx-hex] [flags]
```

## spartan query evm storage

Gets storage for an account with a given key and height.

```bash
spartan query evm storage [address] [key] [flags]
```
