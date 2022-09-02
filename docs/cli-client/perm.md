# Perm

The Perm module controls the role of an account or a contract.

## Available Commands

| Type  | Name                                                           | Description                    |
| ----- | -------------------------------------------------------------- | ------------------------------ |
| Tx    | [assign-roles](#spartan-tx-perm-assign-roles)                  | Assign roles to an account     |
| Tx    | [block-account](#spartan-tx-perm-block-account)                | Block an account               |
| Tx    | [block-contract](#spartan-tx-perm-block-contract)              | Block a contract               |
| Tx    | [unassign-roles](#spartan-tx-perm-unassign-roles)              | Unassign roles from an account |
| Tx    | [unblock-account](#spartan-tx-perm-unblock-account)            | Unblock an account             |
| Tx    | [unblock-contract](#spartan-tx-perm-unblock-contracts)         | Unblock a contract             |
| Query | [block-list-account](#spartan-query-perm-block-list-account)   | Query account blockList        |
| Query | [block-list-contract](#spartan-query-perm-block-list-contract) | Query contract blockList       |
| Query | [roles](#spartan-query-perm-roles)                             | Query a account roles          |

## spartan tx perm assign-roles

Assign roles to an account.

Auth options: `PERM_ADMIN`, `BLACKLIST_ADMIN`, `NODE_ADMIN`, `PARAM_ADMIN`, `ID_ADMIN`, `BASE_M1_ADMIN`, `POWER_USER`, `RELAYER_USER`, `PLATFORM_USER`, `POWER_USER_ADMIN`

```bash
spartan tx perm assign-roles [address] [roles] [flags]
```

## spartan tx perm block-account

Block an account.

```bash
spartan tx perm block-account [address] [flags]
```

## spartan tx perm block-contract

Block a contract.

```bash
spartan tx perm block-contract [contractAddress] [flags]
```

## spartan tx perm unassign-roles

Unassign roles from an account.

```bash
spartan tx perm unassign-roles [address] [roles] [flags]
```

## spartan tx perm unblock-account

Unblock an account.

```bash
spartan tx perm unblock-account [address] [flags]
```

## spartan tx perm unblock-contract

Unblock a contract.

```bash
spartan tx perm unblock-contract [contractAddress] [flags]
```

## spartan query perm block-list-account

Query account blockList

```bash
spartan query perm block-list-account [flags]
```

## spartan query perm block-list-contract

Query contract blockList.

```bash
spartan query perm block-list-contract  [flags]
```

## spartan query perm roles

Query a account roles.

```bash
spartan query perm roles [account] [flags]
```
