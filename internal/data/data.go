package data

import (
	"Link-Kratos/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRealTimeTrafficRepo, NewInterceptorRepo, NewStatisticsRepo, NewUserRepo, NewProfileRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	// 连接数据库,传入配置文件里的数据库配置
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
	})
	if err != nil {
		panic(err)
	}
	// InitDB(db)
	return db
}

// AutoMigrate用于自动创建表
// func InitDB(db *gorm.DB) {
// 	if err := db.AutoMigrate(&User{}); err != nil {
// 		panic(err)
// 	}
// }
