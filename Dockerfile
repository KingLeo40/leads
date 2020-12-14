FROM scratch

COPY cmd/cmd .

EXPOSE 8087

ENTRYPOINT [ "./cmd" ]