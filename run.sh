#!/bin/bash

echo "------------$(date +%F' '%T)------------"
# 开发重启脚本

file=game-control-api

getPid() {
  docmd=$(ps aux | grep ${file} | grep ${file} | grep -v 'grep' | grep -v '\.sh' | awk '{print $2}')
  echo $docmd
}

start() {
  pidstr=$(getPid)
  if [ -n "$pidstr" ]; then
    echo "running with pids $pidstr"
  else
    rm -rf material-mis
    echo "正在编译中..."
    go build -race
    sleep 0.5
    printf "\n"
    printf "正在执行启动...稍候"
    printf "\n"
#    nohup ./material-mis >/dev/null 2>&1 &
    ./game-control-api -e dev
    pidstr=$(getPid)
    echo "start with pids $pidstr Successful"
  fi
}

stop() {

  pidstr=$(getPid)
  if [ ! -n "$pidstr" ]; then
    echo "Not Executed!"
    return
  fi

  echo "kill $pidstr done"
  kill $pidstr

}

restart() {
  stop
  start
}

case "$1" in
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  restart
  ;;
getpid)
  getPid
  ;;
esac