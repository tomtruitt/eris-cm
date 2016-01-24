package definitions

type MintPermFlag uint64

// Base permission references are like unix (the index is already bit shifted)
const (
	// chain permissions
	MintRoot           MintPermFlag = 1 << iota // 1
	MintSend                                    // 2
	MintCall                                    // 4
	MintCreateContract                          // 8
	MintCreateAccount                           // 16
	MintBond                                    // 32
	MintName                                    // 64

	// application permissions
	MintHasBase
	MintSetBase
	MintUnsetBase
	MintSetGlobal
	MintHasRole
	MintAddRole
	MintRmRole

	MintNumPermissions uint = 14 // NOTE Adjust this too. We can support upto 64

	MintTopPermFlag  MintPermFlag = 1 << (MintNumPermissions - 1)
	MintAllPermFlags MintPermFlag = MintTopPermFlag | (MintTopPermFlag - 1)
)

type MintGenesis struct {
	ChainID    string           `json:"base"`
	Accounts   []*MintAccount   `json:"accounts"`
	Validators []*MintValidator `json:"validators"`
}

type MintAccount struct {
	Address     string                  `json:"address"`
	Amount      int                     `json:"amount"`
	Name        string                  `json:"name"`
	Permissions *MintAccountPermissions `json:"permissions"`
}

type MintAccountPermissions struct {
	MintBase  MintBasePermissions `json:"base"`
	MintRoles []string            `json:"roles"`
}

type MintBasePermissions struct {
	MintPerms  MintPermFlag `json:"perms"`
	MintSetBit MintPermFlag `json:"set"`
}

type MintValidator struct {
	PubKey   []interface{}   `json:"pub_key"`
	Name     string          `json:"name"`
	Amount   int             `json:"amount"`
	UnbondTo []*MintTxOutput `json:"unbond_to"`
}

type MintTxOutput struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}
