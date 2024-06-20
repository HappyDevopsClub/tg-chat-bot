FROM golang:1.22.3 as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build /go/bin/app /
CMD ["/app"]

COPY
