package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eris-ltd/eris-chainmaker/chainmaker"
	"github.com/eris-ltd/eris-chainmaker/definitions"
	"github.com/eris-ltd/eris-chainmaker/util"
	"github.com/eris-ltd/eris-chainmaker/version"

	log "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	. "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	logger "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
	keys "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/eris-keys/eris-keys"
	"github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/spf13/cobra"
)

const VERSION = version.VERSION

// Global Do struct
var do *definitions.Do
var keysAddr string

// Defining the root command
var MakerCmd = &cobra.Command{
	Use:   "eris-chainmaker",
	Short: "The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains",
	Long: `The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains.

Made with <3 by Eris Industries.

Complete documentation is available at https://docs.erisindustries.com
` + "\nVersion:\n  " + VERSION,
	Example: `$ eris-chainmaker myChain -- will use the chain-making wizard and make your chain named myChain using eris-keys defaults (available via localhost) (interactive)
$ eris-chainmaker myChain --chain-type=simplechain --  will use the chain type definition files to make your chain named myChain using eris-keys defaults (non-interactive)
$ eris-chainmaker myChain --account-types=Root:1,Developer:0,Validator:0,Participant:1 -- will use the flag to make your chain named myChain using eris-keys defaults (non-interactive)
$ eris-chainmaker myChain --account-types=Root:1,Developer:0,Validator:0,Participant:1 --chain-type=simplechain -- account types trump chain types, this command will use the flags to make the chain (non-interactive)
$ eris-chainmaker myChain --csv /path/to/csv -- will use the csv file to make your chain named myChain using eris-keys defaults (non-interactive)`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// logger stuff
		log.SetFormatter(logger.ErisFormatter{})
		log.SetLevel(log.WarnLevel)
		if do.Verbose {
			log.SetLevel(log.InfoLevel)
		} else if do.Debug {
			log.SetLevel(log.DebugLevel)
		}

		// loop through chains directories to make sure they exist
		for _, d := range ChainsDirs {
			if _, err := os.Stat(d); os.IsNotExist(err) {
				os.MkdirAll(d, 0755)
			}
		}

		// drop default tomls into eris' location
		util.CheckDefaultTypes(AccountsTypePath, "account_types")
		util.CheckDefaultTypes(ChainTypePath, "chain_types")

		keys.DaemonAddr = keysAddr

		// Welcomer....
		log.Info("Hello! I'm the marmot who makes eris chains.")
	},
	Run:               MakeChain,
	PersistentPostRun: Archive,
}

func Execute() {
	InitErisChainMaker()
	AddGlobalFlags()
	MakerCmd.Execute()
}

func InitErisChainMaker() {
	do = definitions.NowDo()
}

// Flags that are to be used by commands are handled by the Do struct
// Define the persistent commands (globals)
func AddGlobalFlags() {
	MakerCmd.PersistentFlags().BoolVarP(&do.Verbose, "verbose", "v", defaultVerbose(), "verbose output; more output than no output flags; less output than debug level; default respects $ERIS_CHAINMAKER_VERBOSE")
	MakerCmd.PersistentFlags().BoolVarP(&do.Debug, "debug", "d", defaultDebug(), "debug level output; the most output available for eris-chainmaker; if it is too chatty use verbose flag; default respects $ERIS_CHAINMAKER_DEBUG")
	MakerCmd.PersistentFlags().StringVarP(&keysAddr, "keys-server", "k", defaultKeys(), "keys server which should be used to generate keys; default respects $ERIS_KEYS_PATH")
	MakerCmd.PersistentFlags().StringSliceVarP(&do.AccountTypes, "account-types", "t", defaultActTypes(), "what number of account types should we use? find these in ~/.eris/chains/account_types; incompatible with and overrides chain-type; default respects $ERIS_CHAINMAKER_ACCOUNTTYPES")
	MakerCmd.PersistentFlags().StringVarP(&do.ChainType, "chain-type", "c", defaultChainType(), "which chain type definition should we use? find these in ~/.eris/chains/chain_types; default respects $ERIS_CHAINMAKER_CHAINTYPE")
	MakerCmd.PersistentFlags().StringVarP(&do.CSV, "csv-file", "s", defaultCsvFiles(), "csv file in the form `account-type,number,tokens,toBond,perms; default respects $ERIS_CHAINMAKER_CSVFILE")
	MakerCmd.PersistentFlags().BoolVarP(&do.Tarball, "tar", "r", defaultTarball(), "instead of making directories in ~/.chains, make tarballs; incompatible with and overrides zip; default respects $ERIS_CHAINMAKER_TARBALLS")
	MakerCmd.PersistentFlags().BoolVarP(&do.Zip, "zip", "z", defaultZip(), "instead of making directories in ~/.chains, make zip files; default respects $ERIS_CHAINMAKER_ZIPFILES")
}

//----------------------------------------------------
func MakeChain(cmd *cobra.Command, args []string) {
	argsMin := 1
	if len(args) < argsMin {
		cmd.Help()
		IfExit(fmt.Errorf("\n**Note** you sent our marmots the wrong number of arguments.\nPlease send the marmots at least %d argument(s).", argsMin))
	}
	do.Name = args[0]
	IfExit(chainmaker.MakeChain(do))
}

func Archive(cmd *cobra.Command, args []string) {
	if do.Tarball {
		// IfExit(util.Tarball(do))
	} else if do.Zip {
		// IfExit(util.Zip(do))
	}
}

// ---------------------------------------------------
// Defaults

func defaultVerbose() bool {
	return setDefaultBool("ERIS_CHAINMAKER_VERBOSE", false)
}

func defaultDebug() bool {
	return setDefaultBool("ERIS_CHAINMAKER_DEBUG", false)
}

func defaultKeys() string {
	return setDefaultString("ERIS_KEYS_PATH", fmt.Sprintf("http://localhost:4767"))
}

func defaultChainType() string {
	return setDefaultString("ERIS_CHAINMAKER_CHAINTYPE", "")
}

func defaultActTypes() []string {
	return setDefaultStringSlice("ERIS_CHAINMAKER_ACCOUNTTYPES", []string{})
}

func defaultCsvFiles() string {
	return setDefaultString("ERIS_CHAINMAKER_CSVFILE", "")
}

func defaultTarball() bool {
	return setDefaultBool("ERIS_CHAINMAKER_TARBALLS", false)
}

func defaultZip() bool {
	return setDefaultBool("ERIS_CHAINMAKER_ZIPFILES", false)
}

func setDefaultBool(envVar string, def bool) bool {
	env := os.Getenv(envVar)
	if env != "" {
		i, _ := strconv.ParseBool(env)
		return i
	}
	return def
}

func setDefaultString(envVar, def string) string {
	env := os.Getenv(envVar)
	if env != "" {
		return env
	}
	return def
}

func setDefaultStringSlice(envVar string, def []string) []string {
	env := os.Getenv(envVar)
	if env != "" {
		return strings.Split(env, ";")
	}
	return def
}
