package api

import (
	"database/sql"

	"github.com/doemoor/moci/internal/database"
)

type ApiCfg struct {
	DbQ *database.Queries
	DB  *sql.DB
}
