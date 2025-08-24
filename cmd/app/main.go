package main

import (
	"flag"
	"fmt"

	"github.com/Aivan-Productions/users/tools/configuration"
)

func main() {

	envFile := flag.String("env", ".env.local", "path to env file")
	flag.Parse()

	cfg := configuration.LoadConfig(*envFile)

	fmt.Println(cfg.PORT)
	fmt.Println(cfg.LOG_FORMAT)
}
