FROM registry.fedoraproject.org/fedora-minimal:33

RUN microdnf install -y wkhtmltopdf unzip && microdnf clean all

WORKDIR /resume

# variable fonts seem to have poor support in wkhtmltopdf, so we're using this. :(
RUN curl https://github.com/rsms/inter/releases/download/v3.15/Inter-3.15.zip -L -o ./inter.zip         && unzip ./inter.zip         -d /usr/share/fonts && rm inter.zip
RUN curl https://fonts.google.com/download?family=IBM%20Plex%20Sans           -L -o ./ibm_plex_sans.zip && unzip ./ibm_plex_sans.zip -d /usr/share/fonts && rm ibm_plex_sans.zip

RUN mkdir target
COPY ./bin ./bin
CMD './bin/build'
