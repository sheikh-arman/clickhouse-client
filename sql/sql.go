package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func SqlConnect() {
	log.Println("SQL Connection started -> ")
	conn, err := sql.Open("clickhouse", fmt.Sprintf("clickhouse://%s:%d?username=%s&password=%s", "click-sample.click.svc.cluster.local", 9000, "default", "lgE_kkgMcyUDTQnz"))
	if err != nil {
		log.Println("sql Connection error ->>")
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Println("sql Ping error ")
		log.Fatal(err)
	}

	_, err = conn.Exec("SELECT 1")
	if err != nil {
		log.Println("sql select eroro")
		log.Fatal(err)
	}
	log.Println("Connection checked Successful")
	// Execute a test query to fetch databases
	rows, err := conn.Query("SHOW DATABASES")
	if err != nil {
		log.Println("Error fetching databases:")
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and print each database name
	var databaseName string
	for rows.Next() {
		if err := rows.Scan(&databaseName); err != nil {
			log.Println("Error scanning database row:")
			log.Fatal(err)
		}
		fmt.Println(databaseName)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error iterating over database rows:")
		log.Fatal(err)
	}
	log.Println("Database list fetched successfully")
	log.Println("SqlConnect passd")
}
