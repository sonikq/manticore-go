# Skdf-manticore-go instructions
***

## Set your environment variables

#### Add your application configuration to your `.env` file in the root of your project. Here is an example:

```

# The address where your web application will be run.
HTTP_HOST=127.0.0.1
HTTP_PORT=8080

# Postgres Configurations
POSTGRES_HOST=127.0.0.1
POSTGRES_PORT=5432
POSTGRES_USERNAME=user
POSTGRES_PASSWORD=psswd
POSTGRES_DB=dbname
POSTGRES_SSLMODE=disable

# Manticoresearch Configurations
MANTICORE_HOST=127.0.0.1
MANTICORE_PORT=9308
MANTICORE_INDEXER_PORT=5000
MANTICORE_LIMIT=10000
MAX_MATCHES=10000

# context timeout in seconds
CTX_TIMEOUT=600

# log level
LOG_LEVEL=debug

```

***

## Post-installation steps for Manticore

### 1. Create the `searchd` and `indexer` groups.
```
sudo groupadd searchd && sudo groupadd indexer
```

### 2. Add your user to the `searchd` and `indexer` groups.
```
sudo usermod -aG searchd $USER && sudo usermod -aG indexer $USER
```

### 3. Give your user permissions.
```
sudo mkdir -p /var/run/manticore && \
sudo chown -hR $USER: /etc/manticoresearch /var/log/manticore /var/run/manticore /var/lib/manticore
```

#### Now you can run `searchd` and `indexer` without sudo.

***

## APIs to rotate index and merge indexes

### 1. Rotate Index

##### HTTP Request
```http request
POST $MANTICORE_HOST:$MANTICORE_INDEXER_PORT/indexer/rotateindex
``` 

##### Example
```json
{
  "index": "idx"
}
```

### 2. Merge Indexes

##### HTTP Request
```http request
POST $MANTICORE_HOST:$MANTICORE_INDEXER_PORT/indexer/mergeindexes
```

##### Example
```json
{
  "main_index": "idx_main",
  "delta_index": "idx_delta"
}
```
