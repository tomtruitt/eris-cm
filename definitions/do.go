package definitions

type Do struct {
	Debug        bool
	Verbose      bool
	Name         string
	ChainType    string
	CSV          string
	AccountTypes []string
	Zip          bool
	Tarball      bool
	Output       bool
	Accounts     []*Account
	Result       string
}

func NowDo() *Do {
	return &Do{}
}
