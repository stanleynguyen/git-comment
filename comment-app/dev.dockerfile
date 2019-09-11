FROM golang:1.13.0-alpine3.10
WORKDIR /go/src/github.com/stanleynguyen/git-comment/comment-app
RUN apk update && apk upgrade && apk add --no-cache git && go get github.com/pilu/fresh
CMD ["fresh"]
ENV PORT 5000
EXPOSE ${PORT}
