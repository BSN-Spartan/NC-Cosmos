# Upgrade

This Upgrade module provides the basic functions of software version upgrade.

**NOTE:** Upgrade is only allowed by submitting a proposal throng the [Gov](./gov.md) module,  so here we only introduces query commands.

## Available Commands

| Type  | Name                                                   | Description                                                      |
| ----- | ------------------------------------------------------ | ---------------------------------------------------------------- |
| Query | [applied](#spartan-tx-upgrade-applied)                 | Block header for height at which a completed upgrade was applied |
| Query | [module_versions](#spartan-tx-upgrade-module_versions) | Get the list of module versions                                  |
| Query | [plan](#spartan-tx-upgrade-plan)                       | Get upgrade plan if one exists                                   |

## spartan tx upgrade applied

If upgrade-name was previously executed on the chain, this returns the header for the block at which it was applied.

This helps a client determine which binary was valid over a given range of blocks, as well as more context to understand past migrations.

```bash
spartan query upgrade applied [upgrade-name] [flags]
```

## spartan tx upgrade module_versions

Gets a list of module names and their respective consensus versions.

Following the command with a specific module name will return only that module's information.

```bash
spartan query upgrade module_versions [optional module_name] [flags]
```

## spartan tx upgrade plan

Gets the currently scheduled upgrade plan, if one exists

```bash
spartan query upgrade plan [flags]
```

