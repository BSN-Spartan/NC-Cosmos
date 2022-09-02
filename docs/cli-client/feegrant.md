# Feegrant

The Feegrant module allows accounts to grant fee allowances and to use fees from their accounts. Grantees can execute any transaction without the need to maintain sufficient fees.

## Availale Commands

| Type  | Name                                     | Description                                       |
| ----- | ---------------------------------------- | ------------------------------------------------- |
| Tx    | [grant](#spartan-tx-feegrant-grant)      | Grant authorization to pay fees from your address |
| Tx    | [revoke](#spartan-tx-feegrant-revoke)    | Revoke fee grant from a granter to a grantee      |
| Query | [grant](#spartan-query-feegrant-grant)   | Query details for a grant                         |
| Query | [grants](#spartan-query-feegrant-grants) | Queries all the grants for a grantee address      |

## spartan tx feegrant grant

Grant authorization to pay fees from your address.

```bash
spartan tx feegrant grant [granter_key_or_address] [grantee] [flags]
```

## spartan tx feegrant revoke

Revoke fee grant from a granter to a grantee.

```bash
spartan tx feegrant revoke [granter] [grantee] [flags]
```

## spartan query feegrant grant

Query details for a grant.

```bash
spartan query feegrant grant [granter] [grantee]
```

## spartan query feegrant grants

Queries all the grants for a grantee address.

```bash
spartan query feegrant grants [grantee]
```
