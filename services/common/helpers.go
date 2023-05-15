package common

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p/core/pnet"
)

// Removes line from the data, and returns the value of the key
func PopYamlEntry(data []byte, key string) (string, []byte) {
	if !strings.HasSuffix(key, ":") {
		key += ":"
	}

	var value string

	keyBytes := []byte(key)
	if bytes.Contains(data, keyBytes) {
		lines := bytes.Split(data, []byte("\n"))
		for i, line := range lines {
			if bytes.Contains(line, keyBytes) {
				value = string(line[len(keyBytes)+1:])
				lines = append(lines[:i], lines[i+1:]...)
				break
			}
		}
		data = bytes.Join(lines, []byte("\n"))
	}

	return value, data
}

func generatePublicKey(privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing private key failed with: %s", err)
	}

	publicKeyDer, err := x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, fmt.Errorf("marshalling PKIX pub key failed with: %s", err)
	}
	pubKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}

	pubKeyPem := pem.EncodeToMemory(&pubKeyBlock)

	return pubKeyPem, nil
}

func parseSwarmKey(key string) (pnet.PSK, error) {
	_key := strings.Split(key, "/")
	_key = deleteEmpty(_key)

	if len(_key) != expectedKeyLength {
		return nil, errors.New("swarm key is not correctly formatted")
	}

	format := fmt.Sprintf(`/%s/%s/%s/%s/
/%s/
%s`, _key[0], _key[1], _key[2], _key[3], _key[4], _key[5])

	return []byte(format), nil
}

func deleteEmpty(s []string) []string {
	if len(s) == 0 {
		return nil
	}

	r := make([]string, 0, len(s))
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
