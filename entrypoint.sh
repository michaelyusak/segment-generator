#!/bin/bash

/startup/docker-entrypoint.sh neo4j &
neo4j_pid=$!

echo "Waiting for Neo4j..."

until cypher-shell \
  -a neo4j://localhost:7687 \
  -u "$SEED_NEO4J_USER" \
  -p "$SEED_NEO4J_PASSWORD" \
  "RETURN 1" >/dev/null 2>&1
do
  sleep 2
done

echo "Neo4j is ready. Seeding..."

cypher-shell \
  -a neo4j://localhost:7687 \
  -u "$SEED_NEO4J_USER" \
  -p "$SEED_NEO4J_PASSWORD" \
  -f /schema/schema.cypher

cypher-shell \
  -a neo4j://localhost:7687 \
  -u "$SEED_NEO4J_USER" \
  -p "$SEED_NEO4J_PASSWORD" \
  -f /schema/seed.cypher

echo "Seed complete."

wait "$neo4j_pid"