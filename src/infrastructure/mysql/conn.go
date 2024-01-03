package mysql

import (
	"fmt"
	"guide_go/src/domain"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var repository *Repository

func InitMysql(mysql *domain.EnvironmentMysql) *Repository {
	fmt.Println("1")
	baseRepo := &BaseRepository{}
	baseRepo.ORMMysql = mysqlDb(mysql)
	repository = &Repository{
		baseRepo,
		&UserRepositoryImpl{
			baseRepo,
		},
	}
	return repository
}

func mysqlDb(Env *domain.EnvironmentMysql) *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        "root:1111@/guide?parseTime=true",
	}), func() gorm.Option {
		if os.Getenv("MODE") == "1" {
			return &gorm.Config{PrepareStmt: true, Logger: gormLogger.Default.LogMode(gormLogger.Info)}
		}
		return &gorm.Config{PrepareStmt: true}
	}())
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		os.Exit(0)
	}

	fmt.Println("Successfully connected to the Mysql")
	return db
}
