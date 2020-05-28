FROM pysyun/tor:alpine

ENV INTERVAL 1
ENV TARGET https://check.torproject.org

COPY main.go .

# ${PWD} - path to our work dir
CMD ["/bin/sh", "-c", "go run ${PWD}/main.go --INTERVAL ${INTERVAL} --TARGET ${TARGET}"]

# docker run --name tor --privileged --tty --interactive name_your_container
