# ==== level 0 ====
FROM golang:alpine

WORKDIR /build

COPY . .

RUN go build -o nameDemo .

WORKDIR /dist

RUN cp /build/nameDemo .

EXPOSE 8009

CMD ["/dist/nameDemo"]

# ==== level 1 ====
FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o nameDemo .

FROM scratch

COPY --from=builder /build/nameDemo /nameDemo

ENTRYPOINT ["/nameDemo"]

# ==== level 2 ====
FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /crm.shopdev.com

ENTRYPOINT ["/crm.shopdev.com", "-config", "/config/config.yaml"]

