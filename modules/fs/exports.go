package fs

func (fs *FsModule) Exports() map[string]interface{} {
	return map[string]interface{}{
		"read": fs.Read,
	}
}
