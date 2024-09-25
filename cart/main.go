package main

import (
	"git.imooc.com/coding-447/cart/domain/repository"
	service2 "git.imooc.com/coding-447/cart/domain/service"
	"git.imooc.com/coding-447/cart/handler"
	"git.imooc.com/coding-447/common"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"

	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/opentracing/opentracing-go"

	cart "git.imooc.com/coding-447/cart/proto/cart"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var QPS = 100

func main() {
	// Configuration center
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// Registration center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Distributed tracing
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Database connection
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	// Create database connection
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// Disable plural table names
	db.SingularTable(true)

	// First time initialization
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		log.Error(err)
	}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		// Exposed service address
		micro.Address("0.0.0.0:8087"),
		// Registration center
		micro.Registry(consul),
		// Distributed tracing
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// Add rate limiting
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialize service
	service.Init()

	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))

	// Register Handler
	cart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
