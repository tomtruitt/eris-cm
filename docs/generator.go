package main

import (
	"fmt"
	"os"

	commands "github.com/eris-ltd/eris-chainmaker/chainmaker"
	"github.com/eris-ltd/eris-chainmaker/version"

	"github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

var RENDER_DIR = fmt.Sprintf("./docs/eris-chainmaker/%s/", version.VERSION)

var SPECS_DIR = "./docs/"

var BASE_URL = fmt.Sprintf("https://docs.erisindustries.com/documentation/eris-chainmaker/%s/", version.VERSION)

const FRONT_MATTER = `---

layout:     documentation
title:      "Documentation | eris:chainmaker | {{}}"

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
