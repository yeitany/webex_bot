FROM debian:stretch-slim

WORKDIR /
COPY ./.builds ./
COPY ./.env /.env

CMD ["/bin/sh"]
