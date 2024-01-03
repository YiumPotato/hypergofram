FROM golang:1.21-alpine

WORKDIR /app


# Install air
RUN go install github.com/cosmtrek/air@latest

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Install sqlc as a binary as it's not available in alpine
RUN apk --no-cache add curl
RUN curl -L https://downloads.sqlc.dev/sqlc_1.24.0_linux_amd64.tar.gz | tar -xz -C /usr/local/bin sqlc


COPY go.mod go.sum ./
RUN go mod download


CMD ["air", "-c", ".air.toml"]