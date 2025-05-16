package memory

import (
	"context"

	v1 "github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
	"github.com/Remrinrin/idol-memory-api-go/internal/model/do"
	"github.com/Remrinrin/idol-memory-api-go/internal/service/idol"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) IdolCreate(ctx context.Context, req *v1.IdolCreateReq) (res *v1.IdolCreateRes, err error) {
	// 参数校验
	if req.Idol.Name == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "偶像名称不能为空")
	}

	// 转换为 DO 对象
	doIdol := &do.IdolIdols{
		Name:        req.Idol.Name,
		EnglishName: req.Idol.EnglishName,
		Nickname:    req.Idol.Nickname,
		Birthday:    &req.Idol.Birthday,
		Height:      req.Idol.Height,
		Weight:      req.Idol.Weight,
		BloodType:   req.Idol.BloodType,
		Zodiac:      req.Idol.Zodiac,
		Position:    req.Idol.Position,
		Status:      req.Idol.Status,
	}

	// 调用 service 层创建偶像
	id, err := idol.Idol().Create(ctx, doIdol)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.IdolCreateRes{
		Id: id,
	}, nil
}
