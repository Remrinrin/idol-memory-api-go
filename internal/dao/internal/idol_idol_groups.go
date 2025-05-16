// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IdolIdolGroupsDao is the data access object for the table idol_groups.
type IdolIdolGroupsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  IdolIdolGroupsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// IdolIdolGroupsColumns defines and stores column names for the table idol_groups.
type IdolIdolGroupsColumns struct {
	Id          string //
	Name        string //
	EnglishName string //
	Nickname    string //
	Company     string //
	FoundedDate string //
	Status      string //
	CreatedAt   string //
	UpdatedAt   string //
}

// idolIdolGroupsColumns holds the columns for the table idol_groups.
var idolIdolGroupsColumns = IdolIdolGroupsColumns{
	Id:          "id",
	Name:        "name",
	EnglishName: "english_name",
	Nickname:    "nickname",
	Company:     "company",
	FoundedDate: "founded_date",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewIdolIdolGroupsDao creates and returns a new DAO object for table data access.
func NewIdolIdolGroupsDao(handlers ...gdb.ModelHandler) *IdolIdolGroupsDao {
	return &IdolIdolGroupsDao{
		group:    "idol",
		table:    "idol_groups",
		columns:  idolIdolGroupsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IdolIdolGroupsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IdolIdolGroupsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IdolIdolGroupsDao) Columns() IdolIdolGroupsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IdolIdolGroupsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IdolIdolGroupsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *IdolIdolGroupsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
