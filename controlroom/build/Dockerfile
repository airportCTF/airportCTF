FROM golang:1.22 as build
LABEL authors="Danila Bogdanov https://t.me/DBugdanov"

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app controlroom/cmd/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/app /
COPY art.txt /
EXPOSE 8012
CMD ["/app"]
