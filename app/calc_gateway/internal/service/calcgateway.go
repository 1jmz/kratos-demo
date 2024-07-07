package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	pb "kratos-project/api/calc_gateway/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	cacul_pb "kratos-project/api/calc_engine/v1"
)

type CalcGatewayService struct {
	pb.UnimplementedCalcGatewayServer
}

func NewCalcGatewayService() *CalcGatewayService {
	return &CalcGatewayService{}
}

/*
calc_gateway 调用 calc_engine
*/
func (s *CalcGatewayService) GetSum(ctx context.Context, req *pb.GetSumRequest) (*pb.GetSumReply, error) {
	// 连接 Consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	dis := consul.New(client)
	endpoint := "discovery:///cacul"
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 调用calc_engine提供的GetSum方法
	c := cacul_pb.NewCalcEngineClient(conn)
	request := &cacul_pb.GetSumRequest{A: req.GetA(), B: req.GetB()}
	r, err := c.GetSum(context.Background(), request)
	if err != nil {
		panic(err)
		return nil, nil
	}
	return &pb.GetSumReply{A: req.GetA(), B: req.GetB(), Sum: r.GetSum()}, nil
}
