FROM golang:1.16-alpine

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y mysql-client build-base

RUN mkdir /app/src
WORKDIR /app/src

# COPY ../src/go.mod .
# COPY ../src/go.sum .
COPY ../src/ .
RUN go mod download
RUN /bin/chmod +x /usr/local/bin/grpc_entrypoint.sh

RUN go build cmd/main.go
RUN mv main /usr/local/bin

CMD [ "main" ]
ENTRYPOINT [ "grpc_entrypoint.sh" ]