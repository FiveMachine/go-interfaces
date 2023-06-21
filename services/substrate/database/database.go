package database

import structureSpec "github.com/taubyte/go-specs/structure"

type Database interface {
	KV() KV
	DBContext() Context
	SetConfig(*structureSpec.Database)
	Close()
	Config() *structureSpec.Database
}
