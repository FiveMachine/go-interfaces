package monkey

import "github.com/taubyte/go-interfaces/services/patrick"

type StatusResponse struct {
	Jid    string
	Status patrick.JobStatus
	Logs   string
}
