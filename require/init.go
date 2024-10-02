package require

import "github.com/rabbit-backend/cloud-runtime/modules"

func (require *Require) Init() {
	for id, module := range modules.CreateFactory(require.VM) {
		require.RegisterModule(id, module)
	}

	require.VM.Set("require", require.GetModule)
}
