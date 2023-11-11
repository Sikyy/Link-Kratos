package interceptinformation

// import (
// 	"context"

// 	"github.com/go-kratos/kratos/v2/middleware"
// 	"github.com/go-kratos/kratos/v2/transport"
// )

// // InterceptorInformation 是一个 Kratos 中间件，用于处理流量拦截信息。
// func InterceptorInformation() middleware.Middleware {
// 	return func(handler middleware.Handler) middleware.Handler {
// 		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
// 			if tr, ok := transport.FromServerContext(ctx); ok {
// 				user_agent := tr.RequestHeader().Get("User-Agent")
// 				method := tr.RequestHeader().Get("Method")
// 				host := tr.RequestHeader().Get("Host")
// 				date := tr.RequestHeader().Get("Date")
// 			}
// 			return handler(ctx, req)
// 		}
// 	}

// }
