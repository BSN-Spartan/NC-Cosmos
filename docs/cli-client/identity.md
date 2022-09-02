# Identity

The Identity module builds a decentralized identity system (`DID`) that implements and extends the [W3C DID specification](https://www.w3.org/TR/did-core/).

## Available Commands

| Type  | Name                                        | Description                                              |
| ----- | ------------------------------------------- | -------------------------------------------------------- |
| Tx    | [create](#spartan-tx-identity-create)        | CreateValidator a new identity based on the given params |
| Tx    | [update](#spartan-tx-identity-update)        | UpdateValidator an existing identity                     |
| Query | [identity](#spartan-query-identity-identity) | Query details of an identity with the specified ID       |

## spartan tx identity create

CreateValidator a new identity based on the given params.

```bash
spartan tx identity create [flags]
```

## spartan tx identity update

UpdateValidator an existing identity.

```bash
spartan tx identity update [id] [flags]
```

## spartan query identity identity

Query details of an identity with the specified ID.

```bash
spartan query identity identity [id] [flags]
```