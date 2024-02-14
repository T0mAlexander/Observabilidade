FROM alpine:3.19.1

COPY ./setup.sh /scripts/

ENV GF_PATHS_CONFIG=/etc/grafana/grafana.ini

RUN chmod +x /scripts/setup.sh
RUN echo -e "https://dl-cdn.alpinelinux.org/alpine/edge/main\nhttps://dl-cdn.alpinelinux.org/alpine/edge/community" > /etc/apk/repositories
RUN apk update && apk add curl openrc grafana
RUN mkdir /run/openrc/
RUN touch /run/openrc/softlevel
RUN rc-update add grafana default

EXPOSE 3000

ENTRYPOINT [ "/scripts/setup.sh" ]