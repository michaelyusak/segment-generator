CREATE CONSTRAINT node_id_unique
FOR (n:Node)
REQUIRE n.id IS UNIQUE;

CREATE CONSTRAINT port_id_unique
FOR (p:Port)
REQUIRE p.id IS UNIQUE;
