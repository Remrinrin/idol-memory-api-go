// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdolGroups is the golang structure of table idol_groups for DAO operations like Where/Data.
type IdolIdolGroups struct {
	g.Meta      `orm:"table:idol_groups, do:true"`
	Id          interface{} //
	Name        interface{} //
	EnglishName interface{} //
	Nickname    interface{} //
	Company     interface{} //
	FoundedDate *gtime.Time //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
