# Feegrant

## Oveview

The Feegrant module allows accounts to grant fee allowances and to use fees from their accounts. Grantees can execute any transaction without the need to maintain sufficient fees.

_For Feegrant commands, refer to [Feegrant CLI client](../cli-client/feegrant.md)_

## Concepts

### Grant

`Grant` is stored in the KVStore to record a grant with full context. Every grant will contain `granter`, `grantee` and what kind of `allowance` is granted. `granter` is an account address who is giving permission to `grantee` (the beneficiary account address) to pay for some or all of `grantee`'s transaction fees. `allowance` defines what kind of fee allowance (`BasicAllowance` or `PeriodicAllowance`, see below) is granted to `grantee`. `allowance` accepts an interface which implements `FeeAllowanceI`, encoded as `Any` type. There can be only one existing fee grant allowed for a `grantee` and `granter`, self grants are not allowed.

### Fee Allowance types

There are two types of fee allowances present at the moment:

- `BasicAllowance`
- `PeriodicAllowance`
- `AllowedMsgAllowance`

### BasicAllowance

`BasicAllowance` is permission for `grantee` to use fee from a `granter`'s account. If any of the `spend_limit` or `expiration` reaches its limit, the grant will be removed from the state.

- `spend_limit` is the limit of coins that are allowed to be used from the `granter` account. If it is empty, it assumes there's no spend limit, `grantee` can use any number of available coins from `granter` account address before the expiration.

- `expiration` specifies an optional time when this allowance expires. If the value is left empty, there is no expiry for the grant.

- When a grant is created with empty values for `spend_limit` and `expiration`, it is still a valid grant. It won't restrict the `grantee` to use any number of coins from `granter` and it won't have any expiration. The only way to restrict the `grantee` is by revoking the grant.

### PeriodicAllowance

`PeriodicAllowance` is a repeating fee allowance for the mentioned period, we can mention when the grant can expire as well as when a period can reset. We can also define the maximum number of coins that can be used in a mentioned period of time.

- `basic` is the instance of `BasicAllowance` which is optional for periodic fee allowance. If empty, the grant will have no `expiration` and no `spend_limit`.

- `period` is the specific period of time, after each period passes, `period_can_spend` will be reset.

- `period_spend_limit` specifies the maximum number of coins that can be spent in the period.

- `period_can_spend` is the number of coins left to be spent before the period_reset time.

- `period_reset` keeps track of when a next period reset should happen.

### AllowedMsgAllowance

`AllowedMsgAllowance` is a fee allowance, it can be any of `BasicFeeAllowance`, `PeriodicAllowance` but restricted only to the allowed messages mentioned by the granter.

- `allowance` is either `BasicAllowance` or `PeriodicAllowance`.

- `allowed_messages` is array of messages allowed to execute the given allowance.
