/*
	Async Job for collecting host metrics
*/

package main

import (
	"log"
	"math"
	"time"

	queue "github.com/mink0/go-dashboard/lib"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	statInterval  = 3 * time.Second
	statQueueSize = int(24 * time.Hour / statInterval)
)

type Stats struct {
	cpu    *queue.Queue
	memory *queue.Queue
	ticker *time.Ticker
}

func statsJobRunner() (stats Stats) {
	stats.ticker = time.NewTicker(1 * time.Second)
	stats.cpu = queue.New(statQueueSize)
	stats.memory = queue.New(statQueueSize)

	// to stop it call: ticker.Stop()
	go func() {
		for range stats.ticker.C {
			p, err := cpu.Percent(statInterval, false)
			if err != nil {
				log.Println("Error fetching cpu usage info", err)
				return
			}

			stats.cpu.Push(math.Ceil(p[0]))

			v, err := mem.VirtualMemory()
			if err != nil {
				log.Println("Error fetching memory usage info", err)
				return
			}

			stats.memory.Push(math.Ceil(v.UsedPercent))
		}
	}()

	return
}
