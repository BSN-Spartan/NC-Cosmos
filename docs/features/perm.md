# Perm

## Overview

The Perm module restricts the permissions of an account. This module is currently forbidden and reserved.

_For Perm commands, refer to [Perm CLI client](../cli-client/perm.md)_

## Concepts

Accounts in Spartan-Cosmos have their roles. Roles limit what an account can and cannot do.
- **ROOT_ADMIN:** There are no restrictions on this type of account.
- **PERM_ADMIN:** Accounts of this role can assign and unassign roles for other accounts.
- **BLOCKLIST_ADMIN:** If an account has the permission of BLOCKLIST_ADMIN, it can remove/add an account from/into the blocklist just like the Administer.
- **NODE_ADMIN:** Accounts of this role have access to the Node module, the Slashing module, and the Upgrade module.
- **PARAM_ADMIN:** Accounts of this role have access to the Parmas module.
- **ID_ADMIN:** Accounts of this role have access to the Identity module.
- **BASE_M1_ADMIN:** Accounts of this role can mint and reclaim Token.
- **POWER_USER:** Account of this role can issue and create NFT.
- **RELAYER_USER:** This role is reserved for a cross-chain relayer.
- **PLATFORM_USER:** Accounts of this role can transfer to any account. 
- **POWER_USER_ADMIN:** In order to create an NFT, an account must have the permission of POWER_USER, which PERM_ADMIN can only grant. For convenience, we add the role of POWER_USER_ADMIN, who can grant an account the permission of POWER_USER. Notice that the POWER_USER_ADMIN role can not grant other permissions.
