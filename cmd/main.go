package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"stock-service/conf"
	grpc2 "stock-service/delivery/grpc"
	serviceHttp "stock-service/delivery/http"
	"stock-service/migrations"
	proto "stock-service/proto/stock"
	"stock-service/repository"
	"stock-service/service/rabbitmq"
	"stock-service/usecase"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	conf.SetEnv()

	confMysql := conf.GetConfig().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort, confMysql.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()
	repo := repository.New(db)
	uc := usecase.New(repo)

	consum := RunConsumer(uc)
	defer consum.Close()

	go RunGRPC(uc)
	//migration
	migrations.Up(db)

	h := serviceHttp.NewHTTPHandler(uc)
	//go func() {
	//	h.Listener = httpL
	//	errs <- h.Start("")
	//}()
	if err := h.Start("0.0.0.0:8090"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
func RunConsumer(uc *usecase.UseCase) *rabbitmq.Consumer {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	consumer := rabbitmq.NewConsumer(conn, uc)
	consumer.StartConsumer()
	return consumer
}
func RunGRPC(uc *usecase.UseCase) {
	port := ":8091"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterStockServiceServer(s, &grpc2.ServerGRPC{UseCase: uc})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
