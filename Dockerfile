FROM gliderlabs/alpine
RUN apk-install bash
EXPOSE 1338
COPY lowercase /
