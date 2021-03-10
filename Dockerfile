FROM node:12 AS build-js
ADD . /app
WORKDIR /app/www/scripts

RUN npm install
RUN npm run build


# Start by building the application.
FROM golang:1.16-alpine as build-go

RUN apk add --no-cache ca-certificates make

WORKDIR /go/src/app
ADD . /go/src/app

COPY --from=build-js /app/www/scripts/dist/index.js /go/src/app/pkg/hyperdash/ui/scripts/app.js

RUN CGO_ENABLED=0 make docker-build

FROM alpine:3.13
RUN apk add --no-cache ca-certificates
COPY --from=build-go /go/src/app/hyperdash /bin
CMD [ "hyperdash", "run", "demo/dashboard.hcl" ]
