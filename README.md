# docker-firebase-emulator-auth

## Usage

Build docker image
```
$ docker build -t docker-firebase-emulator-auth .
```

Run firebase emulator
```
$ docker run -it -p 4000:4000 -p 9099:9099 docker-firebase-emulator-auth:latest
```
