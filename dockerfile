FROM golang:alpine3.15 AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main ./src/main.go

FROM alpine:3.14.4 AS production
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 3000
CMD ["./main"]