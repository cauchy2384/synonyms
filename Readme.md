# Readme

## Build & Run
* Local: 
    ```
    make build
    ./app/synonyms
    ```
* Dockerized: 
    ```
    make build-image
    docker-compose -f ./deployment/docker-compose.yml up
    ```

## Tests
* Lint:
    ```
    make lint
    ```
* Unit: 
    ```
    make test
    ```
* Manual integration: in ./test/synonyms.http