package mysqldb

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var ErrNoTable = errors.New("no table")

func getTableInfo(ctx context.Context, db sqlx.ExtContext, name, schema string) (*Table, error) {
	var t []*Table
	const s = "SELECT TABLE_SCHEMA, TABLE_NAME, TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_NAME = ? AND TABLE_SCHEMA = ?;"
	err := sqlx.SelectContext(ctx, db, &t, s, name, schema)
	if err != nil {
		return nil, err
	}
	if len(t) == 0 {
		return nil, ErrNoTable
	}
	return t[0], nil
}

func getTableColumnInfo(ctx context.Context, db sqlx.ExtContext, name, schema string) ([]*TableColumn, error) {
	var cols []*TableColumn
	const s = "SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_COMMENT FROM information_schema.COLUMNS " +
		"WHERE TABLE_NAME = ? AND TABLE_SCHEMA = ?;"
	err := sqlx.SelectContext(ctx, db, &cols, s, name, schema)
	if err != nil {
		return nil, err
	}
	for _, v := range cols {
		t, ok := dataType2GolangType[v.DataType]
		if !ok {
			return nil, fmt.Errorf("unsupported golang type %s", v.DataType)
		}
		if v.IsNullable == "YES" {
			t = "*" + t
		}
		v.DataType = t
	}
	return cols, nil
}

func NewDatabaseTableInfo(name string) (*DatabaseTable, error) {
	var ctx = context.Background()
	tableInfo, err := getTableInfo(ctx, Pool, name, Schema)
	if err != nil {
		return nil, err
	}
	columnInfo, err := getTableColumnInfo(ctx, Pool, name, Schema)
	if err != nil {
		return nil, err
	}
	table := &DatabaseTable{}
	table.Table = tableInfo
	table.Columns = columnInfo
	return table, nil
}
