package main

import (
        "flag"
        "fmt"
        "log"
        "os"
        "os/exec"
        "strings"
)

var opts []string

func exec_shell(command string) string {
out, err := exec.Command("/bin/bash","-c",command).Output()
    if err != nil {
        log.Fatal(err)
    }
    return string(out)
}

func main() {
        binPath := flag.String("bin", "/usr/bin/nvidia-smi", "nvidia-smi full path")

        flag.Parse()

        if _, err := os.Stat(*binPath); os.IsNotExist(err) {
                fmt.Fprintf(os.Stderr, "Bin path does not exists: %s", *binPath)
                return // exit
        }

        results := exec_shell("nvidia-smi --format=noheader,nounits,csv --query-gpu=fan.speed,memory.total,memory.used,memory.free,pstate,temperature.gpu,name,uuid,compute_mode -i 0,1,2,3")
        if results == "" {
                return // exit
        }

for _, line := range strings.Split(strings.TrimSuffix(results, "\n"), "\n") {
    splitResults := strings.Split(line, ",")

        fmt.Printf("nvidiasmi,uuid=%s ", strings.TrimSpace(splitResults[7]))

        fmt.Printf("gpu_name=\"%s\",", strings.TrimSpace(splitResults[6]))
        fmt.Printf("gpu_compute_mode=\"%s\",", strings.TrimSpace(splitResults[8]))

        fmt.Printf("fan_speed=%s,", strings.TrimSpace(splitResults[0]))

        fmt.Printf("memory_total=%s,", strings.TrimSpace(splitResults[1]))
        fmt.Printf("memory_used=%s,", strings.TrimSpace(splitResults[2]))
        fmt.Printf("memory_free=%s,", strings.TrimSpace(splitResults[3]))

        fmt.Printf("pstate=%s,", strings.TrimSpace(strings.Replace(splitResults[4], "P", "", -1)))
        fmt.Printf("temperature=%s\n", strings.TrimSpace(splitResults[5]))

                }

}
