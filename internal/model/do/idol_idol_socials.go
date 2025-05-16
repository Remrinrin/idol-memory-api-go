// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdolSocials is the golang structure of table idol_socials for DAO operations like Where/Data.
type IdolIdolSocials struct {
	g.Meta    `orm:"table:idol_socials, do:true"`
	Id        interface{} //
	IdolId    interface{} //
	Platform  interface{} //
	AccountId interface{} //
	Username  interface{} //
	Url       interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
