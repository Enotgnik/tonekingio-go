FROM golang

VOLUME /var/log/backend

COPY src /go/src

RUN go install go_files
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors

CMD /go/bin/go_files

EXPOSE 8000
