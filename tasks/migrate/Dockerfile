FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY ./ ./

WORKDIR /app/tasks/migrate

RUN apk add --update make && make compile

# ---

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/tasks/migrate/bin/backend-go-exercise-srv-migrate-linux-amd64 /

CMD ["/backend-go-exercise-srv-migrate-linux-amd64"] 
