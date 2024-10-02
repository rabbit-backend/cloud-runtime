package require

func (require *Require) GetModule(id string) interface{} {
	return require.module_map[id]
}
