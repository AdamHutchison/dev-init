FROM mysql:5.5

ADD database/create-databases.sql /docker-entrypoint-initdb.d
ADD database/database-schema.sql /docker-entrypoint-initdb.d
ADD database/database-db_logs-schema.sql /docker-entrypoint-initdb.d
ADD database/database-db_reports-schema.sql /docker-entrypoint-initdb.d
ADD database/database-db_reporting-schema.sql /docker-entrypoint-initdb.d
