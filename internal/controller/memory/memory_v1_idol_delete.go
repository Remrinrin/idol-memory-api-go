package memory

import (
	"context"

	v1 "github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
	serviceIdol "github.com/Remrinrin/idol-memory-api-go/internal/service/idol"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) IdolDelete(ctx context.Context, req *v1.IdolDeleteReq) (res *v1.IdolDeleteRes, err error) {
	// 参数校验
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "偶像ID不能为空")
	}

	// 检查偶像是否存在
	exist, err := serviceIdol.Idol().GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if exist == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "偶像不存在")
	}

	// 调用 service 层删除偶像
	err = serviceIdol.Idol().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.IdolDeleteRes{}, nil
}
