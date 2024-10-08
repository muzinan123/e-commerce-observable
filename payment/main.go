package main

import (
	"git.imooc.com/coding-447/common"
	"git.imooc.com/coding-447/payment/domain/repository"
	service2 "git.imooc.com/coding-447/payment/domain/service"
	"git.imooc.com/coding-447/payment/handler"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	payment "git.imooc.com/coding-447/payment/proto/payment"
)

func main() {

	consulConfig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		common.Error(err)
	}

	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.service.payment", "localhost:6831")
	if err != nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		common.Error(err)
	}
	defer db.Close()

	db.SingularTable(true)

	tableInit := repository.NewPaymentRepository(db)
	tableInit.InitTable()

	common.PrometheusBoot(9089)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),

		micro.Registry(consul),

		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),

		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),

		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	paymentDataService := service2.NewPaymentDataService(repository.NewPaymentRepository(db))

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), &handler.Payment{PaymentDataService: paymentDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
