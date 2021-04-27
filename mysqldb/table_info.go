package mysqldb

type DatabaseTable struct {
	*Table
	Columns []*TableColumn
}

type Table struct {
	TableName    string `db:"TABLE_NAME"`
	TableSchema  string `db:"TABLE_SCHEMA"`
	TableComment string `db:"TABLE_COMMENT"`
}

type TableColumn struct {
	ColumnName string `db:"COLUMN_NAME"`
	DataType   string `db:"DATA_TYPE"`
	IsNullable string `db:"IS_NULLABLE"`
}

var dataType2GolangType = map[string]string{
	"tinyint":  "int",
	"bigint":   "int64",
	"varchar":  "string",
	"datetime": "time.Time",
	"text":     "string",
}
