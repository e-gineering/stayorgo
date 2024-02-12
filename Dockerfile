FROM golang:1.22 as builder
LABEL authors="tylermendenhall"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o stayorgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/stayorgo .
COPY --from=builder /app/static ./static
CMD ["./stayorgo"]