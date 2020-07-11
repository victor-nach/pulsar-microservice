package handler

import (
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/repo"
)

type Handler struct {
	Repo repo.RepoInterface
}
