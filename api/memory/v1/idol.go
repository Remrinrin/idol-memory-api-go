package v1

import (
	"github.com/Remrinrin/idol-memory-api-go/internal/model/idol"
	"github.com/gogf/gf/v2/frame/g"
)

// IdolCreateReq 创建偶像请求
type IdolCreateReq struct {
	g.Meta `path:"/idol" method:"post" tags:"偶像管理" summary:"创建偶像"`
	Idol   idol.Idol `v:"required#偶像信息不能为空" json:"idol" dc:"偶像信息"`
}

// IdolCreateRes 创建偶像响应
type IdolCreateRes struct {
	Id int64 `json:"id" dc:"偶像ID"`
}

// IdolUpdateReq 更新偶像请求
type IdolUpdateReq struct {
	g.Meta `path:"/idol/{id}" method:"put" tags:"偶像管理" summary:"更新偶像"`
	Id     int64     `v:"required#偶像ID不能为空" in:"path" json:"id" dc:"偶像ID"`
	Idol   idol.Idol `v:"required#偶像信息不能为空" json:"idol" dc:"偶像信息"`
}

// IdolUpdateRes 更新偶像响应
type IdolUpdateRes struct{}

// IdolDeleteReq 删除偶像请求
type IdolDeleteReq struct {
	g.Meta `path:"/idol/{id}" method:"delete" tags:"偶像管理" summary:"删除偶像"`
	Id     int64 `v:"required#偶像ID不能为空" in:"path" json:"id" dc:"偶像ID"`
}

// IdolDeleteRes 删除偶像响应
type IdolDeleteRes struct{}

// IdolGetReq 获取偶像详情请求
type IdolGetReq struct {
	g.Meta `path:"/idol/{id}" method:"get" tags:"偶像管理" summary:"获取偶像详情"`
	Id     int64 `v:"required#偶像ID不能为空" in:"path" json:"id" dc:"偶像ID"`
}

// IdolGetRes 获取偶像详情响应
type IdolGetRes struct {
	Detail idol.IdolDetail `json:"detail" dc:"偶像详情"`
}

// IdolListReq 获取偶像列表请求
type IdolListReq struct {
	g.Meta   `path:"/idols" method:"get" tags:"偶像管理" summary:"获取偶像列表"`
	Page     int    `v:"min:1#页码最小值为1" d:"1" in:"query" json:"page" dc:"页码"`
	PageSize int    `v:"max:50#每页数量最大为50" d:"10" in:"query" json:"page_size" dc:"每页数量"`
	Name     string `in:"query" json:"name" dc:"偶像名称"`
	GroupId  int64  `in:"query" json:"group_id" dc:"团体ID"`
	Status   string `in:"query" json:"status" dc:"状态"`
}

// IdolListRes 获取偶像列表响应
type IdolListRes struct {
	List idol.IdolList `json:"list" dc:"偶像列表"`
}
