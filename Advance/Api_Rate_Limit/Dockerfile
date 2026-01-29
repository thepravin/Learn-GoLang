FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_app_pravin

# Build state ends here

FROM alpine:latest

WORKDIR /app

COPY  --from=builder /app/go_app_pravin .

EXPOSE 8090

CMD [ "./go_app_pravin" ]