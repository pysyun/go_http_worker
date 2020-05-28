FROM pysyun/tor:alpine

ENV TARGET https://check.torproject.org

CMD ["/bin/sh", "-c", "torsocks curl ${TARGET}"]

# docker run --name tor --privileged --tty --interactive name_your_container
