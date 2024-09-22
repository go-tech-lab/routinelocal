package impl

import (
  tls "github.com/go-tech-lab/framework_go_tls"
)

//以下代码属于RoutineLocal接口实现部分，对外不可见
type goroutineLocalImpl struct {
	name string
}

func (impl *goroutineLocalImpl) Name() string {
	return impl.name
}

func (impl *goroutineLocalImpl) Value() interface{} {
	ret, ok := tls.Get(impl.name)
	if !ok || ret == nil {
		return nil
	}
	return ret.Value()
}

func (impl *goroutineLocalImpl) Put(value interface{}) {
	tls.Set(impl.name, tls.MakeData(value))
}

func (impl *goroutineLocalImpl) Exist() bool {
	_, ok := tls.Get(impl.name)
	return ok
}

func (impl *goroutineLocalImpl) Remove() {
	tls.Del(impl.name)
}
