package app

import (
	"bnqkl/chain-cms/config"
	"bnqkl/chain-cms/database/datasource"
	"bnqkl/chain-cms/docs"
	"bnqkl/chain-cms/helper"
	"bnqkl/chain-cms/logger"
	"bnqkl/chain-cms/middleware"
	"bnqkl/chain-cms/modules/attach"
	"bnqkl/chain-cms/modules/entity"
	"bnqkl/chain-cms/redis"
	"bnqkl/chain-cms/storage"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/gorm"
)

func NewApp() *gin.Engine {
	// 设置运行模式
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// 初始化根目录
	err := helper.InitRootPath()
	if err != nil {
		log.Println(err)
		panic("init root path error")
	}
	rootPath := helper.GetRootPath()
	// 初始化配置文件
	err = config.InitConfig(rootPath)
	if err != nil {
		log.Println(err)
		panic("init config error")
	}
	// 初始化日志读写器
	err = logger.InitLogger()
	if err != nil {
		panic(err)
	}
	log := logger.GetLogger()
	// 初始化附件仓库
	err = storage.InitStorage(log)
	if err != nil {
		log.Error(err)
		panic("init attach storage error")
	}
	// 初始化 redis
	err = redis.InitRedisDb(log)
	if err != nil {
		log.Error(err)
		panic("init redis error")
	}
	// 初始化 redis 分布式锁服务
	redis.InitRedisSync(log)
	// 初始化数据库
	db, err := datasource.InitDB(log)
	if err != nil {
		log.Error(err)
		panic("init db error")
	}

	// 注册中间件
	// Gin 中间件的执行顺序是按注册顺序执行的。中间件在路由分组之后添加，分组内的路由将不会继承该中间件。
	registerMiddleware(router, log)

	// 初始化模块
	initModule(db, log)
	group := router.Group("/api")
	// 注册 API Endpoint
	registerApi(group)

	// 注册静态文件服务
	registerStaticServer(router)

	// 注册 swagger
	registerSwagger(router)

	return router
}

func initModule(db *gorm.DB, log *logger.Logger) {
	attach.InitAttachModule(db, log)
	entity.InitEntityModule(db, log)
}

func registerApi(routerGroup *gin.RouterGroup) {
	attach.RegisterAttachApi(routerGroup)
	entity.RegisterEntityApi(routerGroup)
}

func registerMiddleware(router *gin.Engine, log *logger.Logger) {
	// 使用 zap 日志
	// router.Use(middleware.GinLogger(log), middleware.GinRecovery(log, true))
	// 跨域问题
	router.Use(cors.Default())
	// 打印耗时接口
	router.Use(middleware.NewApiTimmerMiddleware(log))
	// 限流
	router.Use(middleware.NewRateLimiterMiddleware())
}

func registerStaticServer(router *gin.Engine) {
	// rootPath := helper.GetRootPath()
	// 网站图标
	// router.Static("/favicon.ico", filepath.Join(rootPath, "favicon.ico"))
	// 静态文件夹
	// blob 仓库
	blobStorage := storage.GetBlobStorageDir()
	router.StaticFS("/blob", http.Dir(blobStorage))
}

func registerSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	mode := gin.Mode()
	if mode != gin.ReleaseMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func Run() {
	router := NewApp()
	config := config.GetConfig()
	srv := &http.Server{
		Addr:    config.Port,
		Handler: router,
	}
	log := logger.GetLogger()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server run fail: %s\n", err.Error())
		}
	}()
	// 等待中断信号以优雅的关闭服务器（设置 5 秒钟的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("server shutdown fail: ", err.Error())
	}
	log.Info("server shutdown success")
}
