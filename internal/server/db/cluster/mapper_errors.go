package cluster

import (
	"net/http"

	"github.com/lxc/incus/v6/shared/api"
)

func init() {
	errMap[errTypeNotFound] = func(_ error, entity string) error {
		return api.StatusErrorf(http.StatusNotFound, "%s not found", entity)
	}

	errMap[errTypeConflict] = func(_ error, entity string) error {
		return api.StatusErrorf(http.StatusConflict, "This entry already exists for %s", entity)
	}
}
