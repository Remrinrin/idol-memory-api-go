// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdolGroups is the golang structure for table idol_idol_groups.
type IdolIdolGroups struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	Name        string      `json:"name"        orm:"name"         description:""` //
	EnglishName string      `json:"englishName" orm:"english_name" description:""` //
	Nickname    string      `json:"nickname"    orm:"nickname"     description:""` //
	Company     string      `json:"company"     orm:"company"      description:""` //
	FoundedDate *gtime.Time `json:"foundedDate" orm:"founded_date" description:""` //
	Status      string      `json:"status"      orm:"status"       description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
}
