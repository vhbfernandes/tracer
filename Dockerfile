FROM golang:onbuild AS build

WORKDIR /build
COPY . /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM scratch
COPY --from=build /build/app /app
CMD ["./app"]