package patrick

import "github.com/taubyte/go-interfaces/services"

type Client interface {
	services.Client
	Lock(jid string, eta uint32) error
	IsLocked(jid string) (bool, error)
	Unlock(jid string) error
	Done(jid string, cid_log map[string]string, assetCid map[string]string) error
	Failed(jid string, cid_log map[string]string, assetCid map[string]string) error
	List() ([]string, error)
	Get(jid string) (*Job, error)
	Timeout(jid string) error
	Cancel(jid string, cid_log map[string]string) (interface{}, error)
	Close()
}
