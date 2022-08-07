FROM scratch
COPY getcurrencyinfo /usr/local/bin/getcurrencyinfo
ENTRYPOINT [ "/usr/local/bin/getcurrencyinfo" ]