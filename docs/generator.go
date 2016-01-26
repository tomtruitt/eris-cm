package main

import (
	"fmt"
	"os"

	commands "github.com/eris-ltd/eris-cm/chain_manager"
	"github.com/eris-ltd/eris-cm/version"

	"github.com/eris-ltd/eris-cm/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

var RENDER_DIR = fmt.Sprintf("./docs/eris-cm/%s/", version.VERSION)

var SPECS_DIR = "./docs/"

var BASE_URL = fmt.Sprintf("https://docs.erisindustries.com/documentation/eris-cm/%s/", version.VERSION)

const FRONT_MATTER = `---

layout:     documentation
title:      "Documentation | eris:chain_manager | {{}}"

---

`

func main() {
	os.MkdirAll(RENDER_DIR, 0755)
	epm := commands.MakerCmd
	commands.InitErisChainMaker()
	commands.AddGlobalFlags()
	specs := common.GenerateSpecs(SPECS_DIR, RENDER_DIR, FRONT_MATTER)
	common.GenerateTree(epm, RENDER_DIR, specs, FRONT_MATTER, BASE_URL)
}
