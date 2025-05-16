package idol

// MediaType 媒体类型枚举
type MediaType string

const (
	MediaTypeTwitter MediaType = "twitter" // 推特链接
	MediaTypeVideo   MediaType = "video"   // 视频链接
	MediaTypeImage   MediaType = "image"   // 图片链接
	MediaTypeText    MediaType = "text"    // 纯文本
)

// MemoryType 记忆类型
type MemoryType string

const (
	MemoryTypeEvent   MemoryType = "event"   // 事件
	MemoryTypeMoment  MemoryType = "moment"  // 瞬间
	MemoryTypeThought MemoryType = "thought" // 想法
)

// MemoryStatus 记忆状态
type MemoryStatus string

const (
	MemoryStatusDraft     MemoryStatus = "draft"     // 草稿
	MemoryStatusPublished MemoryStatus = "published" // 已发布
	MemoryStatusArchived  MemoryStatus = "archived"  // 已归档
)
