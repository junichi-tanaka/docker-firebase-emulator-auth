FROM node:18-alpine

RUN npm install -g firebase-tools
RUN firebase setup:emulators:ui

EXPOSE 9099
EXPOSE 4000

RUN mkdir -p /work
WORKDIR /work
COPY firebase.json .

CMD ["firebase", "emulators:start", "--project", "test", "--only", "auth"]
