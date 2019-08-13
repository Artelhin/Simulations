package main

import "runtime"

const (
	width = 500
	height = 500

	rows = 50
	columns = 50
)

func init() {
	runtime.LockOSThread()
}

func main() {}