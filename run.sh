#!/bin/bash

docker compose up -d cadvisor
docker compose up -d prometheus
docker compose up -d grafana

sleep 10

docker compose up -d c

while :
do
    if docker compose ps -a | grep -q "c.*Exit"; then
        docker compose up -d cpp
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "cpp.*Exit"; then
        docker compose up -d csharp
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "csharp.*Exit"; then
        docker compose up -d go
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "go.*Exit"; then
        docker compose up -d java
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "java.*Exit"; then
        docker compose up -d lua
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "lua.*Exit"; then
        docker compose up -d node
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "node.*Exit"; then
        docker compose up -d perl
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "perl.*Exit"; then
        docker compose up -d python
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "python.*Exit"; then
        docker compose up -d ruby
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "ruby.*Exit"; then
        docker compose up -d rust
        break
    fi
    sleep 10
done

while :
do
    if docker compose ps -a | grep -q "rust.*Exit"; then
        docker compose up -d scala
        break
    fi
    sleep 10
done