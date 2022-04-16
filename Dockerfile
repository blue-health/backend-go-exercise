FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY ./ ./

RUN apk add --update make && make compile

# ---

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/bin/backend-go-exercise-srv-linux-amd64 /

EXPOSE 8080

ENTRYPOINT ["/backend-go-exercise-srv-linux-amd64"]
