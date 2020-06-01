FROM pysyun/golang-tor-runtime

WORKDIR /home/golang-tor-runtime
COPY main.go .

ENV INTERVAL 1
ENV TARGET https://check.torproject.org

CMD ["/bin/sh", "-c", "go run ${PWD}/main.go --INTERVAL ${INTERVAL} --TARGET ${TARGET}"]
