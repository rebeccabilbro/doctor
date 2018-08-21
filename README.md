# doctor

Doctor provides a simple machine health status report.

Using the [`gopsutil`](https://github.com/shirou/gopsutil) and [`check-go-version`](https://github.com/opalmer/check-go-version) libraries, `doctor` constructs `Status`, a custom `struct` with the following information:

- Os              
- PlatformVersion
- ActiveProcesses
- Uptime         
- TotalRAM       
- AvailableRAM   
- UsedRAM        
- UsedRAMPercent  
- TotalDisk   
- FreeDisk
- UsedDisk  
- UsedDiskPercent
- CPUModel  
- CPUCores  
- GoVersion   
- GoPlatform    
- GoArchitecture

The `getMachineStatus()` function returns a `Status` with all of the above information populated.

## Usage

```bash
$ go get github.com/rebeccabilbro/doctor
$ go get # to get the other dependencies
$ go run doctor.go
```

## Sample Output

```bash
OS: darwin Version 10.11.6
OS Active Processes: 334
OS Uptime: 8374703
RAM Total: 17179869184
RAM Available: 5863411712
RAM Used: 11316457472
RAM Used Percent:65.870452%
Disk Total: 499055067136
Disk Available: 61215113216
Disk Used: 437577809920
Disk Used Percent:87.727349%
CPU Model: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
CPU Cores: 4
Go Version: 1.10.3
Go Platform: darwin
Go Architecture: amd64
```

Note: Output can also be retrieved as a JSON object using `String()`:

```
{"os":"darwin","platformVersion":"10.11.6","activeProcesses":334,"uptime":8374703,"totalRAM":17179869184,"availableRAM":5863411712,"usedRAM":11316457472,"usedRAMPercent":65.87045192718506,"totalDisk":499055067136,"freeDisk":61215113216,"usedDisk":437577809920,"usedDiskPercent":87.72734929133924,"CPUModel":"Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz","CPUcores":4,"goVersion":"1.10.3","goPlatform":"darwin","goArchitecture":"amd64"}
```
