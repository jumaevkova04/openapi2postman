FROM golang:1.19-alpine
RUN apk add git

COPY . /home/src
WORKDIR /home/src
RUN go build -o /bin/action ./main.go

ENTRYPOINT [ "/bin/action" ]