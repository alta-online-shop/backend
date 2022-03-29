# Usage
## Migration
Generate config using:
```bash
$ soda g config
```

Run migrations:
```bash
$ soda m -c database/database.yml -p database/migration/
```

## Run
Use this command:
```bash
$ APP_DB_TYPE=pg APP_DB_DSN="<PostgreSQL DSN>" go run main.go
```
