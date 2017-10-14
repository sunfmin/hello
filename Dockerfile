FROM alpine
ADD front/build /www
ADD helloapp /bin/
CMD /bin/helloapp
