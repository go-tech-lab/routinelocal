package impl

import (
	tls "github.com/go-tech-lab/framework_go_tls"
	routinelocal "github.com/go-tech-lab/routinelocal/src"
)

/**
*  RoutineLocal  工厂接口实现，可参考其他语言的ThreadLocal
*  author guang
* date 2019-08-22
 */
func init() {
	factory := goroutineLocalFactory{allRoutineLocals: make([]routinelocal.IRoutineLocal, 0)}
	routinelocal.InjectRoutineLocalFactory(&factory)
}

//IRoutineLocalFactory 实现
type goroutineLocalFactory struct {
	allRoutineLocals []routinelocal.IRoutineLocal
}

func (factory *goroutineLocalFactory) NewRoutineLocal(_name string) routinelocal.IRoutineLocal {
	local := &goroutineLocalImpl{
		name: _name,
	}
	//factory.allRoutineLocals = append(factory.allRoutineLocals, local)
	////协程退出必须清理资源，否则会造成内存资源泄漏 TODO
	//registerRoutineLocalExitCallback(local)
	return local
}

func (factory *goroutineLocalFactory) ClearCurrent() {
	tls.Reset()
}
