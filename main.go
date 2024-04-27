package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	klog "log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {

	klog.Println("\n\nStarted  ->>>>>>>>>> 1")
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT name,toString(uuid) as uuid_str FROM system.tables LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			name, uuid string
		)
		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("name: %s, uuid: %s",
			name, uuid)
	}
	shw, err := conn.Query(ctx, "SHOW DATABASES")
	if err != nil {
		log.Println("Database showing errorError -> 00000000")
		log.Fatal(err)
	}
	var databaseName string
	for shw.Next() {
		log.Println("Show database -> ")
		if err := shw.Scan(&databaseName); err != nil {
			log.Println("Error scanning database row:")
			log.Fatal(err)
		}
		fmt.Println(databaseName)
	}
	sqlConnect()
}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"click-sample.click.svc.cluster.local:9000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "lgE_kkgMcyUDTQnz",
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			//TLS: &tls.Config{
			//	InsecureSkipVerify: false,
			//},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}

func sqlConnect() {
	log.Println("SQL Connection -> ")
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
	time.Sleep(time.Hour * 24)
}
