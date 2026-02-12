package main

import (
	"fmt"

	"github.com/stephenz22/suangongshi/api"
	"github.com/stephenz22/suangongshi/config"
	"github.com/stephenz22/suangongshi/internal/database"
	"github.com/stephenz22/suangongshi/server"
)

func main() {
	fmt.Println("Welcome to 算工时 Server.")
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("Failed to initialize configuration: %s\n", err)
		return
	}

	fmt.Printf("Configuration loaded successfully: %+v\n", config.GlobalConfig)

	// init database

	db := database.InitDB(config.GlobalConfig.Database.DSN)
	gin_engine := api.RegisterRouters(db)
	srv := server.New(db, gin_engine)

	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	if err := srv.Run(addr); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}

}
