package storage

import (
	"context"
	"cushon/ent"
	"cushon/pkg/config"
	"log"
)

func newClient(cfg *config.Config) *ent.Client {
	connStr := "host=" + cfg.PostgresHost +
		" port=" + cfg.PostgresPort +
		" user=" + cfg.PostgresUser +
		" dbname=" + cfg.PostgresDBName +
		" password=" + cfg.PostgresPass +
		" sslmode=disable"
	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
