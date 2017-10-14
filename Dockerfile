FROM alpine
ADD front/build /www
ADD hello /bin/
CMD /bin/hello

