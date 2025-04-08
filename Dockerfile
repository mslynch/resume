FROM gotenberg/gotenberg:8

USER root

RUN echo "deb http://http.us.debian.org/debian stable main contrib" >> /etc/apt/sources.list

RUN apt-get update && apt-get install -y fonts-ibm-plex fonts-inter unzip

# TODO: Geist is kind of nice I guess?
# RUN curl "https://github.com/vercel/geist-font/releases/download/1.4.01/Geist-v1.4.01.zip" -L -o /tmp/geist.zip \
#     && unzip /tmp/geist.zip -d /tmp/geist/ \
#     && mkdir -p /home/gotenberg/.local/share/fonts/opentype/geist \
#     && cp -R /tmp/geist/Geist-v1.4.01/otf/ /home/gotenberg/.local/share/fonts/ \
#     && chmod -R 644 /home/gotenberg/.local/share/fonts/opentype/geist \
#     && rm -rf /tmp/geist.zip /tmp/geist
    
USER gotenberg
