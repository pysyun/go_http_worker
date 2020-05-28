FROM pysyun/tor:alpine

ENV TARGET https://check.torproject.org

CMD ["/bin/sh", "-c", "torsocks curl $${TARGET}"]
