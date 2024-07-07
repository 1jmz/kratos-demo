package service

import (
	"context"
	"fmt"
	pb "kratos-project/api/calc_engine/v1"
)

type CalcEngineService struct {
	pb.UnimplementedCalcEngineServer
}

func NewCalcEngineService() *CalcEngineService {
	return &CalcEngineService{}
}

func (s *CalcEngineService) GetSum(ctx context.Context, req *pb.GetSumRequest) (*pb.GetSumReply, error) {
	// 业务逻辑
	sum := req.GetA() + req.GetB()
	fmt.Printf("调用了[内部gRPC服务: CalcEngineService.GetSum] %d + %d = %d \n", req.GetA(), req.GetB(), sum)
	return &pb.GetSumReply{Sum: sum}, nil
}
