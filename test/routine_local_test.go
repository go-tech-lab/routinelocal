package routinelocal

import (
	"fmt"
	"github.com/go-tech-lab/framework_common/src/routine"
	routinelocal "github.com/go-tech-lab/routinelocal/src"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGoroutineGid(t *testing.T) {
	asserts := require.New(t)
	gid := routine.RoutineUniqNo()
	for i := 0; i < 100000; i++ {
		newGid := routine.RoutineUniqNo()
		asserts.Equal(gid, newGid, fmt.Sprintf("RoutineUniqNo not equal in the same goroutine: expected %d, actual: %d", gid, newGid))
	}
}

func TestGoroutineLocal(t *testing.T) {
	asserts := require.New(t)
	localName := "test_local_name"
	local := routinelocal.NewGoroutineLocal(localName)
	asserts.Nil(local.Value(), "local.Value() should be nil")
	asserts.False(local.Exist(), "local.Exist() should be false")
	asserts.Equal(localName, local.Name(), fmt.Sprintf("Routinelocal name not correct: expected %s, actual: %s", localName, local.Name()))
	localValue := "localValue"
	local.Put(localValue)
	asserts.NotNil(local.Value(), "local.Value() should be not nil")
	asserts.True(local.Exist(), "local.Exist() should be true")
	asserts.Equal(localValue, local.Value(), fmt.Sprintf("Routinelocal name not correct: expected %s, actual: %s", localValue, local.Value()))
	local.Remove()
	asserts.Nil(local.Value(), "local.Value() should be nil")
	asserts.False(local.Exist(), "local.Exist() should be false")
	asserts.Equal(nil, local.Value(), fmt.Sprintf("Routinelocal name not correct: expected %s, actual: %s", localValue, local.Value()))
}
