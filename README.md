<!--
parent:
  order: false
-->

<div align="center">
  <img src="https://" />
  <h1> Enron </h1>
</div>

<!-- TODO: add banner -->
<!-- ![banner](docs/ethermint.jpg) -->

<div align="center">
  <a href="https://github.com/enronchain/enron/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/enronchain/enron.svg" />
  </a>
  <a href="https://github.com/enronchain/enron/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/enronchain/enron.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/enronchain/enron">
    <img alt="GoDoc" src="https://godoc.org/github.com/enronchain/enron?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/enronchain/enron">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/enronchain/enron"/>
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/enron">
    <img alt="Discord" src="https://img.shields.io/discord/962917488180490250.svg" />
  </a>
  <a href="https://github.com/enronchain/enron/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/enronchain/enron/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/enronchain/enron">
    <img alt="Code Coverage" src="https://codecov.io/gh/enronchain/enron/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/EnronNetwork">
    <img alt="Twitter Follow Enron" src="https://img.shields.io/twitter/follow/EnronNetwork"/>
  </a>
</div>

Enron is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum and Cosmos. It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

**Note**: Requires [Go 1.17.5+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions please read the [Installation](https://docs.enron.network) instructions. Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/enronchain/enron/releases).

## Genesis
To get onto our mainnet genesis download the genesis.json here

`wget https://gist.githubusercontent.com/EnronChain/52d111aff20fda340370ad81d69001c4/raw/48514a80cdf5e679e110d074fa3b77e987e3a090/genesis.json`

and then move it into the enrond config (after you have initilized your node)

`mv genesis.json ~/.enrond/config/`

## Quick Start

To learn how the Enron works from a high-level perspective, go to the [Introduction](https://docs.enron.network) section from the documentation. You can also check the instructions to [Run a Node](https://docs.enron.network).

## Community

The following chat channels and forums are a great spot to ask questions about Enron:

- [Enron Twitter](https://twitter.com/EnronNetwork)
- [Enron Discord](https://discord.gg/enron)
- [Enron Telegram](https://t.me/enronnetwork)

## Contributing

Looking for a good place to start contributing? Check out some [`good first issues`](https://github.com/enronchain/enron/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.
