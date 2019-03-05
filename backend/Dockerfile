FROM golang

VOLUME /var/log/backend

COPY src /go/src

RUN go install go_files

CMD /go/bin/go_files

EXPOSE 8000
