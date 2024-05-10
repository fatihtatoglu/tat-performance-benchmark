package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Container struct {
	Id          string         `json:"id"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	StartedAt   time.Time      `json:"started_at"`
	FinishedAt  time.Time      `json:"finished_at"`
	Logs        []ExecutionLog `json:"logs"`
	ElapsedTime int64          `json:"elapsed_time_second"`
}

type ExecutionLog struct {
	Step         int       `json:"step"`
	Result       uint64    `json:"result"`
	CalculatedAt time.Time `json:"calculated_at"`
	ElapsedTime  int64     `json:"elapsed_time_millisecond"`
}

func main() {
	list, err := getContainers()
	if err != nil {
		panic(err)
	}

	containers := make([]Container, 0)

	for _, id := range list {
		if id == "" {
			continue
		}

		container := Container{
			Id: id,
		}

		dates, err := getContainerExecutionDates(id)
		if err != nil {
			panic(err)
		}

		container.CreatedAt = dates["created"]
		container.StartedAt = dates["started"]
		container.FinishedAt = dates["finished"]

		container.ElapsedTime = int64(container.FinishedAt.Sub(container.StartedAt).Seconds())

		logs, err := getExecutionLogs(id)
		if err != nil {
			panic(err)
		}

		container.Logs = logs

		imageName, err := getContainerImageName(id)
		if err != nil {
			panic(err)
		}

		container.Image = imageName

		containers = append(containers, container)
	}

	jsonData, err := json.MarshalIndent(containers, "", "  ")
	if err != nil {
		fmt.Println("JSON formatına dönüştürülürken bir hata oluştu:", err)
		return
	}

	// JSON verisini konsola yazdır
	fmt.Println(string(jsonData))
}

func getContainers() ([]string, error) {
	command := "docker container list --all --format \"{{.ID}}\" --filter \"name=^/tat-perf-bench\" --filter \"status=exited\""
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	list := strings.Split(string(out), "\n")

	return list, nil
}

func getContainerExecutionDates(id string) (map[string]time.Time, error) {
	command := "docker container inspect " + id + " --format \"{{.Created}}\t{{.State.StartedAt}}\t{{.State.FinishedAt}}\""
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	list := strings.Split(strings.ReplaceAll(string(out), "\n", ""), "\t")

	dates := make(map[string]time.Time)
	for i, date := range list {
		if i == 0 {
			dates["created"] = parseDateFromRFC3339(date)
		} else if i == 1 {
			dates["started"] = parseDateFromRFC3339(date)
		} else if i == 2 {
			dates["finished"] = parseDateFromRFC3339(date)
		}
	}

	return dates, nil
}

func getExecutionLogs(id string) ([]ExecutionLog, error) {
	command := "docker logs " + id
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	logs := make([]ExecutionLog, 0)
	var prevLog ExecutionLog
	for i, line := range lines {
		if line == "" {
			continue
		}

		log := getExecutionLog(line)
		if i == 0 {
			log.ElapsedTime = 0
		} else {
			log.ElapsedTime = int64(log.CalculatedAt.Sub(prevLog.CalculatedAt).Seconds())
		}

		logs = append(logs, log)
		prevLog = log
	}

	return logs, nil
}

func parseDateFromRFC3339(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}

	return t
}

func getExecutionLog(line string) ExecutionLog {
	log := ExecutionLog{}

	pattern := `(\d+)\s*:\s*(\d+)\s*\(\s*(\d+)\s*\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(line)
	if len(matches) == 4 {
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.ParseUint(matches[2], 10, 64)
		z, _ := strconv.ParseUint(matches[3], 10, 64)

		log.Step = x
		log.Result = y

		if len(matches[3]) == 13 {
			log.CalculatedAt = time.UnixMilli(int64(z))
		} else if len(matches[3]) == 10 {
			log.CalculatedAt = time.Unix(int64(z), 0)
		} else if len(matches[3]) == 16 {
			log.CalculatedAt = time.UnixMicro(int64(z))
		}
	}

	return log
}

func getContainerImageName(id string) (string, error) {
	command := "docker container inspect " + id + " --format \"{{.Config.Image}}\""
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		return "", err
	}

	imageName := strings.ReplaceAll(string(out), "\n", "")

	return imageName, nil
}
