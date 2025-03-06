package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/MohamedMosalm/banque/api"
	db "github.com/MohamedMosalm/banque/db/sqlc"
	"github.com/MohamedMosalm/banque/gapi"
	"github.com/MohamedMosalm/banque/pb"
	"github.com/MohamedMosalm/banque/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	conn, err := sql.Open(config.DBDeriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB: ", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBanqueServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("gRPC server failed to serve")
	}
}
