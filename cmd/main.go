package main

import (
	"fmt"
	"os"

	"github.com/jeffleon1/ionix/config"
	"github.com/jeffleon1/ionix/internal/app"
)

// @title     Ionix API
// @version   1.0

// @host      localhost:8080
// @BasePath  /

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
		os.Exit(1)
	}
	myApp := &app.App{}
	myApp.Initialize(cfg)
	myApp.Run()
}
