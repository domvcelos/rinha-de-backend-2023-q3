ALTER SYSTEM SET max_connections = 300;
-- Use an extension to enable trigram similarity search and improve LIKE performance
-- https://www.postgresql.org/docs/current/runtime-config-connection.htmlhttps://mazeez.dev/posts/pg-trgm-similarity-search-and-fast-like
CREATE EXTENSION pg_trgm;

CREATE TABLE IF NOT EXISTS PEOPLE 
 (
    ID uuid PRIMARY KEY UNIQUE,
    APELIDO VARCHAR(32) UNIQUE NOT NULL,
    NOME VARCHAR(100) NOT NULL,
    NASCIMENTO TIMESTAMP,
    STACK VARCHAR(1024),
    BUSCA_TGRM TEXT GENERATED ALWAYS AS (
        LOWER(NOME || ' ' || APELIDO || ' ' || STACK)
    ) STORED
);
-- CREATE INDEX CONCURRENTLY IF NOT EXISTS IDX_PEOPLE_BUSCA_TGRM ON PEOPLE USING GIST (BUSCA_TGRM GIST_TRGM_OPS(SIGLEN=64));
CREATE INDEX IDX_PEOPLE_BUSCA_TGRM ON PEOPLE USING gin (BUSCA_TGRM gin_trgm_ops);
-- INSERT INTO "people" VALUES ('5ce4668c-4710-4cfb-ae5f-38988d6d49cb','ana','Ana Barbosa','1985-09-23 00:00:00',ARRAY['Node']),
-- 	('bc6b02e2-82a1-485b-82c8-09daf8cdb6fd','ana','Ana Barbosa','1985-09-23 00:00:00',ARRAY['Postgres', 'Node']);

