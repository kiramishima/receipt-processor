FROM golang:1.21.1-alpine as build
ENV API_NAME=receipt_api
ENV MODE=Development
# HTTP
ENV HTTP_SERVER_IDLE_TIMEOUT=60s
ENV PORT=8080
ENV HTTP_SERVER_READ_TIMEOUT=1s
ENV HTTP_SERVER_WRITE_TIMEOUT=2s
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
ENV GOPROXY https://proxy.golang.org,direct
RUN go mod download
ENV CGO_ENABLED=0
RUN GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo -o $API_NAME ./cmd/api/main.go

FROM scratch as serve
WORKDIR /app
COPY --from=build /app/$API_NAME .
CMD ["/app/receipt_api"]