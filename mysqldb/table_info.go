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
	ColumnName    string `db:"COLUMN_NAME"`
	DataType      string `db:"DATA_TYPE"`
	IsNullable    string `db:"IS_NULLABLE"`
	ColumnComment string `db:"COLUMN_COMMENT"`
}

var dataType2GolangType = map[string]string{
	"bigint":     "int64",
	"binary":     "string",
	"blob":       "string",
	"char":       "string",
	"datetime":   "time.Time",
	"decimal":    "float64",
	"double":     "float64",
	"enum":       "string",
	"float":      "float64",
	"int":        "int",
	"json":       "string",
	"longblob":   "string",
	"longtext":   "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"set":        "string",
	"smallint":   "int16",
	"text":       "string",
	"time":       "time.Time",
	"timestamp":  "time.Time",
	"tinyint":    "int8",
	"varchar":    "string",
	"varbinary":  "string",
}
