package configuration

import "github.com/jmoiron/sqlx"

func NewPgDb(conf Config) *sqlx.DB {
	sdn := conf.Get("SDN")
	db := sqlx.MustConnect("postgres", sdn)
	return db
}
