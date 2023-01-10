FROM golang:alpine AS build

WORKDIR /srv/app

COPY . .

RUN go get

RUN go build -o cqrs

FROM alpine

COPY --from=build /srv/app/cqrs /usr/local/bin

CMD ["/usr/local/bin/cqrs"]

