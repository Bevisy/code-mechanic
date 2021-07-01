from docker.io/library/debian:buster

run apt update && \
    apt install -y curl vim wget net-tools tcpdump dig && \
    apt clean all

entrypoint ["/bin/sh", "-c", "sleep 7d"]