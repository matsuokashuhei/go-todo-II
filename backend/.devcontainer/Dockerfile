FROM golang:1.18 as dev

ENV GO111MODULE on

WORKDIR /usr/src/todo/backend

RUN apt-get update \
 && apt install -y protobuf-compiler \
 && go install github.com/cweill/gotests/gotests@latest \
 && go install github.com/fatih/gomodifytags@latest \
 && go install github.com/josharian/impl@latest \
 && go install github.com/haya14busa/goplay/cmd/goplay@latest \
 && go install honnef.co/go/tools/cmd/staticcheck@latest \
 && go install golang.org/x/tools/gopls@latest \
 && go install github.com/ramya-rao-a/go-outline@latest \
 && go install entgo.io/ent/cmd/ent@latest \
 && go install -tags "mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2 \
 && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY . .
