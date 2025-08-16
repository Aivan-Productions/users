package main

import (
	"fmt"

	"github.com/Aivan-Productions/users/tools/configuration"
)

func main() {
	cfg := configuration.Load()

	fmt.Print(cfg.PORT)
}
