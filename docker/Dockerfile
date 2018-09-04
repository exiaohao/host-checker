FROM ubuntu:16.04

ADD checker /usr/local/bin
RUN chmod +x /usr/local/bin/checker

ENTRYPOINT ["/usr/local/bin/checker", "hosts-check"]