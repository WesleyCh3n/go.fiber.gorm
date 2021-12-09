## PostgreSQL

```shell
docker exec -it gofiberrestful_db_1 psql -U user -h 127.0.0.1 -p 5432
```

```sql
\l -- list database
\c <DBname> -- use db
CREATE DATABASE DB_name;
DROP DATABASE DB_name;

\d -- list table
\d TABLE_NAME -- shown table detail(relation)

INSERT INTO TABLE_NAME (column1, column2, column3,...columnN) VALUES (value1, value2, value3,...valueN); -- insert query
SELECT (column1, column2, columnN) FROM TABLE_NAME -- select query
```
