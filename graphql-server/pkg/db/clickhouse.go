package db

import (
	"database/sql"
	"graphql-server/pkg/utils"
	"log"
)

type eventCategory string

const (
	standup   eventCategory = "standup"
	manzai    eventCategory = "manzai"
	impromptu eventCategory = "impromptu"
)

type Show struct {
	name      string
	cover     string
	thumbnail string
	briefly   string
	events    []Event
	category  eventCategory
}

type Event struct {
	date     int64
	location string
	price    int32
}

var Clickhouse *sql.DB

func ClickhoustInit() {
	Clickhouse, err := sql.Open("clickhouse", "clickhouse://127.0.0.1:9000?dial_timeout=1s&compress=true&max_execution_time=60")

	if err != nil {
		log.Printf("%#v\n", err) // minioClient is now set up
	}
	utils.Use(Clickhouse)
}
