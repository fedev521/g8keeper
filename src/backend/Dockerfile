# Since backend needs packages from tinksrv, which is not publicly available, we
# need to mount tinksrv code in the build container. This Dockerfile is supposed
# to be evaluated from src/, which contains both backend and tinksrv.

# build stage
FROM golang:1.22 AS builder

WORKDIR /go/src/app

RUN --mount=type=cache,mode=0755,target=/go/pkg/mod/ \
    --mount=type=bind,source=backend/go.mod,target=go.mod \
    --mount=type=bind,source=backend/go.sum,target=go.sum \
    --mount=type=bind,source=tinksrv,target=../tinksrv \
    go mod download

ENV CGO_ENABLED=0
RUN --mount=type=cache,mode=0755,target=/root/.cache/go-build \
    --mount=type=bind,source=backend,target=. \
    --mount=type=bind,source=tinksrv,target=../tinksrv \
    go build -o /go/bin/backend /go/src/app/cmd/backend
# RUN --mount=type=cache,mode=0755,target=/root/.cache/go-build \
#     go test -v
# RUN --mount=type=cache,mode=0755,target=/root/.cache/go-build \
#     go vet -v

# package stage
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=builder --chown=nonroot --chmod=500 /go/bin/backend .

CMD [ "/backend" ]
