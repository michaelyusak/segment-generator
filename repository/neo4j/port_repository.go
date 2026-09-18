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
		RETURN source.id AS source_id, source.value AS source_value, sourceNode.id as source_node_id, target.id AS target_id, target.value AS target_value, targetNode.id AS target_node_id
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

func (r *portRepository) GetSegmentHeads(ctx context.Context) ([]entity.Port, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (p:Port)
		MATCH (n:Node)-[ :HAS_PORT]->(p)
		WHERE NOT (p)-[:NEXT]->()
		return p.id as id, p.value as value, n.id as node_id;
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetHeads] failed to execute queries: %w", err)
	}

	res := make([]entity.Port, 0, len(result.Records))

	for _, record := range result.Records {
		port, err := parsePort(record)
		if err != nil {
			return nil, fmt.Errorf("[repository][neo4j][GetHeads] failed to parse segment heads from record: %w", err)
		}

		res = append(res, port)
	}

	return res, nil
}

func (r *portRepository) GetAllConnections(ctx context.Context) ([]entity.PortConnection, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (source:Port)-[:NEXT]->(target:Port)
		MATCH (sourceNode:Node)-[:HAS_PORT]->(source)
		MATCH (targetNode:Node)-[:HAS_PORT]->(target)
		RETURN source.id AS source_id, source.value AS source_value, sourceNode.id as source_node_id, target.id AS target_id, target.value AS target_value, targetNode.id AS target_node_id
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

func (r *portRepository) GetPaths(ctx context.Context) (map[string][][]entity.Port, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (head:Port)
		WHERE NOT (head)-[:NEXT]->()

		MATCH (tail:Port)
		WHERE NOT ()-[:NEXT]->(tail)

		MATCH path = (tail)-[:NEXT*]->(head)

		RETURN [p IN reverse(nodes(path)) | {id: p.id, value: p.value}] AS path
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetPaths] failed to execute queries: %w", err)
	}

	res := map[string][][]entity.Port{}

	for _, record := range result.Records {
		path, err := parsePath(record)
		if err != nil {
			return nil, fmt.Errorf("[repository][neo4j][GetPaths] failed to parse path from record: %w", err)
		}

		if len(path) < 1 {
			continue
		}

		res[path[0].ID] = append(res[path[0].ID], path)
	}

	return res, nil
}
