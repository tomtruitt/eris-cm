package db

import (
	cfg "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/tendermint/tendermint/config"
)

var config cfg.Config = nil

func init() {
	cfg.OnConfig(func(newConfig cfg.Config) {
		config = newConfig
	})
}
