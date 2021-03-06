package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gomods/athens/pkg/payloads"
	"github.com/gomods/athens/pkg/storage"
	"github.com/pkg/errors"
)

func uploadHandler(store storage.Saver) func(c buffalo.Context) error {
	return func(c buffalo.Context) error {
		stdParams, err := getStandardParams(c)
		if err != nil {
			return errors.WithStack(err)
		}
		version := c.Param("version")
		payload := new(payloads.Upload)
		if c.Bind(payload); err != nil {
			return errors.WithStack(err)
		}
		saveErr := store.Save(stdParams.baseURL, stdParams.module, version, payload.Module, payload.Zip)
		if storage.IsVersionAlreadyExistsErr(saveErr) {
			return c.Error(http.StatusConflict, saveErr)
		} else if err != nil {
			return errors.WithStack(err)
		}
		return c.Render(http.StatusOK, r.JSON(nil))
	}
}
