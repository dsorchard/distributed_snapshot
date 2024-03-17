# Distributed Snapshot

Chandy Lamport Algorithm.

### Reference
- [Project 1](https://github.com/FreddyMartinez/chandy-lamport-go)
- [Project 2](https://github.com/mariuscrsn/GlobalStateSnapshot)

### Vector Clock Trace

![Vector Clock Trace](misc/shiviz.png)

### Command To Generate Shiviz Log

1. Run the following to generate the log file
```shell
go install github.com/DistributedClocks/GoVector

GoVector --log_type shiviz --log_dir /Users/arjunsunilkumar/GolandProjects/0sysdev_dec/chandy-lamport-go/govector --outfile final.log
```

2. Open the log file in [ShiViz](https://bestchai.bitbucket.io/shiviz/)

3. Use the following regular expression to parse the log file
```shell
(?<host>\S*)\s+(?<clock>\{.*?\})\s+(?<event>.*)
```

4. Input would look like this

![Shiviz Input](misc/shiviz2.png)