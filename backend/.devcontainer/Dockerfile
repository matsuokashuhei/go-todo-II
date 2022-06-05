FROM golang:1.18 as base
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /usr/src/todo/backend

FROM base as dev
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

FROM base as builder
COPY backend .
RUN go build -o todo .

FROM scratch as prod
COPY --from=builder /usr/src/todo/backend/todo /usr/local/bin/todo
EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/todo"]
