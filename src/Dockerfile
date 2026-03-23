ARG GOVERSION=1.25.8

FROM docker.io/golang:${GOVERSION}-alpine AS build
ARG GOVERSION
WORKDIR /
COPY ./ /
RUN go mod edit -C main -go $GOVERSION
RUN go build -C main main.go

FROM docker.io/golang:${GOVERSION}-alpine
COPY --from=build /main/main /main
CMD [ "/main" ]

