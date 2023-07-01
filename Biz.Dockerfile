FROM golang:alpine AS build

RUN apk --update add ca-certificates git

WORKDIR /srv/build

COPY go.mod .
COPY go.sum .
COPY BizServer/server/server.go .
COPY BizServer ./BizServer
RUN mkdir -p ./pbGenerated
COPY pbGenerated/BizServer.pb.go ./pbGenerated
COPY pbGenerated/BizServer_grpc.pb.go ./pbGenerated

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o server .

FROM golang:alpine AS runtime

WORKDIR /srv/app

COPY --from=build /srv/build/server /bin/

CMD ["/bin/server"]
