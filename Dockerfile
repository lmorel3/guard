FROM iron/go:dev
MAINTAINER Laurent Morel <hello@lmorel3.fr>

WORKDIR /app
EXPOSE 80

RUN go get github.com/codegangsta/gin \
 && go get github.com/xyproto/permissionbolt \
 && go get github.com/gin-contrib/static  \
 && go get github.com/foolin/gin-template \
 && go get github.com/spf13/viper

ENV SRC_DIR=/go/src/github.com/lmorel3/guard-go/app

ADD app/ $SRC_DIR


RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/; cp -r views /app; cp -r assets /app
VOLUME /config

ENTRYPOINT ["./myapp"]
