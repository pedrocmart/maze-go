package consts

type Migration string

const (
	AppName         = "DATASTORE"
	HttpHost        = "HTTP_HOST"
	HttpPort        = "HTTP_PORT"
	DatabaseURL     = "DB_URL"
	LogOutput       = "LOG_OUTPUT"
	LogFormat       = "LOG_FORMAT"
	DatabaseTimeout = "DB_TIMEOUT_SEC"
	DatabaseRefresh = "DB_REFRESH_SEC"

	RequiredParam = "Required param"

	MigrateUp    Migration = "up"
	MigrateDown  Migration = "down"
	MigratePrint Migration = "print"
)
