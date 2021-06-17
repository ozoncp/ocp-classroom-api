package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-classroom-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-classroom-api/internal/producer LogProducer
