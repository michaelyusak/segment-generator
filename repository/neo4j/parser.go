package neo4j

import (
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"

	neo4jDriver "github.com/neo4j/neo4j-go-driver/v6/neo4j"
)

func parseID[T any](record *neo4jDriver.Record) (T, error) {
	var id T

	idAny, ok := record.Get("id")
	if !ok {
		return id, fmt.Errorf("missing port id")
	}

	id, ok = idAny.(T)
	if !ok {
		return id, fmt.Errorf("port id has type %T, want string", idAny)
	}

	return id, nil
}

func parsePort(record *neo4jDriver.Record) (entity.Port, error) {
	var port entity.Port

	id, err := parseID[string](record)
	if err != nil {
		return port, err
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

func parseNode(record *neo4jDriver.Record) (entity.Node, error) {
	var node entity.Node

	id, err := parseID[int64](record)
	if err != nil {
		return node, err
	}

	node.ID = id

	portsAny, ok := record.Get("ports")
	if !ok {
		return node, fmt.Errorf("missing port list")
	}

	portsMap, ok := portsAny.([]any)
	if !ok {
		return node, fmt.Errorf("invalid ports: %T", portsAny)
	}

	node.Ports = make([]entity.Port, 0, len(portsMap))

	for _, portAny := range portsMap {
		portMap, ok := portAny.(map[string]any)
		if !ok {
			return node, fmt.Errorf("invalid port")
		}

		portIDAny, ok := portMap["id"]
		if !ok {
			return node, fmt.Errorf("missing port id")
		}

		portID, ok := portIDAny.(string)
		if !ok {
			return node, fmt.Errorf(
				"port id has type %T, want string",
				portIDAny,
			)
		}

		port := entity.Port{
			ID: portID,
		}

		valueAny, ok := portMap["value"]
		if !ok {
			return node, fmt.Errorf(
				"missing port value for port %q",
				portID,
			)
		}

		if valueAny != nil {
			value, ok := valueAny.(int64)
			if !ok {
				return node, fmt.Errorf(
					"port %q value has type %T, want int64",
					portID,
					valueAny,
				)
			}

			port.Value = &value
		}

		node.Ports = append(node.Ports, port)
	}

	return node, nil
}
