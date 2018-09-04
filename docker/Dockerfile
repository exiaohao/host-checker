FROM ubuntu:16.04

ADD host-checker /usr/local/bin
RUN chmod +x /usr/local/bin/host-checker

ENTRYPOINT ["/usr/local/bin/host-checker"]