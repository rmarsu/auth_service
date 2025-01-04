FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x wait-for-postgres.sh
   
RUN go build -o auth-service cmd/main.go
RUN go build -o migrator cmd/migrate/main.go

EXPOSE 4200

CMD ["sh", "-c", "./migrator up && ./auth-service"]
