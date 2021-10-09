package model

import (
	"fmt"
	"time"

	// "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
	// Docker *gorm.DB
}

var DB *Database
var DBIsReady bool

// func (db *Database) Init() error{
func Init(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname string) error {
	DBIsReady = false
	dbc, err := GetSelfDB(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	if err != nil {
		return err
	}
	DB = &Database{
		Self: dbc,
	}
	return nil
}

func New(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname string) (*Database, error) {
	// func New() error {
	DBIsReady = false
	dbc, err := GetSelfDB(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
	if err != nil {
		return nil, err
	}
	db := &Database{
		Self: dbc,
	}
	return db, nil
}

func GetSelfDB(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname string) (*gorm.DB, error) {
	return InitSelfDB(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname)
}

func InitSelfDB(dbtype, dbfile, dbuser, dbpass, dbhost, dbport, dbname string) (*gorm.DB, error) {
	switch { //dbtype;
	case dbtype == "mysql":
		return openDBMysql(dbuser, dbpass, dbhost, dbport, dbname)
	case dbtype == "postgres":
		return openDBPostgreSQL(dbuser, dbpass, dbhost, dbport, dbname)
	default:
		return openDBSqlite(dbfile)
	}
}

func openDBSqlite(dbfile string) (*gorm.DB, error) { //dbFileName string
	// filename := viper.GetString("db.filename")
	filename := "db.kns"
	if len(dbfile) > 0 {
		filename = dbfile
	}
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		// log.Errorf(err, "Database connection failed.:(%s)", filename)
		return nil, err
	}
	//set for db connection
	setupDB(db)
	return db, nil
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}
func openDBMysql(dbuser, dbpass, dbhost, dbport, dbname string) (*gorm.DB, error) { //user, pass, host, port, dbname string
	// user := viper.GetString("db.username")
	// pass := viper.GetString("db.password")
	// host := viper.GetString("db.host")
	// port := viper.GetString("db.port")
	// dbname := viper.GetString("db.name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Errorf(err, "Database connection failed:(%s)", dsn)
		return nil, err
	}
	//set for db connection
	setupDB(db)
	return db, nil
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func openDBPostgreSQL(dbuser, dbpass, dbhost, dbport, dbname string) (*gorm.DB, error) { //username, password, host, port, name string
	// dbuser := viper.GetString("db.username")
	// dbpass := viper.GetString("db.password")
	// dbhost := viper.GetString("db.host")
	// dbport := viper.GetString("db.port")
	// dbname := viper.GetString("db.name")
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=disable TimeZone=Asia/Shanghai", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Errorf(err, "Database connection failed.:(%s)", dsn)
		return nil, err
	}
	//set for db connection
	setupDB(db)
	return db, nil
	// dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func setupDB(db *gorm.DB) error {
	DBIsReady = true
	sqlDB, err := db.DB()
	// sqlDB, err := db.DB()
	if err != nil {
		// log.Errorf(err, "Database connection setupDB failed.")
		return err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func (db *Database) Close() error {
	DBIsReady = false
	sqlDB, err := db.Self.DB()
	if err != nil {
		// log.Errorf(err, "Database connection DB() failed.")
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		// log.Errorf(err, "Database connection close failed.")
		return err
	}
	// fmt.Println("DB close success")
	return nil
}

func Close() error {
	DBIsReady = false
	sqlDB, err := DB.Self.DB()
	if err != nil {
		// log.Errorf(err, "Database connection DB() failed.")
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		// log.Errorf(err, "Database connection close failed.")
		return err
	}
	// fmt.Println("DB close success")
	return nil
}
