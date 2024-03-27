FROM golang:1.22.1 as build
LABEL authors="Danila Bogdanov https://t.me/DBugdanov"

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app usermanager/cmd/main.go
EXPOSE 8011

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
