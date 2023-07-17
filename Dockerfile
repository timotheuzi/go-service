FROM golang:alpine as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine:latest
COPY --from=builder /app/app .
ARG build_arg_mongodb_atlas_uri
ENV MONGODB_ATLAS_URI=$build_arg_mongodb_atlas_uri
ENV PORT 8080
CMD ["./app"]
