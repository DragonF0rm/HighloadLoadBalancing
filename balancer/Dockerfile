FROM envoyproxy/envoy-alpine
LABEL maintiner="Uimin Maksim"
COPY envoy.yaml /etc/envoy.yaml
VOLUME /var/log/envoy
EXPOSE 80/tcp
EXPOSE 9901/tcp
CMD ["/etc/envoy.yaml"]
ENTRYPOINT ["envoy", "-c"]