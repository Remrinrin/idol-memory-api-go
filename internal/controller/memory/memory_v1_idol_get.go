package memory

import (
	"context"

	v1 "github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
	modelIdol "github.com/Remrinrin/idol-memory-api-go/internal/model/idol"
	serviceIdol "github.com/Remrinrin/idol-memory-api-go/internal/service/idol"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) IdolGet(ctx context.Context, req *v1.IdolGetReq) (res *v1.IdolGetRes, err error) {
	// 参数校验
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "偶像ID不能为空")
	}

	// 调用 service 层获取偶像详情
	doIdol, err := serviceIdol.Idol().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if doIdol == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "偶像不存在")
	}

	// 转换为响应对象
	detail := &modelIdol.IdolDetail{
		Idol: modelIdol.Idol{
			Id:          gconv.Int64(doIdol.Id),
			Name:        gconv.String(doIdol.Name),
			EnglishName: gconv.String(doIdol.EnglishName),
			Nickname:    gconv.String(doIdol.Nickname),
			Birthday:    *doIdol.Birthday,
			Height:      gconv.Int(doIdol.Height),
			Weight:      gconv.Int(doIdol.Weight),
			BloodType:   gconv.String(doIdol.BloodType),
			Zodiac:      gconv.String(doIdol.Zodiac),
			Position:    gconv.String(doIdol.Position),
			Status:      gconv.String(doIdol.Status),
			CreatedAt:   *doIdol.CreatedAt,
			UpdatedAt:   *doIdol.UpdatedAt,
		},
	}

	// 返回结果
	return &v1.IdolGetRes{
		Detail: *detail,
	}, nil
}
