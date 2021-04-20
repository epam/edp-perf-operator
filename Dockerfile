FROM alpine:3.11.10

ENV OPERATOR=/usr/local/bin/perf-operator \
    USER_UID=1001 \
    USER_NAME=perf-operator

# install operator binary
COPY perf-operator ${OPERATOR}

COPY build/bin /usr/local/bin
COPY build/configs /usr/local/configs

RUN  chmod u+x /usr/local/bin/user_setup && chmod ugo+x /usr/local/bin/entrypoint && /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
