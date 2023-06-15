package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/taubyte/go-interfaces/p2p/keypair"
	seerIface "github.com/taubyte/go-interfaces/services/seer"

	spec "github.com/taubyte/go-specs/common"
)

type ConfigBuilder struct {
	// generic
	DefaultP2PListenPort int
	DevP2PListenFormat   string

	// http
	DevHttpListenPort int
}

func (config *GenericConfig) Build(builder ConfigBuilder) error {
	if config == nil {
		config = &GenericConfig{}
		config.Bootstrap = true
		config.PrivateKey = nil
		config.SwarmKey = nil
		config.DevMode = false
	}

	if config.Root == "" {
		config.Root = DatabasePath
	}

	config.buildHttpAndTls(builder)
	config.buildP2P(builder)

	err := config.buildLocation(builder)
	if err != nil {
		return fmt.Errorf("building location failed with: %s", err)
	}

	err = config.buildKeys(builder)
	if err != nil {
		return fmt.Errorf("building keys failed with: %s", err)
	}

	if config.Branch == "" {
		config.Branch = spec.DefaultBranch
	}

	return nil
}

func (config *GenericConfig) buildHttpAndTls(builder ConfigBuilder) {
	if config.DevMode {
		if config.HttpListen == "" {
			config.HttpListen = fmt.Sprintf("0.0.0.0:%d", builder.DevHttpListenPort)
		}
	}

	if config.HttpListen == "" {
		config.HttpListen = DefaultHTTPListen
	}

	// Most likely not needed anymore since we are using odo
	// if config.TLS.Certificate == "" {
	// 	config.TLS.Certificate = DefaultCAFileName
	// }

	// if config.TLS.Key == "" {
	// 	config.TLS.Key = DefaultKeyFileName
	// }

}

func (config *GenericConfig) buildP2P(builder ConfigBuilder) {
	if len(config.P2PListen) == 0 {
		config.P2PListen = []string{fmt.Sprintf(DefaultP2PListenFormat, builder.DefaultP2PListenPort)}
	}

	if config.P2PAnnounce == nil {
		if config.DevMode {
			listenAddrFmt := builder.DevP2PListenFormat
			config.P2PAnnounce = []string{fmt.Sprintf(listenAddrFmt, builder.DefaultP2PListenPort)}

		} else {

			listenAddrFmt := os.Getenv("TAUBYTE_P2P_LISTEN")
			if len(listenAddrFmt) < 1 {
				panic("No Address to announce")
			}

			config.P2PAnnounce = []string{fmt.Sprintf(listenAddrFmt, builder.DefaultP2PListenPort)}
		}
	}

}

func (config *GenericConfig) buildKeys(builder ConfigBuilder) error {
	if len(config.PrivateKey) == 0 {
		if config.DevMode {
			config.PrivateKey = keypair.NewRaw()
		}
	}

	envKey := keypair.LoadRawFromEnv()
	if envKey != nil {
		config.PrivateKey = envKey
	}

	if len(config.SwarmKey) == 0 {
		return errors.New("swarm key is needed. Generate one using spore-drive if you don't have one")
	}

	return nil
}

func (config *GenericConfig) buildLocation(builder ConfigBuilder) error {
	if config.Location == nil {
		_locationJSON := os.Getenv("TAUBYTE_GEO_LOCATION")
		if len(_locationJSON) > 0 {
			config.Location = &seerIface.Location{}
			if err := json.Unmarshal([]byte(_locationJSON), config.Location); err != nil {
				return fmt.Errorf("parsing location failed with: %s", err)
			}
		}
	}

	return nil
}
