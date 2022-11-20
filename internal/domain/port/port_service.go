package port

import (
	"github.com/0x9p/coding_task_1/internal/domain"
)

type portService struct {
	repo *Repo
}

func (s portService) UpsertPort(port *domain.Port) (*domain.Port, error) {
	postgresPort := &PostgresPort{
		Id:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}

	err := s.repo.UpsertPort(postgresPort)

	if err != nil {
		return nil, err
	}

	return mapToPort(postgresPort), nil
}

func mapToPort(postgresPort *PostgresPort) *domain.Port {
	port := &domain.Port{
		Id:          postgresPort.Id,
		Name:        postgresPort.Name,
		City:        postgresPort.City,
		Country:     postgresPort.Country,
		Alias:       postgresPort.Alias,
		Regions:     postgresPort.Regions,
		Coordinates: postgresPort.Coordinates,
		Province:    postgresPort.Province,
		Timezone:    postgresPort.Timezone,
		Unlocs:      postgresPort.Unlocs,
		Code:        postgresPort.Code,
	}
	return port
}

func NewService(repo *Repo) domain.PortService {
	return &portService{
		repo: repo}
}
