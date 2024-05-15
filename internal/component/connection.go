package component

import (
	"database/sql"
	"fmt"
	"golang-clean-architecture/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func GetDbConnection(conf config.Config) *sql.DB {
	strInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Pass, conf.DB.Name)
	db, err := sql.Open("postgres", strInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
