package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var (
	suffix        = " :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::"
	end_suffix    = "================ <<: :>> ================"
	chk_words     = [3]string{"Failed", "Error", "panic"}
	lastStartDate time.Time
)

func getDate(line string) time.Time {
	line_s := strings.TrimSuffix(line, suffix)
	dateTimeString := strings.TrimPrefix(line_s, "\x02\x00\x00\x00\x00\x00\x00R")

	layout := "2006/01/02 15:04:05"

	t, err := time.Parse(layout, dateTimeString)

	if err != nil {
		fmt.Println("Error parsing date-time:", err)
	}

	return t
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("Error creating Docker client: %v\n", err)
		return
	}

	containerID := "mfins"

	logOptions := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: false,
	}

	logs, err := cli.ContainerLogs(context.Background(), containerID, logOptions)
	if err != nil {
		fmt.Printf("Error fetching logs: %v\n", err)
		return
	}
	defer logs.Close()

	scanner := bufio.NewScanner(logs)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasSuffix(line, suffix) {
			lastStartDate = getDate(line)
		}

		for _, word := range chk_words {
			if strings.Contains(line, word) {
				err_trim := strings.TrimPrefix(line, "\x02\x00\x00\x00\x00\x00\x00R")
				log.Printf("Error: MFINS service failed: %s", err_trim)
			}
		}

		if strings.HasSuffix(line, end_suffix) {
			log.Printf("Last successful run: %s", lastStartDate.String())
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Printf("Error reading logs: %v\n", err)
	}
}
