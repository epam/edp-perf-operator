FROM alpine:3.11.6

ENV OPERATOR=/usr/local/bin/perf-operator \
    USER_UID=1001 \
    USER_NAME=perf-operator

# install operator binary
COPY entrypoint ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  chmod u+x /usr/local/bin/user_setup && chmod ugo+x /usr/local/bin/entrypoint && /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
