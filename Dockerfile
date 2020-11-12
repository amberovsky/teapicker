FROM golang:1.15.4-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o /out/service ./cmd/main.go

FROM scratch AS bin
COPY --from=build /out/service /
ENTRYPOINT ["/service"]
