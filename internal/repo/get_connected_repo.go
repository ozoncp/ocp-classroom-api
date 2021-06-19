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
	RepoEndpoint = "postgres:5432"
	RepoDbName   = "postgres"
	RepoSslMode  = "disable"
)

type RepoArgs struct {
	User     *string
	Password *string
	Endpoint *string
	DbName   *string
	SslMode  *string
}

func NewRepoArgs() *RepoArgs {
	return &RepoArgs{
		User:     flag.String("repo-user", RepoUser, "PostgreSQL username"),
		Password: flag.String("repo-password", RepoPassword, "PostgreSQL user's password"),
		Endpoint: flag.String("repo-address", RepoEndpoint, "PostgreSQL server endpoint"),
		DbName:   flag.String("repo-name", RepoDbName, "PostgreSQL DB name"),
		SslMode:  flag.String("repo-sslmode", RepoSslMode, "PostgreSQL ssl mode"),
	}
}

func GetConnectedRepo(ctx context.Context, repoArgs *RepoArgs) (Repo, error) {

	db, err := sql.Open("pgx", repoArgsToDataSourceName(repoArgs))
	if err != nil {
		log.Error().Err(err).Msg("GetConnectedRepo: Failed to open postgres")
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		log.Error().Err(err).Msg("GetConnectedRepo: Failed to ping postgres")
		return nil, err
	}

	log.Debug().Msgf("GetConnectedRepo: Connected to DB %v", *repoArgs.DbName)

	classroomRepo := New(db)

	return classroomRepo, nil
}

func repoArgsToDataSourceName(repoArgs *RepoArgs) string {

	dataSourceName := "postgres://" + *repoArgs.User + ":" +
		*repoArgs.Password + "@" + *repoArgs.Endpoint + "/" +
		*repoArgs.DbName + "?sslmode=" + *repoArgs.SslMode

	return dataSourceName
}
