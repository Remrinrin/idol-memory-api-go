// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IdolIdolSocialsDao is the data access object for the table idol_socials.
type IdolIdolSocialsDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  IdolIdolSocialsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// IdolIdolSocialsColumns defines and stores column names for the table idol_socials.
type IdolIdolSocialsColumns struct {
	Id        string //
	IdolId    string //
	Platform  string //
	AccountId string //
	Username  string //
	Url       string //
	CreatedAt string //
	UpdatedAt string //
}

// idolIdolSocialsColumns holds the columns for the table idol_socials.
var idolIdolSocialsColumns = IdolIdolSocialsColumns{
	Id:        "id",
	IdolId:    "idol_id",
	Platform:  "platform",
	AccountId: "account_id",
	Username:  "username",
	Url:       "url",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewIdolIdolSocialsDao creates and returns a new DAO object for table data access.
func NewIdolIdolSocialsDao(handlers ...gdb.ModelHandler) *IdolIdolSocialsDao {
	return &IdolIdolSocialsDao{
		group:    "idol",
		table:    "idol_socials",
		columns:  idolIdolSocialsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IdolIdolSocialsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IdolIdolSocialsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IdolIdolSocialsDao) Columns() IdolIdolSocialsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IdolIdolSocialsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IdolIdolSocialsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *IdolIdolSocialsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
