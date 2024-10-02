package require

func (require *Require) RegisterModule(id string, module Module) {
	require.module_map[id] = module.Exports()
}
