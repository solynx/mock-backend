package database

import (
	"fmt"

	// "app/models/courses"
	"app/database/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME_STR = "root"
const DB_PASSWORD_STR = "my-secret-pw"
const DB_NAME_STR = "db_mockapiserver"
const DB_HOST_NUM = "localhost"
const DB_PORT_NUM = "3306"

var Database *gorm.DB

func InitDb() *gorm.DB {
	Database = ConnectDBLocal()
	return Database
}

func ConnectDBLocal() *gorm.DB {
	var err error
	// timeZone := "Asia/Ho_Chi_Minh"

	dsn := DB_USERNAME_STR + ":" + DB_PASSWORD_STR + "@tcp(" + DB_HOST_NUM + ":" + DB_PORT_NUM + ")/" + DB_NAME_STR + "?loc=Asia%2FBangkok&parseTime=true&time_zone=%27Asia%2FBangkok%27"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Err:  %v", err)
		return nil
	}
	migrations.Migrate(db)
	return db
}
