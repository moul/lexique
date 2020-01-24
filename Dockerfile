# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.13-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/lexique
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.11
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="lexique" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/lexique/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/lexique" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/lexique" \
                org.label-schema.help="docker exec -it $CONTAINER lexique --help"
COPY            --from=builder /go/bin/lexique /bin/
ENTRYPOINT      ["/bin/lexique"]
#CMD             []
