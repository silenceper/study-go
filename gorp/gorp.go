package main

import(
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"fmt"
)

type Category struct {
	Id          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Utime       int64  `db:"utime"`
	Ctime       int64  `db:"ctime"`
}


func main() {
	dbmap := initDb()
	var cates []Category
	_, err := dbmap.Select(&cates, `select * from category where id=:id`,map[string]interface{}{
		"id":1,
	})
	fmt.Println(cates)
	fmt.Println(err)
}

func initDb() *gorp.DbMap {
    db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/goBlog?charset=utf8")

    dbmap :=&gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-blog:", log.Lmicroseconds))

    dbmap.AddTableWithName(Category{}, "category").SetKeys(true, "Id")

    //err = dbmap.CreateTablesIfNotExists()

    if err!=nil{
    	fmt.Println(err)
    }
    return dbmap
}