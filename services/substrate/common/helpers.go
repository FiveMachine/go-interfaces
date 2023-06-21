package common

import (
	"fmt"
	"path"

	"github.com/taubyte/go-interfaces/services/substrate/counters"
)

func NewPathFromServiceable(serviceable Serviceable) (counters.Path, error) {
	projectCid, err := serviceable.Project()
	if err != nil {
		return nil, fmt.Errorf("getting project cid from serviceable failed with: %s", err)
	}

	return counters.NewPath(path.Join(projectCid.String(), serviceable.Id())), nil
}
