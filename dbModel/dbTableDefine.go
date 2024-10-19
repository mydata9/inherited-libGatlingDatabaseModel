package dbModel

//import "time"

// id primary, Key string, ValueStr string, ValueInt int, ValueFloat float64,date default now;
// type CWJDBTableIDKVD struct {
// 	ID         int64
// 	Key        string
// 	ValueStr   string
// 	ValueInt   int64
// 	ValueFloat float64
// 	Date1      time.Time
// }

type DBMWJTableType int

const (
	DBMWJTT_Unknow DBMWJTableType = iota
	DBMWJTT_KeyAll1

	DBMWJTT_MAXINDEX
)

type DBMWJDatabaseType int

const (
	DBMWJDT_Unknow DBMWJDatabaseType = iota
	DBMWJDT_Sqlite
	DBMWJDT_Postgres

	DBMWJDT_MAXINDEX
)
