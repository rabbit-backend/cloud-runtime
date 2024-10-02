package http

import "github.com/go-zoox/fetch"

func (http *HttpModule) Post(url string, config *fetch.Config) *fetch.Response {
	response, err := fetch.Post(url, config)
	if err != nil {
		panic(http.vm.ToValue(err))
	}

	return response
}
