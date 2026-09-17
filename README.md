# biaenergi-segment-generator

## Stacks
### Neo4j
* Graph structures 
* Fast traversals for linked data

schema:
* [schema.cypher](./schema/neo4j/schema.cypher)
* [seed.cypher](./schema/neo4j/seed.cypher)

### Go Programming Language

## Endpoints
* `/v1/canvas/ports` - list all ports and values
* `/v1/canvas/ports/:port_id` - get port detail
* `/v1/canvas/ports/connections` - list all port connections
* `/v1/canvas/nodes` - list all nodes with their ports
* `/v1/canvas/nodes/:node_id` - get node detail
* `/v1/segments` - get segments count, results, and details