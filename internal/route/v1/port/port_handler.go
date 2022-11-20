package port

import (
	"encoding/json"
	"errors"
	"github.com/0x9p/coding_task_1/internal/domain"
	"github.com/0x9p/coding_task_1/internal/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Handler struct {
	portService domain.PortService
}

func (h *Handler) HandleUpsertPorts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mr, err := r.MultipartReader()

		if err != nil {
			util.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		tmpFile, err := ioutil.TempFile(os.TempDir(), "spa")
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if err != nil {
			util.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		for {
			part, err := mr.NextPart()

			if err == io.EOF {
				break
			}

			_, err = io.Copy(tmpFile, part)

			if err != nil {
				util.Error(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		_, err = tmpFile.Seek(0, io.SeekStart)

		if err != nil {
			util.Error(w, r, http.StatusInternalServerError, err)
			return
		}

		dec := json.NewDecoder(tmpFile)

		// read root object open curly brace
		_, err = dec.Token()

		var upsertedPortIds []string

		for dec.More() {
			portIdToken, err := dec.Token()

			if err != nil {
				util.PartialError(w, r, http.StatusInternalServerError, err, upsertedPortIds)
				return
			}

			var port domain.Port

			err = dec.Decode(&port)

			portId, ok := portIdToken.(string)

			if !ok {
				util.PartialError(w, r, http.StatusInternalServerError, errors.New("unable to resolve port id"), upsertedPortIds)
				return
			}

			port.Id = portId

			_, err = h.portService.UpsertPort(&port)

			if err != nil {
				util.PartialError(w, r, http.StatusInternalServerError, err, upsertedPortIds)
				return
			}

			upsertedPortIds = append(upsertedPortIds, port.Id)
		}

		util.Response(w, http.StatusNoContent)
	}
}

func NewHandler(portService domain.PortService) *Handler {
	return &Handler{
		portService: portService}
}
