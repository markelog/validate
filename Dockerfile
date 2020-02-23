FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app/validate .

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT ["/app/validate"]