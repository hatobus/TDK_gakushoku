FROM gliderlabs/alpine:3.4

ENV GOROOT /go
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /go/lib/time/zoneinfo.zip
RUN apk --no-cache add ca-certificates wget
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk
RUN apk add glibc-2.28-r0.apk
WORKDIR /TDK_gakushoku
COPY ./bin /TDK_gakushoku/bin

CMD /TDK_gakushoku/bin/TDKserver
