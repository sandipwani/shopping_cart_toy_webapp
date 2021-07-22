package model

import "database/sql"

type DBAccessor interface {
	ProductAccessor
	ProductCartAccessor
}

type WriteDbAccess struct {
	Db *sql.DB
}
