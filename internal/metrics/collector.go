package metrics

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func Collect() (float64, float64, error) {
	cpu, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, 0, err
	}

	vm, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, err
	}

	return cpu[0], vm.UsedPercent, nil
}
