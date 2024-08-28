FROM golang:1.22.2

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make clean

RUN make build

EXPOSE 6500

CMD ["./ms-api"]
