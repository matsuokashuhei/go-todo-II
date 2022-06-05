FROM golang:1.18

ENV GO111MODULE on

WORKDIR /usr/src/todo/backend

RUN go install github.com/cweill/gotests/gotests@latest \
 && go install github.com/fatih/gomodifytags@latest \
 && go install github.com/josharian/impl@latest \
 && go install github.com/haya14busa/goplay/cmd/goplay@latest \
 && go install honnef.co/go/tools/cmd/staticcheck@latest \
 && go install golang.org/x/tools/gopls@latest \
 && go install github.com/ramya-rao-a/go-outline@latest \
 && go install entgo.io/ent/cmd/ent@latest

COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY . .
