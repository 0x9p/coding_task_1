package port

import (
	"github.com/0x9p/coding_task_1/internal/nap"
	"github.com/lib/pq"
)

type Repo struct {
	db nap.SqlDb
}

func (r *Repo) UpsertPort(postgresPort *PostgresPort) error {
	if _, err := r.db.Exec(`
INSERT INTO port (
	id,
	name,
	city,
	country,
	alias,
	regions,
	coordinates,
	province,
	timezone,
	unlocs,
	code
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT(id) DO UPDATE SET
  name = EXCLUDED.name,
  city = EXCLUDED.city,
  country = EXCLUDED.country,
  alias = EXCLUDED.alias,
  regions = EXCLUDED.regions,
  coordinates = EXCLUDED.coordinates,
  province = EXCLUDED.province,
  timezone = EXCLUDED.timezone,
  unlocs = EXCLUDED.unlocs,
  code = EXCLUDED.code;
`,
		&postgresPort.Id,
		&postgresPort.Name,
		&postgresPort.City,
		&postgresPort.Country,
		pq.Array(&postgresPort.Alias),
		pq.Array(&postgresPort.Regions),
		pq.Array(&postgresPort.Coordinates),
		&postgresPort.Province,
		&postgresPort.Timezone,
		pq.Array(&postgresPort.Unlocs),
		&postgresPort.Code,
	); err != nil {
		return err
	}

	return nil
}

func NewRepo(db nap.SqlDb) *Repo {
	return &Repo{
		db: db}
}
