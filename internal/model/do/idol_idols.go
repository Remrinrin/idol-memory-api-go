// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdols is the golang structure of table idols for DAO operations like Where/Data.
type IdolIdols struct {
	g.Meta      `orm:"table:idols, do:true"`
	Id          interface{} //
	Name        interface{} //
	EnglishName interface{} //
	Nickname    interface{} //
	Birthday    *gtime.Time //
	Height      interface{} //
	Weight      interface{} //
	BloodType   interface{} //
	Zodiac      interface{} //
	GroupId     interface{} //
	GroupName   interface{} //
	Position    interface{} //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
