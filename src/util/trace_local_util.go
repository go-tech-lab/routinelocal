package trace

import (
	"fmt"
	routinelocal "github.com/go-tech-lab/routinelocal/src"
	"time"
)

var traceContext = routinelocal.NewGoroutineLocal("traceId")

//InitTraceId 初始化traceid
func InitTraceId(traceId string) {
	var _traceId = traceId
	traceContext.Put(_traceId)
}

//GetTraceId 获取traceid，没有则返回空字符串
func GetTraceId() string {
	ret := traceContext.Value()
	if ret == nil {
		traceId := fmt.Sprintf("%d", time.Now().UnixNano())
		traceContext.Put(traceId)
		return traceId
	}
	return ret.(string)
}
