package neo4j

import (
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"

	neo4jDriver "github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func parsePort(record *neo4jDriver.Record) (entity.Port, error) {
	var port entity.Port

	idAny, ok := record.Get("id")
	if !ok {
		return port, fmt.Errorf("missing port id")
	}

	id, ok := idAny.(string)
	if !ok {
		return port, fmt.Errorf("port id has type %T, want string", idAny)
	}

	port.ID = id

	valueAny, ok := record.Get("value")
	if !ok {
		return port, fmt.Errorf("missing port value")
	}

	if valueAny != nil {
		value, ok := valueAny.(int64)
		if !ok {
			return port, fmt.Errorf("port %q value has type %T, want int64", id, valueAny)
		}

		port.Value = &value
	}

	return port, nil
}
