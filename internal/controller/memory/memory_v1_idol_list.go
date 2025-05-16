package memory

import (
	"context"

	v1 "github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
	modelIdol "github.com/Remrinrin/idol-memory-api-go/internal/model/idol"
	serviceIdol "github.com/Remrinrin/idol-memory-api-go/internal/service/idol"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) IdolList(ctx context.Context, req *v1.IdolListReq) (res *v1.IdolListRes, err error) {
	// 参数校验
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 || req.PageSize > 50 {
		req.PageSize = 10
	}

	// 构建查询参数
	input := &serviceIdol.ListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	}

	// 调用 service 层获取列表
	output, err := serviceIdol.Idol().List(ctx, input)
	if err != nil {
		return nil, err
	}

	// 转换为响应对象
	items := make([]modelIdol.Idol, 0, len(output.Items))
	for _, item := range output.Items {
		items = append(items, modelIdol.Idol{
			Id:          gconv.Int64(item.Id),
			Name:        gconv.String(item.Name),
			EnglishName: gconv.String(item.EnglishName),
			Nickname:    gconv.String(item.Nickname),
			Birthday:    *item.Birthday,
			Height:      gconv.Int(item.Height),
			Weight:      gconv.Int(item.Weight),
			BloodType:   gconv.String(item.BloodType),
			Zodiac:      gconv.String(item.Zodiac),
			Position:    gconv.String(item.Position),
			Status:      gconv.String(item.Status),
			CreatedAt:   *item.CreatedAt,
			UpdatedAt:   *item.UpdatedAt,
		})
	}

	// 返回结果
	return &v1.IdolListRes{
		List: modelIdol.IdolList{
			Total:    output.Total,
			Page:     output.Page,
			PageSize: output.PageSize,
			Items:    items,
		},
	}, nil
}
