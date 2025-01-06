FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download && go build -o main .

# Go関連ツールのインストール
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest