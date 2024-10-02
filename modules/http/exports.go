package http

func (http *HttpModule) Exports() map[string]interface{} {
	json := &JSON{
		vm: http.vm,
	}

	return map[string]interface{}{
		"get":    http.Get,
		"post":   http.Post,
		"put":    http.Put,
		"delete": http.Delete,
		"json":   json.Json,
	}
}
