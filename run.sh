#!/bin/bash

docker compose up -d cadvisor
docker compose up -d prometheus
docker compose up -d grafana

sleep 60

docker compose up -d fibonacci-c

while :
do
    if docker compose ps -a | grep -q "fibonacci-c.*Exit"; then
        docker compose up -d fibonacci-cpp
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-cpp.*Exit"; then
        docker compose up -d fibonacci-csharp
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-csharp.*Exit"; then
        docker compose up -d fibonacci-go
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-go.*Exit"; then
        docker compose up -d fibonacci-java
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-java.*Exit"; then
        docker compose up -d fibonacci-lua
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-lua.*Exit"; then
        docker compose up -d fibonacci-node
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-node.*Exit"; then
        docker compose up -d fibonacci-perl
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-perl.*Exit"; then
        docker compose up -d fibonacci-python
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-python.*Exit"; then
        docker compose up -d fibonacci-ruby
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-ruby.*Exit"; then
        docker compose up -d fibonacci-rust
        break
    fi
    sleep 60
done

while :
do
    if docker compose ps -a | grep -q "fibonacci-rust.*Exit"; then
        docker compose up -d fibonacci-scala
        break
    fi
    sleep 60
done