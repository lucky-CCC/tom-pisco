#!/bin/bash
#使用demo ./kill.sh 4222 根据端口杀死进程

if [ -z $1 ]; then
        echo "you must input a port"
        exit 0
fi

PID=$(lsof -i | grep ":$1" | awk '{print $2}' | awk -F '[ / ]' '{print $1}')

if [ $? == 0 ]; then
        echo "process id is:${PID}"
else
        echo "process $1 no exit"
        exit 0
fi

kill -9 ${PID}

if [ $? == 0 ]; then
        echo "kill $1 success"
else
        echo "kill $1 fail"
fi