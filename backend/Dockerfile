FROM golang

VOLUME /var/log/backend

COPY src /go/src

RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go install go_files


CMD /go/bin/go_files

EXPOSE 8000
