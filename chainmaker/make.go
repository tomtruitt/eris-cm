package chains

import (
	"encoding/json"
	"os"

	"github.com/eris-ltd/eris-cli/definitions"
	"github.com/eris-ltd/eris-cli/util"

	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/code.google.com/p/go-uuid/uuid"
	log "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

var (
	tokens    int      = 9999999999
	rootGroup int      = 3
	devsGroup int      = 6
	valsGroup int      = 7
	partGroup int      = 25
	reader    *os.File = os.Stdin
)

type MintGenesis struct {
	chain_id   string           `json:"base"`
	accounts   []*MintAccount   `json:"accounts"`
	validators []*MintValidator `json:"validators"`
}

type MintAccount struct {
	address     string                  `json:"address"`
	amount      int                     `json:"amount"`
	name        string                  `json:"name"`
	permissions *MintAccountPermissions `json:"permissions"`
}

type MintValidator struct {
	PubKey   []interface{}   `json:"pub_key"`
	Name     string          `json:"name"`
	Amount   int             `json:"amount"`
	UnbondTo []*MintTxOutput `json:"unbond_to"`
}

type MintTxOutput struct {
	address string `json:"address"`
	amount  int    `json:"amount"`
}

func MakeChain(do *definitions.Do) error {
	return makeWizard(do)
}

func makeWizard(do *definitions.Do) error {
	proceed, err := util.GetBoolResponse(ChainsMakeWelcome(), true, os.Stdin)
	if err != nil {
		return err
	}

	if !proceed {
		log.Warn("The marmots will not proceed without your authorization. Exiting.")
		return nil
	}

	prelims := make(map[string]bool)
	for e, q := range ChainsMakePrelimQuestions() {
		prelims[e], err = util.GetBoolResponse(q, false, os.Stdin)
		if err != nil {
			return err
		}
	}

	if prelims["dryrun"] {
		return dryRun()
	}

	// get information for the root group
	rootGroup, rootGroupTokens, err := assembleGroup(ChainsMakeRoot(), ChainsMakeRootTokens(), rootGroup, prelims["tokens"])

	// get information for the developer group
	devsGroup, devsGroupTokens, err := assembleGroup(ChainsMakeDevelopers(), ChainsMakeDevelopersTokens(), devsGroup, prelims["tokens"])

	// get information for the validators group
	valsGroup, valsGroupTokens, err := assembleGroup(ChainsMakeValidators(), ChainsMakeValidatorsTokens(), valsGroup, prelims["tokens"])

	// get information for the participant group
	partGroup, partGroupTokens, err := assembleGroup(ChainsMakeParticipants(), ChainsMakeParticipantsTokens(), partGroup, prelims["tokens"])

	if !prelims["manual"] {
		return makeWithoutManual(do.Name, rootGroup, rootGroupTokens, devsGroup, devsGroupTokens, valsGroup, valsGroupTokens, partGroup, partGroupTokens)
	}

	return makeWithManual(do.Name, rootGroup, rootGroupTokens, devsGroup, devsGroupTokens, valsGroup, valsGroupTokens, partGroup, partGroupTokens)
}

func makeWithoutManual(name string, rootGroup, rootGroupTokens, devsGroup, devsGroupTokens, valsGroup, valsGroupTokens, partGroup, partGroupTokens int) error {
	genesis := &MintGenesis{}
	genesis.chain_id = name

	// assemble

	return writeFile(name, genesis)
}

func makeAccount() (string, *MintAccount, error) {
	account := &MintAccount{}
	// reset the writers....
	// keys :=
	doRoot := definitions.NowDo()
	doRoot.Name = strings.Split(uuid.New(), "-")[0]
	// doRoot.Service.Image =
}

func makeWithManual(name string, rootGroup, rootGroupTokens, devsGroup, devsGroupTokens, valsGroup, valsGroupTokens, partGroup, partGroupTokens int) error {

	return nil
}

func assembleGroup(question, tokenQuestion string, defalt int, tokenIze bool) (int, int, error) {
	groupNumber, err := util.GetIntResponse(question, defalt, reader)
	var groupTokens int = tokens
	if tokenIze {
		groupTokens, err = util.GetIntResponse(tokenQuestion, tokens, reader)
		if err != nil {
			return 0, 0, err
		}
	}
	return groupNumber, groupTokens, nil
}

func writeFile(name string, genesis *MintGenesis) error {
	fileBytes := json.MarshalIndent(genesis, "", "    ")
	// write
}

func dryRun() error {
	_, _ = util.GetIntResponse(ChainsMakeRoot(), 3, os.Stdin)
	_, _ = util.GetIntResponse(ChainsMakeDevelopers(), 6, os.Stdin)
	_, _ = util.GetIntResponse(ChainsMakeValidators(), 7, os.Stdin)
	_, _ = util.GetIntResponse(ChainsMakeParticipants(), 25, os.Stdin)
	_, _ = util.GetIntResponse(ChainsMakeManual(), 0, os.Stdin)
	return nil
}
