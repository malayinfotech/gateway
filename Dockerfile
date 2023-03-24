ARG DOCKER_ARCH
FROM ${DOCKER_ARCH:-amd64}/alpine
ARG TAG
ARG GOARCH
ENV GOARCH ${GOARCH}
EXPOSE 7777
WORKDIR /app
VOLUME /root/.local/share/storx/gateway
COPY release/${TAG}/gateway_linux_${GOARCH:-amd64} /app/gateway
COPY entrypoint /entrypoint
ENTRYPOINT ["/entrypoint"]
ENV STORX_CLIENT_ADDITIONAL_USER_AGENT=DockerOfficialImage/1.0.0
ENV STORX_CONFIG_DIR=/root/.local/share/storx/gateway
ENV STORX_SERVER_ADDRESS=0.0.0.0:7777
