package http

func (http *HttpModule) Exports() map[string]interface{} {
	return map[string]interface{}{
		"get":    http.Get,
		"post":   http.Post,
		"put":    http.Put,
		"delete": http.Delete,
	}
}
