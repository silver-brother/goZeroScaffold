package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"goZeroScaffold/common/gormx"
	"goZeroScaffold/internal/config"
	"goZeroScaffold/internal/middleware"
	"goZeroScaffold/internal/model/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	AuthInterceptor rest.Middleware

	DB *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: gormx.NewLog(true),
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:          c,
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,

		DB: query.Use(dbConn),
	}
}
