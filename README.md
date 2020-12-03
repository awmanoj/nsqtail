# nsqtail

## Introduction 
nsqtail helps tail upto last 10 messages on any topic in the NSQ message queue.

## Docker 
```
$ docker pull awmanoj/nsqtail-dock:latest
$ ... 
$ docker run .. 
```

## Build 
```
parallels$ docker build . -t nsqtail-dock
```
```
Sending build context to Docker daemon  11.38MB
Step 1/12 : FROM golang:alpine
 ---> 7fce7b3ec0f9
Step 2/12 : ENV GO111MODULE=on 	CGO_ENABLED=0 	GOOS=linux 	GOARCH=amd64
 ---> Using cache
 ---> 04f73d7cf435
Step 3/12 : WORKDIR /build
 ---> Using cache
 ---> 757851eb8eab
Step 4/12 : COPY go.mod .
 ---> Using cache
 ---> 5273c44a2245
Step 5/12 : COPY go.sum .
 ---> Using cache
 ---> 5d4f8d14816a
Step 6/12 : RUN go mod download
 ---> Using cache
 ---> 88a123e0a534
Step 7/12 : COPY . .
 ---> Using cache
 ---> 75a61c6d6dab
Step 8/12 : RUN go build -o main .
 ---> Running in 363dcaff56bb
...
Successfully built 40764fb38b6c
Successfully tagged nsqtail-dock:latest
```

## Run 
```
parallels$ docker run --net=host -p 8080:8080 nsqtail-dock
```
```
WARNING: Published ports are discarded when using host network mode
2020/11/10 11:11:34 INF    1 [test/nsqtail-taxicab-1729] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=test
2020/11/10 11:11:34 INF    1 [test/nsqtail-taxicab-1729] (parallels-Parallels-Virtual-Platform:4150) connecting to nsqd
2020/11/10 11:11:34 INF    2 [answers/nsqtail-taxicab-1729] querying nsqlookupd http://127.0.0.1:4161/lookup?topic=answers
2020/11/10 11:11:34 INF    2 [answers/nsqtail-taxicab-1729] (parallels-Parallels-Virtual-Platform:4150) connecting to nsqd
2020/11/10 11:11:34 Starting server at port 8080
```
## Interface

#### / (all topics) 
<img width="1376" alt="topics" src="https://user-images.githubusercontent.com/1171470/100529898-7e0b2780-321e-11eb-8c51-62ab2ad1d99e.png">

#### /nsqtail/test1 (last N messages from topic=test1) 
<img width="1376" alt="messages" src="https://user-images.githubusercontent.com/1171470/100529895-79467380-321e-11eb-999c-2a50dd10fdf9.png">
