<!--
order: 1
-->

# Joining Mainnet

This document outlines the steps to join an existing testnet {synopsis}

## Pre-requisite Readings

- [Validator Security](./security.md) {prereq}

## Mainnet

You need to set the **genesis file** and **seeds**. If you need more information about past networks, check our [mainnet repo](https://github.com/tharsis/mainnet).

| Chain ID       | Description   | Site                                                               | Version                                               | Status     |
| -------------- | ------------- | ------------------------------------------------------------------ | ----------------------------------------------------- | ---------- |
| `enron_9001-2` | Enron Mainnet 2 | [Enron](https://github.com/tharsis/mainnet/tree/main/enron_9001-2) | [`v3.0.x`](https://github.com/enronchain/enron/releases) | `Not Live` |
| `enron_9001-1` | Enron Mainnet 1 | [Enron](https://github.com/tharsis/mainnet/tree/main/enron_9001-1) | [`v2.0.1`](https://github.com/enronchain/enron/releases/v2.0.1) | `Stale` |

## Install `enrond`

Follow the [installation](./../quickstart/installation) document to install the {{ $themeConfig.project.name }} binary `{{ $themeConfig.project.binary }}`.

:::warning
Make sure you have the right version of `{{ $themeConfig.project.binary }}` installed.
:::

### Save Chain ID

We recommend saving the mainnet `chain-id` into your `{{ $themeConfig.project.binary }}`'s `client.toml`. This will make it so you do not have to manually pass in the `chain-id` flag for every CLI command.

::: tip
See the Official [Chain IDs](./../basics/chain_id.md#official-chain-ids) for reference.
:::

```bash
enrond config chain-id enron_9001-2
```

## Initialize Node

We need to initialize the node to create all the necessary validator and node configuration files:

```bash
enrond init <your_custom_moniker> --chain-id enron_9001-2
```

::: danger
Monikers can contain only ASCII characters. Using Unicode characters will render your node unreachable.
:::

By default, the `init` command creates your `~/.enrond` (i.e `$HOME`) directory with subfolders `config/` and `data/`.
In the `config` directory, the most important files for configuration are `app.toml` and `config.toml`.

## Genesis & Seeds

### Copy the Genesis File

Download the `genesis.json` file from the [`archive`](https://archive.enron.dev/genesis/genesis.json) and copy it over to the `config` directory: `~/.enrond/config/genesis.json`. This is a genesis file with the chain-id and genesis accounts balances.

```bash
wget https://archive.enron.dev/genesis/genesis.json
mv genesis.json ~/.enrond/config/
```

Then verify the correctness of the genesis configuration file:

```bash
enrond validate-genesis
```

### Add Seed Nodes

Your node needs to know how to find [peers](https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#peers). You'll need to add healthy [seed nodes](https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#seed) to `$HOME/.enrond/config/config.toml`. The [`mainnet`](https://github.com/tharsis/mainnet) repo contains links to some seed nodes.

Edit the file located in `~/.enrond/config/config.toml` and the `seeds` to the following:

```toml
#######################################################
###           P2P Configuration Options             ###
#######################################################
[p2p]

# ...

# Comma separated list of seed nodes to connect to
seeds = "<node-id>@<ip>:<p2p port>"
```

You can use the following code to get seeds from the repo and add it to your config:

```bash
SEEDS=`curl -sL https://raw.githubusercontent.com/tharsis/mainnet/main/enron_9001-2/seeds.txt | awk '{print $1}' | paste -s -d, -`
sed -i.bak -e "s/^seeds =.*/seeds = \"$SEEDS\"/" ~/.enrond/config/config.toml
```

:::tip
For more information on seeds and peers, you can the Tendermint [P2P documentation](https://docs.tendermint.com/master/spec/p2p/peer.html).
:::

### Add Persistent Peers

We can set the [`persistent_peers`](https://docs.tendermint.com/master/tendermint-core/using-tendermint.html#persistent-peer) field in `~/.enrond/config/config.toml` to specify peers that your node will maintain persistent connections with. You can retrieve them from the list of
available peers on the [`mainnet`](https://github.com/tharsis/mainnet) repo.

A list of available persistent peers is also available in the `#find-peers` channel in the [Enron Discord](https://discord.gg/enron). You can get a random 10 entries from the `peers.txt` file in the `PEERS` variable by running the following command:

```bash
PEERS=`curl -sL https://raw.githubusercontent.com/tharsis/mainnet/main/enron_9001-2/peers.txt | sort -R | head -n 10 | awk '{print $1}' | paste -s -d, -`
```

Use `sed` to include them into the configuration. You can also add them manually:

```bash
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" ~/.enrond/config/config.toml
```

## Run a Mainnet Validator

::: tip
For more details on how to configure your validator, follow the validator [setup](./../guides/validators/setup.md) instructions.
:::

```bash
enrond tx staking create-validator \
  --amount=1000000000000aenron \
  --pubkey=$(enrond tendermint show-validator) \
  --moniker="EnronWhale" \
  --chain-id=<chain_id> \
  --commission-rate="0.05" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="0.025aenron" \
  --from=<key_name>
```

::: danger
🚨 **DANGER**: <u>Never</u> create your validator keys using a [`test`](./../guides/keys-wallets/keyring.md#testing) keyring backend. Doing so might result in a loss of funds by making your funds remotely accessible via the `eth_sendTransaction` JSON-RPC endpoint.

Ref: [Security Advisory: Insecurely configured geth can make funds remotely accessible](https://blog.ethereum.org/2015/08/29/security-alert-insecurely-configured-geth-can-make-funds-remotely-accessible/)
:::

## Start mainnet

The final step is to [start the nodes](./../quickstart/run_node#start-node). Once enough voting power (+2/3) from the genesis validators is up-and-running, the node will start producing blocks.

```bash
enrond start
```

## Upgrading Your Node

::: tip
These instructions are for full nodes that have ran on previous versions of and would like to upgrade to the latest mainnet version.
:::

### Reset Data

::: warning
If the version <new_version> you are upgrading to is not breaking from the previous one, you **should not** reset the data. If this is the case you can skip to [Restart](#restart)
:::

First, remove the outdated files and reset the data.

```bash
rm $HOME/.enrond/config/addrbook.json $HOME/.enrond/config/genesis.json
enrond unsafe-reset-all
```

Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before,
your node will still try to connect to them, but may fail if they haven't also
been upgraded.

::: danger Warning
Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to double sign.
:::

### Restart

To restart your node, just type:

```bash
enrond start
```

## Share your Peer

You can share your peer to posting it in the `#find-peers` channel in the [Enron Discord](https://discord.gg/enron).

::: tip
To get your Node ID use

```bash
enrond tendermint show-node-id
```

:::
