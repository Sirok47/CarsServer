package main

import (
	"context"
	"fmt"
	protocol "github.com/Sirok47/CarsServer/protocol"
	"github.com/Sirok47/CarsServer/service"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"net"
)

func main() {
	dbconn, err := pgx.Connect(context.Background(), "postgres://maks:glazirovanniisirok@127.0.0.1:5432/cars")
	if err != nil {
		fmt.Print(err)
		return
	}
	srv := grpc.NewServer()
	srvobj := service.NewService(dbconn)
	protocol.RegisterCarsServer(srv, srvobj)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Print(err)
	}
	err = srv.Serve(l)
	if err != nil {
		fmt.Print(err)
	}
}
