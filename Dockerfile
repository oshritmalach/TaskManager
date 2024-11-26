FROM golang:1.23 AS builder
COPY . .
RUN go build -o /server .

FROM ubuntu
COPY --from=builder /server /server
CMD ["/server"]

