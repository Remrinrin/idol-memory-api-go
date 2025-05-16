// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package memory

import (
	"context"

	"github.com/Remrinrin/idol-memory-api-go/api/memory/v1"
)

type IMemoryV1 interface {
	IdolCreate(ctx context.Context, req *v1.IdolCreateReq) (res *v1.IdolCreateRes, err error)
	IdolUpdate(ctx context.Context, req *v1.IdolUpdateReq) (res *v1.IdolUpdateRes, err error)
	IdolDelete(ctx context.Context, req *v1.IdolDeleteReq) (res *v1.IdolDeleteRes, err error)
	IdolGet(ctx context.Context, req *v1.IdolGetReq) (res *v1.IdolGetRes, err error)
	IdolList(ctx context.Context, req *v1.IdolListReq) (res *v1.IdolListRes, err error)
}
