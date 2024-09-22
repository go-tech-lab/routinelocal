package main

import (
	routinelocal "github.com/go-tech-lab/routinelocal/src"
	_ "github.com/go-tech-lab/routinelocal/src/impl"
	"time"
)

func main() {
	factory := routinelocal.RefRoutineLocalFactory()
	go func() {
		local1 := factory.NewRoutineLocal("local1")
		local2 := factory.NewRoutineLocal("local2")
		local1.Put("aaaaaaa1")
		local1.Put("aaaaaaa2")
		local2.Put("bbbbbbbb1")
		local2.Put("bbbbbbbb2")
	}()
	time.Sleep(5 * time.Second)
}
