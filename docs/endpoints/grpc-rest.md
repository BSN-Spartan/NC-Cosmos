---
order: 2
---

# gRPC Gateway

## API Port, Activation and Configuration

All routes are configured under the following fields in `~/.spartan/config/app.toml`:

- `api.enable = true|false` field defines if the REST server should be enabled. Defaults to `true`.
- `api.address = {string}` field defines the address (really, the port, since the host should be kept at `0.0.0.0`) the server should bind to. Defaults to `tcp://0.0.0.0:1317`.
- some additional API configuration options are defined in `~/.spartan/config/app.toml`, along with comments, please refer to that file directly.

### gRPC-gateway REST Routes

If, for various reasons, you cannot use gRPC (for example, you are building a web application, and browsers don't support HTTP2 on which gRPC is built), then the Spartan-Cosmos offers REST routes via gRPC-gateway.

[gRPC-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) is a tool to expose gRPC endpoints as REST endpoints. For each RPC endpoint defined in a Protobuf service, the SDK offers a REST equivalent. 

For application developers, gRPC-gateway REST routes needs to be wired up to the REST server, this is done by calling the `RegisterGRPCGatewayRoutes` function on the ModuleManager.

## API Endpoints

**Spartan-Cosmos API Endpoints**

| API Endpoints                                                            | Description                                                                                      |
| :----------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------- |
| `GET` `/cosmos/auth/v1beta1/accounts/{address}`                          | Return account details based on address                                                          |
| `GET` `/cosmos/auth/v1beta1/params`                                      | Query all parameters                                                                             |
| `GET` `/cosmos/bank/v1beta1/balances/{address}`                          | Query the balance of all coins for a single account                                              |
| `GET` `/cosmos/bank/v1beta1/balances/{address}/{denom}`                  | Query the balance of a single coin for a single account                                          |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata`                             | Query the client metadata for all registered coin denominations                                  |
| `GET` `/cosmos/bank/v1beta1/denoms_metadata/{denom}`                     | Query the client metadata of a given coin denomination                                           |
| `GET` `/cosmos/bank/v1beta1/params`                                      | Query the parameters of bank module                                                              |
| `GET` `/cosmos/bank/v1beta1/supply`                                      | Query the total supply of all coins                                                              |
| `GET` `/cosmos/bank/v1beta1/supply/{denom}`                              | Query the supply of a single coin                                                                |
| `GET` `/cosmos/gov/v1beta1/params/{params_type}`                         | Query all parameters of the gov module                                                           |
| `GET` `/cosmos/gov/v1beta1/proposals`                                    | Query all proposals based on given status                                                        |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}`                      | Query proposal details based on ProposalID                                                       |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits`             | Query all deposits of a single proposal                                                          |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}` | Query single deposit information based proposalID, depositAddr                                   |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/tally`                | Query the tally of a proposal vote                                                               |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes`                | Query votes of a given proposal                                                                  |
| `GET` `/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}`        | Query voted information based on proposalID, voterAddr                                           |
| `GET` `/cosmos/params/v1beta1/params`                                    | Query a specific parameter of a module, given its subspace and key                               |
| `GET` `/cosmos/slashing/v1beta1/params`                                  | Query the parameters of slashing module                                                          |
| `GET` `/cosmos/slashing/v1beta1/signing_infos`                           | Query signing info of all validators                                                             |
| `GET` `/cosmos/slashing/v1beta1/signing_infos/{cons_address}`            | Query the signing info of given cons address                                                     |
| `GET` `/cosmos/upgrade/v1beta1/applied_plan/{name}`                      | Query a previously applied upgrade plan by its name                                              |
| `GET` `/cosmos/upgrade/v1beta1/current_plan`                             | Query the current upgrade plan                                                                   |
| `GET` `/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}`   | Query the consensus state that will serve as a trusted kernel for the next version of this chain |
| `GET` `/irismod/mt/denoms/supply`                                        | Supply queries the total supply of a given denom or owner                                        |
| `GET` `/irismod/mt/denoms`                                               | Denoms queries all the denoms                                                                    |
| `GET` `/irismod/mt/denoms/{denom_id}`                                    | Denom queries the definition of a given denom ID                                                 |
| `GET` `/irismod/mt/mts/{denom_id}/{mt_id}/supply`                        | MTSupply queries the total supply of given denom and mt ID                                       |
| `GET` `/irismod/mt/mts/{denom_id}`                                       | MTs queries all the MTs of a given denom ID                                                      |
| `GET` `/irismod/mt/mts/{denom_id}/{mt_id}`                               | MT queries the MT of the given denom and mt ID                                                   |
| `GET` `/irismod/mt/mts/{owner}/balances`                                 | Balances queries the MT balances of a specified owner                                            |
| `GET` `/irismod/nft/collections/{denom_id}`                              | Query the NFTs by the specified denom                                                            |
| `GET` `/irismod/nft/collections/{denom_id}/supply`                       | Query the total supply by a given denom                                                          |
| `GET` `/irismod/nft/denoms`                                              | Query all the denoms                                                                             |
| `GET` `/irismod/nft/denoms/{denom_id}`                                   | Query the definition by a given denom ID                                                         |
| `GET` `/irismod/nft/nfts`                                                | Query the NFTs by the specified owner                                                            |
| `GET` `/irismod/nft/nfts/{denom_id}/{token_id}`                          | Query the NFT by the given denom ID and token ID                                                 |
| `GET` `/irismod/service/bindings/{service_name}`                         | Return all service Bindings with service name and owner                                          |
| `GET` `/irismod/service/bindings/{service_name}/{provider}`              | Return service Binding with service name and provider                                            |
| `GET` `/irismod/service/contexts/{request_context_id}`                   | Return the request context                                                                       |
| `GET` `/irismod/service/definitions/{service_name}`                      | Return service definition                                                                        |
| `GET` `/irismod/service/fees/{provider}`                                 | Return the earned service fee of one provider                                                    |
| `GET` `/irismod/service/owners/{owner}/withdraw-address`                 | Return the withdraw address of the binding owner                                                 |
| `GET` `/irismod/service/params`                                          | Query the service parameters                                                                     |
| `GET` `/irismod/service/requests/{request_context_id}/{batch_counter}`   | Return all requests of one service call batch                                                    |
| `GET` `/irismod/service/requests/{request_id}`                           | Return the request                                                                               |
| `GET` `/irismod/service/requests/{service_name}/{provider}`              | Return all requests of one service with provider                                                 |
| `GET` `/irismod/service/responses/{request_context_id}/{batch_counter}`  | Return all responses of one service call batch                                                   |
| `GET` `/irismod/service/responses/{request_id}`                          | Return the response of one request                                                               |
| `GET` `/irismod/service/schemas/{schema_name}`                           | Return the schema                                                                                |

**Tendermint API Endpoints**

| API Endpoints                                                  | Description                                                |
|:-------------------------------------------------------------- |:---------------------------------------------------------- |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/latest`          | Return the latest block.                                   |
| `GET` `/cosmos/base/tendermint/v1beta1/blocks/{height}`        | Query block for given height.                              |
| `GET` `/cosmos/base/tendermint/v1beta1/node_info`              | Query the current node info.                               |
| `GET` `/cosmos/base/tendermint/v1beta1/syncing`                | Query node syncing.                                        |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/latest`   | Query latest validator-set.                                |
| `GET` `/cosmos/base/tendermint/v1beta1/validatorsets/{height}` | Query validator-set at a given height.                     |
| `POST` `/cosmos/tx/v1beta1/simulate`                           | Simulate executing a transaction for estimating gas usage. |
| `GET` `/cosmos/tx/v1beta1/txs`                                 | Fetch txs by event.                                        |
| `POST` `/cosmos/tx/v1beta1/txs`                                | Broadcast transaction.                                     |
| `GET` `/cosmos/tx/v1beta1/txs/{hash}`                          | Fetch a tx by hash.                                        |

## Generating and Signing Transactions

It is not possible to generate or sign a transaction using REST, only to broadcast one. You can generating and signing transactions using [gRPC Client](./grpc-client.md).

## Broadcasting Transactions

Broadcasting a transaction using the gRPC-gateway REST endpoint `cosmos/tx/v1beta1/txs` can be done by sending a POST request as follows, where the `txBytes` are the protobuf-encoded bytes of a signed transaction:

```bash
curl -X POST \
    -H "Content-Type: application/json" \
    -d'{"tx_bytes":"{{txBytes}}","mode":"BROADCAST_MODE_SYNC"}' \
    "localhost:1317/cosmos/tx/v1beta1/txs"
```

## Querying Transactions

Querying transactions using the gRPC-gateway REST endpoint can be done by sending a GET request as follows:

- **Query tx by hash:** `/cosmos/tx/v1beta1/txs/{hash}`

  ```bash
  curl -X GET \
      -H "accept: application/json" \
      "http://localhost:1317/cosmos/tx/v1beta1/txs/{hash}"
  ```

- **Query tx by events:** `/cosmos/tx/v1beta1/txs`

  ```bash
  curl -X GET \
      -H "accept: application/json" \
      "http://localhost:1317/cosmos/tx/v1beta1/txs?events={event_content}"
  ```
