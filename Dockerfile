FROM scratch
MAINTAINER Jason Luck
ADD go-microservice /go-microservice
ENTRYPOINT ["/go-microservice"]
