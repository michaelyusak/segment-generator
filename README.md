# biaenergi-segment-generator

## Stacks

### Neo4j

* Graph structure and relationship storage
* Fast traversal of linked data

Schema and seed data:

* [schema.cypher](./schema/neo4j/schema.cypher)
* [seed.cypher](./schema/neo4j/seed.cypher)

### Go

* REST API
* Service and business logic

## API Documentation

For detailed API documentation, start the service and visit:

[Swagger UI](http://localhost:8080/swagger/index.html)

## Endpoints

### Canvas

* `GET /v1/canvas/ports` — List all ports and their values
* `GET /v1/canvas/ports/{port_id}` — Get port details
* `GET /v1/canvas/ports/connections` — List all port connections
* `GET /v1/canvas/nodes` — List all nodes with their ports
* `GET /v1/canvas/nodes/{node_id}` — Get node details

### Segments

* `GET /v1/segments` — Get the segment count and segment results
