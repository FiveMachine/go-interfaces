package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

var (
	_                 Config = &GenericConfig{}
	expectedKeyLength        = 6
)

func (cnf *GenericConfig) Parse(ctx *cli.Context) ([]byte, error) {
	// Parse from yaml
	if ctx.IsSet("config") {
		data, err := os.ReadFile(ctx.Path("config"))
		if err != nil {
			return nil, fmt.Errorf("reading config file path `%s` failed with: %s", ctx.Path("config"), err)
		}

		var (
			privateKey string
			swarmKey   string
			secure     string
			verbose    string
		)

		privateKey, data = PopYamlEntry(data, "privatekey")
		swarmKey, data = PopYamlEntry(data, "swarmkey")
		secure, data = PopYamlEntry(data, "http-secure")
		if len(secure) > 0 {
			_secure, err := strconv.ParseBool(secure)
			if err != nil {
				return nil, fmt.Errorf("parsing bool `%s` failed with: %s", secure, err)
			}
			cnf.HttpSecure = _secure
		}

		verbose, data = PopYamlEntry(data, "verbose")
		if verbose != "" {
			_verbose, err := strconv.ParseBool(verbose)
			if err != nil {
				return nil, fmt.Errorf("parsing `%s` failed with: %v", verbose, err)
			}
			cnf.Verbose = _verbose
		}

		if err = yaml.Unmarshal(data, &cnf); err != nil {
			return nil, fmt.Errorf("yaml unmarshal in GenericConfig failed with: %s", err)
		}

		if len(privateKey) > 0 {
			base64Key, err := base64.StdEncoding.DecodeString(privateKey)
			if err != nil {
				return nil, fmt.Errorf("converting private key to base 64 failed with: %s", err)
			}

			cnf.PrivateKey = []byte(base64Key)
		}

		if len(swarmKey) > 0 {
			swarmKeyBytes, err := os.ReadFile(swarmKey)
			if err != nil {
				return nil, fmt.Errorf("reading swarm key `%s` failed with: %s", cnf.Domains.Key.Private, err)
			}

			cnf.SwarmKey, err = parseSwarmKey(string(swarmKeyBytes))
			if err != nil {
				return nil, fmt.Errorf("parse swarm key `%s` failed with: %s", swarmKey, err)
			}
		}

		if err = cnf.DVKeyHandler(); err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, errors.New("config path was not set")
}

func (cnf *GenericConfig) Empty() (string, error) {
	bytes, err := yaml.Marshal(new(GenericConfig))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (_cnf *GenericConfig) String() string {
	cnf := *_cnf
	privateKey := string(cnf.PrivateKey)
	cnf.PrivateKey = nil
	swarmKey := string(cnf.SwarmKey)
	cnf.SwarmKey = nil

	bytes, err := yaml.Marshal(cnf)
	if err != nil {
		return err.Error()
	}

	display := string(bytes)
	display = strings.Replace(display, "privatekey: []", "privatekey: "+privateKey, 1)
	display = strings.Replace(display, "swarmkey: []", "swarmkey: "+swarmKey, 1)

	return display

}

func (cnf *GenericConfig) DVKeyHandler() error {
	var err error

	if err = cnf.validateKeys(); err != nil {
		return err
	}

	// Private Key
	if cnf.DVPrivateKey, err = os.ReadFile(cnf.Domains.Key.Private); err != nil {
		return fmt.Errorf("reading private key `%s` failed with: %s", cnf.Domains.Key.Private, err)
	}

	// Public Key
	if cnf.Domains.Key.Public != "" {
		if cnf.DVPublicKey, err = os.ReadFile(cnf.Domains.Key.Public); err != nil {
			return fmt.Errorf("reading public key `%s` failed with: %s", cnf.Domains.Key.Public, err)
		}
	} else {
		if cnf.DVPublicKey, err = generatePublicKey(cnf.DVPrivateKey); err != nil {
			return fmt.Errorf("generating public key failed with: %s", err)
		}
	}

	return nil
}

/*
	Cases

1. If private key is provided and no public key ->  generate public key
2. If both are provided read from both
3. If only public key and
  - has auth then return an error
  - no auth then its ok

4. No public key but runs monkey and node -> return an error
*/
func (cnf *GenericConfig) validateKeys() error {
	if slices.Contains(cnf.Protocols, "auth") && cnf.Domains.Key.Private == "" {
		return errors.New("domains private key cannot be empty when running auth")
	}

	for _, srv := range cnf.Protocols {
		if (srv == "monkey" || srv == "node") && (cnf.Domains.Key.Public == "" && cnf.Domains.Key.Private == "") {
			return errors.New("domains public key cannot be empty when running monkey or node")
		}
	}

	return nil
}
