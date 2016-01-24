package definitions

type Do struct {
	Debug   bool   `mapstructure:"," json:"," yaml:"," toml:","`
	Verbose bool   `mapstructure:"," json:"," yaml:"," toml:","`
	Name    string `mapstructure:"," json:"," yaml:"," toml:","`

	Chain  *Chain
	Result string
}

func NowDo() *Do {
	return &Do{}
}
