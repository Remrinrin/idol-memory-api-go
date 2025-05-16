package idol

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ResourceType 资源类型
type ResourceType string

const (
	ResourceTypeImage    ResourceType = "image"    // 图片
	ResourceTypeVideo    ResourceType = "video"    // 视频
	ResourceTypeAudio    ResourceType = "audio"    // 音频
	ResourceTypeArticle  ResourceType = "article"  // 文章
	ResourceTypeSocial   ResourceType = "social"   // 社交媒体
	ResourceTypeWebsite  ResourceType = "website"  // 网站
	ResourceTypeDocument ResourceType = "document" // 文档
)

// ResourceLink 资源链接结构
type ResourceLink struct {
	Id          string       `json:"id"`          // 资源唯一ID
	MemoryId    string       `json:"memory_id"`   // 关联的记忆ID
	Type        ResourceType `json:"type"`        // 资源类型
	Title       string       `json:"title"`       // 资源标题
	Description string       `json:"description"` // 资源描述
	Url         string       `json:"url"`         // 资源链接
	Source      string       `json:"source"`      // 来源平台（如：微博、Twitter、Instagram等）
	SourceId    string       `json:"source_id"`   // 来源平台上的ID
	Author      string       `json:"author"`      // 资源作者
	CreatedAt   gtime.Time   `json:"created_at"`  // 创建时间
	UpdatedAt   gtime.Time   `json:"updated_at"`  // 更新时间

	// 元数据
	Tags       []string `json:"tags"`        // 资源标签
	IsVerified bool     `json:"is_verified"` // 是否已验证
	IsOfficial bool     `json:"is_official"` // 是否官方资源

	// 统计信息
	ViewCount  int `json:"view_count"`  // 查看次数
	ShareCount int `json:"share_count"` // 分享次数
}

// ResourceList 资源列表响应
type ResourceList struct {
	Total    int64          `json:"total"`     // 总数
	Page     int            `json:"page"`      // 当前页码
	PageSize int            `json:"page_size"` // 每页数量
	Items    []ResourceLink `json:"items"`     // 资源列表
}
