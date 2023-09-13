package monkey

import "github.com/taubyte/go-interfaces/services"

type Client interface {
	services.Client
	Status(jid string) (*StatusResponse, error)
	Update(jid string, body map[string]interface{}) (string, error)
	List() ([]string, error)
}
