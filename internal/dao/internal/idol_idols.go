// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IdolIdolsDao is the data access object for the table idols.
type IdolIdolsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  IdolIdolsColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// IdolIdolsColumns defines and stores column names for the table idols.
type IdolIdolsColumns struct {
	Id          string //
	Name        string //
	EnglishName string //
	Nickname    string //
	Birthday    string //
	Height      string //
	Weight      string //
	BloodType   string //
	Zodiac      string //
	Position    string //
	Status      string //
	CreatedAt   string //
	UpdatedAt   string //
}

// idolIdolsColumns holds the columns for the table idols.
var idolIdolsColumns = IdolIdolsColumns{
	Id:          "id",
	Name:        "name",
	EnglishName: "english_name",
	Nickname:    "nickname",
	Birthday:    "birthday",
	Height:      "height",
	Weight:      "weight",
	BloodType:   "blood_type",
	Zodiac:      "zodiac",
	Position:    "position",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewIdolIdolsDao creates and returns a new DAO object for table data access.
func NewIdolIdolsDao(handlers ...gdb.ModelHandler) *IdolIdolsDao {
	return &IdolIdolsDao{
		group:    "idol",
		table:    "idols",
		columns:  idolIdolsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *IdolIdolsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *IdolIdolsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *IdolIdolsDao) Columns() IdolIdolsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *IdolIdolsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *IdolIdolsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *IdolIdolsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
