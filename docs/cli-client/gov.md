# Gov

The Gov module provides the basic functionalities for governance.

## Available Commands

| Type  | Name                                               | Description                                                                  |
| ----- | -------------------------------------------------- | ---------------------------------------------------------------------------- |
| Tx    | [deposit](#spartan-tx-gov-deposit)                 | Deposit tokens for an active proposal                                        |
| Tx    | [submit-proposal](#spartan-tx-gov-submit-proposal) | Submit a proposal along with an initial deposit                              |
| Tx    | [vote](#spartan-tx-gov-vote)                       | Vote for an active proposal, options: yes/no/no_with_veto/abstain            |
| Tx    | [weighted-vote ](#spartan-tx-gov-weighted-vote)    | Vote for an active proposal, options: yes/no/no_with_veto/abstain            |
| Query | [deposit](#spartan-query-gov-deposit)              | Query details of a deposit                                                   |
| Query | [deposits](#spartan-query-gov-deposits)            | Query deposits on a proposal                                                 |
| Query | [param](#spartan-query-gov-param)                  | Query the parameters (voting, tallying or deposit) of the governance process |
| Query | [params](#spartan-query-gov-params)                | Query the parameters of the governance process                               |
| Query | [proposal](#spartan-query-gov-proposal)            | Query details of a single proposal                                           |
| Query | [proposals](#spartan-query-gov-proposals)          | Query proposals with optional filter                                         |
| Qeury | [proposer](#spartan-query-gov-properser)           | Query the proposer of a governance proposal                                  |
| Query | [vote](#spartan-query-gov-vote)                    | Query details of a single vote                                               |
| Query | [votes](#spartan-query-gov-votes)                  | Query votes on a proposal                                                    |
| Query | [tally](#spartan-query-gov-tally)                  | Get the tally of a proposal vote                                             |

## spartan tx gov deposit

Submit a deposit for an active proposal.

```bash
spartan tx gov deposit [proposal-id] [deposit] [flags]
```

## spartan tx gov submit-proposal

Submit a proposal along with an initial deposit. Proposal title, description, type and deposit can be given directly or through a proposal JSON file.

**Example**

```
$ spartan tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey

Where proposal.json contains:

{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10test"
}
```

### Available Commands

| Type | Name                                                                               | Description                                  |
| ---- | ---------------------------------------------------------------------------------- | -------------------------------------------- |
| Tx   | [cancel-software-upgrade](#spartan-tx-gov-submit-proposal-cancel-software-upgrade) | Cancel the current software upgrade proposal |
| Tx   | [create-validator](#spartan-tx-gov-submit-proposal-create-validator)               | Submit a validator create proposal           |
| Tx   | [param-change ](#spartan-tx-gov-submit-proposal-param-change)                      | Submit a parameter change proposal           |
| Tx   | [remove-validator](#spartan-tx-gov-submit-proposal-remove-validator)               | Submit a validator remove proposal           |
| Tx   | [software-upgrade](#spartan-tx-gov-submit-proposal-software-upgrade)               | Submit a software upgrade proposal           |
| Tx   | [update-validator](#spartan-tx-gov-submit-proposal-update-validator)               | Submit a validator update proposal           |

### spartan tx gov submit-proposal cancel-software-upgrade

Cancel a software upgrade along with an initial deposit.

```bash
spartan tx gov submit-proposal cancel-software-upgrade [flags]
```

### spartan tx gov submit-proposal software-upgrade

Submit a software upgrade along with an initial deposit.

Please specify a unique name and height for the upgrade to take effect.

You may include info to reference a binary download link, in a format compatible with: <https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor>

```bash
spartan tx gov submit-proposal software-upgrade [name] (--upgrade-height [height]) (--upgrade-info [info]) [flags]
```

### spartan tx gov submit-proposal param-change

Submit a parameter proposal along with an initial deposit.

The proposal details must be supplied via a JSON file. For values that containsobjects, only non-empty fields will be updated.

IMPORTANT: Currently parameter changes are evaluated but not validated, so it isvery important that any "value" change is valid (ie. correct type and within bounds) for its respective parameter, eg. "MaxValidators" should be an integer and not a decimal.

Proper vetting of a parameter change proposal should prevent this from happening(no deposits should occur during the governance process), but it should be noted regardless.

```bash
spartan tx gov submit-proposal param-change [proposal-file] [flags]
```

Where proposal.json contains:

```json
{
  "title": "Staking Param Change",
  "description": "Update max validators",
  "changes": [
    {
      "subspace": "staking",
      "key": "MaxValidators",
      "value": 105
    }
  ],
  "deposit": "1000stake"
}
```

### spartan tx gov submit-proposal create-validator

Submit a proposal along with an initial deposit.

```bash
spartan tx gov submit-proposal create-validator [file] [flags]
```

Where proposal.json contains:

```json
{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"name":"node1",
		"pubkey":"icp1zcjduepq0c8s6tgqy77emkxv5jkw4eu99ggwk5uyqty6vcxuvn8cvr3j8pkqlxnk76",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

```

### spartan tx gov submit-proposal remove-validator

Submit a proposal along with an initial deposit.

```bash
spartan tx gov submit-proposal remove-validator [file] [flags]
```

Where proposal.json contains:

```json
{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"xxxxxx",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}
```

### spartan tx gov submit-proposal update-validator

Submit a proposal along with an initial deposit.

```bash
spartan tx gov submit-proposal update-validator [file] [flags]
```

Where proposal.json contains:

```json
{
	"deposit":"1000usnmt",
	"content":{
		"title":"proposal title",
		"summary":"proposal description",
		"id":"xxxxxx",
		"name":"node1",
	    "power":"100",
	    "description":"my node1",
	    "operator":"iaa104hrdtdkk5lfh8c3nc3pf20ad0sgdvselg0vxs"
	}
}

```

## spartan tx gov vote

Submit a vote for an active proposal.

```bash
spartan tx gov vote [proposal-id] [option] [flags]
```

## spartan tx gov weighted-vote

Submit a vote for an active proposal.

```bash
spartan tx gov weighted-vote [proposal-id] [weighted-options] [flags]
```

## spartan query gov deposit

Query details for a single proposal deposit on a proposal by its identifier.

```bash
spartan query gov deposit [proposal-id] [depositer-addr] [flags]
```

## spartan query gov deposits

Query details for all deposits on a proposal.

```bash
spartan query gov deposits [proposal-id] [flags]
```

## spartan query gov param

Query the parameters (voting|tallying|deposit) of the governance process.

```bash
spartan query gov param [param-type] [flags]
```

## spartan query gov params

Query the all the parameters for the governance process.

```bash
spartan query gov params [flags]
```

## spartan query gov proposal

Query details for a proposal.

```bash
spartan query gov proposal [proposal-id] [flags]
```

## spartan query gov proposals

Query for a all paginated proposals that match optional filter.

```bash
spartan query gov proposals [flags]
```

### Query proposals by conditions

```bash
spartan query gov proposals --limit=3 --status=Passed --depositor=<iaa...>
```

## spartan query gov proposer

Query which address proposed a proposal with a given ID.

```bash
spartan query gov proposer [proposal-id] [flags]
```

## spartan query gov vote

Query details for a single vote on a proposal given its identifier.

```bash
spartan query gov vote [proposal-id] [voter-addr] [flags]
```

## spartan query gov votes

Query vote details for a single proposal by its identifier.

```bash
spartan query gov votes [proposal-id] [flags]
```

## spartan query gov tally

Query tally of votes on a proposal.

```bash
spartan query gov tally [proposal-id] [flags]
```
