package runtime

import (
	"github.com/dop251/goja"
	"github.com/rabbit-backend/cloud-runtime/require"
)

func New() *goja.Runtime {
	vm := goja.New()

	require := require.New(vm)
	require.Init()

	return vm
}
