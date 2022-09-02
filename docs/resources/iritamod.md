---
order: 4
---
# IRITA Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [genutil/genesis.proto](#genutil_genesis-proto)
    - [GenesisState](#cschain-genutil-GenesisState)
  
- [identity/genesis.proto](#identity_genesis-proto)
    - [GenesisState](#iritamod-identity-GenesisState)
  
- [identity/identity.proto](#identity_identity-proto)
    - [Identity](#iritamod-identity-Identity)
    - [PubKeyInfo](#iritamod-identity-PubKeyInfo)
  
    - [PubKeyAlgorithm](#iritamod-identity-PubKeyAlgorithm)
  
- [identity/query.proto](#identity_query-proto)
    - [QueryIdentityRequest](#iritamod-identity-QueryIdentityRequest)
    - [QueryIdentityResponse](#iritamod-identity-QueryIdentityResponse)
  
    - [Query](#iritamod-identity-Query)
  
- [identity/tx.proto](#identity_tx-proto)
    - [MsgCreateIdentity](#iritamod-identity-MsgCreateIdentity)
    - [MsgCreateIdentityResponse](#iritamod-identity-MsgCreateIdentityResponse)
    - [MsgUpdateIdentity](#iritamod-identity-MsgUpdateIdentity)
    - [MsgUpdateIdentityResponse](#iritamod-identity-MsgUpdateIdentityResponse)
  
    - [Msg](#iritamod-identity-Msg)
  
- [node/genesis.proto](#node_genesis-proto)
    - [GenesisState](#iritamod-node-GenesisState)
  
- [node/node.proto](#node_node-proto)
    - [HistoricalInfo](#iritamod-node-HistoricalInfo)
    - [Node](#iritamod-node-Node)
    - [Params](#iritamod-node-Params)
    - [Validator](#iritamod-node-Validator)
  
- [node/query.proto](#node_query-proto)
    - [QueryNodeRequest](#iritamod-node-QueryNodeRequest)
    - [QueryNodeResponse](#iritamod-node-QueryNodeResponse)
    - [QueryNodesRequest](#iritamod-node-QueryNodesRequest)
    - [QueryNodesResponse](#iritamod-node-QueryNodesResponse)
    - [QueryParamsRequest](#iritamod-node-QueryParamsRequest)
    - [QueryParamsResponse](#iritamod-node-QueryParamsResponse)
    - [QueryValidatorRequest](#iritamod-node-QueryValidatorRequest)
    - [QueryValidatorResponse](#iritamod-node-QueryValidatorResponse)
    - [QueryValidatorsRequest](#iritamod-node-QueryValidatorsRequest)
    - [QueryValidatorsResponse](#iritamod-node-QueryValidatorsResponse)
  
    - [Query](#iritamod-node-Query)
  
- [node/tx.proto](#node_tx-proto)
    - [MsgCreateValidator](#iritamod-node-MsgCreateValidator)
    - [MsgCreateValidatorResponse](#iritamod-node-MsgCreateValidatorResponse)
    - [MsgGrantNode](#iritamod-node-MsgGrantNode)
    - [MsgGrantNodeResponse](#iritamod-node-MsgGrantNodeResponse)
    - [MsgRemoveValidator](#iritamod-node-MsgRemoveValidator)
    - [MsgRemoveValidatorResponse](#iritamod-node-MsgRemoveValidatorResponse)
    - [MsgRevokeNode](#iritamod-node-MsgRevokeNode)
    - [MsgRevokeNodeResponse](#iritamod-node-MsgRevokeNodeResponse)
    - [MsgUpdateValidator](#iritamod-node-MsgUpdateValidator)
    - [MsgUpdateValidatorResponse](#iritamod-node-MsgUpdateValidatorResponse)
  
    - [Msg](#iritamod-node-Msg)
  
- [opb/genesis.proto](#opb_genesis-proto)
    - [GenesisState](#iritamod-opb-GenesisState)
  
- [opb/opb.proto](#opb_opb-proto)
    - [Params](#iritamod-opb-Params)
  
- [opb/query.proto](#opb_query-proto)
    - [QueryParamsRequest](#iritamod-opb-QueryParamsRequest)
    - [QueryParamsResponse](#iritamod-opb-QueryParamsResponse)
  
    - [Query](#iritamod-opb-Query)
  
- [opb/tx.proto](#opb_tx-proto)
    - [MsgMint](#iritamod-opb-MsgMint)
    - [MsgMintResponse](#iritamod-opb-MsgMintResponse)
    - [MsgReclaim](#iritamod-opb-MsgReclaim)
    - [MsgReclaimResponse](#iritamod-opb-MsgReclaimResponse)
  
    - [Msg](#iritamod-opb-Msg)
  
- [params/params.proto](#params_params-proto)
    - [ParamChange](#iritamod-params-ParamChange)
  
- [params/tx.proto](#params_tx-proto)
    - [MsgUpdateParams](#iritamod-params-MsgUpdateParams)
    - [MsgUpdateParamsResponse](#iritamod-params-MsgUpdateParamsResponse)
  
    - [Msg](#iritamod-params-Msg)
  
- [perm/genesis.proto](#perm_genesis-proto)
    - [GenesisState](#iritamod-perm-GenesisState)
    - [RoleAccount](#iritamod-perm-RoleAccount)
  
- [perm/perm.proto](#perm_perm-proto)
    - [Role](#iritamod-perm-Role)
  
- [perm/query.proto](#perm_query-proto)
    - [QueryBlockListRequest](#iritamod-perm-QueryBlockListRequest)
    - [QueryBlockListResponse](#iritamod-perm-QueryBlockListResponse)
    - [QueryContractDenyList](#iritamod-perm-QueryContractDenyList)
    - [QueryContractDenyListResponse](#iritamod-perm-QueryContractDenyListResponse)
    - [QueryRolesRequest](#iritamod-perm-QueryRolesRequest)
    - [QueryRolesResponse](#iritamod-perm-QueryRolesResponse)
  
    - [Query](#iritamod-perm-Query)
  
- [perm/tx.proto](#perm_tx-proto)
    - [MsgAssignRoles](#iritamod-perm-MsgAssignRoles)
    - [MsgAssignRolesResponse](#iritamod-perm-MsgAssignRolesResponse)
    - [MsgBlockAccount](#iritamod-perm-MsgBlockAccount)
    - [MsgBlockAccountResponse](#iritamod-perm-MsgBlockAccountResponse)
    - [MsgBlockContract](#iritamod-perm-MsgBlockContract)
    - [MsgBlockContractResponse](#iritamod-perm-MsgBlockContractResponse)
    - [MsgUnassignRoles](#iritamod-perm-MsgUnassignRoles)
    - [MsgUnassignRolesResponse](#iritamod-perm-MsgUnassignRolesResponse)
    - [MsgUnblockAccount](#iritamod-perm-MsgUnblockAccount)
    - [MsgUnblockAccountResponse](#iritamod-perm-MsgUnblockAccountResponse)
    - [MsgUnblockContract](#iritamod-perm-MsgUnblockContract)
    - [MsgUnblockContractResponse](#iritamod-perm-MsgUnblockContractResponse)
  
    - [Msg](#iritamod-perm-Msg)
  
- [slashing/slashing.proto](#slashing_slashing-proto)
    - [Params](#iritamod-slashing-Params)
    - [ValidatorSigningInfo](#iritamod-slashing-ValidatorSigningInfo)
  
- [slashing/tx.proto](#slashing_tx-proto)
    - [MsgUnjailValidator](#iritamod-slashing-MsgUnjailValidator)
    - [MsgUnjailValidatorResponse](#iritamod-slashing-MsgUnjailValidatorResponse)
  
    - [Msg](#iritamod-slashing-Msg)
  
- [upgrade/tx.proto](#upgrade_tx-proto)
    - [MsgCancelUpgrade](#iritamod-upgrade-MsgCancelUpgrade)
    - [MsgCancelUpgradeResponse](#iritamod-upgrade-MsgCancelUpgradeResponse)
    - [MsgUpgradeSoftware](#iritamod-upgrade-MsgUpgradeSoftware)
    - [MsgUpgradeSoftwareResponse](#iritamod-upgrade-MsgUpgradeSoftwareResponse)
  
    - [Msg](#iritamod-upgrade-Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="genutil_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## genutil/genesis.proto



<a name="cschain-genutil-GenesisState"></a>

### GenesisState
GenesisState defines the raw genesis transaction in JSON


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| gen_txs | [bytes](#bytes) | repeated |  |





 

 

 

 



<a name="identity_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/genesis.proto



<a name="iritamod-identity-GenesisState"></a>

### GenesisState
GenesisState defines the identity module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identities | [Identity](#iritamod-identity-Identity) | repeated |  |





 

 

 

 



<a name="identity_identity-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/identity.proto



<a name="iritamod-identity-Identity"></a>

### Identity
Identity defines a struct for an identity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pub_keys | [PubKeyInfo](#iritamod-identity-PubKeyInfo) | repeated |  |
| certificates | [string](#string) | repeated |  |
| credentials | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="iritamod-identity-PubKeyInfo"></a>

### PubKeyInfo
PubKey represents a public key along with the corresponding algorithm


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pub_key | [string](#string) |  |  |
| algorithm | [PubKeyAlgorithm](#iritamod-identity-PubKeyAlgorithm) |  |  |





 


<a name="iritamod-identity-PubKeyAlgorithm"></a>

### PubKeyAlgorithm
PubKeyAlgorithm defines the algorithm names for the public key

| Name | Number | Description |
| ---- | ------ | ----------- |
| UnknownPubKeyAlgorithm | 0 | UnknownPubKeyAlgorithm defines an unknown algorithm name |
| RSA | 1 | RSA defines a RSA algorithm name |
| DSA | 2 | DSA defines a DSA algorithm name. |
| ECDSA | 3 | ECDSA defines an ECDSA algorithm name. |
| ED25519 | 4 | ED25519 defines an ED25519 algorithm name. |
| SM2 | 5 | SM2 defines an SM2 algorithm name. |


 

 

 



<a name="identity_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/query.proto



<a name="iritamod-identity-QueryIdentityRequest"></a>

### QueryIdentityRequest
QueryIdentityRequest is request type for the Query/Identity RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="iritamod-identity-QueryIdentityResponse"></a>

### QueryIdentityResponse
QueryIdentityResponse is response type for the Query/Identity RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| identity | [Identity](#iritamod-identity-Identity) |  |  |





 

 

 


<a name="iritamod-identity-Query"></a>

### Query
Query defines the gRPC querier service for identity module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Identity | [QueryIdentityRequest](#iritamod-identity-QueryIdentityRequest) | [QueryIdentityResponse](#iritamod-identity-QueryIdentityResponse) | Identity queries the identity by the given id |

 



<a name="identity_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity/tx.proto



<a name="iritamod-identity-MsgCreateIdentity"></a>

### MsgCreateIdentity
MsgCreateIdentity defines a message to create an identity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pub_key | [PubKeyInfo](#iritamod-identity-PubKeyInfo) |  |  |
| certificate | [string](#string) |  |  |
| credentials | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="iritamod-identity-MsgCreateIdentityResponse"></a>

### MsgCreateIdentityResponse
MsgCreateIdentityResponse defines the Msg/Create response type.






<a name="iritamod-identity-MsgUpdateIdentity"></a>

### MsgUpdateIdentity
MsgUpdateIdentity defines a message to update an identity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pub_key | [PubKeyInfo](#iritamod-identity-PubKeyInfo) |  |  |
| certificate | [string](#string) |  |  |
| credentials | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="iritamod-identity-MsgUpdateIdentityResponse"></a>

### MsgUpdateIdentityResponse
MsgUpdateIdentityResponse defines the Msg/Update response type.





 

 

 


<a name="iritamod-identity-Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateIdentity | [MsgCreateIdentity](#iritamod-identity-MsgCreateIdentity) | [MsgCreateIdentityResponse](#iritamod-identity-MsgCreateIdentityResponse) | CreateIdentity defines a method for creating a new identity. |
| UpdateIdentity | [MsgUpdateIdentity](#iritamod-identity-MsgUpdateIdentity) | [MsgUpdateIdentityResponse](#iritamod-identity-MsgUpdateIdentityResponse) | UpdateIdentity defines a method for Updating a identity. |

 



<a name="node_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## node/genesis.proto



<a name="iritamod-node-GenesisState"></a>

### GenesisState
GenesisState defines the node module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root_cert | [string](#string) |  |  |
| params | [Params](#iritamod-node-Params) |  |  |
| validators | [Validator](#iritamod-node-Validator) | repeated |  |
| nodes | [Node](#iritamod-node-Node) | repeated |  |





 

 

 

 



<a name="node_node-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## node/node.proto



<a name="iritamod-node-HistoricalInfo"></a>

### HistoricalInfo
HistoricalInfo contains the historical information that gets stored at
each height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| header | [tendermint.types.Header](#tendermint-types-Header) |  |  |
| valset | [Validator](#iritamod-node-Validator) | repeated |  |






<a name="iritamod-node-Node"></a>

### Node
Node defines a struct to represent a node identity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| certificate | [string](#string) |  |  |






<a name="iritamod-node-Params"></a>

### Params
Params defines the parameters for the node module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| historical_entries | [uint32](#uint32) |  |  |






<a name="iritamod-node-Validator"></a>

### Validator
Request defines a standard for validator. The validator will participate the
blockchain consensus, power determines the probability of proposing a new block.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| pubkey | [string](#string) |  |  |
| certificate | [string](#string) |  |  |
| power | [int64](#int64) |  |  |
| description | [string](#string) |  |  |
| jailed | [bool](#bool) |  |  |
| operator | [string](#string) |  |  |





 

 

 

 



<a name="node_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## node/query.proto



<a name="iritamod-node-QueryNodeRequest"></a>

### QueryNodeRequest
QueryNodeRequest is the request type for the Query/Node RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="iritamod-node-QueryNodeResponse"></a>

### QueryNodeResponse
QueryNodeResponse is the response type for the Query/Node RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#iritamod-node-Node) |  |  |






<a name="iritamod-node-QueryNodesRequest"></a>

### QueryNodesRequest
QueryNodesRequest is the request type for the Query/Nodes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.query.PageRequest](#cosmos-query-PageRequest) |  |  |






<a name="iritamod-node-QueryNodesResponse"></a>

### QueryNodesResponse
QueryNodesResponse is the response type for the Query/Nodes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#iritamod-node-Node) | repeated |  |
| pagination | [cosmos.query.PageResponse](#cosmos-query-PageResponse) |  |  |






<a name="iritamod-node-QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method






<a name="iritamod-node-QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#iritamod-node-Params) |  |  |






<a name="iritamod-node-QueryValidatorRequest"></a>

### QueryValidatorRequest
QueryValidatorRequest is the request type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="iritamod-node-QueryValidatorResponse"></a>

### QueryValidatorResponse
QueryValidatorResponse is the response type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validator | [Validator](#iritamod-node-Validator) |  |  |






<a name="iritamod-node-QueryValidatorsRequest"></a>

### QueryValidatorsRequest
QueryValidatorsRequest is the request type for the Query/Validators RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.query.PageRequest](#cosmos-query-PageRequest) |  |  |






<a name="iritamod-node-QueryValidatorsResponse"></a>

### QueryValidatorsResponse
QueryValidatorsResponse is the response type for the Query/Validators RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validators | [Validator](#iritamod-node-Validator) | repeated |  |
| pagination | [cosmos.query.PageResponse](#cosmos-query-PageResponse) |  |  |





 

 

 


<a name="iritamod-node-Query"></a>

### Query
Query defines the gRPC querier service for node module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Validator | [QueryValidatorRequest](#iritamod-node-QueryValidatorRequest) | [QueryValidatorResponse](#iritamod-node-QueryValidatorResponse) | Validator queries the validator by the given id |
| Validators | [QueryValidatorsRequest](#iritamod-node-QueryValidatorsRequest) | [QueryValidatorsResponse](#iritamod-node-QueryValidatorsResponse) | Validators queries the validators |
| Node | [QueryNodeRequest](#iritamod-node-QueryNodeRequest) | [QueryNodeResponse](#iritamod-node-QueryNodeResponse) | Node queries the node by the given id |
| Nodes | [QueryNodesRequest](#iritamod-node-QueryNodesRequest) | [QueryNodesResponse](#iritamod-node-QueryNodesResponse) | Nodes queries the nodes |
| Params | [QueryParamsRequest](#iritamod-node-QueryParamsRequest) | [QueryParamsResponse](#iritamod-node-QueryParamsResponse) | Params queries the parameters of the node module |

 



<a name="node_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## node/tx.proto



<a name="iritamod-node-MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator defines an SDK message for creating a new validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| certificate | [string](#string) |  |  |
| power | [int64](#int64) |  |  |
| description | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-node-MsgCreateValidatorResponse"></a>

### MsgCreateValidatorResponse
MsgCreateValidatorResponse defines the Msg/CreateValidator response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="iritamod-node-MsgGrantNode"></a>

### MsgGrantNode
MsgGrantNode defines a message to grant a node access


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| certificate | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-node-MsgGrantNodeResponse"></a>

### MsgGrantNodeResponse
MsgGrantNodeResponse defines the Msg/GrantNode response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="iritamod-node-MsgRemoveValidator"></a>

### MsgRemoveValidator
MsgRemoveValidator defines an SDK message for removing an existing validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-node-MsgRemoveValidatorResponse"></a>

### MsgRemoveValidatorResponse
MsgRemoveValidatorResponse defines the Msg/RemoveValidator response type.






<a name="iritamod-node-MsgRevokeNode"></a>

### MsgRevokeNode
MsgRevokeNode defines a message to revoke access from a node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-node-MsgRevokeNodeResponse"></a>

### MsgRevokeNodeResponse
MsgRevokeNodeResponse defines the Msg/RevokeNode response type.






<a name="iritamod-node-MsgUpdateValidator"></a>

### MsgUpdateValidator
MsgUpdateValidator defines an SDK message for updating an existing validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| certificate | [string](#string) |  |  |
| power | [int64](#int64) |  |  |
| description | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-node-MsgUpdateValidatorResponse"></a>

### MsgUpdateValidatorResponse
MsgUpdateValidatorResponse defines the Msg/UpdateValidator response type.





 

 

 


<a name="iritamod-node-Msg"></a>

### Msg
Msg defines the node Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateValidator | [MsgCreateValidator](#iritamod-node-MsgCreateValidator) | [MsgCreateValidatorResponse](#iritamod-node-MsgCreateValidatorResponse) | CreateValidator defines a method for creating a validator. |
| UpdateValidator | [MsgUpdateValidator](#iritamod-node-MsgUpdateValidator) | [MsgUpdateValidatorResponse](#iritamod-node-MsgUpdateValidatorResponse) | UpdateValidator defines a method for updating a validator. |
| RemoveValidator | [MsgRemoveValidator](#iritamod-node-MsgRemoveValidator) | [MsgRemoveValidatorResponse](#iritamod-node-MsgRemoveValidatorResponse) | RemoveValidator defines a method for removing a validator. |
| GrantNode | [MsgGrantNode](#iritamod-node-MsgGrantNode) | [MsgGrantNodeResponse](#iritamod-node-MsgGrantNodeResponse) | GrantNode defines a method for granting a node access. |
| RevokeNode | [MsgRevokeNode](#iritamod-node-MsgRevokeNode) | [MsgRevokeNodeResponse](#iritamod-node-MsgRevokeNodeResponse) | RevokeNode defines a method for revoking access from a node. |

 



<a name="opb_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## opb/genesis.proto



<a name="iritamod-opb-GenesisState"></a>

### GenesisState
GenesisState defines the OPB module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#iritamod-opb-Params) |  |  |





 

 

 

 



<a name="opb_opb-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## opb/opb.proto



<a name="iritamod-opb-Params"></a>

### Params
Params defines the parameters for the OPB module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_token_denom | [string](#string) |  |  |
| point_token_denom | [string](#string) |  |  |
| base_token_manager | [string](#string) |  |  |
| unrestricted_token_transfer | [bool](#bool) |  |  |





 

 

 

 



<a name="opb_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## opb/query.proto



<a name="iritamod-opb-QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method






<a name="iritamod-opb-QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#iritamod-opb-Params) |  |  |





 

 

 


<a name="iritamod-opb-Query"></a>

### Query
Query defines the gRPC querier service for the OPB module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Params | [QueryParamsRequest](#iritamod-opb-QueryParamsRequest) | [QueryParamsResponse](#iritamod-opb-QueryParamsResponse) | Params queries the parameters of the OPB module |

 



<a name="opb_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## opb/tx.proto



<a name="iritamod-opb-MsgMint"></a>

### MsgMint
MsgMint defines a message to mint the base native token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |
| recipient | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-opb-MsgMintResponse"></a>

### MsgMintResponse
MsgMintResponse defines the Msg/Mint response type.






<a name="iritamod-opb-MsgReclaim"></a>

### MsgReclaim
MsgReclaim defines a message to reclaim the specified native token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-opb-MsgReclaimResponse"></a>

### MsgReclaimResponse
MsgReclaimResponse defines the Msg/Reclaim response type.





 

 

 


<a name="iritamod-opb-Msg"></a>

### Msg
Msg defines the OPB Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Mint | [MsgMint](#iritamod-opb-MsgMint) | [MsgMintResponse](#iritamod-opb-MsgMintResponse) | Mint defines a method for minting the base native token. |
| Reclaim | [MsgReclaim](#iritamod-opb-MsgReclaim) | [MsgReclaimResponse](#iritamod-opb-MsgReclaimResponse) | Reclaim defines a method for reclaiming the specified native token from the corresponding escrow. |

 



<a name="params_params-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## params/params.proto



<a name="iritamod-params-ParamChange"></a>

### ParamChange
ParamChange defines a parameter change.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subspace | [string](#string) |  |  |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="params_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## params/tx.proto



<a name="iritamod-params-MsgUpdateParams"></a>

### MsgUpdateParams
MsgUpdateParams defines an SDK message for updating params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| changes | [ParamChange](#iritamod-params-ParamChange) | repeated |  |
| operator | [string](#string) |  |  |






<a name="iritamod-params-MsgUpdateParamsResponse"></a>

### MsgUpdateParamsResponse
MsgUpdateParamsResponse defines the Msg/UpdateParams response type.





 

 

 


<a name="iritamod-params-Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateParams | [MsgUpdateParams](#iritamod-params-MsgUpdateParams) | [MsgUpdateParamsResponse](#iritamod-params-MsgUpdateParamsResponse) | UpdateParams defines a method for update a set of system params. |

 



<a name="perm_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## perm/genesis.proto



<a name="iritamod-perm-GenesisState"></a>

### GenesisState
GenesisState defines the perm module&#39;s genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_accounts | [RoleAccount](#iritamod-perm-RoleAccount) | repeated |  |
| black_list | [string](#string) | repeated |  |
| contract_deny_list | [string](#string) | repeated |  |






<a name="iritamod-perm-RoleAccount"></a>

### RoleAccount
RoleAccount represents an account with roles.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| roles | [Role](#iritamod-perm-Role) | repeated |  |





 

 

 

 



<a name="perm_perm-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## perm/perm.proto


 


<a name="iritamod-perm-Role"></a>

### Role
Role represents a role

| Name | Number | Description |
| ---- | ------ | ----------- |
| ROOT_ADMIN | 0 | ROOT_ADMIN defines the root admin role index. |
| PERM_ADMIN | 1 | PERM_ADMIN defines the permission admin role index. |
| BLACKLIST_ADMIN | 2 | BLACKLIST_ADMIN defines the blacklist admin role index. |
| NODE_ADMIN | 3 | NODE_ADMIN defines the node admin role index. |
| PARAM_ADMIN | 4 | PARAM_ADMIN defines the param admin role index. |
| POWER_USER | 5 | POWER_USER defines the power user role index. |
| RELAYER_USER | 6 | RELAYER_USER defines the relayer user role index. |
| ID_ADMIN | 7 | ID_ADMIN defines the identity admin role index. |
| BASE_M1_ADMIN | 8 | BASE_M1_ADMIN defines the base M1 admin role index. |
| PLATFORM_USER | 9 | Chain_Account_Role defines the platform admin role index. |
| POWER_USER_ADMIN | 10 | POWER_USER_ADMIN defines the power admin role index. |


 

 

 



<a name="perm_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## perm/query.proto



<a name="iritamod-perm-QueryBlockListRequest"></a>

### QueryBlockListRequest
QueryBlacklistRequest is request type for the Query/Blacklist RPC method






<a name="iritamod-perm-QueryBlockListResponse"></a>

### QueryBlockListResponse
QueryBlacklistResponse is response type for the Query/Blacklist RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="iritamod-perm-QueryContractDenyList"></a>

### QueryContractDenyList
QueryBlacklistRequest is request type for the Query/Blacklist RPC method






<a name="iritamod-perm-QueryContractDenyListResponse"></a>

### QueryContractDenyListResponse
QueryBlacklistResponse is response type for the Query/Blacklist RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="iritamod-perm-QueryRolesRequest"></a>

### QueryRolesRequest
QueryRolesRequest is request type for the Query/Roles RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="iritamod-perm-QueryRolesResponse"></a>

### QueryRolesResponse
QueryRolesResponse is response type for the Query/Roles RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [Role](#iritamod-perm-Role) | repeated |  |





 

 

 


<a name="iritamod-perm-Query"></a>

### Query
Query defines the gRPC querier service for perm module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Roles | [QueryRolesRequest](#iritamod-perm-QueryRolesRequest) | [QueryRolesResponse](#iritamod-perm-QueryRolesResponse) | Roles queries the roles of a given address |
| AccountBlockList | [QueryBlockListRequest](#iritamod-perm-QueryBlockListRequest) | [QueryBlockListResponse](#iritamod-perm-QueryBlockListResponse) | Blacklist queries the black list |
| ContractDenyList | [QueryContractDenyList](#iritamod-perm-QueryContractDenyList) | [QueryContractDenyListResponse](#iritamod-perm-QueryContractDenyListResponse) | ContractDenyList queries the contract deny list |

 



<a name="perm_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## perm/tx.proto



<a name="iritamod-perm-MsgAssignRoles"></a>

### MsgAssignRoles
MsgAssignRoles defines an SDK message for assigning roles to an address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| roles | [Role](#iritamod-perm-Role) | repeated |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgAssignRolesResponse"></a>

### MsgAssignRolesResponse
MsgAssignRolesResponse defines the Msg/AssignRoles response type.






<a name="iritamod-perm-MsgBlockAccount"></a>

### MsgBlockAccount
MsgBlockAccount defines an SDK message for blocking an account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgBlockAccountResponse"></a>

### MsgBlockAccountResponse
MsgBlockAccountResponse defines the Msg/BlockAccount response type.






<a name="iritamod-perm-MsgBlockContract"></a>

### MsgBlockContract
MsgBlockContract defines an SDK message for blocking an contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgBlockContractResponse"></a>

### MsgBlockContractResponse
MsgBlockContractResponse defines the Msg/MsgBlockContract response type.






<a name="iritamod-perm-MsgUnassignRoles"></a>

### MsgUnassignRoles
MsgUnassignRoles defines an SDK message for unassigning roles from an address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| roles | [Role](#iritamod-perm-Role) | repeated |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgUnassignRolesResponse"></a>

### MsgUnassignRolesResponse
MsgUnassignRolesResponse defines the Msg/UnassignRoles response type.






<a name="iritamod-perm-MsgUnblockAccount"></a>

### MsgUnblockAccount
MsgUnblockAccount defines an SDK message for unblocking an account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgUnblockAccountResponse"></a>

### MsgUnblockAccountResponse
MsgUnblockAccountResponse defines the Msg/UnblockAccount response type.






<a name="iritamod-perm-MsgUnblockContract"></a>

### MsgUnblockContract
MsgUnblockContract defines an SDK message for unblocking an contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-perm-MsgUnblockContractResponse"></a>

### MsgUnblockContractResponse
MsgUnblockAccountResponse defines the Msg/MsgUnblockContract response type.





 

 

 


<a name="iritamod-perm-Msg"></a>

### Msg
Msg defines the perm Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AssignRoles | [MsgAssignRoles](#iritamod-perm-MsgAssignRoles) | [MsgAssignRolesResponse](#iritamod-perm-MsgAssignRolesResponse) | AssignRoles defines a method for assigning roles for the operator. |
| UnassignRoles | [MsgUnassignRoles](#iritamod-perm-MsgUnassignRoles) | [MsgUnassignRolesResponse](#iritamod-perm-MsgUnassignRolesResponse) | UnassignRoles defines a method for unassigning roles from the operator. |
| BlockAccount | [MsgBlockAccount](#iritamod-perm-MsgBlockAccount) | [MsgBlockAccountResponse](#iritamod-perm-MsgBlockAccountResponse) | BlockAccount defines a method for blocking an account |
| UnblockAccount | [MsgUnblockAccount](#iritamod-perm-MsgUnblockAccount) | [MsgUnblockAccountResponse](#iritamod-perm-MsgUnblockAccountResponse) | UnblockAccount defines a method for unblocking a blocked account |
| BlockContract | [MsgBlockContract](#iritamod-perm-MsgBlockContract) | [MsgBlockContractResponse](#iritamod-perm-MsgBlockContractResponse) | BlockContract defines a method for blocking an contract |
| UnblockContract | [MsgUnblockContract](#iritamod-perm-MsgUnblockContract) | [MsgUnblockContractResponse](#iritamod-perm-MsgUnblockContractResponse) | UnblockContract defines a method for unblocking a blocked contract |

 



<a name="slashing_slashing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## slashing/slashing.proto



<a name="iritamod-slashing-Params"></a>

### Params
Params represents the parameters used for by the slashing module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signed_blocks_window | [int64](#int64) |  |  |
| min_signed_per_window | [bytes](#bytes) |  |  |
| downtime_jail_duration | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| slash_fraction_double_sign | [bytes](#bytes) |  |  |
| slash_fraction_downtime | [bytes](#bytes) |  |  |






<a name="iritamod-slashing-ValidatorSigningInfo"></a>

### ValidatorSigningInfo
ValidatorSigningInfo defines a validator&#39;s signing info for monitoring their liveness activity.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| start_height | [int64](#int64) |  | height at which validator was first a candidate OR was unjailed |
| index_offset | [int64](#int64) |  | index offset into signed block bit array |
| jailed_until | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | timestamp validator cannot be unjailed until |
| tombstoned | [bool](#bool) |  | whether or not a validator has been tombstoned (killed out of validator set) |
| missed_blocks_counter | [int64](#int64) |  | missed blocks counter (to avoid scanning the array every time) |





 

 

 

 



<a name="slashing_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## slashing/tx.proto



<a name="iritamod-slashing-MsgUnjailValidator"></a>

### MsgUnjailValidator
MsgUnjailValidator - struct for unjailing jailed validator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| operator | [string](#string) |  |  |






<a name="iritamod-slashing-MsgUnjailValidatorResponse"></a>

### MsgUnjailValidatorResponse
MsgUnjailValidatorResponse defines the Msg/Unjail response type.





 

 

 


<a name="iritamod-slashing-Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UnjailValidator | [MsgUnjailValidator](#iritamod-slashing-MsgUnjailValidator) | [MsgUnjailValidatorResponse](#iritamod-slashing-MsgUnjailValidatorResponse) | Unjail defines a method for unjail a validator. |

 



<a name="upgrade_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## upgrade/tx.proto



<a name="iritamod-upgrade-MsgCancelUpgrade"></a>

### MsgCancelUpgrade
MsgCancelUpgrade - struct for cancel software upgrade


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operator | [string](#string) |  |  |






<a name="iritamod-upgrade-MsgCancelUpgradeResponse"></a>

### MsgCancelUpgradeResponse
MsgCancelUpgradeResponse defines the Msg/CancelUpgrade response type.






<a name="iritamod-upgrade-MsgUpgradeSoftware"></a>

### MsgUpgradeSoftware
MsgUpgradeSoftware - struct for upgrade software


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| height | [int64](#int64) |  | The height at which the upgrade must be performed. Only used if Time is not set. |
| info | [string](#string) |  | Any application specific upgrade info to be included on-chain such as a git commit that validators could automatically upgrade to |
| operator | [string](#string) |  |  |






<a name="iritamod-upgrade-MsgUpgradeSoftwareResponse"></a>

### MsgUpgradeSoftwareResponse
MsgUpgradeSoftwareResponse defines the Msg/UpgradeSoftware response type.





 

 

 


<a name="iritamod-upgrade-Msg"></a>

### Msg
Msg defines the ibc Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpgradeSoftware | [MsgUpgradeSoftware](#iritamod-upgrade-MsgUpgradeSoftware) | [MsgUpgradeSoftwareResponse](#iritamod-upgrade-MsgUpgradeSoftwareResponse) | CreateClient defines a method for creating a light client. |
| CancelUpgrade | [MsgCancelUpgrade](#iritamod-upgrade-MsgCancelUpgrade) | [MsgCancelUpgradeResponse](#iritamod-upgrade-MsgCancelUpgradeResponse) | CancelUpgrade defines a method for updating a light client state. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

