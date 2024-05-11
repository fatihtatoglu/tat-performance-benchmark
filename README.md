# Tat Performance Benchmark

After reading an amazing post from Piotr Kołaczkowski ([How Much Memory Do You Need to Run 1 Million Concurrent Tasks?](https://pkolaczk.github.io/memory-consumption-of-async/)), I got curious about that Measurements of memory and CPU usage and execution times of common computing tasks in different language runtimes. In addition, I hypothesize that changing programming languages from X to Y does not provide beneficial differences.

My approach contains the following steps.

- Select the most popular programming languages.
- Find a challenging problem.
- Decide test environment.
- Decide how to measure the performance.
- Implement the related codes for every programming language.
- Execute.
- Collect the results.
- Evaluate the results.

## Programming Languages

Selecting the most popular programming languages was the easiest step in this project. I used the [Stackoverflow 2023 Developer Survey](https://survey.stackoverflow.co/2023/#most-popular-technologies-language) for choosing the below programming languages.

1. C
2. C++
3. C#
4. Go
5. Java
6. Lua
7. NodeJs
8. Perl
9. Python
10. Ruby
11. Rust
12. Scala

I'm not experienced in Scala, Rust, Ruby, and Lua programming languages. However, another assumption of mine has worked.

**Assumption:** If you know the structure of the programming languages and terms, you can learn a programming language in one day. Learning means you can understand the business flow of the program.

Without using ChatGPT or any other AI assistant (by using Google Search to find the programming language documentation) I could develop the solution to the problem.

## Challenging Problem

The selected problem should take time to find the solution and use a high CPU and enough memory resources. Otherwise, measuring performance wasn't meaningful. So, I decided to select a mathematical problem to solve by code.

In the beginning, I selected factorial calculation, nevertheless overflowing the 64-bit integer's limit blocked the implementations. Then I selected the calculation of the [Fibonacci Sequence](https://en.wikipedia.org/wiki/Fibonacci_sequence). It takes time and is also suitable for recursive implementation. To make the challenge a little bit harder, I preferred the recursive implementation of the sequence.

By the way, the Fibonacci sequence is one of the challenges. I will add some additional challenges, such as string operations, and different mathematical calculations.

## Fair Testing Environment

To be fair while measuring the performance of the programming languages, providing an isolated environment is essential. For this reason, I use docker by limiting CPU and memory. Moreover, I have started the images in order.

```yml
  fibonacci-c:
    build: ./fibonacci/c
    cpus: "2"
    mem_reservation: "1G"
    mem_limit: "4G"
```

On the other hand, isolating the environment cannot be enough many times. The operation systems also use the resources. At this point, Alpine Linux comes to my rescue. Implementing Dockerfiles from the Alpine-based images reduced unnecessary resource usage.

I have used the multi-stage image build to prevent additional resource usage, for suitable programming languages, such as C, C++, C#, Go, etc.

## Monitoring Resources

After impleting the applications, the correct measurement is the key point of the project. As I know `docker stats` doesn't provide enough information. So, I used cAdvisor, Prometheus, and Grafana.

The implementation was so easy, but there was some tricky points, like selecting correct metrics and connecting the tools to each others. The helpful resources is below.

- [Collect Docker metrics with Prometheus](https://docs.docker.com/config/daemon/prometheus/)
- [Docker Container Monitoring with cAdvisor, Prometheus, and Grafana using Docker Compose](https://medium.com/@sohammohite/docker-container-monitoring-with-cadvisor-prometheus-and-grafana-using-docker-compose-b47ec78efbc)
- [Measure execution time of docker container](https://stackoverflow.com/questions/51650405/measure-execution-time-of-docker-container)

## Execution

### Prerequisites

- Docker Engine (I tested 36.1.2)
- Docker Compose v2.27.0 or newer
- Go 1.18 or higher (for reporting)

```bash
docker compose build
```

```bash
chmod +x run.sh
./run.sh
```

By running the `run.sh` all the containers work in an order.

## Collecting Measurements

In the project, there are two parts to collect the measurements. One is the Grafana dashboard and raw data. Another one is the execution logs of the application that is implemented for every programming language's code.

In addition, to find out the general information I used the `docker container inspect` command.

```bash
go run ./reporting/main.go > out.json
```

With the above command general information and application execution logs are extracted into the JSON file.

## Evaluation

A long story short, I was shocked at the first sight. Then, I decided to share those measurements with other developers and curious people without comments. Maybe, my testing environment was wrong or the selected problem wasn't suitable for this purpose.

According to the contributors' comments, I should evaluate the results. I need your help to complete the evaluation step.

- Validate the implementation of the applications.
  - Validate the compiler-level optimizations.
- Validate the test environments.
  - Limiting CPU, Memory, or other resources.
  - Validate the isolation is done.
- Validate the metric measurements.
  - CPU
  - Memory
  - Execution Time
  - Step level throughput

## Contribution

Please share your ideas on repository issues. Moreover, you can share your results (Grafana CSV exports and out.json file) over the issue template if you want.

### Exporting Data From Garafana

Please track these steps after completing all scenarios.

- Select **Docker Monitoring Test** dashboard.
- From dashboard select **Last 24 Hours**
- Open the menu of the **CPU Usage** Chart.
- Select **Inspect**.
- Select **JSON** and (from ***Select source*** list) **Panel data**.
- Copy and paste that JSON.
- Repeat the same steps for the **Memory Usage** Chart.

For **CPU Usage** data, name the file **grafana-cpu.json**, and for **Memory Usage** data, name the file **grafana-memory.json**.

### Docker Logs Export

Please track these steps after completing all scenarios.

- Execute the following code.

```bash
go run ./reporting/main.go > out.json
```

- out.json contains all related data and application logs.

After exporting Grafana and Log data, please create an issue and upload those files (out.json, grafana-cpu.json, grafana-memory.json).
