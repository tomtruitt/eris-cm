package chains

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

// A particular permission
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

	// moderator permissions
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
	// MintDefaultPermFlags MintPermFlag = Send | Call | CreateContract | CreateAccount | Bond | Name | HasBase | HasRole
)

// var (
// 	MintZeroBasePermissions    = MintBasePermissions{0, 0}
// 	MintZeroAccountPermissions = MintAccountPermissions{
// 		Base: MintZeroBasePermissions,
// 	}
// 	DefaultAccountPermissions = MintAccountPermissions{
// 		Base: MintBasePermissions{
// 			MintPerms:  DefaultPermFlags,
// 			MintSetBit: AllPermFlags,
// 		},
// 		Roles: []string{},
// 	}
// )

// Base chain permissions struct
type MintBasePermissions struct {
	// bit array with "has"/"doesn't have" for each permission
	MintPerms MintPermFlag `json:"perms"`

	// bit array with "set"/"not set" for each permission (not-set should fall back to global)
	MintSetBit MintPermFlag `json:"set"`
}

// Set a permission bit. Will set the permission's set bit to true.
func (p *MintBasePermissions) Set(ty MintPermFlag, value bool) error {
	if ty == 0 {
		// return ErrInvalidPermission(ty)
	}
	p.MintSetBit |= ty
	if value {
		p.MintPerms |= ty
	} else {
		p.MintPerms &= ^ty
	}
	return nil
}

type MintAccountPermissions struct {
	MintBase  MintBasePermissions `json:"base"`
	MintRoles []string            `json:"roles"`
}

func MintPermStringToFlag(perm string) (pf PermFlag, err error) {
	switch perm {
	case "root":
		pf = MintRoot
	case "send":
		pf = MintSend
	case "call":
		pf = MintCall
	case "create_contract":
		pf = MintCreateContract
	case "create_account":
		pf = MintCreateAccount
	case "bond":
		pf = MintBond
	case "name":
		pf = MintName
	case "has_base":
		pf = MintHasBase
	case "set_base":
		pf = MintSetBase
	case "unset_base":
		pf = MintUnsetBase
	case "set_global":
		pf = MintSetGlobal
	case "has_role":
		pf = MintHasRole
	case "add_role":
		pf = MintAddRole
	case "rm_role":
		pf = MintRmRole
	default:
		err = fmt.Errorf("Unknown permission %s", perm)
	}
	return
}

func MintPermsStringsToPerm(args map[string]int) (*MintPermFlag, error) {
	bp := ZeroBasePermissions

	for name, val := range args {
		pf, err := MintPermStringToFlag(name)
		if err != nil {
			return 0, err
		}
		bp.Set(pf, val > 0)
	}

	return bp, nil
}

func MintPermsRoot() (*MintPermFlag, error) {
	perms := make(map[string]int)

	perms["root"] = 1
	perms["send"] = 1
	perms["call"] = 1
	perms["create_contract"] = 1
	perms["create_account"] = 1
	perms["bond"] = 1
	perms["name"] = 1
	perms["has_base"] = 1
	perms["set_base"] = 1
	perms["unset_base"] = 1
	perms["set_global"] = 1
	perms["has_role"] = 1
	perms["add_role"] = 1
	perms["rm_role"] = 1

	return MintPermsStringsToPerm(m)
}

func MintPermsDevelopers() (*MintPermFlag, error) {
	perms := make(map[string]int)

	perms["root"] = 0
	perms["send"] = 1
	perms["call"] = 1
	perms["create_contract"] = 1
	perms["create_account"] = 1
	perms["bond"] = 0
	perms["name"] = 1
	perms["has_base"] = 0
	perms["set_base"] = 0
	perms["unset_base"] = 0
	perms["set_global"] = 0
	perms["has_role"] = 1
	perms["add_role"] = 1
	perms["rm_role"] = 1

	return MintPermsStringsToPerm(m)
}

func MintPermsValidators() (*MintPermFlag, error) {
	perms := make(map[string]int)

	perms["root"] = 0
	perms["send"] = 0
	perms["call"] = 0
	perms["create_contract"] = 0
	perms["create_account"] = 0
	perms["bond"] = 1
	perms["name"] = 0
	perms["has_base"] = 0
	perms["set_base"] = 0
	perms["unset_base"] = 0
	perms["set_global"] = 0
	perms["has_role"] = 0
	perms["add_role"] = 0
	perms["rm_role"] = 0

	return MintPermsStringsToPerm(m)
}

func MintPermsParticipants() (*MintPermFlag, error) {
	perms := make(map[string]int)

	perms["root"] = 0
	perms["send"] = 1
	perms["call"] = 1
	perms["create_contract"] = 0
	perms["create_account"] = 0
	perms["bond"] = 0
	perms["name"] = 1
	perms["has_base"] = 0
	perms["set_base"] = 0
	perms["unset_base"] = 0
	perms["set_global"] = 0
	perms["has_role"] = 1
	perms["add_role"] = 0
	perms["rm_role"] = 0

	return MintPermsStringsToPerm(m)
}
