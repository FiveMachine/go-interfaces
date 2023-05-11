package seer

import "github.com/taubyte/go-interfaces/p2p/streams"

type Usage interface {
	Beacon(hostname, nodeId, clientNodeId string, signature []byte) UsageBeacon
	Heartbeat(usage *UsageData, hostname, nodeId, clientNodeId string, signature []byte) (streams.Response, error)
	Announce(services Services, nodeId, clientNodeId string, signature []byte) (streams.Response, error)
	AddService(svrType ServiceType, meta map[string]string)
	List() ([]string, error)
	ListServiceId(name string) (streams.Response, error)
	Get(id string) (*UsageReturn, error)
}

type UsageBeacon interface {
	Start()
}
