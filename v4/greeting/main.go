package main

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/mwitkow/go-grpc-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/golovers/grpc/v4/api-golang/gr/greeting/v1"
	"github.com/golovers/grpc/v4/greeting/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"github.com/golovers/grpc/v4/greeting/conf"
	"google.golang.org/grpc/credentials"
)




func main() {
	conf := conf.Config{
		HTTPAddress:":8080",
		GRPCAddress:":8000",
		PrometheusAddress:":8081",
		TLSCert:"certs/server.crt",
		TLSKey: "certs/server.key",
		TLSHostName:"localhost",
	}
	// gRPC
	lis, err := net.Listen("tcp", conf.GRPCAddress)
	if err != nil {
		log.Fatal("Failed to start greeting", err)
	}
	// TLS
	creds, err := credentials.NewServerTLSFromFile(conf.TLSCert, conf.TLSKey);
	if err != nil {
		log.Println("Failed to load TLS", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
			),
		),
		grpc.Creds(creds),
	)
	greetings := service.NewGreetingService()
	pb.RegisterGreetingsServer(server, greetings)

	// HTTP
	mux := runtime.NewServeMux()
	dopts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, conf.TLSHostName)),
	}
	if err := pb.RegisterGreetingsHandlerFromEndpoint(context.Background(), mux, conf.GRPCAddress, dopts); err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	go server.Serve(lis)

	// Start Metrics server
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(conf.PrometheusAddress, nil)

	// Start REST API GW
	if err := http.ListenAndServeTLS(conf.HTTPAddress, conf.TLSCert, conf.TLSKey, mux); err != nil {
		log.Fatal(err)
	}

}
