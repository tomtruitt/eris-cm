package commands

import (
	"os"
	"strconv"
	"strings"

	"github.com/eris-ltd/eris-chainmaker/chainmaker"
	"github.com/eris-ltd/eris-chainmaker/definitions"
	"github.com/eris-ltd/eris-chainmaker/version"

	log "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	. "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	logger "github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
	"github.com/eris-ltd/eris-chainmaker/Godeps/_workspace/src/github.com/spf13/cobra"
)

const VERSION = version.VERSION

// Global Do struct
var do *definitions.Do

// Defining the root command
var MakerCmd = &cobra.Command{
	Use:   "eris-chainmaker",
	Short: "The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains",
	Long: `The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains

Made with <3 by Eris Industries.

Complete documentation is available at https://docs.erisindustries.com
` + "\nVersion:\n  " + VERSION,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.SetFormatter(logger.ErisFormatter{})

		log.SetLevel(log.WarnLevel)
		if do.Verbose {
			log.SetLevel(log.InfoLevel)
		} else if do.Debug {
			log.SetLevel(log.DebugLevel)
		}

		// Welcomer....
		log.Info("Hello! I'm eris-chainmaker.")
	},

	Run: MakeChain,

	PersistentPostRun: func(cmd *cobra.Command, args []string) {},
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
	MakerCmd.PersistentFlags().BoolVarP(&do.Verbose, "verbose", "v", defaultVerbose(), "verbose output; more output than no output flags; less output than debug level; default respects $EPM_VERBOSE")
	MakerCmd.PersistentFlags().BoolVarP(&do.Debug, "debug", "d", defaultDebug(), "debug level output; the most output available for epm; if it is too chatty use verbose flag; default respects $EPM_DEBUG")
	// MakerCmd.PersistentFlags().StringVarP(&do.DefaultOutput, "output", "o", defaultOutput(), "output format which epm should use [csv,json]; default respects $EPM_OUTPUT_FORMAT")
	// MakerCmd.PersistentFlags().BoolVarP(&do.SummaryTable, "summary", "t", defaultSummaryTable(), "output a table summarizing epm jobs; default respects $EPM_SUMMARY_TABLE")
}

//----------------------------------------------------
func MakeChain(cmd *cobra.Command, args []string) {
	IfExit(chains.MakeChain(do))
}

// ---------------------------------------------------
// Defaults

func defaultVerbose() bool {
	return setDefaultBool("EPM_VERBOSE", false)
}

func defaultDebug() bool {
	return setDefaultBool("EPM_DEBUG", false)
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
