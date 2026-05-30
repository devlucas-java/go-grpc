package main

import (
	"database/sql"
	"net"

	"github.com/devlucas-java/go-grpc/internal/delivery/grpc/interceptor"
	"github.com/devlucas-java/go-grpc/internal/delivery/grpc/pb"
	"github.com/devlucas-java/go-grpc/internal/infra/database"
	"github.com/devlucas-java/go-grpc/internal/infra/migration"
	"github.com/devlucas-java/go-grpc/internal/service"
	grpcgo "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m := migration.NewMigration(db)
	if err := m.Run(); err != nil {
		panic(err)
	}

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)

	grpcServer := grpcgo.NewServer(
		grpcgo.ChainUnaryInterceptor(
			interceptor.LogIntercepto,
			interceptor.RateLimiter,
		),
	)
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
