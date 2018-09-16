FROM golang:alpine AS build-env

RUN apk --update add git
ENV DIR=/go/src/github.com/ernst01/sudoku-solver
RUN go get -u github.com/kardianos/govendor

ADD . $DIR

RUN cd $DIR && govendor sync
RUN cd $DIR && go build cmd/sudoku/sudoku.go && cp sudoku /tmp/.

FROM alpine

WORKDIR /app

RUN apk --update add ca-certificates

ARG APP_ENV

ENV APP_ENV=$APP_ENV

COPY --from=build-env /tmp/sudoku /app/.

ENTRYPOINT  ["/app/sudoku"]