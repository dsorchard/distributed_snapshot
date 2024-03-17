```shell
go build
```

```bash
#!/bin/bash

# Define the directory where your project is located
project_dir="/Users/arjunsunilkumar/GolandProjects/0sysdev_dec/snapshot/cmd/server"

# Ports to be checked and processes to be killed if found
ports=(8100 8101 8102)

# Kill process running on each port
for port in "${ports[@]}"; do
  pid=$(lsof -ti tcp:$port)
  if [ ! -z "$pid" ]; then
    echo "Killing process on port $port"
    kill -9 $pid
  fi
done

# Use AppleScript to open a new Terminal tab and run each command
osascript <<EOF
tell application "Terminal"
  activate
  do script "cd $project_dir && ./server 0 network.json"
end tell

delay 1
tell application "System Events" to tell process "Terminal" to keystroke "t" using command down
delay 1
tell application "Terminal" to do script "cd $project_dir && ./server 1 network.json" in front window

delay 1
tell application "System Events" to tell process "Terminal" to keystroke "t" using command down
delay 1
tell application "Terminal" to do script "cd $project_dir && ./server 2 network.json" in front window
EOF
```