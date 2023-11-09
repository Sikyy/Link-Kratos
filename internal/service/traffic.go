package service

import (
	v1 "Link-Kratos/api/link/v1"
	"context"
)

// 拦截流量并返回拦截信息
func (s *LinkService) Interceptor(c context.Context, req *v1.InterceptorRequest) (*v1.InterceptorReply, error) {
	return &v1.InterceptorReply{
		Interceptor: &v1.InterceptorReply_Interceptor{
			Date:       "Now",
			ClientName: "Yuan Shen",
			Status:     "Start",
			Method:     "GET",
		},
	}, nil
}

// 统计流量信息，今日，今周，今月，今年
func (s *LinkService) Statistics(c context.Context, req *v1.StatisticsRequest) (*v1.StatisticsReply, error) {
	return &v1.StatisticsReply{
		Statistics: &v1.StatisticsReply_Statistics{
			Upflow:   "100M",
			Downflow: "1000M",
		},
	}, nil
}

// 显示实时流量信息，实时上下行
func (s *LinkService) RealTimeTraffic(c context.Context, req *v1.RealTimeTrafficRequest) (*v1.RealTimeTrafficReply, error) {
	return &v1.RealTimeTrafficReply{
		RealtimeTraffic: &v1.RealTimeTrafficReply_RealTimeTraffic{
			Upflow:   "1000M",
			Downflow: "1000M",
		},
	}, nil
}
