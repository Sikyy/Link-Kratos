package data

import (
	"Link-Kratos/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// realTimeTrafficRepo .实现RealTimeTrafficRepo接口
type realTimeTrafficRepo struct {
	data *Data
	log  *log.Helper
}

func NewRealTimeTrafficRepo(data *Data, logger log.Logger) biz.RealTimeTrafficRepo {
	return &realTimeTrafficRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建实时流量信息
func (r *realTimeTrafficRepo) CreateRealTimeTraffic(ctx context.Context, re *biz.RealTimeTraffic) error {
	return nil
}

// interceptorRepo 实现InterceptorRepo接口
type interceptorRepo struct {
	data *Data
	log  *log.Helper
}

func NewInterceptorRepo(data *Data, logger log.Logger) biz.InterceptorRepo {
	return &interceptorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建拦截信息，进数据库
func (r *interceptorRepo) CreateInterceptor(ctx context.Context, in *biz.Interceptor) error {
	return nil
}

// 根据id获取拦截信息
func (r *interceptorRepo) GetInterceptorById(ctx context.Context, i *biz.Interceptor) error {
	return nil
}

// statisticsRepo 实现StatisticsRepo接口

type statisticsRepo struct {
	data *Data
	log  *log.Helper
}

func NewStatisticsRepo(data *Data, logger log.Logger) biz.StatisticsRepo {
	return &statisticsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建统计信息，进数据库
func (r *statisticsRepo) CreateStatistics(ctx context.Context, st *biz.Statistic) error {
	return nil
}

// 根据日期获取统计信息
func (r *statisticsRepo) GetStatisticsByDate(ctx context.Context, date string) (*biz.Statistic, error) {
	return nil, nil
}
