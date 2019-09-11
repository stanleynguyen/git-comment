FROM golang:1.13.0-alpine3.10 as builder
WORKDIR /go/src/github.com/stanleynguyen/git-comment/comment-app
RUN apk update && apk upgrade
COPY . .
RUN GOOS=linux go build -o commentapp.out .

FROM alpine:latest
RUN apk update && apk upgrade
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/stanleynguyen/git-comment/comment-app/commentapp.out .
ENV PORT 5000
EXPOSE ${PORT}
CMD ["./commentapp.out"]
