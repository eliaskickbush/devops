FROM golang:1.15

WORKDIR /app
COPY . .
RUN go get
RUN go build
CMD ./kubertest