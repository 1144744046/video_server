#! /bin/bash

cp -R ./template ./bin/
mkdir ./bin/video

cd bin
nohup ./web &
nohup ./api &
nohup ./scheduler &
nohup ./stream &