FROM scratch

WORKDIR /opt
COPY docker-excel .
COPY web/ /opt/web
COPY upload/ /opt/upload
EXPOSE 9001
CMD ["/opt/docker-excel"]