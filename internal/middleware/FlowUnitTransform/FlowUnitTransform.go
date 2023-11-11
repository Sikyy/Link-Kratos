package FlowUnitTransform

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type FlowUnit struct {
	Name   string  //流量单位名称
	Symbol string  //流量单位符号
	Factor float64 //流量单位换算因子
	Rate   float64 //流量单位速率

}

// FlowUnitTransform 是一个 Kratos 中间件，用于处理流量单位转换逻辑。
func FlowUnitTransform() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 获取请求流量大小
				UpflowSize := tr.RequestHeader().Get("upload")
				DownflowSize := tr.RequestHeader().Get("download")

				// 转换流量单位
				UptransformedSize := TransformUnit(UpflowSize)
				DowntransformedSize := TransformUnit(DownflowSize)

				fmt.Println("上传流量大小: ", UptransformedSize, "下载流量大小: ", DowntransformedSize)
			}

			return handler(ctx, req)
		}
	}
}

// 根据需要实现流量单位转换的逻辑
// 2023/11/10 花了三小时,换了https://github.com/shopspring/decimal库还是没搞定,先放着吧，换回原先的逻辑了，能用就行
func TransformUnit(size string) string {

	// 检查 size 是否为空
	if size == "" {
		log.Println("输入流量大小为空")
		return "" // 或者返回默认值
	}

	log.Printf("输入流量大小: %s"+" B", size)

	// 定义一个流量单位的映射表
	unitToFactor := map[string]int64{
		"B":  1,
		"KB": 1024,
		"MB": 1048576,
		"GB": 1073741824,
		"TB": 1099511627776,
		"PB": 1125899906842624,
	}

	// 转换流量大小为int64类型
	newsize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		log.Printf("转换格式错误: %v", err)
	}

	// 根据流量大小确定流量单位
	var unit string
	for k, v := range unitToFactor {
		if newsize >= v {
			unit = k
			break
		}
	}

	// 转换流量大小为 float64 类型
	floatNewSize := float64(newsize)
	// 计算转换后的流量大小,无单位
	newfloatsize := floatNewSize / float64(unitToFactor[unit])

	result := fmt.Sprintf("%s %s", strconv.FormatFloat(newfloatsize, 'f', 2, 64), unit)

	// 返回转换后的流量大小和单位
	return result
}
