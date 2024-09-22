package routinelocal

const (
	DefaultRoutineLocalMapSegmentNum = 1000
)

type IRoutineLocalFactory interface {
	//NewRoutineLocal 创建RoutineLocal: name-名称
	NewRoutineLocal(name string) IRoutineLocal

	//ClearCurrent 协程退出时清除相关信息
	ClearCurrent()
}

//单例模式
var contextLocalFactory IRoutineLocalFactory

func InjectRoutineLocalFactory(_routineLocalFactory IRoutineLocalFactory) {
	contextLocalFactory = _routineLocalFactory
}

func RefRoutineLocalFactory() IRoutineLocalFactory {
	return contextLocalFactory
}

func NewGoroutineLocal(name string) IRoutineLocal {
	return contextLocalFactory.NewRoutineLocal(name)
}
