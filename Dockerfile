ARG HARBOR_PROXY=$HARBOR_PROXY
ARG GOLANG_VER=1.19

FROM ${HARBOR_PROXY}golang:${GOLANG_VER}-alpine AS builder
ARG SVC_NAME

RUN apk --no-cache add ca-certificates make

WORKDIR /build
COPY . .
RUN make build \
    && mv build/app /exe

FROM scratch
COPY --from=builder /exe /
COPY --from=builder /build/config config/
COPY --from=builder /build/docs /docs

ENTRYPOINT ["exe"]
