<!--
parent:
  order: false
-->

<div align="center">
  <h1> ethos </h1>
</div>

Thanks to evmos for bringing evm into the world of cosmos. we respect the great work of evmos very much.

Since evmos has modified the license, commercial use of versions above v13.x requires permission from the evmos team. But I really like the related functions of pre-compiled contracts, I decided based on the LGPLv3 license of evmos v12.1.6, remove some modules that I felt were no longer needed, such as recovery, incentives, and claims, and then implement the functions related to pre-compiled contracts. 

My goal is to implement all modules such as staking, gov, and distribution using precompiled contracts. Since evmos has implemented some related pre-compiled contract functions, in order to avoid legal risks, I will refer to the project [fx-core](https://github.com/FunctionX/fx-core/) whose license is Apache License 2.0 to implement all pre-compiled contract related functions. 

more details see [Evmos Signals Ethereum Alignment By Fully Deprecating Cosmos Transactions](https://medium.com/evmos/evmos-signals-ethereum-alignment-by-fully-deprecating-cosmos-transactions-d1f92c2cd443)

<div align="center">
  <h1> Evmos </h1>
</div>

Evmos is a scalable, high-throughput Proof-of-Stake blockchain
that is fully compatible and interoperable with Ethereum.
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
which runs on top of the [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

## Documentation

Our documentation is hosted in a [separate repository](https://github.com/evmos/docs) and can be found at [docs.evmos.org](https://docs.evmos.org).
Head over there and check it out.

**Note**: Requires [Go 1.20+](https://golang.org/dl/)

## Installation

For prerequisites and detailed build instructions
please read the [Installation](https://docs.evmos.org/protocol/evmos-cli) instructions.
Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/evmos/evmos/releases).

## Quick Start

To learn how the Evmos works from a high-level perspective,
go to the [Protocol Overview](https://docs.evmos.org/protocol) section from the documentation.
You can also check the instructions to [Run a Node](https://docs.evmos.org/protocol/evmos-cli#run-an-evmos-node).

## Community

The following chat channels and forums are a great spot to ask questions about Evmos:

- [Evmos Twitter](https://twitter.com/EvmosOrg)
- [Evmos Discord](https://discord.gg/evmos)
- [Evmos Forum](https://commonwealth.im/evmos)

## Contributing

Looking for a good place to start contributing?
Check out some
[`good first issues`](https://github.com/evmos/evmos/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

## Careers

See our open positions on [Greenhouse](https://boards.eu.greenhouse.io/evmos).
