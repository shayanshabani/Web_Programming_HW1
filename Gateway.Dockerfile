FROM golang:alpine AS build

RUN apk --update add ca-certificates git

WORKDIR /srv/build

COPY go.mod .
COPY go.sum .
COPY Gateway/server/server.go .
COPY Gateway ./Gateway
RUN mkdir -p ./pbGenerated
ADD pbGenerated ./pbGenerated

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .

FROM golang:alpine AS runtime

WORKDIR /srv/app

COPY --from=build /srv/build/server /bin/

CMD ["/bin/server"]
