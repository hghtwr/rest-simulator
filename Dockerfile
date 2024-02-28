FROM golang:1.21.5-alpine as builder
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o ./bin/rest-simulator

FROM gcr.io/distroless/static-debian12 as distroless
COPY --from=builder /src/bin/rest-simulator /rest-simulator
ENTRYPOINT ["/rest-simulator"]

FROM scratch as scratch
COPY --from=builder /src/bin/rest-simulator /rest-simulator
ENTRYPOINT ["/rest-simulator"]