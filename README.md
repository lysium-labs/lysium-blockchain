# Lysium Network Blockchain


__The Lysium Network blockchain is an EVM-compatible blockchain that runs a version of PoU (Proof-of-Usage) consensus mechanism.__

The Lysium blockchain is a decentralized ledger of transactions. No entity is controlling a blockchain, there are only actors participating in its growth and integrity. Based on the Proof of Usage protocol, the system cannot be hacked, modified, or shut down by any of the users.

## Block generation
--- 

__The Lysium Network is a blockchain system that validates transactions through the issuance of blocks. A block contains a list of transactions that are confirmed by the network and verified by the validator that signs the block.__

An innovation of the Lysium Network is the implementation of 3 different types of blocks: Kilo (or normal) blocks, that are issued every 4 seconds, Mega blocks that are issued every 24 hours, and Giga blocks that are issued every 365 days.

## Native token
---
__LSX will run on Lysium Network Chain in the same way as ETH runs on Ethereum so that it remains as native token for Lysium Network. This means, LSX will be used to:__

- pay gas to deploy or invoke Smart Contract on Lysium Network Chain
- stake and receive rewards by becoming a validator on the network

## Proof of Usage (PoU)
---

__To incentivize power users and strategic partners to use our infrastructure for deploying their capital and distributed applications, we developed an algorithm that provides a scalable and low-fee system for the users of the network.__

The Proof-of-Usage (PoU) consensus is built directly into the blockchain infrastructure and the nodes block-building source-code. The validators (the nodes that are creating the blocks), select the best transactions to be included based on the `nonce` of the sender account (address).

## Scalability
---

__Scalability is a major differentiator for blockchain networks and one of the most pressing issues in terms of achieving mass-adoption for a network. Lysium fixes the scalability issue by implementing a system that can process 20,000 – 25,000 transactions and reward validators in a fair way.__

A key difference between Lysium Network and other decentralized networks is the consensus protocol. Over time, people have come to a false understanding that blockchains have to be slow and not scalable.

The Lysium Network protocol employs a novel approach to consensus to achieve its strong safety guarantees, quick finality, and high-throughput without compromising decentralization. Capable of 20,000 – 25,000 real transactions per second (up to 200,000 TPS with LSX 2.0 Upgrade) - an order of magnitude greater than existing blockchains (fastest EVM-Compatible Chain).

## Security
---

__The Lysium Network added a feature in which the node generates a hash of the state of the blockchain and publishes it on the public Bitcon blockchain. This is used in order to secure the network against blockchain rewrites and redeployments. By having the blockchain secured by the most trusted blockchain in the world (Bitcoin blockchain), gives the users and investors an enhanced security.__

The current state of the blockchain is hashed using `md5` cryptographic function. The resulted hash is then attached to a Bitcoin transaction with 0 value. This transaction is published on the Bitcoin network. This action occurs roughly every 30 days.

## Build the source-code
---

Building `geth` requires both a Go (version 1.14 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, run

`make geth`

## Run the blockchain
--- 

In order to easily run the blockchain, you need to download the `Dockerfile` and execute locally on your computer.
The server will sync automatically with the rest of the blockchain network.

## Command line tool

`geth` = The Lysium Network client (similar to Ethereum-go)

`bootnode` = Stripped down version of our Lysium Network client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.

`evm` = Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes



