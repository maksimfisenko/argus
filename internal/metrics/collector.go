package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
)

type Snapshot struct {
	CPU       float64
	Memory    float64
	DiskUsage float64
	AvgLoad   float64
	Uptime    uint64
}

func snapCPU() (float64, error) {
	cpu, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return cpu[0], nil
}

func snapMemory() (float64, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vm.UsedPercent, nil
}

func snapDiskUsage() (float64, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return usage.UsedPercent, nil
}

func snapAvgLoad() (float64, error) {
	avg, err := load.Avg()
	if err != nil {
		return 0, err
	}
	return avg.Load1, nil
}

func snapUptime() (uint64, error) {
	return host.Uptime()
}

func Collect() (Snapshot, error) {
	cpu, err := snapCPU()
	if err != nil {
		return Snapshot{}, err
	}

	memory, err := snapMemory()
	if err != nil {
		return Snapshot{}, err
	}

	diskUsage, err := snapDiskUsage()
	if err != nil {
		return Snapshot{}, err
	}

	avgLoad, err := snapAvgLoad()
	if err != nil {
		return Snapshot{}, err
	}

	uptime, err := snapUptime()
	if err != nil {
		return Snapshot{}, err
	}

	return Snapshot{
		CPU:       cpu,
		Memory:    memory,
		DiskUsage: diskUsage,
		AvgLoad:   avgLoad,
		Uptime:    uptime,
	}, nil
}
