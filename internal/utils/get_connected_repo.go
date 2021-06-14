package utils

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/ozoncp/ocp-classroom-api/internal/repo"
	"github.com/rs/zerolog/log"
)

const (
	user     = "postgres"
	password = "postgres"
	address  = "localhost:5432"
	dbName   = "ozon"
	sslmode  = "sslmode=disable"

	dataSourceName = "postgres://" + user + ":" + password + "@" + address + "/" + dbName + "?" + sslmode
)

func GetConnectedRepo(ctx context.Context) *repo.Repo {

	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal().Err(err).Msg("GetConnectedRepo: Failed to open postgres")
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("GetConnectedRepo: Failed to ping postgres")
	}

	log.Debug().Msgf("GetConnectedRepo: Connected to DB %v", dbName)

	classroomRepo := repo.New(db)

	return &classroomRepo
}
