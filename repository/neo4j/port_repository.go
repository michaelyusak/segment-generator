package neo4j

import (
	"context"
	"fmt"

	"michaelyusak/biaenergi-segment-generator.git/entity"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type portRepository struct {
	driver neo4j.Driver
	dbName string
}

func NewPortRepository(driver neo4j.Driver, dbName string) *portRepository {
	return &portRepository{
		driver: driver,
		dbName: dbName,
	}
}

func (r *portRepository) GetPorts(ctx context.Context) ([]entity.Port, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (p: Port)
		RETURN p.id as id, p.value as value;
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetPorts] failed to execute queries: %w", err)
	}

	res := make([]entity.Port, 0, len(result.Records))

	for _, record := range result.Records {
		idAny, ok := record.Get("id")
		if !ok {
			return nil, fmt.Errorf("missing port id")
		}

		id, ok := idAny.(string)
		if !ok {
			return nil, fmt.Errorf("port id has type %T, want string", idAny)
		}

		port := entity.Port{
			ID: id,
		}

		valueAny, ok := record.Get("value")
		if !ok {
			return nil, fmt.Errorf("missing port value")
		}

		if valueAny != nil {
			value, ok := valueAny.(int64)
			if !ok {
				return nil, fmt.Errorf("port %q value has type %T, want int64", id, valueAny)
			}

			port.Value = &value
		}

		res = append(res, port)
	}

	return res, nil
}
