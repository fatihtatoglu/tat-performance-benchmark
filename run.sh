#!/bin/bash

docker compose up -d cadvisor

sleep 120

docker compose up -d prometheus

sleep 120

docker compose up -d grafana

sleep 30

# C
docker compose up -d fibonacci-c
while :
do
    if docker compose ps -a | grep -q "fibonacci-c.*Exit"; then
        docker compose up -d fibonacci-c-optimized
        break
    fi
    sleep 30
done

# CPP
while :
do
    if docker compose ps -a | grep -q "fibonacci-c-optimized.*Exit"; then
        docker compose up -d fibonacci-cpp
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-cpp.*Exit"; then
        docker compose up -d fibonacci-cpp-optimized
        break
    fi
    sleep 30
done

# C#
while :
do
    if docker compose ps -a | grep -q "fibonacci-cpp-optimized.*Exit"; then
        docker compose up -d fibonacci-csharp
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-csharp.*Exit"; then
        docker compose up -d fibonacci-csharp-optimized
        break
    fi
    sleep 30
done

# Go
while :
do
    if docker compose ps -a | grep -q "fibonacci-csharp-optimized.*Exit"; then
        docker compose up -d fibonacci-go
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-go.*Exit"; then
        docker compose up -d fibonacci-go-optimized
        break
    fi
    sleep 30
done

# NodeJS
while :
do
    if docker compose ps -a | grep -q "fibonacci-go-optimized.*Exit"; then
        docker compose up -d fibonacci-node
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-node.*Exit"; then
        docker compose up -d fibonacci-node-optimized
        break
    fi
    sleep 30
done

# Python
while :
do
    if docker compose ps -a | grep -q "fibonacci-node-optimized.*Exit"; then
        docker compose up -d fibonacci-python
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-python.*Exit"; then
        docker compose up -d fibonacci-python-optimized
        break
    fi
    sleep 30
done

# Rust
while :
do
    if docker compose ps -a | grep -q "fibonacci-python-optimized.*Exit"; then
        docker compose up -d fibonacci-rust
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-rust.*Exit"; then
        docker compose up -d fibonacci-rust-optimized
        break
    fi
    sleep 30
done

# Scala
while :
do
    if docker compose ps -a | grep -q "fibonacci-rust-optimized.*Exit"; then
        docker compose up -d fibonacci-scala
        break
    fi
    sleep 30
done
while :
do
    if docker compose ps -a | grep -q "fibonacci-scala.*Exit"; then
        docker compose up -d fibonacci-scala-optimized
        break
    fi
    sleep 30
done

# Java
while :
do
    if docker compose ps -a | grep -q "fibonacci-scala-optimized.*Exit"; then
        docker compose up -d fibonacci-java
        break
    fi
    sleep 30
done

# Ruby
while :
do
    if docker compose ps -a | grep -q "fibonacci-java.*Exit"; then
        docker compose up -d fibonacci-ruby
        break
    fi
    sleep 30
done

# Perl
while :
do
    if docker compose ps -a | grep -q "fibonacci-ruby.*Exit"; then
        docker compose up -d fibonacci-perl
        break
    fi
    sleep 30
done

# Lua
while :
do
    if docker compose ps -a | grep -q "fibonacci-perl.*Exit"; then
        docker compose up -d fibonacci-lua
        break
    fi
    sleep 30
done