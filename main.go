package main

import (
	busRoute "github.com/kmhtoo/busappoint/routes"
)

func main() {
	busRoute := busRoute.Router{}
	busRoute.Init()
}
