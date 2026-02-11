package main

import (
	"fmt"

	"github.com/stephenz22/suangongshi/config"
)

func main() {
	fmt.Println("Welcome to 算工时 Server.")
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("Failed to initialize configuration: %s\n", err)
		return
	}

	fmt.Printf("Configuration loaded successfully: %+v\n", config.GlobalConfig)
}
