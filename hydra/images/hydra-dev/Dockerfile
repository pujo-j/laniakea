FROM ubuntu:focal-20200925

ENV DEBIAN_FRONTEND=noninteractive

ENV HYDRA_VERSION=0.2.1

# Dist upgrade
RUN apt-get update -qq && apt-get dist-upgrade -y -qq


# Standardize locale
RUN apt-get install -y locales -qq
ENV TZ=Etc/GMT
RUN locale-gen en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

# Base packages
RUN apt install -y -qq \
    tmux \
    unzip \
    ca-certificates \
    libaio1 \
    wget \
    unzip \
    git \
    gnupg \
    software-properties-common \
    curl \
    bzip2 \
    python3-dev \
    build-essential \
    python3-pip \
    python3-venv \
    openssl \
    aria2 \
    libssl-dev \
    htop \
    iftop \
    vim \
    net-tools \
    openjdk-8-jre-headless \
    maven \
    python-is-python3

# Get latest build chain
RUN add-apt-repository -y ppa:ubuntu-toolchain-r/test && apt update && apt upgrade

# Database drivers
RUN apt-get install -y -qq \
    freetds-dev \
    libmariadbclient-dev \
    libpq-dev

# Install Oracle instant client
WORKDIR    /opt/oracle
RUN  wget https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linuxx64.zip \
            && unzip instantclient-basiclite-linuxx64.zip \
            && rm -f instantclient-basiclite-linuxx64.zip \
            && cd /opt/oracle/instantclient* \
            && rm -f *jdbc* *occi* *mysql* *README *jar uidrvci genezi adrci \
            && echo /opt/oracle/instantclient* > /etc/ld.so.conf.d/oracle-instantclient.conf \
            && ldconfig

# Install node
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt install -y nodejs

# Install yarn
RUN curl -sL https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" |  tee /etc/apt/sources.list.d/yarn.list
RUN apt update && apt install yarn

# Install terraform
RUN wget https://releases.hashicorp.com/terraform/0.13.2/terraform_0.13.2_linux_amd64.zip && \
 unzip terraform_0.13.2_linux_amd64.zip && mv ./terraform /usr/local/bin/terraform && \
 chmod a+x /usr/local/bin/terraform

# Install native kubectl
RUN wget -O /usr/local/bin/kubectl "https://storage.googleapis.com/kubernetes-release/release/v1.19.0/bin/linux/amd64/kubectl"
RUN chmod a+x /usr/local/bin/kubectl

# Install chrome (headless)

RUN curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& echo "deb https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list

RUN apt-get update && apt-get install -y \
	google-chrome-stable \
	&& rm -rf /var/lib/apt/lists/*

# Add go
RUN mkdir -p /opt/
RUN wget -O /opt/go1.15.1.linux-amd64.tar.gz "https://golang.org/dl/go1.15.1.linux-amd64.tar.gz"
RUN cd /opt/ && tar zxvf go1.15.1.linux-amd64.tar.gz && rm /opt/go1.15.1.linux-amd64.tar.gz

# Python stuff
RUN pip3 install --upgrade pip

ENV THEIA_DEFAULT_PLUGINS=local-dir:/home/theia/plugins
ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/jre/
ENV JAVA_HOME=/usr/share/maven
ENV PATH=$PATH:/opt/go/bin
ENV PYTHON=/usr/bin/python

# Add Tini
ENV TINI_VERSION v0.19.0

RUN wget -O /tini https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini
RUN chmod +x /tini

# Add poetry
RUN pip3 install "poetry==1.0.10"

# Switch to user

RUN adduser --system theia --shell=/bin/bash --ingroup adm
USER theia
WORKDIR /home/theia
COPY package.json .
RUN mkdir -p /home/theia/root
RUN yarn && yarn theia build && yarn theia download:plugins

# Create a laniakea venv
RUN mkdir /home/theia/.virtualenvs/ && python3 -m venv /home/theia/.virtualenvs/laniakea-venv

# Cache laniakea library and dependencies in the venv by creating a dummy project
RUN mkdir -p /home/theia/root/template
WORKDIR /home/theia/root/template

COPY pyproject.toml .
RUN . /home/theia/.virtualenvs/laniakea-venv/bin/activate && poetry install

# Install python dev tools in the virtualenv too
RUN pip3 install python-language-server flake8 autopep8 "poetry==1.0.10" pylint
WORKDIR /home/theia/root
RUN rm -rf /home/theia/root/template

# And prepopulate settings
RUN mkdir -p /home/theia/.theia
COPY settings.json /home/theia/.theia/settings.json

ENV SHELL=/bin/bash

WORKDIR /home/theia
COPY start.sh .
#RUN chmod +x start.sh

EXPOSE 3000
ENTRYPOINT ["/tini", "--"]
CMD ["bash","start.sh"]