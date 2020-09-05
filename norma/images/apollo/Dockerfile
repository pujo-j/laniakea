FROM node:10.16.0

ENV PREFECT_SERVER_COMMIT=c43868fbab3d35fe340adae562743c8e0044b30c

RUN mkdir -p /opt
WORKDIR /opt
RUN git clone https://github.com/PrefectHQ/server.git
WORKDIR /opt/server
RUN git checkout $PREFECT_SERVER_COMMIT
WORKDIR /opt/server/services/apollo
RUN npm install
RUN npm install --save-dev nodemon@2.0.4
RUN npm run build

RUN adduser --system prefect
USER prefect

CMD ["npm","run","serve"]