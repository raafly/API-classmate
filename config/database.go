package config

import (
	"databsase/sql"
	"time"
	"github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	sql.Open("mysql",)
}