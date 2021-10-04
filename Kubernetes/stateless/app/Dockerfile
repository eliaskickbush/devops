FROM golang:latest AS build
COPY * /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build

FROM alpine:latest as run
WORKDIR /app
COPY --from=build /app/stateless ./
CMD ./stateless
