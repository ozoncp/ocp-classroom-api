FROM golang:latest AS builder

COPY . /ocp-classroom-api
WORKDIR /ocp-classroom-api

RUN apt-get update -q && apt-get install -y protobuf-compiler
RUN make deps 
RUN make build

FROM alpine:latest

COPY --from=builder /ocp-classroom-api/bin/ocp-classroom-api /ocp-classroom-api
# for gRPC client
EXPOSE 7002
# for PostgreSQL
EXPOSE 5432
CMD ["/ocp-classroom-api"]