FROM node:12 AS build-js
ADD . /app
WORKDIR /app/www/scripts

RUN npm install
RUN npm run build


# Start by building the application.
FROM golang:1.16-buster as build-go

WORKDIR /go/src/app
ADD . /go/src/app

COPY --from=build-js /app/www/scripts/dist/index.js /go/src/app/pkg/hyperdash/ui/scripts/app.js

RUN make docker-build

FROM gcr.io/distroless/base-debian10
COPY --from=build-go /go/src/app/hyperdash /
ENTRYPOINT [ "/hyperdash" ]
CMD [ "run", "demo/dashboard.hcl" ]
