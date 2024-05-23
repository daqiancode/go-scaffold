package tables

import (
	"fmt"
	"go-scaffold/internals/drivers"

	"github.com/daqiancode/gormx"
)

func CreateTables() {
	fmt.Println("Create tables")
	db := drivers.GetDB()
	tables := []interface{}{&User{}}
	err := db.AutoMigrate(tables...)
	if err != nil {
		panic(err)
	}
	ddl := gormx.NewDDL(db)
	ddl.DefaultOnDelete = "RESTRICT"
	ddl.DefaultOnUpdate = "CASCADE"
	ddl.AddTables(tables...)
	ddl.MakeFKs()
	fmt.Println("Create tables OK")
}
