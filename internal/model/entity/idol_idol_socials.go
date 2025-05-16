// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdolSocials is the golang structure for table idol_idol_socials.
type IdolIdolSocials struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	IdolId    int64       `json:"idolId"    orm:"idol_id"    description:""` //
	Platform  string      `json:"platform"  orm:"platform"   description:""` //
	AccountId string      `json:"accountId" orm:"account_id" description:""` //
	Username  string      `json:"username"  orm:"username"   description:""` //
	Url       string      `json:"url"       orm:"url"        description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
