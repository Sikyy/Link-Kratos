package biz

import (
	"context"
	"time"

	v1 "Link-Kratos/api/link/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// 实时流量信息
type RealTimeTraffic struct {
	Upflow   string
	Downflow string
}

// 拦截的请求信息
type Interceptor struct {
	Id         int    //自生成id
	User_Agent string //请求头信息
	Method     string //请求方法
	Host       string //请求主机
	Date       string //请求时间
	ClientName string //客户端名称
	Status     string //请求状态
	Upload     int64  //上传流量大小
	Download   int64  //下载流量大小
	Time       string //请求耗时
}

// 统计的流量信息
type Statistic struct {
	AllUpflow   string
	AllDownflow string
}

type RealTimeTrafficRepo interface {
	// 创建实时流量分析
	CreateRealTimeTraffic(ctx context.Context, rt *RealTimeTraffic) error
}

type InterceptorRepo interface {
	//创建流量拦截信息
	CreateInterceptor(ctx context.Context, in *Interceptor) error
	//根据id获取流量拦截信息
	GetInterceptorById(ctx context.Context, i *Interceptor) error
}

type StatisticsRepo interface {
	//创建流量统计信息
	CreateStatistics(ctx context.Context, st *Statistic) error
	//根据日期获取流量统计信息
	GetStatisticsByDate(ctx context.Context, date string) (*Statistic, error)
}

type TrafficUsecase struct {
	re RealTimeTrafficRepo
	in InterceptorRepo
	st StatisticsRepo

	log *log.Helper
}

// 创建一个TrafficUsecase
func NewTrafficUsecase(re RealTimeTrafficRepo, in InterceptorRepo, st StatisticsRepo, logger log.Logger) *TrafficUsecase {
	return &TrafficUsecase{re: re, in: in, st: st, log: log.NewHelper(logger)}
}

// 创建实时流量信息
func (uc *TrafficUsecase) CreateRealTimeTraffic(ctx context.Context, re *RealTimeTraffic) error {
	err := uc.re.CreateRealTimeTraffic(ctx, re)
	if err != nil {
		return err
	}
	return nil
}

// 创建拦截信息
func (uc *TrafficUsecase) CreateInterceptor(ctx context.Context, in *Interceptor) error {
	err := uc.in.CreateInterceptor(ctx, in)
	if err != nil {
		return err
	}
	return nil
}

// 根据id获取拦截信息
func (uc *TrafficUsecase) GetInterceptorById(ctx context.Context, id int) (*Interceptor, error) {
	i := &Interceptor{
		Id:         id,
		User_Agent: "User_Agent",
		Method:     "200",
		Host:       "bilibili.com:443",
		Date:       time.Now().String(),
	}
	if err := uc.in.GetInterceptorById(ctx, i); err != nil {
		return nil, err
	}
	return &Interceptor{
		Id:         i.Id,
		User_Agent: i.User_Agent,
		Method:     i.Method,
		Host:       i.Host,
		Date:       i.Time,
		ClientName: "哔哩哔哩",
		Status:     "活动中",
		Upload:     482348346,
		Download:   649839493275,
		Time:       "3000ms",
	}, nil
}

// 创建统计信息
func (uc *TrafficUsecase) CreateStatistics(ctx context.Context, st *Statistic) error {
	err := uc.st.CreateStatistics(ctx, st)
	if err != nil {
		return err
	}
	return nil
}

// 根据日期获取统计信息
func (uc *TrafficUsecase) GetStatisticsByDate(ctx context.Context, date string) (*Statistic, error) {
	st, err := uc.st.GetStatisticsByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	return st, nil
}
