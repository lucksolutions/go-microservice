#go-microservice
Example of a service that implements a healthz endpoint.

## Create Docker Image
### Build the go binary
`./build`

### Build the Docker image
`docker build -t lucksolutions/go-microservice:1.0.0`
`docker push lucksolutions/go-microservice:1.0.0`