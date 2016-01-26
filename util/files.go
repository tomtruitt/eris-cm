package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/eris-ltd/eris-chainmaker/definitions"
	"github.com/eris-ltd/eris-chainmaker/version"

	"github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/BurntSushi/toml"
	log "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	. "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

// ensures that the files which are included in this repository (`defaultTyps`) are also
// present in the user's .eris/chains/account_types directory.
//
// does not ensure that the contents of the files are the same so will not affect user
// defined settings around these files.
//
// does not check if the user has more account_types files in their .eris/chains/account_types
// directory either so users can safely add additional account_types beyond the marmot
// established defaults.
func CheckDefaultTypes(erisPath, myPath string) error {
	defaultTyps, err := filepath.Glob(filepath.Join(ErisGo, version.NAME, myPath, "*.toml"))
	if err != nil {
		return err
	}

	haveTyps, err := AccountTypesNames(erisPath, true)
	if err != nil {
		return err
	}

	for _, file := range defaultTyps {
		f := filepath.Base(file)
		itsThere := false

		// check if present
		for _, b := range haveTyps {
			if f == b {
				itsThere = true
			}
		}

		if !itsThere {
			Copy(file, filepath.Join(erisPath, f))
		}
	}

	return nil
}

// returns a list of filenames which are the account_types files
// these *should be* absolute paths, but this is not a contract
// with calling functions.
func AccountTypes(erisPath string) ([]string, error) {
	haveTyps, err := filepath.Glob(filepath.Join(erisPath, "*.toml"))
	if err != nil {
		return []string{}, err
	}
	return haveTyps, nil
}

func AccountTypesNames(erisPath string, withExt bool) ([]string, error) {
	files, err := AccountTypes(erisPath)
	if err != nil {
		return []string{}, err
	}
	names := []string{}
	for _, file := range files {
		names = append(names, filepath.Base(file))
	}
	if !withExt {
		for e, name := range names {
			names[e] = strings.Replace(name, ".toml", "", 1)
		}
	}
	return names, nil
}

func WriteGenesisFile(name string, genesis *definitions.MintGenesis, account *definitions.Account, single bool) error {
	return writer(genesis, name, account.Name, "genesis.json", single)
}

func WritePrivVals(name string, account *definitions.Account, single bool) error {
	return writer(account.MintKey, name, account.Name, "priv_validator.json", single)
}

func SaveAccountType(thisActT *definitions.AccountType) error {
	writer, err := os.Create(filepath.Join(AccountsTypePath, fmt.Sprintf("%s.toml", thisActT.Name)))
	defer writer.Close()
	if err != nil {
		return err
	}

	enc := toml.NewEncoder(writer)
	enc.Indent = ""
	err = enc.Encode(thisActT)
	if err != nil {
		return err
	}
	return nil
}

func writer(toWrangle interface{}, chain_name, account_name, fileBase string, single bool) error {
	var file string
	fileBytes, err := json.MarshalIndent(toWrangle, "", "  ")
	if err != nil {
		return err
	}
	if !single {
		file = filepath.Join(ChainsPath, chain_name, account_name, fileBase)
	} else {
		file = filepath.Join(ChainsPath, chain_name, fileBase)
	}
	log.WithField("path", file).Debug("Saving File.")
	err = WriteFile(string(fileBytes), file)
	if err != nil {
		return err
	}
	return nil
}
