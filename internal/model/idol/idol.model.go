package idol

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Idol 偶像基本信息
type Idol struct {
	Id          int64      `json:"id"`           // 偶像ID
	Name        string     `json:"name"`         // 偶像名称
	EnglishName string     `json:"english_name"` // 英文名
	Nickname    string     `json:"nickname"`     // 昵称
	Birthday    gtime.Time `json:"birthday"`     // 生日
	Height      int        `json:"height"`       // 身高(cm)
	Weight      int        `json:"weight"`       // 体重(kg)
	BloodType   string     `json:"blood_type"`   // 血型
	Zodiac      string     `json:"zodiac"`       // 星座
	GroupId     int64      `json:"group_id"`     // 所属团体ID
	GroupName   string     `json:"group_name"`   // 所属团体名称
	Position    string     `json:"position"`     // 在团体中的位置
	Status      string     `json:"status"`       // 状态（活跃、毕业等）
	CreatedAt   gtime.Time `json:"created_at"`   // 创建时间
	UpdatedAt   gtime.Time `json:"updated_at"`   // 更新时间
}

// IdolGroup 偶像团体
type IdolGroup struct {
	Id          int64      `json:"id"`           // 团体ID
	Name        string     `json:"name"`         // 团体名称
	EnglishName string     `json:"english_name"` // 英文名
	Nickname    string     `json:"nickname"`     // 昵称
	Company     string     `json:"company"`      // 所属公司
	FoundedDate gtime.Time `json:"founded_date"` // 成立日期
	Status      string     `json:"status"`       // 状态（活跃、解散等）
	CreatedAt   gtime.Time `json:"created_at"`   // 创建时间
	UpdatedAt   gtime.Time `json:"updated_at"`   // 更新时间
}

// IdolSocial 偶像社交媒体账号
type IdolSocial struct {
	Id        int64      `json:"id"`         // 账号ID
	IdolId    int64      `json:"idol_id"`    // 偶像ID
	Platform  string     `json:"platform"`   // 平台（微博、Twitter等）
	AccountId string     `json:"account_id"` // 平台账号ID
	Username  string     `json:"username"`   // 用户名
	Url       string     `json:"url"`        // 主页链接
	CreatedAt gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt gtime.Time `json:"updated_at"` // 更新时间
}

// IdolList 偶像列表响应
type IdolList struct {
	Total    int64  `json:"total"`     // 总数
	Page     int    `json:"page"`      // 当前页码
	PageSize int    `json:"page_size"` // 每页数量
	Items    []Idol `json:"items"`     // 偶像列表
}

// IdolGroupList 偶像团体列表响应
type IdolGroupList struct {
	Total    int64       `json:"total"`     // 总数
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页数量
	Items    []IdolGroup `json:"items"`     // 团体列表
}

// IdolRelation 偶像关联关系
type IdolRelation struct {
	Id           int64      `json:"id"`            // 关联ID
	SourceId     string     `json:"source_id"`     // 源偶像ID
	TargetId     string     `json:"target_id"`     // 目标偶像ID
	RelationType string     `json:"relation_type"` // 关联类型
	Weight       float64    `json:"weight"`        // 关联权重
	CreatedAt    gtime.Time `json:"created_at"`    // 创建时间
	UpdatedAt    gtime.Time `json:"updated_at"`    // 更新时间
}

// IdolRelationType 偶像关联类型
type IdolRelationType string

const (
	IdolRelationTypeSameGroup     IdolRelationType = "same_group"    // 同一团体
	IdolRelationTypeSameCompany   IdolRelationType = "same_company"  // 同一公司
	IdolRelationTypeCollaboration IdolRelationType = "collaboration" // 合作关系
	IdolRelationTypeFriendship    IdolRelationType = "friendship"    // 朋友关系
	IdolRelationTypeMentor        IdolRelationType = "mentor"        // 师徒关系
)

// IdolDetail 偶像详情响应
type IdolDetail struct {
	Idol
	// 相关记忆
	RelatedMemories struct {
		RecentMemories []SharedMemory `json:"recent_memories"` // 最近记忆
		HotMemories    []SharedMemory `json:"hot_memories"`    // 热门记忆
		EventMemories  []SharedMemory `json:"event_memories"`  // 活动记忆
	} `json:"related_memories"`

	// 相关偶像
	RelatedIdols struct {
		GroupMembers  []Idol `json:"group_members"` // 团体成员
		CompanyIdols  []Idol `json:"company_idols"` // 同公司偶像
		Collaborators []Idol `json:"collaborators"` // 合作偶像
		Friends       []Idol `json:"friends"`       // 朋友关系
	} `json:"related_idols"`
}
