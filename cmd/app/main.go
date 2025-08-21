package main

import (
	"flag"
	"fmt"

	"github.com/Aivan-Productions/users/tools/configuration"
)

func main() {
	// parsing the flags that were passed from the command line
	envFile := flag.String("env", ".env.local", "path to env file")
	flag.Parse()

	// loading variables from files into the structure
	cfg := configuration.Load(*envFile)

	// tests
	fmt.Println(cfg.PORT)
	fmt.Println(cfg.LOG_FORMAT)
}
