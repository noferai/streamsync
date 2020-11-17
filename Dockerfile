FROM golang:1.15 AS builder
ADD . /src
WORKDIR /src/backend
RUN go mod download
RUN CGO_ENABLED=0 go build -o app

FROM node:10.19-alpine AS node_builder
COPY --from=builder /src/frontend .
RUN npm install --silent
RUN npm run build

FROM alpine:3.12.1
WORKDIR /chat
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/backend/app .
COPY --from=node_builder /build ./static
EXPOSE 8080
CMD ./app