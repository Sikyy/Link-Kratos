package server

import (
	v1 "Link-Kratos/api/link/v1"
	"Link-Kratos/internal/conf"
	"Link-Kratos/internal/middleware/jwt"
	"Link-Kratos/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// 初始化一个允许列表，可以在这里添加跳过认证的路由
func NewAllowListMatcher() selector.MatchFunc {

	allowList := make(map[string]struct{})
	allowList["/link.v1.Link/Login"] = struct{}{}
	allowList["/link.v1.Link/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := allowList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, link *service.LinkService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			//设置日志
			recovery.Recovery(),
			//设置认证，token
			selector.Server(jwt.JWTAuth(jwtc.Token)).Match(NewAllowListMatcher()).Build(),
		),
		//jwt.JWTAuth(jwtc.Token),
		//FlowUnitTransform.FlowUnitTransform(),
		//设置跨域
		http.Filter(handlers.CORS(
			//允许的Header的类型
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			//允许的请求方式
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			//允许的域名
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterLinkHTTPServer(srv, link)
	return srv
}
