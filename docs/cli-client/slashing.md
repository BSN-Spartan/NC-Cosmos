# Slashing

The Slashing module can unjail validator previously jailed for downtime.

## Available Commands

| Type  | Name                                                   | Description                                 |
| ----- | ------------------------------------------------------ | ------------------------------------------- |
| Tx    | [unjail](#spartan-tx-slashing-unjail)                  | Unjail an jailed validator                  |
| Query | [params](#spartan-query-slashing-params)               | Query the current slashing parameters       |
| Query | [signing-info](#spartan-query-slashing-signing-info)   | Query a validator's signing information     |
| Query | [signing-infos](#spartan-query-slashing-signing-infos) | Query signing information of all validators |

## spartan tx slashing unjail

Unjail an jailed validator.

```bash
spartan tx slashing unjail [id] [flags]
```

## spartan query slashing params

Query the current slashing parameters.

```bash
spartan query slashing params [flags]
```

## spartan query slashing signing-info

Query a validator's signing information.

```bash
spartan query slashing signing-info [validator-conspub] [flags]
```

## spartan query slashing signing-infos

Query signing information of all validators.

```bash
spartan query slashing signing-infos [flags]
```
