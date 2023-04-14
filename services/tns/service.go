package tns

import (
	"github.com/taubyte/go-interfaces/kvdb"
	"github.com/taubyte/go-interfaces/services"
)

type Service interface {
	services.Service
	KV() kvdb.KVDB
}
