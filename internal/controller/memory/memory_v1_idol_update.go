package memory

import (
	"context"

	v1 "github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
	"github.com/Remrinrin/idol-memory-api-go/internal/model/do"
	serviceIdol "github.com/Remrinrin/idol-memory-api-go/internal/service/idol"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) IdolUpdate(ctx context.Context, req *v1.IdolUpdateReq) (res *v1.IdolUpdateRes, err error) {
	// 参数校验
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "偶像ID不能为空")
	}
	if req.Idol.Name == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "偶像名称不能为空")
	}

	// 检查偶像是否存在
	exist, err := serviceIdol.Idol().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if exist == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "偶像不存在")
	}

	// 转换为 DO 对象
	doIdol := &do.IdolIdols{
		Id:          req.Id,
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

	// 调用 service 层更新偶像
	err = serviceIdol.Idol().Update(ctx, doIdol)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.IdolUpdateRes{}, nil
}
