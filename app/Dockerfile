# syntax=docker/dockerfile:1

FROM --platform=amd64 golang:1.18 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o goke


FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /app/goke /app/goke
COPY static/ /app/static/
COPY templates /app/templates/

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/app/goke"]
