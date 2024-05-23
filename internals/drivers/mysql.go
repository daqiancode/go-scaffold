package drivers

import (
	"go-scaffold/logs"
	"log"
	"os"
	"time"

	"github.com/daqiancode/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _log = logs.GetLogger("drivers/mysql.go")

func CreateMysql() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	var datetimePrecision = 2
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       env.Get("MYSQL_URL", "root:123456@tcp(localhost:3306)/iam?charset=utf8&parseTime=True&loc=Local"), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                                                                                               // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                                              // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision,                                                                                // default datetime precision
		DontSupportRenameIndex:    true,                                                                                              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                              // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                             // smart configure based on used version
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		_log.Error().Err(err).Msg("failed create mysql connection")
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(env.GetIntMust("MYSQL_MAX_IDLE", 5))
	sqlDB.SetMaxOpenConns(env.GetIntMust("MYSQL_MAX_OPEN", 100))
	sqlDB.SetConnMaxLifetime(time.Duration(env.GetIntMust("MYSQL_MAX_LIFE_TIME", 30)) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(env.GetIntMust("MYSQL_MAX_IDLE_TIME", 5)) * time.Minute)

	return db, err
}

var db *gorm.DB

func GetDB() *gorm.DB {
	var err error
	if db == nil {
		db, err = CreateMysql()
		if err != nil {
			panic(err)
		}
	}
	return db
}
