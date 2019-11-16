FROM envoyproxy/envoy-alpine
LABEL maintiner="Uimin Maksim"
COPY envoy.yaml /etc/envoy.yaml
VOLUME /var/log/envoy
CMD ["/etc/envoy.yaml"]
ENTRYPOINT ["envoy", "-c", "-l debug", "--log_path /var/log/envoy.log"]