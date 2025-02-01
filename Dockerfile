FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.21

WORKDIR /app

COPY --from=build /app/main .

RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT ["/app/main"]
