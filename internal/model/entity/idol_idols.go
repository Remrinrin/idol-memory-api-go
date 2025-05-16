// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// IdolIdols is the golang structure for table idol_idols.
type IdolIdols struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	Name        string      `json:"name"        orm:"name"         description:""` //
	EnglishName string      `json:"englishName" orm:"english_name" description:""` //
	Nickname    string      `json:"nickname"    orm:"nickname"     description:""` //
	Birthday    *gtime.Time `json:"birthday"    orm:"birthday"     description:""` //
	Height      int         `json:"height"      orm:"height"       description:""` //
	Weight      int         `json:"weight"      orm:"weight"       description:""` //
	BloodType   string      `json:"bloodType"   orm:"blood_type"   description:""` //
	Zodiac      string      `json:"zodiac"      orm:"zodiac"       description:""` //
	Position    string      `json:"position"    orm:"position"     description:""` //
	Status      string      `json:"status"      orm:"status"       description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
}
