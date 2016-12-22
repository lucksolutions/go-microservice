#go-microservice
Example of a service that implements a healthz endpoint.

## Setup Vault
### Start Vault Server
The server starts in development mode, already unsealed.

`docker-compose up`

### Configure Vault Client
Set the `VAULT_ADDR` environment variable to the address of the Vault server.

`VAULT_ADDR=http://127.0.0.1:8200`

Authenticate to the server using the `Root Token` shown on server start.

`vault auth {Root Token}` 

## Create Docker Image
### Build the go binary
`GOOS=linux bash build`

### Build the Docker image
`docker build -t lucksolutions/go-microservice:1.0.0`

`docker push lucksolutions/go-microservice:1.0.0`