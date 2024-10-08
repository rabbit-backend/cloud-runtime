package main

import (
	"fmt"

	runtime "github.com/rabbit-backend/cloud-runtime"
)

func main() {
	runtime := runtime.New()

	runtime.Set("log", fmt.Println)

	runtime.RunString(`
		const http = require("http")
		const env = require("env")
		const fs = require("fs")

		const response = http.get("https://httpbin.zcorky.com/get")

		log("response", http.json(response));
		log("pwd", env.getEnv("PWD"))

		const buffer = fs.read("go.mod")
		log(buffer)
	`)
}
