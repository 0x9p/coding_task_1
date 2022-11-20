package port

import (
	"errors"
	"github.com/0x9p/coding_task_1/internal/domain"
	"github.com/0x9p/coding_task_1/internal/nap"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpsertPort(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := nap.NewMockSqlDb(ctrl)

	db.EXPECT().Exec(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Times(1)

	repo := NewRepo(db)
	svc := NewService(repo)

	port := &domain.Port{
		Id:          "test_port_id",
		Name:        "test_port_name",
		City:        "test_port_city",
		Country:     "test_port_country",
		Alias:       []string{"test_port_alias"},
		Regions:     []string{"test_port_region"},
		Coordinates: []float64{100.1, 100.2},
		Province:    "test_port_province",
		Timezone:    "test_port_timezone",
		Unlocs:      []string{"test_port_unloc"},
		Code:        "test_port_code",
	}

	upsertedPort, err := svc.UpsertPort(port)

	assert.Equal(t, port.Id, upsertedPort.Id)
	assert.Equal(t, port.Name, upsertedPort.Name)
	assert.Equal(t, port.City, upsertedPort.City)
	assert.Equal(t, port.Country, upsertedPort.Country)
	assert.Equal(t, port.Alias, upsertedPort.Alias)
	assert.Equal(t, port.Regions, upsertedPort.Regions)
	assert.Equal(t, port.Coordinates, upsertedPort.Coordinates)
	assert.Equal(t, port.Province, upsertedPort.Province)
	assert.Equal(t, port.Timezone, upsertedPort.Timezone)
	assert.Equal(t, port.Unlocs, upsertedPort.Unlocs)
	assert.Equal(t, port.Code, upsertedPort.Code)

	assert.Nil(t, err)
}

func TestUpsertPortWhenUpsertError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := nap.NewMockSqlDb(ctrl)

	db.EXPECT().Exec(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(nil, errors.New("something went wrong")).Times(1)

	repo := NewRepo(db)
	svc := NewService(repo)

	port := &domain.Port{
		Id:          "test_port_id",
		Name:        "test_port_name",
		City:        "test_port_city",
		Country:     "test_port_country",
		Alias:       []string{"test_port_alias"},
		Regions:     []string{"test_port_region"},
		Coordinates: []float64{100.1, 100.2},
		Province:    "test_port_province",
		Timezone:    "test_port_timezone",
		Unlocs:      []string{"test_port_unloc"},
		Code:        "test_port_code",
	}

	upsertedPort, err := svc.UpsertPort(port)

	assert.Nil(t, upsertedPort)

	assert.Equal(t, err.Error(), "something went wrong")
}
