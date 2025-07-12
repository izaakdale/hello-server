FROM golang:1-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download

RUN go build -o /hello-server

FROM alpine
COPY --from=builder /hello-server .

EXPOSE 80
CMD [ "/hello-server" ]