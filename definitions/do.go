package definitions

type Do struct {
	Debug        bool     `mapstructure:"," json:"," yaml:"," toml:","`
	Verbose      bool     `mapstructure:"," json:"," yaml:"," toml:","`
	Name         string   `mapstructure:"," json:"," yaml:"," toml:","`
	ChainType    string   `mapstructure:"," json:"," yaml:"," toml:","`
	CSV          string   `mapstructure:"," json:"," yaml:"," toml:","`
	AccountTypes []string `mapstructure:"," json:"," yaml:"," toml:","`
	Zip          bool     `mapstructure:"," json:"," yaml:"," toml:","`
	Tarball      bool     `mapstructure:"," json:"," yaml:"," toml:","`

	Result string
}

func NowDo() *Do {
	return &Do{}
}
