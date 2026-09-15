package neo4j

import (
	"context"
	"fmt"

	"michaelyusak/biaenergi-segment-generator.git/apperror"
	"michaelyusak/biaenergi-segment-generator.git/entity"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type portRepository struct {
	driver neo4jDriver.Driver
	dbName string
}

func NewPortRepository(driver neo4jDriver.Driver, dbName string) *portRepository {
	return &portRepository{
		driver: driver,
		dbName: dbName,
	}
}

func (r *portRepository) GetPorts(ctx context.Context) ([]entity.Port, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (p: Port)
		RETURN p.id as id, p.value as value
		ORDER BY p.id ASC;
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
		port, err := parsePort(record)
		if err != nil {
			return nil, fmt.Errorf("[repository][neo4j][GetPorts] failed to parse port from record: %w", err)
		}

		res = append(res, port)
	}

	return res, nil
}

func (r *portRepository) GetPort(ctx context.Context, portID string) (*entity.Port, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (p: Port)
		WHERE p.id = $id
		RETURN p.id as id, p.value as value;
	`,
		map[string]any{
			"id": portID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetPort] failed to execute queries: %w", err)
	}

	if len(result.Records) < 1 {
		return nil, fmt.Errorf("[repository][neo4j][GetPort] record not found %w", apperror.ErrNotFound)
	}

	port, err := parsePort(result.Records[0])
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetPort] failed to parse port from record %w", err)
	}

	return &port, nil
}

func (r *portRepository) GetConnections(ctx context.Context) ([]entity.PortConnection, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (source:Port)-[:NEXT]->(target:Port)
		MATCH (sourceNode:Node)-[:HAS_PORT]->(source)
		MATCH (targetNode:Node)-[:HAS_PORT]->(target)
		WHERE sourceNode.id <> targetNode.id
		RETURN source.id AS source_id, source.value AS source_value, target.id AS target_id, target.value AS target_value
		ORDER BY targetNode.id, target ASC;
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetConnections] failed to execute queries: %w", err)
	}

	res := make([]entity.PortConnection, 0, len(result.Records))

	for _, record := range result.Records {
		connection, err := parseConnection(record)
		if err != nil {
			return nil, fmt.Errorf("[repository][neo4j][GetPorts] failed to connection port from record: %w", err)
		}

		res = append(res, connection)
	}

	return res, nil
}
