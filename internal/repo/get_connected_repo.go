package repo

import (
	"context"
	"database/sql"
	"flag"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/rs/zerolog/log"
)

const (
	RepoUser     = "postgres"
	RepoPassword = "postgres"
	RepoEndpoint = "127.0.0.1:5432"
	RepoDbName   = "postgres"
	RepoSslMode  = "disable"
)

type RepoArgs struct {
	User     string
	Password string
	Endpoint string
	DbName   string
	SslMode  string
}

func SetRepoArgsFromCommandLine(repoArgs *RepoArgs) {

	flag.StringVar(&repoArgs.User, "repo-user", RepoUser, "PostgreSQL username")
	flag.StringVar(&repoArgs.Password, "repo-password", RepoPassword, "PostgreSQL user's password")
	flag.StringVar(&repoArgs.Endpoint, "repo-endpoint", RepoEndpoint, "PostgreSQL server endpoint")
	flag.StringVar(&repoArgs.DbName, "repo-dbname", RepoDbName, "PostgreSQL DB name")
	flag.StringVar(&repoArgs.SslMode, "repo-sslmode", RepoSslMode, "PostgreSQL ssl mode")
}

func GetConnectedRepo(ctx context.Context, repoArgs *RepoArgs) (Repo, error) {

	dataSourceName := repoArgsToDataSourceName(repoArgs)

	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Error().Err(err).Msg("GetConnectedRepo: Failed to open postgres")
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		log.Error().Err(err).Msgf("GetConnectedRepo: Failed to ping postgres %v", dataSourceName)
		return nil, err
	}

	log.Debug().Msgf("GetConnectedRepo: Connected to DB %v", repoArgs.DbName)

	classroomRepo := New(db)

	return classroomRepo, nil
}

func repoArgsToDataSourceName(repoArgs *RepoArgs) string {

	dataSourceName := "postgres://" + repoArgs.User + ":" +
		repoArgs.Password + "@" + repoArgs.Endpoint + "/" +
		repoArgs.DbName + "?sslmode=" + repoArgs.SslMode

	return dataSourceName
}
