package api

import (
	"github.com/doemoor/moci/internal/database"
)

type ApiConfig struct {
	DbQ *database.Queries
}
