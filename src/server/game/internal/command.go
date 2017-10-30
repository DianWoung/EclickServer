package internal

import (
	"fmt"
)

func init() {
	skeleton.RegisterCommand("echo", "echo user inputs", commandEcho)
	skeleton.RegisterCommand("leaf","test command",commandLeaf)
}

func commandEcho(args []interface{}) interface{} {
	return fmt.Sprintf("%v", args)
}

func commandLeaf(args []interface{}) interface{} {
	return fmt.Sprintf("%s",args)
}