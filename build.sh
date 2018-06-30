#! /bin/bash

cd D:/GoProject/src/MyProject/2017_6_23video_server/web
#env GOOS=linux GOARCH=arm64 go build -o ../bin/web
go build -o ../bin/web.exe
cd D:/GoProject/src/MyProject/2017_6_23video_server/api
go build -o ../bin/api.exe
#env GOOS=linux GOARCH=arm64 go build -o ../bin/api
cd D:/GoProject/src/MyProject/2017_6_23video_server/scheduler
go build -o ../bin/scheduler.exe
#env GOOS=linux GOARCH=arm64 go build -o ../bin/scheduler
cd D:/GoProject/src/MyProject/2017_6_23video_server/stream
go build -o ../bin/stream.exe
#env GOOS=linux GOARCH=arm64 go build -o ../bin/stream

