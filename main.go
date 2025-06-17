package main

import (
	"log"

	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db"
	"github.com/jadermoura/caixa-eletronico-app/internal/ui/routes"
)

func main() {
	conn, err := db.ConnectPOstgress()

	if err != nil {
		log.Fatalf("CRITICAL: Failed to connect to the database: %v", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("Failed to get raw database connection: %v", err)
	}
	defer sqlDB.Close()

	routes.GetRoutes(conn)

}
