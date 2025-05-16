package idol

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SharedMemory 共享记忆结构
type SharedMemory struct {
	Id          string     `json:"id"`           // 记忆唯一ID（UUID格式）
	CreatorId   int64      `json:"creator_id"`   // 创建者ID
	CreatorName string     `json:"creator_name"` // 创建者名称
	Title       string     `json:"title"`        // 标题
	Content     string     `json:"content"`      // 文字描述
	MemoryDate  gtime.Time `json:"memory_date"`  // 记忆发生的时间
	CreatedAt   gtime.Time `json:"created_at"`   // 创建时间
	UpdatedAt   gtime.Time `json:"updated_at"`   // 更新时间

	// 元数据
	Type      MemoryType   `json:"type"`       // 记忆类型（事件、瞬间、想法）
	Status    MemoryStatus `json:"status"`     // 记忆状态
	Tags      []string     `json:"tags"`       // 标签
	Location  string       `json:"location"`   // 地点（如：演唱会、活动等）
	IsPrivate bool         `json:"is_private"` // 是否私密

	// 统计信息
	LikeCount    int `json:"like_count"`    // 点赞数
	DislikeCount int `json:"dislike_count"` // 反对数
	ViewCount    int `json:"view_count"`    // 查看数
	ShareCount   int `json:"share_count"`   // 分享数
	ForwardCount int `json:"forward_count"` // 转发数
	CollectCount int `json:"collect_count"` // 收藏数

	// 情感相关
	EmotionTags []string `json:"emotion_tags"` // 情感标签（如：开心、感动、怀念等）
	MoodScore   float64  `json:"mood_score"`   // 平均心情评分（1-5分）

	// 偶像相关
	IdolId      int64  `json:"idol_id"`      // 关联的偶像ID
	IdolName    string `json:"idol_name"`    // 偶像名称
	EventType   string `json:"event_type"`   // 活动类型（演唱会、见面会、签售会等）
	EventScale  string `json:"event_scale"`  // 活动规模（个人、小型、大型等）
	EventStatus string `json:"event_status"` // 活动状态（已结束、进行中、即将开始等）

	// 用户互动
	Votes         []MemoryVote         `json:"votes"`         // 用户投票/评价列表
	Contributions []MemoryContribution `json:"contributions"` // 用户补充内容列表

	// 图数据库相关
	GraphId   string           `json:"graph_id"`  // 图数据库节点ID
	Relations []MemoryRelation `json:"relations"` // 关联关系列表

	// 资源链接
	Resources []ResourceLink `json:"resources"` // 关联的资源链接列表

	// 个人记忆关联
	PersonalMemories []PersonalMemory `json:"personal_memories"` // 关联的个人记忆列表
}

// PersonalMemory 个人记忆结构
type PersonalMemory struct {
	Id         string     `json:"id"`          // 个人记忆唯一ID（UUID格式）
	UserId     int64      `json:"user_id"`     // 用户ID
	UserName   string     `json:"user_name"`   // 用户名
	SharedId   string     `json:"shared_id"`   // 关联的共享记忆ID
	Title      string     `json:"title"`       // 标题
	Content    string     `json:"content"`     // 个人描述
	MemoryDate gtime.Time `json:"memory_date"` // 记忆发生的时间
	CreatedAt  gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  gtime.Time `json:"updated_at"`  // 更新时间

	// 个人情感
	EmotionTags []string `json:"emotion_tags"` // 个人情感标签
	MoodScore   float64  `json:"mood_score"`   // 个人心情评分

	// 个人标签
	Tags []string `json:"tags"` // 个人标签

	// 个人互动
	IsLiked     bool `json:"is_liked"`     // 是否点赞
	IsCollected bool `json:"is_collected"` // 是否收藏
	IsShared    bool `json:"is_shared"`    // 是否分享

	// 个人贡献
	Contributions []MemoryContribution `json:"contributions"` // 个人补充内容
}

// MemoryList 记忆列表响应
type MemoryList struct {
	Total    int64          `json:"total"`     // 总数
	Page     int            `json:"page"`      // 当前页码
	PageSize int            `json:"page_size"` // 每页数量
	Items    []SharedMemory `json:"items"`     // 记忆列表
}

// MemoryDetail 记忆详情响应
type MemoryDetail struct {
	SharedMemory
	Comments  []Comment `json:"comments"` // 评论列表
	VoteStats struct {
		TotalVotes       int            `json:"total_votes"`       // 总投票数
		ApprovalRate     float64        `json:"approval_rate"`     // 赞成率
		AverageScore     float64        `json:"average_score"`     // 平均评分
		VoteDistribution map[string]int `json:"vote_distribution"` // 投票分布
	} `json:"vote_stats"`

	// 关联记忆
	RelatedMemories struct {
		SameEvent      []SharedMemory `json:"same_event"`      // 同一事件相关记忆
		SameIdol       []SharedMemory `json:"same_idol"`       // 同一偶像相关记忆
		SameLocation   []SharedMemory `json:"same_location"`   // 同一地点相关记忆
		TimeRelated    []SharedMemory `json:"time_related"`    // 时间相关记忆
		UserRelated    []SharedMemory `json:"user_related"`    // 用户相关记忆
		ContentRelated []SharedMemory `json:"content_related"` // 内容相关记忆
	} `json:"related_memories"`
}

// PersonalMemoryList 个人记忆列表响应
type PersonalMemoryList struct {
	Total    int64            `json:"total"`     // 总数
	Page     int              `json:"page"`      // 当前页码
	PageSize int              `json:"page_size"` // 每页数量
	Items    []PersonalMemory `json:"items"`     // 个人记忆列表
}
