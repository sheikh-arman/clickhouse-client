package main

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	sqlconn "github.com/sheikh-arman/clickhouse-client/sql"
	"log"
	klog "log"
	"time"
	// _ "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func main() {

	klog.Println("\n\nStarted  ->>>>>>>>>> 1")
	clickhouseNativeConnect()
	sqlconn.SqlConnect()
	kubeDBConnect()

	time.Sleep(time.Hour * 24)
}

func clickhouseNativeConnect() {
	klog.Println("\n\nStarted clickhouseNative ->>>>>>>>>> ")
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
	klog.Println("\n\nclickhouseNative Passed ->>>>>>>>>> ")
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

func kubeDBConnect() {
	log.Println("KubeDB Connection started -> ")
	log.Println("SqlConnect passd")
}
