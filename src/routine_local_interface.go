package routinelocal

type IRoutineLocal interface {
	//Name key of value
	Name() string

	//Value value from goroutine local context
	Value() interface{}

	//Put value to goroutine local context
	Put(value interface{})

	//Exist value in goroutine local context
	Exist() bool

	//Remove value from goroutine local context
	Remove()
}
