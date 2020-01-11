package main

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	epb "github.com/vladfin1/grpc/proto/service/employeepb"
	pb "github.com/vladfin1/grpc/proto/service/unitpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run(address string, opts ...runtime.ServeMuxOption) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	gw := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterUnitServiceHandlerFromEndpoint(ctx, gw, ":8080", dialOpts)
	if err != nil {
		return err
	}
	err1 := epb.RegisterEmployeeServiceHandlerFromEndpoint(ctx, gw, ":8080", dialOpts)
	if err1 != nil {
		return err1
	}
	return http.ListenAndServe(":8090", gw)
}

func main() {
	defer glog.Flush()
	if err := run("gw:8080"); err != nil {
		glog.Fatal(err)
	}
}
