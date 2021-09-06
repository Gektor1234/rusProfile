package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	grpcRusProfile "rusProfile/internal/grpc"
	"rusProfile/internal/logic"
	"strconv"
)

func main() {
	rusProfileLogic := logic.NewRusProfileLogic()
	grpcServer := grpcRusProfile.NewRPCHandlers(rusProfileLogic)
	StartGRPC(grpcServer, 7898)
	select {}
}

func StartGRPC(s *grpc.Server, port int) {
	fmt.Printf("START GRPC SERVER ON PORT %v", port)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err, "can't listen tcp on port ", port)
	}
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatal(err, "can't serve grpc server")
		}
	}()
}
