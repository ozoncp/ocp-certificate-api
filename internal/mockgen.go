package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-certificate-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-certificate-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/metrics_mock.go -package=mocks github.com/ozoncp/ocp-certificate-api/internal/metrics Metrics
