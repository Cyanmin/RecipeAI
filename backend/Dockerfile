FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git

# go.mod のみをコピーし、go.sum は存在する場合のみコピーする
COPY go.mod ./
COPY go.sum* ./

# 必要な依存関係をダウンロード
RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]