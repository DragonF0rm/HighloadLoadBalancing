FROM prom/prometheus
LABEL maintainer="Uimin Maksim"
COPY prometheus.yaml /etc/prometheus.yaml
CMD ["--config.file=/etc/prometheus.yaml"]
ENTRYPOINT ["prometheus"]