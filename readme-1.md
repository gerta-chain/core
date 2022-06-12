<p>&nbsp;</p>
<p align="center">

<img src="core_logo.svg" width=500>

</p>

<p align="center">
The full-node software implementation of the Gerta blockchain.<br/><br/>

<a href="https://codecov.io/gh/gerta-chain/core">
    <img src="https://codecov.io/gh/gerta-chain/core/branch/main/graph/badge.svg">
</a>
<a href="https://goreportcard.com/report/github.com/gerta-chain/core">
    <img src="https://goreportcard.com/badge/github.com/gerta-chain/core">
</a>

</p>

<p align="center">
  <a href="https://docs.gerta.money/"><strong>Explore the Docs »</strong></a>
  <br />
  <br/>
  <a href="https://docs.gerta.money/docs/develop/module-specifications/README.html">Gerta Core reference</a>
  ·
  <a href="https://pkg.go.dev/github.com/gerta-chain/core?tab=subdirectories">Go API</a>
  ·
  <a href="https://lcd.gerta.dev/swagger/#/">Rest API</a>
  ·
  <a href="https://finder.gerta.money/">Finder</a>
  ·
  <a href="https://station.gerta.money/">Station</a>
</p>

<br/>

## Gerta migration guides

Visit the [migration guide](https://migrate.gerta.money) to learn how to migrate from Gerta Classic to the new Gerta blockchain.

## Table of Contents <!-- omit in toc -->

- [What is Gerta?](#what-is-gerta)
- [Installation](#installation)
  - [From Binary](#from-binary)
  - [From Source](#from-source)
  - [gertad](#gertad)
- [Node Setup](#node-setup)
  - [Gerta node quickstart](#gerta-node-quickstart)
  - [Join the mainnet](#join-the-mainnet)
  - [Join a testnet](#join-a-testnet)
  - [Run a local testnet](#run-a-local-testnet)
  - [Run a single node testnet](#run-a-single-node-testnet)
- [Set up a production environment](#set-up-a-production-environment)
  - [Increase maximum open files](#increase-maximum-open-files)
  - [Create a dedicated user](#create-a-dedicated-user)
  - [Port configuration](#port-configuration)
  - [Run the server as a daemon](#run-the-server-as-a-daemon)
  - [Register gertad as a service](#register-gertad-as-a-service)
  - [Start, stop, or restart service](#start-stop-or-restart-service)
  - [Access logs](#access-logs)
- [Resources](#resources)
- [Community](#community)
- [Contributing](#contributing)
- [License](#license)

## What is Gerta?

**[Gerta](https://gerta.money)** is a public, open-source, proof-of-stake blockchain. **The Gerta Core** is the reference implementation of the Gerta protocol written in Golang. The Gerta Core is powered by the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) and [Tendermint](https://github.com/tendermint/tendermint) BFT consensus.

## Installation

### From Binary

The easiest way to install the Gerta Core is to download a pre-built binary. You can find the latest binaries on the [releases](https://github.com/gerta-chain/core/releases) page.

### From Source

**Step 1: Install Golang**

Go v1.18+ or higher is required for The Gerta Core.

1. Install [Go 1.18+ from the official site](https://go.dev/dl/). Ensure that your `GOPATH` and `GOBIN` environment variables are properly set up by using the following commands:

   For Windows:

   ```sh
   wget <https://golang.org/dl/go1.18.2.linux-amd64.tar.gz>
   sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

   For Mac:

   ```sh
   export PATH=$PATH:/usr/local/go/bin
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

2. Confirm your Go installation by checking the version:

   ```sh
   go version
   ```


**Step 2: Get Gerta Core source code**

Clone the Gerta Core from the [official repo](https://github.com/gerta-chain/core/) and check out the `main` branch for the latest stable release.

```bash
git clone https://github.com/gerta-chain/core/
cd core
git checkout main
```

**Step 3: Build Gerta core**

Run the following command to install `gertad` to your `GOPATH` and build the Gerta Core. `gertad` is the node daemon and CLI for interacting with a Gerta node.

```bash
# COSMOS_BUILD_OPTIONS=rocksdb make install
make install
```

**Step 4: Verify your installation**

Verify your installation with the following command:

```bash
gertad version --long
```

A successful installation will return the following:

```bash
name: gerta
server_name: gertad
version: <x.x.x>
commit: <Commit hash>
build_tags: netgo,ledger
go: go version go1.18.2 darwin/amd64
```

## `gertad`

`gertad` is the all-in-one CLI and node daemon for interacting with the Gerta blockchain. 

To view various subcommands and their expected arguments, use the following command:

``` sh
$ gertad --help
```


```
Stargate Gerta App

Usage:
  gertad [command]

Available Commands:
  add-genesis-account Add a genesis account to genesis.json
  collect-gentxs      Collect genesis txs and output a genesis.json file
  debug               Tool for helping with debugging your application
  export              Export state to JSON
  gentx               Generate a genesis tx carrying a self delegation
  help                Help about any command
  init                Initialize private validator, p2p, genesis, and application configuration files
  keys                Manage your application's keys
  migrate             Migrate genesis to a specified target version
  query               Querying subcommands
  rosetta             spin up a rosetta server
  start               Run the full node
  status              Query remote node for status
  tendermint          Tendermint subcommands
  testnet             Initialize files for a gertad testnet
  tx                  Transactions subcommands
  unsafe-reset-all    Resets the blockchain database, removes address book files, and resets data/priv_validator_state.json to the genesis state
  validate-genesis    validates the genesis file at the default location or at the location passed as an arg
  version             Print the application binary version information

Flags:
  -h, --help                help for gertad
      --home string         directory for config and data (default "/Users/evan/.gerta")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors

Use "gertad [command] --help" for more information about a command.
```

Visit the [gertad documentation page](https://docs.gerta.money/docs/develop/how-to/gertad/README.html) for more info on usage. 


## Node Setup

Once you have `gertad` installed, you will need to set up your node to be part of the network.

### Join the mainnet

The following requirements are recommended for running a mainnet node:

- Four or more CPU cores
- At least 32 GB of memory
- At least 300 mbps of network bandwidth
- At least 2 TB NVME SSD
- A Linux distribution

### Gerta node quickstart
```
gertad init nodename
wget -O ~/.gerta/config/genesis.json https://cloudflare-ipfs.com/ipfs/QmZAMcdu85Qr8saFuNpL9VaxVqqLGWNAs72RVFhchL9jWs
curl https://network.gerta.dev/addrbook.json > ~/.gertad/config/addrbook.json
gertad start
```

### Join a testnet

Several testnets might exist simultaneously. Ensure that your version of `gertad` is compatible with the network you want to join.

To set up a node on the latest testnet, visit [the testnet repo](https://github.com/gerta-chain/testnet).

#### Run a local testnet

The easiest way to set up a local testing environment is to run [LocalGerta](https://github.com/gerta-chain/LocalGerta), a zero-configuration complete testing environment. 

### Run a single node testnet

You can also run a local testnet using a single node. On a local testnet, you will be the sole validator signing blocks.

**Step 1: Create network and account**

First, initialize your genesis file to bootstrap your network. Create a name for your local testnet and provide a moniker to refer to your node:

```bash
gertad init --chain-id=<testnet_name> <node_moniker>
```

Next, create a Gerta account by running the following command:

```bash
gertad keys add <account_name>
```

**Step 2: Add account to genesis**

Add your account to genesis and set an initial balance to start. Run the following commands to add your account and set the initial balance:

```bash
gertad add-genesis-account $(gertad keys show <account_name> -a) 100000000uluna
gertad gentx <account_name> 10000000uluna --chain-id=<testnet_name>
gertad collect-gentxs
```

**Step 3: Run gertad**

Now you can start your private Gerta network:

```bash
gertad start
```

Your `gertad` node will be running a node on `tcp://localhost:26656`, listening for incoming transactions and signing blocks.

## Set up a production environment

**Note**: This guide only covers general settings for a production-level full node. Visit the [Gerta validator's guide](https://docs.gerta.money/docs/full-node/manage-a-gerta-validator/README.html) for more information.

**This guide has been tested against Linux distributions only. To ensure you successfully set up your production environment, consider setting it up on an Linux system.**

### Increase maximum open files

By default, `gertad` can't open more than 1024 files at once.

You can increase this limit by modifying `/etc/security/limits.conf` and raising the `nofile` capability.

```
*                soft    nofile          65535
*                hard    nofile          65535
```

### Create a dedicated user

It is recommended that you run `gertad` as a normal user. Super-user accounts are only recommended during setup to create and modify files.

### Port configuration

`gertad` uses several TCP ports for different purposes.

- `26656`: The default port for the P2P protocol. Use this port to communicate with other nodes. While this port must be open to join a network, it does not have to be open to the public. Validator nodes should configure `persistent_peers` and close this port to the public.

- `26657`: The default port for the RPC protocol. This port is used for querying / sending transactions and must be open to serve queries from `gertad`. **DO NOT** open this port to the public unless you are planning to run a public node.

- `1317`: The default port for [Lite Client Daemon](https://docs.gerta.money/docs/develop/how-to/start-lcd.html) (LCD), which can be enabled in `~/.gerta/config/app.toml`. The LCD provides an HTTP RESTful API layer to allow applications and services to interact with your `gertad` instance through RPC. Check the [Gerta REST API](https://lcd.gerta.dev/swagger/#/) for usage examples. Don't open this port unless you need to use the LCD.

- `26660`: The default port for interacting with the [Prometheus](https://prometheus.io) database. You can use Promethues to monitor an environment. This port is closed by default.

### Run the server as a daemon

**Important**:

Keep `gertad` running at all times. The simplest solution is to register `gertad` as a `systemd` service so that it automatically starts after system reboots and other events.


### Register gertad as a service

First, create a service definition file in `/etc/systemd/system`.

**Sample file: `/etc/systemd/system/gertad.service`**

```
[Unit]
Description=Gerta Daemon
After=network.target

[Service]
Type=simple
User=gerta
ExecStart=/data/gerta/go/bin/gertad start
Restart=on-abort

[Install]
WantedBy=multi-user.target

[Service]
LimitNOFILE=65535
```

Modify the `Service` section from the given sample above to suit your settings.
Note that even if you raised the number of open files for a process, you still need to include `LimitNOFILE`.

After creating a service definition file, you should execute `systemctl daemon-reload`.

### Start, stop, or restart service

Use `systemctl` to control (start, stop, restart)

```bash
# Start
systemctl start gertad
# Stop
systemctl stop gertad
# Restart
systemctl restart gertad
```

### Access logs

```bash
# Entire log
journalctl -t gertad
# Entire log reversed
journalctl -t gertad -r
# Latest and continuous
journalctl -t gertad -f
```

## Resources

Developer Tools:
  - [Gerta docs](https://docs.gerta.money): Developer documentation.
  - [Faucet](https://faucet.gerta.money): Get testnet Luna.
  - [LocalGerta](https://www.github.com/gerta-chain/LocalGerta): A dockerized local blockchain testnet. 

Developer Forums: 
  - [Gerta Developer Discord](https://discord.com/channels/464241079042965516/591812948867940362)
  - [Gerta Developer Telegram room](https://t.me/+gCxCPohmVBkyNDRl)

Block Explorer:
  - [Gerta Finder](https://finder.gerta.money): Gerta's basic block explorer.

Wallets:
  - [Gerta Station](https://station.gerta.money): The official Gerta wallet.
  - Gerta Station Mobile:
    - [iOS](https://apps.apple.com/us/app/gerta-station/id1548434735)
    - [Android](https://play.google.com/store/apps/details?id=money.gerta.station&hl=en_US&gl=US)

Research:
  - [Agora](https://agora.gerta.money): Research forum

## Community

- [Offical Website](https://gerta.money)
- [Discord](https://discord.gg/e29HWwC2Mz)
- [Telegram](https://t.me/gerta_announcements)
- [Twitter](https://twitter.com/gerta_money)
- [YouTube](https://goo.gl/3G4T1z)

## Contributing

If you are interested in contributing to Gerta Core source, please review our [code of conduct](./CODE_OF_CONDUCT.md).

## License

This software is [licensed under the Apache 2.0 license](LICENSE).

© 2022 Gertaform Labs, PTE LTD

<p>&nbsp;</p>
<p align="center">
    <a href="https://gerta.money/"><img src="https://assets.website-files.com/611153e7af981472d8da199c/61794f2b6b1c7a1cb9444489_symbol-gerta-blue.svg" align="center" width=200/></a>
</p>
<div align="center">
</div>


