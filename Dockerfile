FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.16-alpine3.13 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN apk add --no-cache git

ARG GOPROXY
# download deps before gobuild
RUN go mod download -x
ARG TARGETOS
ARG TARGETARCH
RUN source scripts/update_version.sh && \
  GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o static-file-server -ldflags "${LD_FLAGS}" cmd/main.go


FROM alpine:3.13

LABEL maintainer="Jeeva Kandasamy <jkandasa@gmail.com>"

ENV APP_HOME="/app" \
    DATA_HOME="/data" \
    BRAND_NAME="Lightweight Static File Server"

EXPOSE 8080

# install timezone utils
RUN apk --no-cache add tzdata

# create a user and give permission for the locations
RUN mkdir -p ${APP_HOME} && mkdir -p ${DATA_HOME}

# copy application bin file
COPY --from=builder /app/static-file-server ${APP_HOME}/static-file-server

RUN chmod +x ${APP_HOME}/static-file-server

WORKDIR ${APP_HOME}

CMD ["/app/static-file-server"]
