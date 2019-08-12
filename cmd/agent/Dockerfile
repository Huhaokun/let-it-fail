# TODO refine its size

FROM golang:1.12.6-alpine

LABEL maintainer="cloud9 <geekun94@gmail.com>"

WORKDIR /app


COPY lif-agent .

EXPOSE 7999

CMD ["./lif-agent"]

