FROM ubuntu:focal

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/GMT
ENV PREFECT_SERVER_COMMIT=c43868fbab3d35fe340adae562743c8e0044b30c
ENV PREFECT_VERSION=0.13.1

RUN apt-get update && apt-get install -y build-essential git python3-pip


RUN mkdir -p /opt
WORKDIR /opt
RUN git clone https://github.com/PrefectHQ/server.git
WORKDIR /opt/server
RUN git checkout $PREFECT_SERVER_COMMIT
RUN pip3 install -e /opt/server

RUN adduser --system prefect
USER prefect
WORKDIR /home/prefect
