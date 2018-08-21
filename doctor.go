package main

import (
    "fmt"
    "log"
    "encoding/json"

    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/host"
    "github.com/opalmer/check-go-version/api"
)

type Status struct {
    Os              string         `json:"os"`
    PlatformVersion string         `json:"platformVersion"`
    ActiveProcesses uint64         `json:"activeProcesses"`
    Uptime          uint64         `json:"uptime"`
    TotalRAM        uint64         `json:"totalRAM"`
    AvailableRAM    uint64         `json:"availableRAM"`
    UsedRAM         uint64         `json:"usedRAM"`
    UsedRAMPercent  float64        `json:"usedRAMPercent"`
    TotalDisk       uint64         `json:"totalDisk"`
    FreeDisk        uint64         `json:"freeDisk"`
    UsedDisk        uint64         `json:"usedDisk"`
    UsedDiskPercent float64        `json:"usedDiskPercent"`
    CPUModel        string         `json:"CPUModel"`
    CPUCores        int32          `json:"CPUcores"`
    GoVersion       string         `json:"goVersion"`
    GoPlatform      string         `json:"goPlatform"`
    GoArchitecture  string         `json:"goArchitecture"`
}

func getMachineStatus() Status {
    o, err := host.Info()
    if err != nil {
        fmt.Println("Could not retrieve OS details.")
        log.Fatal(err)
	}

    m, err := mem.VirtualMemory()
    if err != nil {
        fmt.Println("Could not retrieve RAM details.")
        log.Fatal(err)
	}

    d, err := disk.Usage("/")
    if err != nil {
        fmt.Println("Could not retrieve disk details.")
        log.Fatal(err)
    }

    c, err := cpu.Info()
    if err != nil {
        fmt.Println("Could not retrieve CPU details.")
        log.Fatal(err)
    }

    r, err := api.GetRunningVersion()
    if err != nil {
        fmt.Println("Could not retrieve Go installation details.")
        log.Fatal(err)
    }

    return Status{
        Os: o.OS,
        PlatformVersion: o.PlatformVersion,
        ActiveProcesses: o.Procs,
        Uptime: o.Uptime,
        TotalRAM: m.Total,
        AvailableRAM: m.Available,
        UsedRAM: m.Used,
        UsedRAMPercent: m.UsedPercent,
        TotalDisk: d.Total,
        FreeDisk: d.Free,
        UsedDisk: d.Used,
        UsedDiskPercent: d.UsedPercent,
        CPUModel: c[0].ModelName,
        CPUCores: c[0].Cores,
        GoVersion: r.FullVersion,
        GoPlatform: r.Platform,
        GoArchitecture: r.Architecture,
    }
}

func (s *Status) String() string {
	t, _ := json.Marshal(s)
	return string(t)
}

func main() {
    s := getMachineStatus()

    fmt.Printf(
        "OS: %v Version %v \nOS Active Processes: %v \nOS Uptime: %v\n" +
        "RAM Total: %v \nRAM Available: %v\nRAM Used: %v\nRAM Used Percent:%f%%\n" +
        "Disk Total: %v \nDisk Available: %v\nDisk Used: %v\nDisk Used Percent:%f%%\n" +
        "CPU Model: %v \nCPU Cores: %v \n" +
        "Go Version: %v\nGo Platform: %v\nGo Architecture: %v\n\n",
        s.Os, s.PlatformVersion, s.ActiveProcesses, s.Uptime,
        s.TotalRAM, s.AvailableRAM, s.UsedRAM, s.UsedRAMPercent,
        s.TotalDisk, s.FreeDisk, s.UsedDisk, s.UsedDiskPercent,
        s.CPUModel, s.CPUCores,
        s.GoVersion, s.GoPlatform, s.GoArchitecture,
    )

    // will return JSON
    fmt.Printf(s.String())

}
