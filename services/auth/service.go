package auth

import (
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.DBService
	services.GitHubAuth
}
