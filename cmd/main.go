package main

import (
	"fmt"

	runtime "github.com/rabbit-backend/cloud-runtime"
)

func main() {
	runtime := runtime.CreateRuntime()

	runtime.Set("log", fmt.Println)

	runtime.RunString(`
		const http = require("rabbit:http")
		const env = require("rabbit:env")

		const response = http.get("https://httpbin.zcorky.com/get")

		log("response", response);
		log("username", env.getEnv("USERNAME"))
	`)
}
