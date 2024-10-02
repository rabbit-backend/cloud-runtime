package runtime

import (
	"github.com/dop251/goja"
	"github.com/rabbit-backend/cloud-runtime/require"
)

func CreateRuntime() *goja.Runtime {
	vm := goja.New()

	require := require.NewRequire(vm)
	require.Init()

	return vm
}
