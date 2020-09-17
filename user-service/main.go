package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	database "user-service/db"
	"user-service/handler"
	pb "user-service/proto/user"
	repository "user-service/repo"
)

func main() {
	//创建数据库连接，程序退出时断开连接
	db, err := database.CreateConnection()
	defer db.Close()

	if err != nil{
		log.Fatalf("Could not connect to DB: %v", err)
	}

	//每次启动服务时都会检查，如果数据表不存在则自动创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})

	//初始化 Repo 实例，用于后续数据库操作
	repo := &repository.UserRepository{db}


	// 创建micro微服务
	service := micro.NewService(
		micro.Name("sam.gxc.service.user"),
		micro.Version("latest"),
	)


	// Initialise service
	service.Init()

	// Register Handler
	pb.RegisterUserServiceHandler(service.Server(),&handler.UserService{repo})

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("sam.gxc.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
