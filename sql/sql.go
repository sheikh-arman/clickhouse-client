package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func sqlConnect() {
	conn, err := sql.Open("clickhouse", fmt.Sprintf("clickhouse://%s:%d?username=%s&password=%s", "click-sample.click.svc.cluster.local", "9000", "default", "lgE_kkgMcyUDTQnz"))
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection check Successful")
}
