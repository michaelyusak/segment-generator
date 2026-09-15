package neo4j

import (
	"context"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/apperror"
	"michaelyusak/biaenergi-segment-generator.git/entity"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
	neo4jDriver "github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

type nodeRepository struct {
	driver neo4jDriver.Driver
	dbName string
}

func NewNodeRepository(driver neo4jDriver.Driver, dbName string) *nodeRepository {
	return &nodeRepository{
		driver: driver,
		dbName: dbName,
	}
}

func (r *nodeRepository) GetNodes(ctx context.Context) ([]entity.Node, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (n:Node)
		OPTIONAL MATCH (n)-[:HAS_PORT]->(p:Port)
		RETURN n.id AS id, collect(
			CASE
				WHEN p IS NULL THEN NULL
				ELSE {
					id: p.id,
					value: p.value
				}
			END
		) AS ports
		ORDER BY n.id ASC;
	`,
		nil,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetNodes] failed to execute queries: %w", err)
	}

	res := make([]entity.Node, 0, len(result.Records))

	for _, record := range result.Records {
		node, err := parseNode(record)
		if err != nil {
			return nil, fmt.Errorf("[repository][neo4j][GetNodes] failed to parse node from record: %w", err)
		}

		res = append(res, node)
	}

	return res, nil
}

func (r *nodeRepository) GetNode(ctx context.Context, nodeID int64) (*entity.Node, error) {
	result, err := neo4j.ExecuteQuery(ctx, r.driver, `
		MATCH (n:Node)
		WHERE n.id = $id
		OPTIONAL MATCH (n)-[:HAS_PORT]->(p:Port)
		RETURN n.id AS id, collect(
			CASE
				WHEN p IS NULL THEN NULL
				ELSE {
					id: p.id,
					value: p.value
				}
			END
		) AS ports
		ORDER BY n.id ASC;
	`,
		map[string]any{
			"id": nodeID,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(r.dbName),
	)
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetNode] failed to execute queries: %w", err)
	}

	if len(result.Records) < 1 {
		return nil, fmt.Errorf("[repository][neo4j][GetNode] record not found %w", apperror.ErrNotFound)
	}

	node, err := parseNode(result.Records[0])
	if err != nil {
		return nil, fmt.Errorf("[repository][neo4j][GetNode] failed to parse node from record %w", err)
	}

	return &node, nil
}
