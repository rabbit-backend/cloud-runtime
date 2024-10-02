/*
Description: invoke the go routines from javascript to leverage concurrency
*/
package task_runner

import "github.com/dop251/goja"

type TaskRunnerModule struct {
	vm *goja.Runtime
}
