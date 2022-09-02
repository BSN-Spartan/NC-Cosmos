---
order: 3
---
# IRISnet Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [mt/genesis.proto](#mt_genesis-proto)
    - [Collection](#irismod-mt-Collection)
    - [DenomBalance](#irismod-mt-DenomBalance)
    - [GenesisState](#irismod-mt-GenesisState)
    - [Owner](#irismod-mt-Owner)
  
- [mt/mt.proto](#mt_mt-proto)
    - [Balance](#irismod-mt-Balance)
    - [Denom](#irismod-mt-Denom)
    - [MT](#irismod-mt-MT)
  
- [mt/query.proto](#mt_query-proto)
    - [QueryBalancesRequest](#irismod-mt-QueryBalancesRequest)
    - [QueryBalancesResponse](#irismod-mt-QueryBalancesResponse)
    - [QueryDenomRequest](#irismod-mt-QueryDenomRequest)
    - [QueryDenomResponse](#irismod-mt-QueryDenomResponse)
    - [QueryDenomsRequest](#irismod-mt-QueryDenomsRequest)
    - [QueryDenomsResponse](#irismod-mt-QueryDenomsResponse)
    - [QueryMTRequest](#irismod-mt-QueryMTRequest)
    - [QueryMTResponse](#irismod-mt-QueryMTResponse)
    - [QueryMTSupplyRequest](#irismod-mt-QueryMTSupplyRequest)
    - [QueryMTSupplyResponse](#irismod-mt-QueryMTSupplyResponse)
    - [QueryMTsRequest](#irismod-mt-QueryMTsRequest)
    - [QueryMTsResponse](#irismod-mt-QueryMTsResponse)
    - [QuerySupplyRequest](#irismod-mt-QuerySupplyRequest)
    - [QuerySupplyResponse](#irismod-mt-QuerySupplyResponse)
  
    - [Query](#irismod-mt-Query)
  
- [mt/tx.proto](#mt_tx-proto)
    - [MsgBurnMT](#irismod-mt-MsgBurnMT)
    - [MsgBurnMTResponse](#irismod-mt-MsgBurnMTResponse)
    - [MsgEditMT](#irismod-mt-MsgEditMT)
    - [MsgEditMTResponse](#irismod-mt-MsgEditMTResponse)
    - [MsgIssueDenom](#irismod-mt-MsgIssueDenom)
    - [MsgIssueDenomResponse](#irismod-mt-MsgIssueDenomResponse)
    - [MsgMintMT](#irismod-mt-MsgMintMT)
    - [MsgMintMTResponse](#irismod-mt-MsgMintMTResponse)
    - [MsgTransferDenom](#irismod-mt-MsgTransferDenom)
    - [MsgTransferDenomResponse](#irismod-mt-MsgTransferDenomResponse)
    - [MsgTransferMT](#irismod-mt-MsgTransferMT)
    - [MsgTransferMTResponse](#irismod-mt-MsgTransferMTResponse)
  
    - [Msg](#irismod-mt-Msg)
  
- [nft/genesis.proto](#nft_genesis-proto)
    - [GenesisState](#irismod-nft-GenesisState)
  
- [nft/nft.proto](#nft_nft-proto)
    - [BaseNFT](#irismod-nft-BaseNFT)
    - [Collection](#irismod-nft-Collection)
    - [Denom](#irismod-nft-Denom)
    - [IDCollection](#irismod-nft-IDCollection)
    - [Owner](#irismod-nft-Owner)
  
- [nft/query.proto](#nft_query-proto)
    - [QueryCollectionRequest](#irismod-nft-QueryCollectionRequest)
    - [QueryCollectionResponse](#irismod-nft-QueryCollectionResponse)
    - [QueryDenomRequest](#irismod-nft-QueryDenomRequest)
    - [QueryDenomResponse](#irismod-nft-QueryDenomResponse)
    - [QueryDenomsRequest](#irismod-nft-QueryDenomsRequest)
    - [QueryDenomsResponse](#irismod-nft-QueryDenomsResponse)
    - [QueryNFTRequest](#irismod-nft-QueryNFTRequest)
    - [QueryNFTResponse](#irismod-nft-QueryNFTResponse)
    - [QueryOwnerRequest](#irismod-nft-QueryOwnerRequest)
    - [QueryOwnerResponse](#irismod-nft-QueryOwnerResponse)
    - [QuerySupplyRequest](#irismod-nft-QuerySupplyRequest)
    - [QuerySupplyResponse](#irismod-nft-QuerySupplyResponse)
  
    - [Query](#irismod-nft-Query)
  
- [nft/tx.proto](#nft_tx-proto)
    - [MsgBurnNFT](#irismod-nft-MsgBurnNFT)
    - [MsgBurnNFTResponse](#irismod-nft-MsgBurnNFTResponse)
    - [MsgEditNFT](#irismod-nft-MsgEditNFT)
    - [MsgEditNFTResponse](#irismod-nft-MsgEditNFTResponse)
    - [MsgIssueDenom](#irismod-nft-MsgIssueDenom)
    - [MsgIssueDenomResponse](#irismod-nft-MsgIssueDenomResponse)
    - [MsgMintNFT](#irismod-nft-MsgMintNFT)
    - [MsgMintNFTResponse](#irismod-nft-MsgMintNFTResponse)
    - [MsgTransferDenom](#irismod-nft-MsgTransferDenom)
    - [MsgTransferDenomResponse](#irismod-nft-MsgTransferDenomResponse)
    - [MsgTransferNFT](#irismod-nft-MsgTransferNFT)
    - [MsgTransferNFTResponse](#irismod-nft-MsgTransferNFTResponse)
  
    - [Msg](#irismod-nft-Msg)
  
- [service/genesis.proto](#service_genesis-proto)
    - [GenesisState](#irismod-service-GenesisState)
    - [GenesisState.RequestContextsEntry](#irismod-service-GenesisState-RequestContextsEntry)
    - [GenesisState.WithdrawAddressesEntry](#irismod-service-GenesisState-WithdrawAddressesEntry)
  
- [service/query.proto](#service_query-proto)
    - [QueryBindingRequest](#irismod-service-QueryBindingRequest)
    - [QueryBindingResponse](#irismod-service-QueryBindingResponse)
    - [QueryBindingsRequest](#irismod-service-QueryBindingsRequest)
    - [QueryBindingsResponse](#irismod-service-QueryBindingsResponse)
    - [QueryDefinitionRequest](#irismod-service-QueryDefinitionRequest)
    - [QueryDefinitionResponse](#irismod-service-QueryDefinitionResponse)
    - [QueryEarnedFeesRequest](#irismod-service-QueryEarnedFeesRequest)
    - [QueryEarnedFeesResponse](#irismod-service-QueryEarnedFeesResponse)
    - [QueryParamsRequest](#irismod-service-QueryParamsRequest)
    - [QueryParamsResponse](#irismod-service-QueryParamsResponse)
    - [QueryRequestContextRequest](#irismod-service-QueryRequestContextRequest)
    - [QueryRequestContextResponse](#irismod-service-QueryRequestContextResponse)
    - [QueryRequestRequest](#irismod-service-QueryRequestRequest)
    - [QueryRequestResponse](#irismod-service-QueryRequestResponse)
    - [QueryRequestsByReqCtxRequest](#irismod-service-QueryRequestsByReqCtxRequest)
    - [QueryRequestsByReqCtxResponse](#irismod-service-QueryRequestsByReqCtxResponse)
    - [QueryRequestsRequest](#irismod-service-QueryRequestsRequest)
    - [QueryRequestsResponse](#irismod-service-QueryRequestsResponse)
    - [QueryResponseRequest](#irismod-service-QueryResponseRequest)
    - [QueryResponseResponse](#irismod-service-QueryResponseResponse)
    - [QueryResponsesRequest](#irismod-service-QueryResponsesRequest)
    - [QueryResponsesResponse](#irismod-service-QueryResponsesResponse)
    - [QuerySchemaRequest](#irismod-service-QuerySchemaRequest)
    - [QuerySchemaResponse](#irismod-service-QuerySchemaResponse)
    - [QueryWithdrawAddressRequest](#irismod-service-QueryWithdrawAddressRequest)
    - [QueryWithdrawAddressResponse](#irismod-service-QueryWithdrawAddressResponse)
  
    - [Query](#irismod-service-Query)
  
- [service/service.proto](#service_service-proto)
    - [CompactRequest](#irismod-service-CompactRequest)
    - [Params](#irismod-service-Params)
    - [Pricing](#irismod-service-Pricing)
    - [PromotionByTime](#irismod-service-PromotionByTime)
    - [PromotionByVolume](#irismod-service-PromotionByVolume)
    - [Request](#irismod-service-Request)
    - [RequestContext](#irismod-service-RequestContext)
    - [Response](#irismod-service-Response)
    - [ServiceBinding](#irismod-service-ServiceBinding)
    - [ServiceDefinition](#irismod-service-ServiceDefinition)
  
    - [RequestContextBatchState](#irismod-service-RequestContextBatchState)
    - [RequestContextState](#irismod-service-RequestContextState)
  
- [service/tx.proto](#service_tx-proto)
    - [MsgBindService](#irismod-service-MsgBindService)
    - [MsgBindServiceResponse](#irismod-service-MsgBindServiceResponse)
    - [MsgCallService](#irismod-service-MsgCallService)
    - [MsgCallServiceResponse](#irismod-service-MsgCallServiceResponse)
    - [MsgDefineService](#irismod-service-MsgDefineService)
    - [MsgDefineServiceResponse](#irismod-service-MsgDefineServiceResponse)
    - [MsgDisableServiceBinding](#irismod-service-MsgDisableServiceBinding)
    - [MsgDisableServiceBindingResponse](#irismod-service-MsgDisableServiceBindingResponse)
    - [MsgEnableServiceBinding](#irismod-service-MsgEnableServiceBinding)
    - [MsgEnableServiceBindingResponse](#irismod-service-MsgEnableServiceBindingResponse)
    - [MsgKillRequestContext](#irismod-service-MsgKillRequestContext)
    - [MsgKillRequestContextResponse](#irismod-service-MsgKillRequestContextResponse)
    - [MsgPauseRequestContext](#irismod-service-MsgPauseRequestContext)
    - [MsgPauseRequestContextResponse](#irismod-service-MsgPauseRequestContextResponse)
    - [MsgRefundServiceDeposit](#irismod-service-MsgRefundServiceDeposit)
    - [MsgRefundServiceDepositResponse](#irismod-service-MsgRefundServiceDepositResponse)
    - [MsgRespondService](#irismod-service-MsgRespondService)
    - [MsgRespondServiceResponse](#irismod-service-MsgRespondServiceResponse)
    - [MsgSetWithdrawAddress](#irismod-service-MsgSetWithdrawAddress)
    - [MsgSetWithdrawAddressResponse](#irismod-service-MsgSetWithdrawAddressResponse)
    - [MsgStartRequestContext](#irismod-service-MsgStartRequestContext)
    - [MsgStartRequestContextResponse](#irismod-service-MsgStartRequestContextResponse)
    - [MsgUpdateRequestContext](#irismod-service-MsgUpdateRequestContext)
    - [MsgUpdateRequestContextResponse](#irismod-service-MsgUpdateRequestContextResponse)
    - [MsgUpdateServiceBinding](#irismod-service-MsgUpdateServiceBinding)
    - [MsgUpdateServiceBindingResponse](#irismod-service-MsgUpdateServiceBindingResponse)
    - [MsgWithdrawEarnedFees](#irismod-service-MsgWithdrawEarnedFees)
    - [MsgWithdrawEarnedFeesResponse](#irismod-service-MsgWithdrawEarnedFeesResponse)
  
    - [Msg](#irismod-service-Msg)
  
- [token/genesis.proto](#token_genesis-proto)
    - [GenesisState](#irismod-token-GenesisState)
  
- [token/query.proto](#token_query-proto)
    - [QueryFeesRequest](#irismod-token-QueryFeesRequest)
    - [QueryFeesResponse](#irismod-token-QueryFeesResponse)
    - [QueryParamsRequest](#irismod-token-QueryParamsRequest)
    - [QueryParamsResponse](#irismod-token-QueryParamsResponse)
    - [QueryTokenRequest](#irismod-token-QueryTokenRequest)
    - [QueryTokenResponse](#irismod-token-QueryTokenResponse)
    - [QueryTokensRequest](#irismod-token-QueryTokensRequest)
    - [QueryTokensResponse](#irismod-token-QueryTokensResponse)
    - [QueryTotalBurnRequest](#irismod-token-QueryTotalBurnRequest)
    - [QueryTotalBurnResponse](#irismod-token-QueryTotalBurnResponse)
  
    - [Query](#irismod-token-Query)
  
- [token/token.proto](#token_token-proto)
    - [Params](#irismod-token-Params)
    - [Token](#irismod-token-Token)
  
- [token/tx.proto](#token_tx-proto)
    - [MsgBurnToken](#irismod-token-MsgBurnToken)
    - [MsgBurnTokenResponse](#irismod-token-MsgBurnTokenResponse)
    - [MsgEditToken](#irismod-token-MsgEditToken)
    - [MsgEditTokenResponse](#irismod-token-MsgEditTokenResponse)
    - [MsgIssueToken](#irismod-token-MsgIssueToken)
    - [MsgIssueTokenResponse](#irismod-token-MsgIssueTokenResponse)
    - [MsgMintToken](#irismod-token-MsgMintToken)
    - [MsgMintTokenResponse](#irismod-token-MsgMintTokenResponse)
    - [MsgTransferTokenOwner](#irismod-token-MsgTransferTokenOwner)
    - [MsgTransferTokenOwnerResponse](#irismod-token-MsgTransferTokenOwnerResponse)
  
    - [Msg](#irismod-token-Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="mt_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mt/genesis.proto



<a name="irismod-mt-Collection"></a>

### Collection
Collection defines a type of collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [Denom](#irismod-mt-Denom) |  |  |
| mts | [MT](#irismod-mt-MT) | repeated |  |






<a name="irismod-mt-DenomBalance"></a>

### DenomBalance
DenomBalance defines a type of denom balances owned by an account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| balances | [Balance](#irismod-mt-Balance) | repeated |  |






<a name="irismod-mt-GenesisState"></a>

### GenesisState
GenesisState defines the MT module&#39;s genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [Collection](#irismod-mt-Collection) | repeated |  |
| owners | [Owner](#irismod-mt-Owner) | repeated |  |






<a name="irismod-mt-Owner"></a>

### Owner
Owner defines a type of account balances


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| denoms | [DenomBalance](#irismod-mt-DenomBalance) | repeated |  |





 

 

 

 



<a name="mt_mt-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mt/mt.proto



<a name="irismod-mt-Balance"></a>

### Balance
Balance defines multi token balance for owners


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mt_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |






<a name="irismod-mt-Denom"></a>

### Denom
Denom defines a class of MTs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-mt-MT"></a>

### MT
MT defines a multi token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| supply | [uint64](#uint64) |  |  |
| data | [bytes](#bytes) |  |  |





 

 

 

 



<a name="mt_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mt/query.proto



<a name="irismod-mt-QueryBalancesRequest"></a>

### QueryBalancesRequest
QueryBalancesRequest is the request type for the Query/Balances RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-mt-QueryBalancesResponse"></a>

### QueryBalancesResponse
QueryBalancesResponse is the response type for the Query/Balances RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [Balance](#irismod-mt-Balance) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-mt-QueryDenomRequest"></a>

### QueryDenomRequest
QueryDenomRequest is the request type for the Query/Denom RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |






<a name="irismod-mt-QueryDenomResponse"></a>

### QueryDenomResponse
QueryDenomResponse is the response type for the Query/Denom RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [Denom](#irismod-mt-Denom) |  |  |






<a name="irismod-mt-QueryDenomsRequest"></a>

### QueryDenomsRequest
QueryDenomsRequest is the request type for the Query/Denoms RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-mt-QueryDenomsResponse"></a>

### QueryDenomsResponse
QueryDenomsResponse is the response type for the Query/Denoms RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denoms | [Denom](#irismod-mt-Denom) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-mt-QueryMTRequest"></a>

### QueryMTRequest
QueryMTRequest is the request type for the Query/MT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| mt_id | [string](#string) |  |  |






<a name="irismod-mt-QueryMTResponse"></a>

### QueryMTResponse
QueryMTResponse is the response type for the Query/MT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mt | [MT](#irismod-mt-MT) |  |  |






<a name="irismod-mt-QueryMTSupplyRequest"></a>

### QueryMTSupplyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| mt_id | [string](#string) |  |  |






<a name="irismod-mt-QueryMTSupplyResponse"></a>

### QueryMTSupplyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |






<a name="irismod-mt-QueryMTsRequest"></a>

### QueryMTsRequest
QueryMTsRequest is the request type for the Query/MTs RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-mt-QueryMTsResponse"></a>

### QueryMTsResponse
QueryMTsResponse is the response type for the Query/MTs RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mts | [MT](#irismod-mt-MT) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-mt-QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the request type for the Query RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-mt-QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the response type for the Query/Supply RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |





 

 

 


<a name="irismod-mt-Query"></a>

### Query
Query defines the gRPC querier service for MT module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Supply | [QuerySupplyRequest](#irismod-mt-QuerySupplyRequest) | [QuerySupplyResponse](#irismod-mt-QuerySupplyResponse) | Supply queries the total supply of a given denom or owner |
| Denoms | [QueryDenomsRequest](#irismod-mt-QueryDenomsRequest) | [QueryDenomsResponse](#irismod-mt-QueryDenomsResponse) | Denoms queries all the denoms |
| Denom | [QueryDenomRequest](#irismod-mt-QueryDenomRequest) | [QueryDenomResponse](#irismod-mt-QueryDenomResponse) | Denom queries the definition of a given denom ID |
| MTSupply | [QueryMTSupplyRequest](#irismod-mt-QueryMTSupplyRequest) | [QueryMTSupplyResponse](#irismod-mt-QueryMTSupplyResponse) | MTSupply queries the total supply of given denom and mt ID |
| MTs | [QueryMTsRequest](#irismod-mt-QueryMTsRequest) | [QueryMTsResponse](#irismod-mt-QueryMTsResponse) | MTs queries all the MTs of a given denom ID |
| MT | [QueryMTRequest](#irismod-mt-QueryMTRequest) | [QueryMTResponse](#irismod-mt-QueryMTResponse) | MT queries the MT of the given denom and mt ID |
| Balances | [QueryBalancesRequest](#irismod-mt-QueryBalancesRequest) | [QueryBalancesResponse](#irismod-mt-QueryBalancesResponse) | Balances queries the MT balances of a specified owner |

 



<a name="mt_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mt/tx.proto



<a name="irismod-mt-MsgBurnMT"></a>

### MsgBurnMT
MsgBurnMT defines an SDK message for burning an MT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |
| sender | [string](#string) |  |  |






<a name="irismod-mt-MsgBurnMTResponse"></a>

### MsgBurnMTResponse
MsgBurnMTResponse defines the Msg/BurnMT response type.






<a name="irismod-mt-MsgEditMT"></a>

### MsgEditMT
MsgEditMT defines an SDK message for editing an MT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| sender | [string](#string) |  |  |






<a name="irismod-mt-MsgEditMTResponse"></a>

### MsgEditMTResponse
MsgEditMTResponse defines the Msg/EditMT response type.






<a name="irismod-mt-MsgIssueDenom"></a>

### MsgIssueDenom
MsgIssueDenom defines an SDK message for creating a new denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |
| sender | [string](#string) |  |  |






<a name="irismod-mt-MsgIssueDenomResponse"></a>

### MsgIssueDenomResponse
MsgIssueDenomResponse defines the Msg/IssueDenom response type.






<a name="irismod-mt-MsgMintMT"></a>

### MsgMintMT
MsgMintMT defines an SDK message for creating a new MT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |
| data | [bytes](#bytes) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="irismod-mt-MsgMintMTResponse"></a>

### MsgMintMTResponse
MsgMintMTResponse defines the Msg/MintMT response type.






<a name="irismod-mt-MsgTransferDenom"></a>

### MsgTransferDenom
MsgTransferDenom defines an SDK message for transferring an denom to recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="irismod-mt-MsgTransferDenomResponse"></a>

### MsgTransferDenomResponse
MsgTransferDenomResponse defines the Msg/TransferDenom response type.






<a name="irismod-mt-MsgTransferMT"></a>

### MsgTransferMT
MsgTransferMT defines an SDK message for transferring an MT to recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="irismod-mt-MsgTransferMTResponse"></a>

### MsgTransferMTResponse
MsgTransferMTResponse defines the Msg/TransferMT response type.





 

 

 


<a name="irismod-mt-Msg"></a>

### Msg
Msg defines the mt Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IssueDenom | [MsgIssueDenom](#irismod-mt-MsgIssueDenom) | [MsgIssueDenomResponse](#irismod-mt-MsgIssueDenomResponse) | IssueDenom defines a method for issuing a denom. |
| TransferDenom | [MsgTransferDenom](#irismod-mt-MsgTransferDenom) | [MsgTransferDenomResponse](#irismod-mt-MsgTransferDenomResponse) | TransferDenom defines a method for transferring a denom. |
| MintMT | [MsgMintMT](#irismod-mt-MsgMintMT) | [MsgMintMTResponse](#irismod-mt-MsgMintMTResponse) | MintMT defines a method for creating a new MT or minting amounts of an existing MT |
| EditMT | [MsgEditMT](#irismod-mt-MsgEditMT) | [MsgEditMTResponse](#irismod-mt-MsgEditMTResponse) | EditMT defines a method for editing an MT. |
| TransferMT | [MsgTransferMT](#irismod-mt-MsgTransferMT) | [MsgTransferMTResponse](#irismod-mt-MsgTransferMTResponse) | TransferMT defines a method for transferring an MT. |
| BurnMT | [MsgBurnMT](#irismod-mt-MsgBurnMT) | [MsgBurnMTResponse](#irismod-mt-MsgBurnMTResponse) | BurnMT defines a method for burning an MT. |

 



<a name="nft_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/genesis.proto



<a name="irismod-nft-GenesisState"></a>

### GenesisState
GenesisState defines the NFT module&#39;s genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [Collection](#irismod-nft-Collection) | repeated |  |





 

 

 

 



<a name="nft_nft-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/nft.proto



<a name="irismod-nft-BaseNFT"></a>

### BaseNFT
BaseNFT defines a non-fungible token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| data | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |






<a name="irismod-nft-Collection"></a>

### Collection
Collection defines a type of collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [Denom](#irismod-nft-Denom) |  |  |
| nfts | [BaseNFT](#irismod-nft-BaseNFT) | repeated |  |






<a name="irismod-nft-Denom"></a>

### Denom
Denom defines a type of NFT


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| schema | [string](#string) |  |  |
| creator | [string](#string) |  |  |
| symbol | [string](#string) |  |  |
| mint_restricted | [bool](#bool) |  |  |
| update_restricted | [bool](#bool) |  |  |
| description | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="irismod-nft-IDCollection"></a>

### IDCollection
IDCollection defines a type of collection with specified ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| token_ids | [string](#string) | repeated |  |






<a name="irismod-nft-Owner"></a>

### Owner
Owner defines a type of owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| id_collections | [IDCollection](#irismod-nft-IDCollection) | repeated |  |





 

 

 

 



<a name="nft_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/query.proto



<a name="irismod-nft-QueryCollectionRequest"></a>

### QueryCollectionRequest
QueryCollectionRequest is the request type for the Query/Collection RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-nft-QueryCollectionResponse"></a>

### QueryCollectionResponse
QueryCollectionResponse is the response type for the Query/Collection RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#irismod-nft-Collection) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-nft-QueryDenomRequest"></a>

### QueryDenomRequest
QueryDenomRequest is the request type for the Query/Denom RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |






<a name="irismod-nft-QueryDenomResponse"></a>

### QueryDenomResponse
QueryDenomResponse is the response type for the Query/Denom RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [Denom](#irismod-nft-Denom) |  |  |






<a name="irismod-nft-QueryDenomsRequest"></a>

### QueryDenomsRequest
QueryDenomsRequest is the request type for the Query/Denoms RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-nft-QueryDenomsResponse"></a>

### QueryDenomsResponse
QueryDenomsResponse is the response type for the Query/Denoms RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denoms | [Denom](#irismod-nft-Denom) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-nft-QueryNFTRequest"></a>

### QueryNFTRequest
QueryNFTRequest is the request type for the Query/NFT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| token_id | [string](#string) |  |  |






<a name="irismod-nft-QueryNFTResponse"></a>

### QueryNFTResponse
QueryNFTResponse is the response type for the Query/NFT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nft | [BaseNFT](#irismod-nft-BaseNFT) |  |  |






<a name="irismod-nft-QueryOwnerRequest"></a>

### QueryOwnerRequest
QueryOwnerRequest is the request type for the Query/Owner RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-nft-QueryOwnerResponse"></a>

### QueryOwnerResponse
QueryOwnerResponse is the response type for the Query/Owner RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [Owner](#irismod-nft-Owner) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-nft-QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the request type for the Query/HTLC RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom_id | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-nft-QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the response type for the Query/Supply RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [uint64](#uint64) |  |  |





 

 

 


<a name="irismod-nft-Query"></a>

### Query
Query defines the gRPC querier service for NFT module

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Supply | [QuerySupplyRequest](#irismod-nft-QuerySupplyRequest) | [QuerySupplyResponse](#irismod-nft-QuerySupplyResponse) | Supply queries the total supply of a given denom or owner |
| Owner | [QueryOwnerRequest](#irismod-nft-QueryOwnerRequest) | [QueryOwnerResponse](#irismod-nft-QueryOwnerResponse) | Owner queries the NFTs of the specified owner |
| Collection | [QueryCollectionRequest](#irismod-nft-QueryCollectionRequest) | [QueryCollectionResponse](#irismod-nft-QueryCollectionResponse) | Collection queries the NFTs of the specified denom |
| Denom | [QueryDenomRequest](#irismod-nft-QueryDenomRequest) | [QueryDenomResponse](#irismod-nft-QueryDenomResponse) | Denom queries the definition of a given denom |
| Denoms | [QueryDenomsRequest](#irismod-nft-QueryDenomsRequest) | [QueryDenomsResponse](#irismod-nft-QueryDenomsResponse) | Denoms queries all the denoms |
| NFT | [QueryNFTRequest](#irismod-nft-QueryNFTRequest) | [QueryNFTResponse](#irismod-nft-QueryNFTResponse) | NFT queries the NFT for the given denom and token ID |

 



<a name="nft_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## nft/tx.proto



<a name="irismod-nft-MsgBurnNFT"></a>

### MsgBurnNFT
MsgBurnNFT defines an SDK message for burning a NFT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| sender | [string](#string) |  |  |






<a name="irismod-nft-MsgBurnNFTResponse"></a>

### MsgBurnNFTResponse
MsgBurnNFTResponse defines the Msg/BurnNFT response type.






<a name="irismod-nft-MsgEditNFT"></a>

### MsgEditNFT
MsgEditNFT defines an SDK message for editing a nft.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| data | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |






<a name="irismod-nft-MsgEditNFTResponse"></a>

### MsgEditNFTResponse
MsgEditNFTResponse defines the Msg/EditNFT response type.






<a name="irismod-nft-MsgIssueDenom"></a>

### MsgIssueDenom
MsgIssueDenom defines an SDK message for creating a new denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| schema | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| symbol | [string](#string) |  |  |
| mint_restricted | [bool](#bool) |  |  |
| update_restricted | [bool](#bool) |  |  |
| description | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |
| data | [string](#string) |  |  |






<a name="irismod-nft-MsgIssueDenomResponse"></a>

### MsgIssueDenomResponse
MsgIssueDenomResponse defines the Msg/IssueDenom response type.






<a name="irismod-nft-MsgMintNFT"></a>

### MsgMintNFT
MsgMintNFT defines an SDK message for creating a new NFT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| data | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |






<a name="irismod-nft-MsgMintNFTResponse"></a>

### MsgMintNFTResponse
MsgMintNFTResponse defines the Msg/MintNFT response type.






<a name="irismod-nft-MsgTransferDenom"></a>

### MsgTransferDenom
MsgTransferDenom defines an SDK message for transferring an denom to recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |






<a name="irismod-nft-MsgTransferDenomResponse"></a>

### MsgTransferDenomResponse
MsgTransferDenomResponse defines the Msg/TransferDenom response type.






<a name="irismod-nft-MsgTransferNFT"></a>

### MsgTransferNFT
MsgTransferNFT defines an SDK message for transferring an NFT to recipient.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| denom_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| uri | [string](#string) |  |  |
| data | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| uri_hash | [string](#string) |  |  |






<a name="irismod-nft-MsgTransferNFTResponse"></a>

### MsgTransferNFTResponse
MsgTransferNFTResponse defines the Msg/TransferNFT response type.





 

 

 


<a name="irismod-nft-Msg"></a>

### Msg
Msg defines the nft Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IssueDenom | [MsgIssueDenom](#irismod-nft-MsgIssueDenom) | [MsgIssueDenomResponse](#irismod-nft-MsgIssueDenomResponse) | IssueDenom defines a method for issue a denom. |
| MintNFT | [MsgMintNFT](#irismod-nft-MsgMintNFT) | [MsgMintNFTResponse](#irismod-nft-MsgMintNFTResponse) | MintNFT defines a method for mint a new nft |
| EditNFT | [MsgEditNFT](#irismod-nft-MsgEditNFT) | [MsgEditNFTResponse](#irismod-nft-MsgEditNFTResponse) | RefundHTLC defines a method for editing a nft. |
| TransferNFT | [MsgTransferNFT](#irismod-nft-MsgTransferNFT) | [MsgTransferNFTResponse](#irismod-nft-MsgTransferNFTResponse) | TransferNFT defines a method for transferring a nft. |
| BurnNFT | [MsgBurnNFT](#irismod-nft-MsgBurnNFT) | [MsgBurnNFTResponse](#irismod-nft-MsgBurnNFTResponse) | BurnNFT defines a method for burning a nft. |
| TransferDenom | [MsgTransferDenom](#irismod-nft-MsgTransferDenom) | [MsgTransferDenomResponse](#irismod-nft-MsgTransferDenomResponse) | TransferDenom defines a method for transferring a denom. |

 



<a name="service_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/genesis.proto



<a name="irismod-service-GenesisState"></a>

### GenesisState
GenesisState defines the service module&#39;s genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#irismod-service-Params) |  |  |
| definitions | [ServiceDefinition](#irismod-service-ServiceDefinition) | repeated |  |
| bindings | [ServiceBinding](#irismod-service-ServiceBinding) | repeated |  |
| withdraw_addresses | [GenesisState.WithdrawAddressesEntry](#irismod-service-GenesisState-WithdrawAddressesEntry) | repeated |  |
| request_contexts | [GenesisState.RequestContextsEntry](#irismod-service-GenesisState-RequestContextsEntry) | repeated |  |






<a name="irismod-service-GenesisState-RequestContextsEntry"></a>

### GenesisState.RequestContextsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [RequestContext](#irismod-service-RequestContext) |  |  |






<a name="irismod-service-GenesisState-WithdrawAddressesEntry"></a>

### GenesisState.WithdrawAddressesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="service_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/query.proto



<a name="irismod-service-QueryBindingRequest"></a>

### QueryBindingRequest
QueryBindingRequest is request type for the Query/Binding RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |






<a name="irismod-service-QueryBindingResponse"></a>

### QueryBindingResponse
QueryDefinitionResponse is response type for the Query/Binding RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_binding | [ServiceBinding](#irismod-service-ServiceBinding) |  |  |






<a name="irismod-service-QueryBindingsRequest"></a>

### QueryBindingsRequest
QueryBindingsRequest is request type for the Query/Bindings RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request |






<a name="irismod-service-QueryBindingsResponse"></a>

### QueryBindingsResponse
QueryDefinitionsResponse is response type for the Query/Bindings RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_bindings | [ServiceBinding](#irismod-service-ServiceBinding) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-service-QueryDefinitionRequest"></a>

### QueryDefinitionRequest
QueryDefinitionRequest is request type for the Query/Definition RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |






<a name="irismod-service-QueryDefinitionResponse"></a>

### QueryDefinitionResponse
QueryDefinitionResponse is response type for the Query/Definition RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_definition | [ServiceDefinition](#irismod-service-ServiceDefinition) |  |  |






<a name="irismod-service-QueryEarnedFeesRequest"></a>

### QueryEarnedFeesRequest
QueryEarnedFeesRequest is request type for the Query/EarnedFees RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |






<a name="irismod-service-QueryEarnedFeesResponse"></a>

### QueryEarnedFeesResponse
QueryEarnedFeesResponse is response type for the Query/EarnedFees RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fees | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |






<a name="irismod-service-QueryParamsRequest"></a>

### QueryParamsRequest
QueryParametersRequest is request type for the Query/Parameters RPC method






<a name="irismod-service-QueryParamsResponse"></a>

### QueryParamsResponse
QueryParametersResponse is response type for the Query/Parameters RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#irismod-service-Params) |  |  |
| res | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-service-QueryRequestContextRequest"></a>

### QueryRequestContextRequest
QueryRequestContextRequest is request type for the Query/RequestContext RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |






<a name="irismod-service-QueryRequestContextResponse"></a>

### QueryRequestContextResponse
QueryRequestContextResponse is response type for the Query/RequestContext RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context | [RequestContext](#irismod-service-RequestContext) |  |  |






<a name="irismod-service-QueryRequestRequest"></a>

### QueryRequestRequest
QueryRequestRequest is request type for the Query/Request RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  |  |






<a name="irismod-service-QueryRequestResponse"></a>

### QueryRequestResponse
QueryRequestResponse is response type for the Query/Request RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request | [Request](#irismod-service-Request) |  |  |






<a name="irismod-service-QueryRequestsByReqCtxRequest"></a>

### QueryRequestsByReqCtxRequest
QueryRequestsByReqCtxRequest is request type for the Query/RequestsByReqCtx RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| batch_counter | [uint64](#uint64) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  |  |






<a name="irismod-service-QueryRequestsByReqCtxResponse"></a>

### QueryRequestsByReqCtxResponse
QueryRequestsByReqCtxResponse is response type for the Query/RequestsByReqCtx RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requests | [Request](#irismod-service-Request) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-service-QueryRequestsRequest"></a>

### QueryRequestsRequest
QueryRequestsRequest is request type for the Query/Requests RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  |  |






<a name="irismod-service-QueryRequestsResponse"></a>

### QueryRequestsResponse
QueryRequestsResponse is response type for the Query/Requests RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requests | [Request](#irismod-service-Request) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-service-QueryResponseRequest"></a>

### QueryResponseRequest
QueryResponseRequest is request type for the Query/Response RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  |  |






<a name="irismod-service-QueryResponseResponse"></a>

### QueryResponseResponse
QueryResponseResponse is response type for the Query/Response RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| response | [Response](#irismod-service-Response) |  |  |






<a name="irismod-service-QueryResponsesRequest"></a>

### QueryResponsesRequest
QueryResponsesRequest is request type for the Query/Responses RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| batch_counter | [uint64](#uint64) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  |  |






<a name="irismod-service-QueryResponsesResponse"></a>

### QueryResponsesResponse
QueryResponsesResponse is response type for the Query/Responses RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| responses | [Response](#irismod-service-Response) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-service-QuerySchemaRequest"></a>

### QuerySchemaRequest
QuerySchemaRequest is request type for the Query/Schema RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema_name | [string](#string) |  |  |






<a name="irismod-service-QuerySchemaResponse"></a>

### QuerySchemaResponse
QuerySchemaResponse is response type for the Query/Schema RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schema | [string](#string) |  |  |






<a name="irismod-service-QueryWithdrawAddressRequest"></a>

### QueryWithdrawAddressRequest
QueryWithdrawAddressRequest is request type for the Query/WithdrawAddress RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |






<a name="irismod-service-QueryWithdrawAddressResponse"></a>

### QueryWithdrawAddressResponse
QueryWithdrawAddressResponse is response type for the Query/WithdrawAddress RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| withdraw_address | [string](#string) |  |  |





 

 

 


<a name="irismod-service-Query"></a>

### Query
Query creates service with iservice as rpc

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Definition | [QueryDefinitionRequest](#irismod-service-QueryDefinitionRequest) | [QueryDefinitionResponse](#irismod-service-QueryDefinitionResponse) | Definition returns service definition |
| Binding | [QueryBindingRequest](#irismod-service-QueryBindingRequest) | [QueryBindingResponse](#irismod-service-QueryBindingResponse) | Binding returns service Binding with service name and provider |
| Bindings | [QueryBindingsRequest](#irismod-service-QueryBindingsRequest) | [QueryBindingsResponse](#irismod-service-QueryBindingsResponse) | Bindings returns all service Bindings with service name and owner |
| WithdrawAddress | [QueryWithdrawAddressRequest](#irismod-service-QueryWithdrawAddressRequest) | [QueryWithdrawAddressResponse](#irismod-service-QueryWithdrawAddressResponse) | WithdrawAddress returns the withdraw address of the binding owner |
| RequestContext | [QueryRequestContextRequest](#irismod-service-QueryRequestContextRequest) | [QueryRequestContextResponse](#irismod-service-QueryRequestContextResponse) | RequestContext returns the request context |
| Request | [QueryRequestRequest](#irismod-service-QueryRequestRequest) | [QueryRequestResponse](#irismod-service-QueryRequestResponse) | Request returns the request |
| Requests | [QueryRequestsRequest](#irismod-service-QueryRequestsRequest) | [QueryRequestsResponse](#irismod-service-QueryRequestsResponse) | Request returns all requests of one service with provider |
| RequestsByReqCtx | [QueryRequestsByReqCtxRequest](#irismod-service-QueryRequestsByReqCtxRequest) | [QueryRequestsByReqCtxResponse](#irismod-service-QueryRequestsByReqCtxResponse) | RequestsByReqCtx returns all requests of one service call batch |
| Response | [QueryResponseRequest](#irismod-service-QueryResponseRequest) | [QueryResponseResponse](#irismod-service-QueryResponseResponse) | Response returns the response of request |
| Responses | [QueryResponsesRequest](#irismod-service-QueryResponsesRequest) | [QueryResponsesResponse](#irismod-service-QueryResponsesResponse) | Responses returns all responses of one service call batch |
| EarnedFees | [QueryEarnedFeesRequest](#irismod-service-QueryEarnedFeesRequest) | [QueryEarnedFeesResponse](#irismod-service-QueryEarnedFeesResponse) | EarnedFees returns the earned service fee of one provider |
| Schema | [QuerySchemaRequest](#irismod-service-QuerySchemaRequest) | [QuerySchemaResponse](#irismod-service-QuerySchemaResponse) | Schema returns the schema |
| Params | [QueryParamsRequest](#irismod-service-QueryParamsRequest) | [QueryParamsResponse](#irismod-service-QueryParamsResponse) | Params queries the service parameters |

 



<a name="service_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/service.proto



<a name="irismod-service-CompactRequest"></a>

### CompactRequest
CompactRequest defines a standard for compact request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| request_context_batch_counter | [uint64](#uint64) |  |  |
| provider | [string](#string) |  |  |
| service_fee | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| request_height | [int64](#int64) |  |  |
| expiration_height | [int64](#int64) |  |  |






<a name="irismod-service-Params"></a>

### Params
Params defines service module&#39;s parameters


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_request_timeout | [int64](#int64) |  |  |
| min_deposit_multiple | [int64](#int64) |  |  |
| min_deposit | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| service_fee_tax | [string](#string) |  |  |
| slash_fraction | [string](#string) |  |  |
| complaint_retrospect | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| arbitration_time_limit | [google.protobuf.Duration](#google-protobuf-Duration) |  |  |
| tx_size_limit | [uint64](#uint64) |  |  |
| base_denom | [string](#string) |  |  |
| restricted_service_fee_denom | [bool](#bool) |  |  |






<a name="irismod-service-Pricing"></a>

### Pricing
Pricing defines a standard for service pricing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| price | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| promotions_by_time | [PromotionByTime](#irismod-service-PromotionByTime) | repeated |  |
| promotions_by_volume | [PromotionByVolume](#irismod-service-PromotionByVolume) | repeated |  |






<a name="irismod-service-PromotionByTime"></a>

### PromotionByTime
PromotionByTime defines a standard for service promotion by time


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| discount | [string](#string) |  |  |






<a name="irismod-service-PromotionByVolume"></a>

### PromotionByVolume
PromotionByVolume defines a standard for service promotion by volume


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| volume | [uint64](#uint64) |  |  |
| discount | [string](#string) |  |  |






<a name="irismod-service-Request"></a>

### Request
Request defines a standard for request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| consumer | [string](#string) |  |  |
| input | [string](#string) |  |  |
| service_fee | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| request_height | [int64](#int64) |  |  |
| expiration_height | [int64](#int64) |  |  |
| request_context_id | [string](#string) |  |  |
| request_context_batch_counter | [uint64](#uint64) |  |  |






<a name="irismod-service-RequestContext"></a>

### RequestContext
RequestContext defines a standard for request context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| providers | [string](#string) | repeated |  |
| consumer | [string](#string) |  |  |
| input | [string](#string) |  |  |
| service_fee_cap | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| module_name | [string](#string) |  |  |
| timeout | [int64](#int64) |  |  |
| repeated | [bool](#bool) |  |  |
| repeated_frequency | [uint64](#uint64) |  |  |
| repeated_total | [int64](#int64) |  |  |
| batch_counter | [uint64](#uint64) |  |  |
| batch_request_count | [uint32](#uint32) |  |  |
| batch_response_count | [uint32](#uint32) |  |  |
| batch_response_threshold | [uint32](#uint32) |  |  |
| response_threshold | [uint32](#uint32) |  |  |
| batch_state | [RequestContextBatchState](#irismod-service-RequestContextBatchState) |  |  |
| state | [RequestContextState](#irismod-service-RequestContextState) |  |  |






<a name="irismod-service-Response"></a>

### Response
Response defines a standard for response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [string](#string) |  |  |
| consumer | [string](#string) |  |  |
| result | [string](#string) |  |  |
| output | [string](#string) |  |  |
| request_context_id | [string](#string) |  |  |
| request_context_batch_counter | [uint64](#uint64) |  |  |






<a name="irismod-service-ServiceBinding"></a>

### ServiceBinding
ServiceBinding defines a standard for service binding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| deposit | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| pricing | [string](#string) |  |  |
| qos | [uint64](#uint64) |  |  |
| options | [string](#string) |  |  |
| available | [bool](#bool) |  |  |
| disabled_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-ServiceDefinition"></a>

### ServiceDefinition
ServiceDefinition defines a standard for service definition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| tags | [string](#string) | repeated |  |
| author | [string](#string) |  |  |
| author_description | [string](#string) |  |  |
| schemas | [string](#string) |  |  |





 


<a name="irismod-service-RequestContextBatchState"></a>

### RequestContextBatchState
RequestContextBatchState is a type alias that represents a request batch status as a byte

| Name | Number | Description |
| ---- | ------ | ----------- |
| BATCH_RUNNING | 0 | BATCH_RUNNING defines the running batch status. |
| BATCH_COMPLETED | 1 | BATCH_COMPLETED defines the completed batch status. |



<a name="irismod-service-RequestContextState"></a>

### RequestContextState
RequestContextState is a type alias that represents a request status as a byte

| Name | Number | Description |
| ---- | ------ | ----------- |
| RUNNING | 0 | RUNNING defines the running request context status |
| PAUSED | 1 | PAUSED defines the paused request context status |
| COMPLETED | 2 | COMPLETED defines the completed request context status |


 

 

 



<a name="service_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service/tx.proto



<a name="irismod-service-MsgBindService"></a>

### MsgBindService
MsgBindService defines an SDK message for binding to an existing service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| deposit | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| pricing | [string](#string) |  |  |
| qos | [uint64](#uint64) |  |  |
| options | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-MsgBindServiceResponse"></a>

### MsgBindServiceResponse
MsgBindServiceResponse defines the Msg/BindService response type






<a name="irismod-service-MsgCallService"></a>

### MsgCallService
MsgCallService defines an SDK message to initiate a service request context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| providers | [string](#string) | repeated |  |
| consumer | [string](#string) |  |  |
| input | [string](#string) |  |  |
| service_fee_cap | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| timeout | [int64](#int64) |  |  |
| repeated | [bool](#bool) |  |  |
| repeated_frequency | [uint64](#uint64) |  |  |
| repeated_total | [int64](#int64) |  |  |






<a name="irismod-service-MsgCallServiceResponse"></a>

### MsgCallServiceResponse
MsgCallServiceResponse defines the Msg/CallService response type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |






<a name="irismod-service-MsgDefineService"></a>

### MsgDefineService
MsgDefineService defines an SDK message for defining a new service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| tags | [string](#string) | repeated |  |
| author | [string](#string) |  |  |
| author_description | [string](#string) |  |  |
| schemas | [string](#string) |  |  |






<a name="irismod-service-MsgDefineServiceResponse"></a>

### MsgDefineServiceResponse
MsgDefineServiceResponse defines the Msg/DefineService response type






<a name="irismod-service-MsgDisableServiceBinding"></a>

### MsgDisableServiceBinding
MsgDisableServiceBinding defines an SDK message to disable a service binding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-MsgDisableServiceBindingResponse"></a>

### MsgDisableServiceBindingResponse
MsgDisableServiceBindingResponse defines the Msg/DisableServiceBinding response type






<a name="irismod-service-MsgEnableServiceBinding"></a>

### MsgEnableServiceBinding
MsgEnableServiceBinding defines an SDK message to enable a service binding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| deposit | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-MsgEnableServiceBindingResponse"></a>

### MsgEnableServiceBindingResponse
MsgEnableServiceBindingResponse defines the Msg/EnableServiceBinding response type






<a name="irismod-service-MsgKillRequestContext"></a>

### MsgKillRequestContext
MsgKillRequestContext defines an SDK message to terminate a service request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| consumer | [string](#string) |  |  |






<a name="irismod-service-MsgKillRequestContextResponse"></a>

### MsgKillRequestContextResponse
MsgKillRequestContextResponse defines the Msg/KillRequestContext response type






<a name="irismod-service-MsgPauseRequestContext"></a>

### MsgPauseRequestContext
MsgPauseRequestContext defines an SDK message to pause a service request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| consumer | [string](#string) |  |  |






<a name="irismod-service-MsgPauseRequestContextResponse"></a>

### MsgPauseRequestContextResponse
MsgPauseRequestContextResponse defines the Msg/PauseRequestContext response type






<a name="irismod-service-MsgRefundServiceDeposit"></a>

### MsgRefundServiceDeposit
MsgRefundServiceDeposit defines an SDK message to refund deposit from a service binding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-MsgRefundServiceDepositResponse"></a>

### MsgRefundServiceDepositResponse
MsgRefundServiceDepositResponse defines the Msg/RefundServiceDeposit response type






<a name="irismod-service-MsgRespondService"></a>

### MsgRespondService
MsgRespondService defines an SDK message to respond a service request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| result | [string](#string) |  |  |
| output | [string](#string) |  |  |






<a name="irismod-service-MsgRespondServiceResponse"></a>

### MsgRespondServiceResponse
MsgRespondServiceResponse defines the Msg/RespondService response type






<a name="irismod-service-MsgSetWithdrawAddress"></a>

### MsgSetWithdrawAddress
MsgSetWithdrawAddress defines an SDK message to set the withdrawal address for a provider


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| withdraw_address | [string](#string) |  |  |






<a name="irismod-service-MsgSetWithdrawAddressResponse"></a>

### MsgSetWithdrawAddressResponse
MsgSetWithdrawAddressResponse defines the Msg/SetWithdrawAddress response type






<a name="irismod-service-MsgStartRequestContext"></a>

### MsgStartRequestContext
MsgStartRequestContext defines an SDK message to resume a service request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| consumer | [string](#string) |  |  |






<a name="irismod-service-MsgStartRequestContextResponse"></a>

### MsgStartRequestContextResponse
MsgStartRequestContextResponse defines the Msg/StartRequestContext response type






<a name="irismod-service-MsgUpdateRequestContext"></a>

### MsgUpdateRequestContext
MsgUpdateRequestContext defines an SDK message to update a service request context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_context_id | [string](#string) |  |  |
| providers | [string](#string) | repeated |  |
| consumer | [string](#string) |  |  |
| service_fee_cap | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| timeout | [int64](#int64) |  |  |
| repeated_frequency | [uint64](#uint64) |  |  |
| repeated_total | [int64](#int64) |  |  |






<a name="irismod-service-MsgUpdateRequestContextResponse"></a>

### MsgUpdateRequestContextResponse
MsgUpdateRequestContextResponse defines the Msg/UpdateRequestContext response type






<a name="irismod-service-MsgUpdateServiceBinding"></a>

### MsgUpdateServiceBinding
MsgUpdateServiceBinding defines an SDK message for updating an existing service binding


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service_name | [string](#string) |  |  |
| provider | [string](#string) |  |  |
| deposit | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |
| pricing | [string](#string) |  |  |
| qos | [uint64](#uint64) |  |  |
| options | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-service-MsgUpdateServiceBindingResponse"></a>

### MsgUpdateServiceBindingResponse
MsgUpdateServiceBindingResponse defines the Msg/UpdateServiceBinding response type






<a name="irismod-service-MsgWithdrawEarnedFees"></a>

### MsgWithdrawEarnedFees
MsgWithdrawEarnedFees defines an SDK message to withdraw the fees earned by the provider or owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| provider | [string](#string) |  |  |






<a name="irismod-service-MsgWithdrawEarnedFeesResponse"></a>

### MsgWithdrawEarnedFeesResponse
MsgWithdrawEarnedFeesResponse defines the Msg/WithdrawEarnedFees response type





 

 

 


<a name="irismod-service-Msg"></a>

### Msg
Msg defines the oracle Msg service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DefineService | [MsgDefineService](#irismod-service-MsgDefineService) | [MsgDefineServiceResponse](#irismod-service-MsgDefineServiceResponse) | DefineService defines a method for define a new service |
| BindService | [MsgBindService](#irismod-service-MsgBindService) | [MsgBindServiceResponse](#irismod-service-MsgBindServiceResponse) | BindService defines a method for bind a server |
| UpdateServiceBinding | [MsgUpdateServiceBinding](#irismod-service-MsgUpdateServiceBinding) | [MsgUpdateServiceBindingResponse](#irismod-service-MsgUpdateServiceBindingResponse) | UpdateServiceBinding defines a method for update a service binding |
| SetWithdrawAddress | [MsgSetWithdrawAddress](#irismod-service-MsgSetWithdrawAddress) | [MsgSetWithdrawAddressResponse](#irismod-service-MsgSetWithdrawAddressResponse) | SetWithdrawAddress defines a method for setting a withdraw address |
| EnableServiceBinding | [MsgEnableServiceBinding](#irismod-service-MsgEnableServiceBinding) | [MsgEnableServiceBindingResponse](#irismod-service-MsgEnableServiceBindingResponse) | EnableServiceBinding defines a method for enabling a service binding |
| DisableServiceBinding | [MsgDisableServiceBinding](#irismod-service-MsgDisableServiceBinding) | [MsgDisableServiceBindingResponse](#irismod-service-MsgDisableServiceBindingResponse) | DisableServiceBinding defines a method for disabling a service binding |
| RefundServiceDeposit | [MsgRefundServiceDeposit](#irismod-service-MsgRefundServiceDeposit) | [MsgRefundServiceDepositResponse](#irismod-service-MsgRefundServiceDepositResponse) | RefundServiceDeposit defines a method for refunding a fee |
| CallService | [MsgCallService](#irismod-service-MsgCallService) | [MsgCallServiceResponse](#irismod-service-MsgCallServiceResponse) | CallService defines a method for calling a service |
| RespondService | [MsgRespondService](#irismod-service-MsgRespondService) | [MsgRespondServiceResponse](#irismod-service-MsgRespondServiceResponse) | RespondService defines a method for responding a service |
| PauseRequestContext | [MsgPauseRequestContext](#irismod-service-MsgPauseRequestContext) | [MsgPauseRequestContextResponse](#irismod-service-MsgPauseRequestContextResponse) | PauseRequestContext defines a method for pausing a service call |
| StartRequestContext | [MsgStartRequestContext](#irismod-service-MsgStartRequestContext) | [MsgStartRequestContextResponse](#irismod-service-MsgStartRequestContextResponse) | StartRequestContext defines a method for starting a service call |
| KillRequestContext | [MsgKillRequestContext](#irismod-service-MsgKillRequestContext) | [MsgKillRequestContextResponse](#irismod-service-MsgKillRequestContextResponse) | KillRequestContext defines a method for killing a service call |
| UpdateRequestContext | [MsgUpdateRequestContext](#irismod-service-MsgUpdateRequestContext) | [MsgUpdateRequestContextResponse](#irismod-service-MsgUpdateRequestContextResponse) | UpdateRequestContext defines a method for updating a service call |
| WithdrawEarnedFees | [MsgWithdrawEarnedFees](#irismod-service-MsgWithdrawEarnedFees) | [MsgWithdrawEarnedFeesResponse](#irismod-service-MsgWithdrawEarnedFeesResponse) | WithdrawEarnedFees defines a method for Withdrawing a earned fees |

 



<a name="token_genesis-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/genesis.proto



<a name="irismod-token-GenesisState"></a>

### GenesisState
GenesisState defines the token module&#39;s genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#irismod-token-Params) |  |  |
| tokens | [Token](#irismod-token-Token) | repeated |  |
| burned_coins | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |





 

 

 

 



<a name="token_query-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/query.proto



<a name="irismod-token-QueryFeesRequest"></a>

### QueryFeesRequest
QueryFeesRequest is request type for the Query/Fees RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |






<a name="irismod-token-QueryFeesResponse"></a>

### QueryFeesResponse
QueryFeesResponse is response type for the Query/Fees RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| issue_fee | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) |  |  |
| mint_fee | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) |  |  |






<a name="irismod-token-QueryParamsRequest"></a>

### QueryParamsRequest
QueryParametersRequest is request type for the Query/Parameters RPC method






<a name="irismod-token-QueryParamsResponse"></a>

### QueryParamsResponse
QueryParametersResponse is response type for the Query/Parameters RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#irismod-token-Params) |  |  |
| res | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-token-QueryTokenRequest"></a>

### QueryTokenRequest
QueryTokenRequest is request type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| denom | [string](#string) |  |  |






<a name="irismod-token-QueryTokenResponse"></a>

### QueryTokenResponse
QueryTokenResponse is response type for the Query/Token RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Token | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="irismod-token-QueryTokensRequest"></a>

### QueryTokensRequest
QueryTokensRequest is request type for the Query/Tokens RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos-base-query-v1beta1-PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="irismod-token-QueryTokensResponse"></a>

### QueryTokensResponse
QueryTokensResponse is response type for the Query/Tokens RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Tokens | [google.protobuf.Any](#google-protobuf-Any) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos-base-query-v1beta1-PageResponse) |  |  |






<a name="irismod-token-QueryTotalBurnRequest"></a>

### QueryTotalBurnRequest
QueryTokenRequest is request type for the Query/TotalBurn RPC method






<a name="irismod-token-QueryTotalBurnResponse"></a>

### QueryTotalBurnResponse
QueryTotalBurnResponse is response type for the Query/TotalBurn RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| burned_coins | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) | repeated |  |





 

 

 


<a name="irismod-token-Query"></a>

### Query
Query creates service with token as RPC

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Token | [QueryTokenRequest](#irismod-token-QueryTokenRequest) | [QueryTokenResponse](#irismod-token-QueryTokenResponse) | Token returns token with token name |
| Tokens | [QueryTokensRequest](#irismod-token-QueryTokensRequest) | [QueryTokensResponse](#irismod-token-QueryTokensResponse) | Tokens returns the token list |
| Fees | [QueryFeesRequest](#irismod-token-QueryFeesRequest) | [QueryFeesResponse](#irismod-token-QueryFeesResponse) | Fees returns the fees to issue or mint a token |
| Params | [QueryParamsRequest](#irismod-token-QueryParamsRequest) | [QueryParamsResponse](#irismod-token-QueryParamsResponse) | Params queries the token parameters |
| TotalBurn | [QueryTotalBurnRequest](#irismod-token-QueryTotalBurnRequest) | [QueryTotalBurnResponse](#irismod-token-QueryTotalBurnResponse) | TotalBurn queries all the burnt coins |

 



<a name="token_token-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/token.proto



<a name="irismod-token-Params"></a>

### Params
Params defines token module&#39;s parameters


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_tax_rate | [string](#string) |  |  |
| issue_token_base_fee | [cosmos.base.v1beta1.Coin](#cosmos-base-v1beta1-Coin) |  |  |
| mint_token_fee_ratio | [string](#string) |  |  |






<a name="irismod-token-Token"></a>

### Token
Token defines a standard for the fungible token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |
| name | [string](#string) |  |  |
| scale | [uint32](#uint32) |  |  |
| min_unit | [string](#string) |  |  |
| initial_supply | [uint64](#uint64) |  |  |
| max_supply | [uint64](#uint64) |  |  |
| mintable | [bool](#bool) |  |  |
| owner | [string](#string) |  |  |





 

 

 

 



<a name="token_tx-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token/tx.proto



<a name="irismod-token-MsgBurnToken"></a>

### MsgBurnToken
MsgBurnToken defines an SDK message for burning some tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |
| sender | [string](#string) |  |  |






<a name="irismod-token-MsgBurnTokenResponse"></a>

### MsgBurnTokenResponse
MsgBurnTokenResponse defines the Msg/BurnToken response type






<a name="irismod-token-MsgEditToken"></a>

### MsgEditToken
MsgEditToken defines an SDK message for editing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |
| name | [string](#string) |  |  |
| max_supply | [uint64](#uint64) |  |  |
| mintable | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-token-MsgEditTokenResponse"></a>

### MsgEditTokenResponse
MsgEditTokenResponse defines the Msg/EditToken response type






<a name="irismod-token-MsgIssueToken"></a>

### MsgIssueToken
MsgIssueToken defines an SDK message for issuing a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |
| name | [string](#string) |  |  |
| scale | [uint32](#uint32) |  |  |
| min_unit | [string](#string) |  |  |
| initial_supply | [uint64](#uint64) |  |  |
| max_supply | [uint64](#uint64) |  |  |
| mintable | [bool](#bool) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-token-MsgIssueTokenResponse"></a>

### MsgIssueTokenResponse
MsgIssueTokenResponse defines the Msg/IssueToken response type






<a name="irismod-token-MsgMintToken"></a>

### MsgMintToken
MsgMintToken defines an SDK message for minting a new token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| symbol | [string](#string) |  |  |
| amount | [uint64](#uint64) |  |  |
| to | [string](#string) |  |  |
| owner | [string](#string) |  |  |






<a name="irismod-token-MsgMintTokenResponse"></a>

### MsgMintTokenResponse
MsgMintTokenResponse defines the Msg/MintToken response type






<a name="irismod-token-MsgTransferTokenOwner"></a>

### MsgTransferTokenOwner
MsgTransferTokenOwner defines an SDK message for transferring the token owner


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| src_owner | [string](#string) |  |  |
| dst_owner | [string](#string) |  |  |
| symbol | [string](#string) |  |  |






<a name="irismod-token-MsgTransferTokenOwnerResponse"></a>

### MsgTransferTokenOwnerResponse
MsgTransferTokenOwnerResponse defines the Msg/TransferTokenOwner response type





 

 

 


<a name="irismod-token-Msg"></a>

### Msg
Msg defines the oracle Msg service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IssueToken | [MsgIssueToken](#irismod-token-MsgIssueToken) | [MsgIssueTokenResponse](#irismod-token-MsgIssueTokenResponse) | IssueToken defines a method for issuing a new token |
| EditToken | [MsgEditToken](#irismod-token-MsgEditToken) | [MsgEditTokenResponse](#irismod-token-MsgEditTokenResponse) | EditToken defines a method for editing a token |
| MintToken | [MsgMintToken](#irismod-token-MsgMintToken) | [MsgMintTokenResponse](#irismod-token-MsgMintTokenResponse) | MintToken defines a method for minting some tokens |
| BurnToken | [MsgBurnToken](#irismod-token-MsgBurnToken) | [MsgBurnTokenResponse](#irismod-token-MsgBurnTokenResponse) | BurnToken defines a method for burning some tokens |
| TransferTokenOwner | [MsgTransferTokenOwner](#irismod-token-MsgTransferTokenOwner) | [MsgTransferTokenOwnerResponse](#irismod-token-MsgTransferTokenOwnerResponse) | TransferTokenOwner defines a method for minting some tokens |

 



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

