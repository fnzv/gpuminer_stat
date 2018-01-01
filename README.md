# gpuminer_stat
Telegraf exporter for Nvidia GPU mining stats

### Usage
1) Git clone, enter directory and go build 
2) Install telegraf and configure it to collect metrics from the binary you just built nvidia-smi-export
  Add those lines on conf (telegraf.conf):
 [[inputs.exec]]
  command = "/home/ubuntu/nvidia-smi-export"
  data_format = "influx"
 
 3) Start telegraf, i made public also a Grafana dashboard https://grafana.com/dashboards/4214
  
