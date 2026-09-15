package adaptor

import (
	"context"
	"fmt"
	"time"

	"michaelyusak/biaenergi-segment-generator.git/config"

	"github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func ConnectNeo4j(conf config.Neo4jConfig) (neo4j.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.ConnectTimeout))
	defer cancel()

	driver, err := neo4j.NewDriver(
		conf.Uri,
		neo4j.BasicAuth(conf.Username, conf.Password, ""))
	if err != nil {
		return nil, fmt.Errorf("[adaptor][ConnectNeo4j] failed to create neo4j driver: %w", err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, fmt.Errorf("[adaptor][ConnectNeo4j] failed to verify neo4j connectivity: %w", err)
	}

	return driver, nil
}
