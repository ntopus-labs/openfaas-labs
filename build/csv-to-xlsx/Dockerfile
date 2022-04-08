FROM ghcr.io/openfaas/classic-watchdog:0.2.0 as watchdog

FROM linuxserver/libreoffice:7.2.2

RUN mkdir -p /home/app

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Add non root user
# RUN addgroup -S app && adduser app -S -G app
# RUN chown app /home/app

WORKDIR /home/app

# USER app

RUN echo "#!/bin/bash" > /home/app/start.sh
RUN echo "cat - > stdin-file.csv" >> /home/app/start.sh
RUN echo 'libreoffice --headless --convert-to xlsx --infilter="CSV:59/44,34,76,1" stdin-file.csv > /dev/null' >> /home/app/start.sh
RUN echo "cat stdin-file.xlsx" >> /home/app/start.sh
RUN chmod +x start.sh

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="/home/app/start.sh"
# Set to true to see request in function logs
ENV write_debug="false"



EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

ENTRYPOINT ["fwatchdog"]
