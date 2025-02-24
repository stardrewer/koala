package main

var config_template = `port: 8080
prometheus:
  switch_on: true
  port: 8081
service_name: {{.Package.Name}}
register:
  switch_on: true
  register_path: /ibinarytree/koala/service/
  timeout: 1s
  heart_beat: 10
  register_name: etcd
  register_addr: 127.0.0.1:2379
log:
  level: debug
  path: ./logs/
  chan_size: 10000
  console_log: true
limit:
  switch_on: true
  qps: 50000
trace:
  switch_on: true
  report_addr: http://60.205.218.189:9411/api/v1/spans
  sample_type: const
  sample_rate: 1
`
