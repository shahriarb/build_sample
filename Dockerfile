FROM busybox
MAINTAINER Shahriar Boroujerdin

#TEST a CHANGE
RUN mkdir -p /app-tst

#switch to our app directory
RUN mkdir -p /app
WORKDIR /app

#copy the source files
COPY ./build_sample /app

CMD ["/app/build_sample"]
