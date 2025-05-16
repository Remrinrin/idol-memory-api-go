package idol

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment 评论结构
type Comment struct {
	Id        int64      `json:"id"`
	MemoryId  int64      `json:"memory_id"`
	UserId    int64      `json:"user_id"`
	UserName  string     `json:"user_name"`
	Content   string     `json:"content"`
	CreatedAt gtime.Time `json:"created_at"`
}

// MemoryVote 记忆投票/评价
type MemoryVote struct {
	Id        int64      `json:"id"`         // 投票ID
	MemoryId  int64      `json:"memory_id"`  // 记忆ID
	UserId    int64      `json:"user_id"`    // 用户ID
	UserName  string     `json:"user_name"`  // 用户名
	VoteType  string     `json:"vote_type"`  // 投票类型（赞成/反对）
	Comment   string     `json:"comment"`    // 评价内容
	CreatedAt gtime.Time `json:"created_at"` // 创建时间
}

// MemoryContribution 记忆贡献
type MemoryContribution struct {
	Id        int64      `json:"id"`         // 贡献ID
	MemoryId  int64      `json:"memory_id"`  // 记忆ID
	UserId    int64      `json:"user_id"`    // 贡献用户ID
	UserName  string     `json:"user_name"`  // 贡献用户名
	Content   string     `json:"content"`    // 补充内容
	CreatedAt gtime.Time `json:"created_at"` // 创建时间
}

// MemoryRelation 记忆关联关系
type MemoryRelation struct {
	Id           int64      `json:"id"`            // 关联ID
	SourceId     string     `json:"source_id"`     // 源记忆ID
	TargetId     string     `json:"target_id"`     // 目标记忆ID
	RelationType string     `json:"relation_type"` // 关联类型
	Weight       float64    `json:"weight"`        // 关联权重
	CreatedAt    gtime.Time `json:"created_at"`    // 创建时间
	UpdatedAt    gtime.Time `json:"updated_at"`    // 更新时间
}

// MemoryRelationType 记忆关联类型
type MemoryRelationType string

const (
	RelationTypeSameEvent      MemoryRelationType = "same_event"      // 同一事件
	RelationTypeSameIdol       MemoryRelationType = "same_idol"       // 同一偶像
	RelationTypeSameLocation   MemoryRelationType = "same_location"   // 同一地点
	RelationTypeTimeRelated    MemoryRelationType = "time_related"    // 时间相关
	RelationTypeUserRelated    MemoryRelationType = "user_related"    // 用户相关
	RelationTypeContentRelated MemoryRelationType = "content_related" // 内容相关
)
