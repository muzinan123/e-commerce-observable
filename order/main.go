package main

import (
	"git.imooc.com/coding-447/common"
	"git.imooc.com/coding-447/order/domain/repository"
	service2 "git.imooc.com/coding-447/order/domain/service"
	"git.imooc.com/coding-447/order/handler"
	order "git.imooc.com/coding-447/order/proto/order"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	//qps = os.Getenv("QPS")
	QPS = 1000
)

func main() {
	// 1. Configuration center
	consulConfig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 2. Registry center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})
	// 3. Jaeger tracing
	t, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 4. Initialize database
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// Disable plural table names
	db.SingularTable(true)

	// Create tables on first run
	//tableInit := repository.NewOrderRepository(db)
	//tableInit.InitTable()

	// Create instance
	orderDataService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	// 5. Expose monitoring address
	common.PrometheusBoot(9092)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		// Service address to expose
		micro.Address(":9085"),
		// Add consul registry
		micro.Registry(consul),
		// Add tracing
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// Add rate limiting
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		// Add monitoring
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialize service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), &handler.Order{OrderDataService: orderDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
