package db
import(
	_ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)
func Connect() (db *xorm.Engine,err error ){
	db, err = xorm.NewEngine("mysql", "test1:test1@tcp(localhost:3306)/test1?charset=utf8")
    if err != nil{
		return
	}
	db.DBMetas()
	return
}