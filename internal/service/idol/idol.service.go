package idol

import (
	"context"

	"github.com/Remrinrin/idol-memory-api-go/internal/dao"
	"github.com/Remrinrin/idol-memory-api-go/internal/model/do"
)

type IIdol interface {
	Create(ctx context.Context, in *do.IdolIdols) (id int64, err error)
	Update(ctx context.Context, in *do.IdolIdols) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*do.IdolIdols, error)
	List(ctx context.Context, in *ListInput) (*ListOutput, error)
}

// ListInput 列表查询参数
type ListInput struct {
	Page     int    `json:"page"`      // 页码
	PageSize int    `json:"page_size"` // 每页数量
	Name     string `json:"name"`      // 偶像名称
	GroupId  int64  `json:"group_id"`  // 团体ID
	Status   string `json:"status"`    // 状态
}

// ListOutput 列表输出
type ListOutput struct {
	Total    int64          `json:"total"`     // 总数
	Page     int            `json:"page"`      // 当前页码
	PageSize int            `json:"page_size"` // 每页数量
	Items    []do.IdolIdols `json:"items"`     // 偶像列表
}

var localIdol IIdol

func Idol() IIdol {
	return localIdol
}

func RegisterIdol(i IIdol) {
	localIdol = i
}

type sIdol struct{}

func init() {
	RegisterIdol(New())
}

func New() *sIdol {
	return &sIdol{}
}

// Create 创建偶像
func (s *sIdol) Create(ctx context.Context, in *do.IdolIdols) (id int64, err error) {
	// 插入数据并返回ID
	result, err := dao.IdolIdols.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// Update 更新偶像
func (s *sIdol) Update(ctx context.Context, in *do.IdolIdols) error {
	// 更新数据
	_, err := dao.IdolIdols.Ctx(ctx).
		Where(dao.IdolIdols.Columns().Id, in.Id).
		Data(in).
		Update()
	return err
}

// Delete 删除偶像
func (s *sIdol) Delete(ctx context.Context, id int64) error {
	// 删除数据
	_, err := dao.IdolIdols.Ctx(ctx).
		Where(dao.IdolIdols.Columns().Id, id).
		Delete()
	return err
}

// GetById 根据ID获取偶像详情
func (s *sIdol) GetById(ctx context.Context, id int64) (*do.IdolIdols, error) {
	var detail do.IdolIdols

	// 获取基本信息
	err := dao.IdolIdols.Ctx(ctx).
		Where(dao.IdolIdols.Columns().Id, id).
		Scan(&detail)
	if err != nil {
		return nil, err
	}

	return &detail, nil
}

// List 获取偶像列表
func (s *sIdol) List(ctx context.Context, in *ListInput) (*ListOutput, error) {
	// 构建查询条件
	model := dao.IdolIdols.Ctx(ctx)
	if in.Name != "" {
		model = model.WhereLike(dao.IdolIdols.Columns().Name, "%"+in.Name+"%")
	}
	if in.Status != "" {
		model = model.Where(dao.IdolIdols.Columns().Status, in.Status)
	}

	// 获取总数
	count, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 获取分页数据
	var items []do.IdolIdols
	err = model.Page(in.Page, in.PageSize).
		OrderDesc(dao.IdolIdols.Columns().CreatedAt).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	return &ListOutput{
		Total:    int64(count),
		Page:     in.Page,
		PageSize: in.PageSize,
		Items:    items,
	}, nil
}
