package main

import (
	"context"
	"fmt"
	"os"
 pgx "github.com/jackc/pgx/v5"

)



func main()  {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
  fmt.Println("db Connect")
	defer conn.Close(context.Background())



  var username string
  var password string


	err = conn.QueryRow(context.Background(), "select username, password from asalcoba where id=$1", 1).Scan(&username, &password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(username, password)

}
