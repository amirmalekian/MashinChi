FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

# Static binary: no libc/musl dependency at runtime, so the builder and
# runtime images do not have to share the exact same Alpine release.
RUN CGO_ENABLED=0 go build -v -o server ./src/cmd/main.go

FROM alpine:latest AS runtime

# gorm.io/driver/postgres calls time.LoadLocation("Asia/Tehran") for the
# DSN's TimeZone parameter. Go resolves that name from the IANA zoneinfo
# database on disk, which Alpine does not ship by default. Setting TZ alone
# cannot satisfy time.LoadLocation, so tzdata is a real runtime dependency.
RUN apk add --no-cache ca-certificates tzdata

# The application builds its docker config name as
# filepath.Join(".", "/app/config/config-docker.yml"), i.e. relative to the
# process working directory. Starting from "/" makes it resolve to
# /app/config/config-docker.yml. Changing this to /app breaks config loading.
WORKDIR /

ENV APP_ENV=docker \
    TZ=Asia/Tehran

COPY --from=builder /app/server /app/server
COPY --from=builder /app/src/config/config-docker.yml /app/config/config-docker.yml
COPY --from=builder /app/src/docs /app/docs
# r.Static("/static", "./uploads") is evaluated from WORKDIR "/".
COPY --from=builder /app/src/cmd/uploads /uploads

CMD ["/app/server"]
