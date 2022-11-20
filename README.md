# Solution for Golang microservices assignment

## Setup environment

1. Install Make `brew install make`
2. Install Go `brew install golang`
3. Create Go home directory `mkdir -p $HOME/go/{bin,src}`
4. Update shell profile `nano ~/.bash_profile`

   ```shell
   # ~/.bash_profile
   
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ````

5. Install Golang migration tool `brew install golang-migrate`

## How to?

1. Use this command to run Postgres database.

    ```shell
    make run-pg
   
    # to down database
    make down 
    ```

2. Use this command to run Postgres migrations.

    ```shell
    make run-pg-migrations
   
    # to rollback migrations
    make down-pg-migrations
    ```

3. Use this command to run service.

    ```shell
    make run-svc
    ```

4. Use this command to run unit tests.

    ```shell
    make test
    ```

## Explanation and CURLs

We are unable to upload whole file aat once because of restricted RAM size. As a result, the client should use multipart
content type to send large file data in chunks with small size.

```shell
curl --location --request POST 'localhost:2106/ports/batch' \
   --form 'part1=@"~/coding_task_1/samples/sample1_part1"' \
   --form 'part2=@"~/coding_task_1/samples/sample1_part2"'
```

## Database migrations

1. Install Golang migration tool `brew install golang-migrate`
2. Use this command to create new migration

   ```shell
   migrate create -ext sql -dir migrations/postgres ${YOUR_MIGRATION_NAME}
   ```

3. Use this command to apply migrations

   ```shell
   migrate -path migrations/postgres -database "postgres://${YOUR_DB_USERNAME}:${YOUR_DB_PASSWORD}@localhost/${YOUR_DB_NAME}?sslmode=disable" up
   ```

4. Use this command to undo migrations

   ```shell
   migrate -path migrations/postgres -database "postgres://${YOUR_DB_USERNAME}:${YOUR_DB_PASSWORD}@localhost/${YOUR_DB_NAME}?sslmode=disable" down
   ```

# Mocks for unit tests

1. Install mockgen

   ```shell
   go install github.com/golang/mock/mockgen@v1.6.0
   ```

2. Generate mocks

   ```shell
   go generate ./...
   ```
