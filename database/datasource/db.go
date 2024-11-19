package datasource

import (
	"bnqkl/chain-cms/config"
	"fmt"
	"log"
	"os"
	"time"

	"bnqkl/chain-cms/database/model"
	"bnqkl/chain-cms/helper"
	fileLogger "bnqkl/chain-cms/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB(_log *fileLogger.Logger) (*gorm.DB, error) {
	mode := gin.Mode()
	config := config.GetConfig()
	mysqlConfig := config.MySql
	// 自定义日志模板，打印 SQl 语句
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 色彩
		})
	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不加 s
		},
	}
	if mode != gin.ReleaseMode {
		gormConfig.Logger = logger
	}
	initDB := func() error {
		// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
		newDb, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dns,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gormConfig)
		if err != nil {
			return err
		}
		sqlDB, err := newDb.DB()
		if err != nil {
			return err
		}
		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(2000)
		sqlDB.SetConnMaxLifetime(time.Hour)
		db = newDb
		_log.Info("init mysql success")
		autoCreateTable()
		_log.Info("auto create table success")
		return nil
	}
	var err error
	for i := 0; i < 3; i++ {
		err = initDB()
		if err == nil {
			return db, err
		}
		_log.Error(err)
		time.Sleep(time.Second)
	}
	return db, err
}

func autoCreateTable() error {
	tables := []model.TableSchema{
		&model.Entity{},
	}
	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			continue
		}
		err := db.Set("gorm:table_options", table.GetTableOptions()).AutoMigrate(&table)
		if err != nil {
			return err
		}
	}
	return nil
}

func FindByPage[T any, U any](model *T, condition U, values []any, page int, pageSize int, list *[]T) (helper.Pagination, error) {
	var total int64 = 0
	result := db.Model(model).Where(condition, values...).Count(&total)
	if result.Error != nil {
		return helper.Pagination{}, result.Error
	}
	pagination := helper.NewPagination(page, pageSize, total)
	if total == 0 {
		return pagination, nil
	}
	result = db.Model(model).Where(condition).Limit(pageSize).Offset((page - 1) * pageSize).Find(list)
	return pagination, result.Error
}
