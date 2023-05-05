package constants

const (
	RowsPerPageEnvKey                   = "DB_ROWS_PERPAGE"
	DatabaseDataFilePathLocation        = "data/apps.db"
	SelectSqliteVersion          string = "SELECT SQLITE_VERSION()"
	UserTableName                string = "user"
	CreateUserTableTemplate      string = `CREATE TABLE IF NOT EXISTS %v (
		id TEXT PRIMARY KEY,
		aggregate_id TEXT,
		name TEXT, 
		type TEXT,
		firstname TEXT, 
		lastname TEXT,
		email TEXT, 
		mobile TEXT, 
		address TEXT, 
		username TEXT, 
		password TEXT, 
		pin TEXT, 
		photo BLOB, 
		roles TEXT,
		status TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
		)`
)
