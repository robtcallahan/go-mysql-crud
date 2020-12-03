package main

import (
	"fmt"
	"os"

	"go-mysql-crud/pkg/driver"
	"go-mysql-crud/pkg/handler"
)

func main() {
	params := &driver.ConnectParams{
		DBName: "register",
		User:   "register",
		Pass:   "dinx9one",
		Host:   "localhost",
		Port:   "3306",
	}

	connection, err := driver.ConnectSQL(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	pHandler := handler.NewQueryHandler(connection)
	h := pHandler.Fetch
	cols := h()
	for _, c := range cols {
		fmt.Println(c.Name)
	}
}
