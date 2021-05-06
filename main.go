package main

import (
	"github.com/Sirok47/CarsServer/protocol"
	"github.com/Sirok47/CarsServer/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"net"
)

func main(){
	dbconn, err := pgx.Connect(context.Background(), "postgres://maks:glazirovanniisirok@127.0.0.1:5432/cars")
	if err!= nil{
		fmt.Print(err)
		return
	}
	srv:=grpc.NewServer()
	srvobj:=service.NewService(dbconn)
	protocol.RegisterCarsServer(srv,srvobj)
	l,_:=net.Listen("tcp",":8080")
	srv.Serve(l)
}