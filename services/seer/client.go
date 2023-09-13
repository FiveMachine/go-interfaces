package seer

import "github.com/taubyte/go-interfaces/services"

type Geo interface {
	All() ([]*Peer, error)
	Set(location Location) error
	Beacon(location Location) GeoBeacon
	Distance(from Location, distance float32) ([]*Peer, error)
}

type GeoBeacon interface {
	Start()
}

type Client interface {
	services.Client
	Geo() Geo
	Usage() Usage
}
