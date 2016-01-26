[![Circle CI](https://circleci.com/gh/eris-ltd/eris-cm/tree/master.svg?style=svg)](https://circleci.com/gh/eris-ltd/eris-cm/tree/master)

[![GoDoc](https://godoc.org/github.com/eris-cm?status.png)](https://godoc.org/github.com/eris-ltd/eris-cm)

# Eris Chain Manager

```
The Eris Chain Manager is a utility for performing complex operations on eris chains
```

`eris:chain_manager` is a set of high-level tooling for working with `eris chains`. It is similar in nature, design, and level as the `eris:package_manager` which is built to handle smart contract packages and other packages necessary for building blockchain backed applications.

# Installation

Usually this repository should be run in a docker container via the [eris-cli](https://docs.erisindustries.com/tutorials/). Should you want/desire/need to install this repository natively on your host make sure you have go installed and then:

```
go get github.com/eris-ltd/eris-cm/cmd/eris-cm
```

# Overview

Currently `eris:chain_manager` provides the following functionality:

* `maker` -- wizard, config file, or csv based utility for creating keys and genesis.json's necessary for a range of different chain types.

(More functionality coming soon.)

#