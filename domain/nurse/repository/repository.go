package repository

import (
	"database/sql"

	"github.com/mrakhaf/halo-suster/domain/nurse/interfaces"
)

type repoHandler struct {
	databaseDB *sql.DB
}

func NewRepository(databaseDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		databaseDB: databaseDB,
	}
}
