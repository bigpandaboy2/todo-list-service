FROM golang:1.16-alpine

WORKDIR /app

# Install build tools
RUN apk add --no-cache gcc musl-dev

COPY . .

WORKDIR /app/cmd

RUN go mod tidy && go build -o /app/main .

CMD ["/app/main"]