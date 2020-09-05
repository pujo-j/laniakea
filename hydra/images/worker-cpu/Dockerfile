FROM ubuntu:focal

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
RUN apt-get install -y -qq \
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
    vim

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

# Install chrome (headless)

RUN curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& echo "deb https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list

RUN apt-get update && apt-get install -y \
	google-chrome-stable \
	&& rm -rf /var/lib/apt/lists/*

# Add Tini
ENV TINI_VERSION v0.19.0
RUN wget -O /tini https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini
RUN chmod +x /tini

# Add poetry
RUN pip3 install --system "poetry==1.0.10"


# Switch to user
RUN adduser --system prefect

COPY laniakea/pyproject.toml /home/prefect
COPY laniakea/poetry.lock /home/prefect
RUN chown -R prefect /home/prefect
USER prefect

WORKDIR /home/prefect

# Preinstall dependencies to cache layer
RUN poetry install --no-root
RUN poetry run pyppeteer-install

# And finally install laniakea
COPY laniakea/ ./

USER root
RUN chown prefect /home/prefect/* && chown -R prefect /home/prefect/laniakea*
USER prefect

RUN poetry install

ENTRYPOINT ["/tini", "--"]