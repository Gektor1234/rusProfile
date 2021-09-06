package grpc

import (
	"context"
	"google.golang.org/grpc"
	"rusProfile/internal/app"
	grpcapi "rusProfile/internal/proto_buf"
)

type rpcHandlers struct {
	rusProfileLogic app.RusProfileLogic
}

func NewRPCHandlers(rusProfileLogic app.RusProfileLogic) *grpc.Server {
	g := rpcHandlers{rusProfileLogic: rusProfileLogic}
	grpcServer := grpc.NewServer()
	grpcapi.RegisterRusProfileServiceServer(grpcServer, g)
	return grpcServer
}

func (r rpcHandlers) GetCompanyByINN(ctx context.Context, req *grpcapi.GetCompanyByINNRequest) (*grpcapi.GetCompanyByINNResponse, error) {
	company, err := r.rusProfileLogic.GetCompanyByINN(req.INN)
	if err != nil {
		return &grpcapi.GetCompanyByINNResponse{}, err
	}
	result := grpcapi.Company{}
	result.INN = company.INN
	result.Name = company.Name
	result.CEOName = company.CEOName
	result.OGRN = company.OGRN
	return &grpcapi.GetCompanyByINNResponse{Company: &result}, nil
}
