package main

import (
	"github.com/andygeiss/tinygo/internal/app/services"
)

func main() {
	services.NewHomeService().Run()
}
