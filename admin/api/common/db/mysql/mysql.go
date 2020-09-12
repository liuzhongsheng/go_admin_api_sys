package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goAdminApi/config"
)
var _db *gorm.DB
func init(){
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	host,err := config.Cfg.Section("database").GetKey("MYSQL_HOST")
	if err != nil {
		fmt.Println("loade config fail")
	}

	port,err := config.Cfg.Section("database").GetKey("MYSQL_PORT")
	if err != nil {
		fmt.Println("loade config fail")
	}

	user,err := config.Cfg.Section("database").GetKey("MYSQL_USER")
	if err != nil {
		fmt.Println("loade config fail")
	}

	password,err := config.Cfg.Section("database").GetKey("MYSQL_PASSWORD")
	if err != nil {
		fmt.Println("loade config fail")
	}
	dbname,err := config.Cfg.Section("database").GetKey("MYSQL_DB_NAME")
	if err != nil {
		fmt.Println("loade config fail")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,password,host,port,dbname)

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Print("连接数据库失败, error=" + err.Error())
	}
	_db.SingularTable(true)
	_db.LogMode(true)
	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}


//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}