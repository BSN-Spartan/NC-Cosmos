# Service

Service module allows you to define, bind, invoke services on the Spartan-Cosmos.

## Available Commands

| Type  | Name                                                       | Description                                                        |
| ----- | ---------------------------------------------------------- | ------------------------------------------------------------------ |
| Tx    | [bind](#spartan-tx-service-bind)                           | Bind an existing service definition                                |
| Tx    | [call](#spartan-tx-service-call)                           | Initiate a service call                                            |
| Tx    | [define](#spartan-tx-service-define)                       | Define a new service                                               |
| Tx    | [disable](#spartan-tx-service-disable)                     | Disable an available service binding                               |
| Tx    | [enable](#spartan-tx-service-enable)                       | Enable an unavailable service binding                              |
| Tx    | [kill](#spartan-tx-service-kill)                           | Terminate a request context                                        |
| Tx    | [pause](#spartan-tx-service-pause)                         | Pause a running request context                                    |
| Tx    | [refund-deposit](#spartan-tx-service-refund-deposit)       | Refund all deposit from a service binding                          |
| Tx    | [respond](#spartan-tx-service-respond)                     | Respond to a service request                                       |
| Tx    | [set-withdraw-addr](#spartan-tx-service-set-withdraw-addr) | Set a withdrawal address for an owner                              |
| Tx    | [start](#spartan-tx-service-start)                         | Start a paused request context                                     |
| Tx    | [update](#spartan-tx-service-update)                       | Update a request context                                           |
| Tx    | [update-binding](#spartan-tx-service-update-binding)       | Update an existing service binding                                 |
| Tx    | [withdraw-fees](#spartan-tx-service-withdraw-fees)         | Withdraw the earned fees of a provider or owner                    |
| Query | [binding](#spartan-query-service-binding)                  | Query a service binding                                            |
| Query | [bindings](#spartan-query-service-bindings)                | Query all bindings of a service definition with an optional owner  |
| Query | [definition](#spartan-query-service-definition)            | Query a service definition                                         |
| Query | [fees](#spartan-query-service-fees)                        | Query the earned fees of a provider                                |
| Query | [params](#spartan-query-service-params)                    | Query the current service parameter values                         |
| Query | [request](#spartan-query-service-request)                  | Query a request by the request ID                                  |
| Query | [request-context](#spartan-query-service-request-context)  | Query a request context                                            |
| Query | [requests](#spartan-query-service-requests)                | Query active requests by the service binding or request context ID |
| Query | [response](#spartan-query-service-response)                | Query a response by the request ID                                 |
| Query | [responses](#spartan-query-service-responses)              | Query active responses by the request context ID and batch counter |
| Query | [schema](#spartan-query-service-schema)                    | Query the system schema by the schema name                         |
| Query | [withdraw-addr](#spartan-query-service-withdraw-addr)      | Query the withdrawal address of an owner                           |

## spartan tx service bind

Bind an existing service definition.

```bash
spartan tx service bind [flags]
```

## spartan tx service call

Initiate a service call.

```bash
spartan tx service call [flags]
```

## spartan tx service define

Define a new service based on the given params.

```bash
spartan tx service define [flags]
```

## spartan tx service disable

Disable an available service binding.

```bash
spartan tx service disable [service-name] [provider-address] [flags]
```

## spartan tx service enable

Enable an unavailable service binding.

```bash
spartan tx service enable [service-name] [provider-address] [flags]
```

## spartan tx service kill

Terminate a request context.

```bash
spartan tx service kill [request-context-id] [flags]
```

## spartan tx service pause

Pause a running request context.

```bash
spartan tx service pause [request-context-id] [flags]
```

## spartan tx service refund-deposit

Refund all deposit from a service binding.

```bash
spartan tx service refund-deposit [service-name] [provider-address] [flags]
```

## spartan tx service respond

Respond to an active service request.

```bash
spartan tx service respond [flags]
```

## spartan tx service set-withdraw-addr

Set a withdrawal address for an owner.

```bash
spartan tx service set-withdraw-addr [withdrawal-address] [flags]
```

## spartan tx service start

Start a paused request context.

```bash
spartan tx service start [request-context-id] [flags]
```

## spartan tx service update

Update a request context.

```bash
spartan tx service update [request-context-id] [flags]
```

## spartan tx service update-binding

Update an existing service binding.

```bash
spartan tx service update-binding [service-name] [provider-address] [flags]
```

## spartan tx service withdraw-fees

Withdraw the earned fees of the specified provider, for all providers of the owner if the provider not given.

```bash
spartan tx service withdraw-fees [provider-address] [flags]
```

## spartan query service binding

Query details of a service binding.

```bash
spartan query service binding [service-name] [provider-address] [flags]
```

## spartan query service bindings

Query all bindings of a service definition with an optional owner.

```bash
spartan query service bindings [service-name] [flags]
```

## spartan query service definition

Query details of a service definition.

```bash
spartan query service definition [service-name] [flags]
```

## spartan query service fees

Query the earned fees of a provider.

```bash
spartan query service fees [provider-address] [flags]
```

## spartan query service params

Query values set as service parameters.

```bash
spartan query service params [flags]
```

## spartan query service request

Query details of a service request.

```bash
spartan query service request [request-id] [flags]
```

## spartan query service request-context

Query a request context.

```bash
spartan query service request-context [request-context-id] [flags]
```

## spartan query service requests

Query active requests by the service binding or request context ID.

```bash
spartan query service requests [service-name] [provider] | [request-context-id] [batch-counter] [flags]
```

## spartan query service response

Query details of a service response.

```bash
spartan query service response [request-id] [flags]
```

## spartan query service responses

Query active responses by the request context ID and batch counter.

```bash
spartan query service responses [request-context-id] [batch-counter] [flags]
```

## spartan query service schema

Query the system schema by the schema name, only pricing and result allowed.

```bash
spartan query service schema [schema-name] [flags]
```

## spartan query service withdraw-addr

Query the withdrawal address of an owner.

```bash
spartan query service withdraw-addr [address] [flags]
```
