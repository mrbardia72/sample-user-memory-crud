package main

import (
	"sample-memory/controllers"
)

func main() {
	controllers.PrintHeader()
	controllers.FillInitialData(4)
	controllers.PrintMenu()
}
