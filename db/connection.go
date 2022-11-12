package db

import (
	"database/sql"
	"fmt"

	"github.com/bertbr/poc-golang/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	con, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = con.Ping()

	return con, err
}
