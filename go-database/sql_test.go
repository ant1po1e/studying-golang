package go_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background() 

	script := "INSERT INTO customer(id, name) VALUES('niga', 'niga')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Inserting Data to Database Succeded")
}


func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background() 

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID: ", id, "||", "Name: ", name)
		
	}

	defer rows.Close()
}