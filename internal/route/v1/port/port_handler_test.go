package port

import (
	"bytes"
	"github.com/0x9p/coding_task_1/internal/domain"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestUpsertPorts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := domain.NewMockPortService(ctrl)

	p := domain.Port{
		Id:          "AEAJM",
		Name:        "Ajman",
		City:        "Ajman",
		Country:     "United Arab Emirates",
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: []float64{55.5136433, 25.4052165},
		Province:    "Ajman",
		Timezone:    "Asia/Dubai",
		Unlocs:      []string{"AEAJM"},
		Code:        "52000",
	}
	svc.EXPECT().UpsertPort(&p).Return(nil, nil).Times(1)

	h := NewHandler(svc)

	f := `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }
}
`

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "chunk")
	part.Write([]byte(f))
	writer.Close()

	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", "/ports/batch", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())

	h.HandleUpsertPorts()(w, r)
}
