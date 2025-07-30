package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Snapshot struct {
	CPU    float64
	Memory float64
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

func Collect() (Snapshot, error) {
	cpu, err := snapCPU()
	if err != nil {
		return Snapshot{}, err
	}

	memory, err := snapMemory()
	if err != nil {
		return Snapshot{}, err
	}

	return Snapshot{
		CPU:    cpu,
		Memory: memory,
	}, nil
}
