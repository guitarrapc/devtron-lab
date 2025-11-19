FROM gcr.io/distroless/static-debian12 AS base
EXPOSE 8080

FROM golang:1.25-alpine AS build
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app .

FROM base AS final
COPY --from=build /go/bin/app /
ENTRYPOINT ["/app"]
