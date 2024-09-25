package main

import (
	"fmt"

	"git.imooc.com/coding-447/user/domain/repository"
	service2 "git.imooc.com/coding-447/user/domain/service"
	"git.imooc.com/coding-447/user/handler"
	user "git.imooc.com/coding-447/user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	srv.Init()

	db, err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

	//rp:=repository.NewUserRepository(db)
	//rp.InitTable()

	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
