package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=casbin sslmode=disable"

	a, err := gormadapter.NewAdapter("postgres", dsn, true)
	if err != nil {
		log.Fatalf("NewAdapter failed: %v\n", err)
	}

	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		log.Fatalf("NewEnforcer failed: %v\n", err)
	}

	if err := e.LoadPolicy(); err != nil {
		log.Fatalf("LoadPolicy failed: %v\n", err)
	}

	e.EnableAutoSave(true)

	_, _ = e.AddPolicy("dajun", "data1", "read")
	_, _ = e.AddPolicy("lizi", "data2", "write")

	ok1, _ := e.Enforce("dajun", "data1", "read")
	ok2, _ := e.Enforce("dajun", "data1", "write")
	ok3, _ := e.Enforce("lizi", "data2", "write")
	ok4, _ := e.Enforce("dajun", "data2", "read")
	fmt.Println("dajun data1 read  =", ok1)
	fmt.Println("dajun data1 write =", ok2)
	fmt.Println("lizi data2 write  =", ok3)
	fmt.Println("dajun data2 read  =", ok4)
}
