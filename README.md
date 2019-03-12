# Profiling a Go application
## How does a profiler work?
A profiler runs the program and configures the OS to send interrupts to the process at regular intervals. 
This is done by sending SIGPROF to the program, which suspends the execution and transfers the execution to the profiler. The profiler then grabs the program counter for executing thread and restarts the program.

## Runtime Profiling
Profiling a program can tell you a lot about the memory and CPU usage of different functions in the program. This can greatly help you find the bottlenecks and optimize the program by hitting the bulls eye. 
### How I did it
1. Import the package "runtime/pprof"

2. Install graphviz [graphviz](https://www.graphviz.org/)

3. Now, I've done profiling for CPU and memory one by one. First I profiled a sample Golang code (prime number calculation) for CPU usage. For that first I added this line of code at the top of the main:
```
var cpuProfile = flag.String("cpu", "", "write cpu profile to file")
var memProfile = flag.String("memory", "", "write memory profile to file")
```

In the `main()`:-
```
flag.Parse()
	if *cpuProfile != "" {
		f, er := os.Create(*cpuProfile)
		if er != nil {
			fmt.Println("Error in creating file for writing cpu profile: ", er)
			return
		}
		defer f.Close()

		if e := pprof.StartCPUProfile(f); e != nil {
			fmt.Println("Error in starting CPU profile: ", e)
			return
		}
		defer pprof.StopCPUProfile()

	}

	prime(100000)

	if *memProfile != "" {
		f, er := os.Create(*memProfile)
		if er != nil {
			fmt.Println("Error in creating file for writing memory profile to: ", er)
			return
		}
		defer f.Close()
		runtime.GC()
		if e := pprof.WriteHeapProfile(f); e != nil {
			fmt.Println("Error in writing memory profile: ", e)
			return
		}
	}
	```

4. Compiled the program using `go install`.

5. Ran the following command to generate the profile :-
For CPU profile => `profiling -cpu <cpu-profile-file>`
For memory profile => `profiling -memory <memory-profile-file>`

6. Now to analyze the generated profile:-
`go tool pprof <profile-file>`
-------------------------------------------------------------------------------------------------------------------------------------


## HTTP Profiling
Just like runtime profiling, HTTP profiling can elucidate the working of the application by just running a web server in the application.

### How I did it

1. Import `_net/http/pprof`
2. Run a webserver if not already running.
```
go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```
3. Now different profiles will be available at the following endpoints:-
*http://localhost:6060/debug/pprof/goroutine
*http://localhost:6060/debug/pprof/heap
*http://localhost:6060/debug/pprof/threadcreate
*http://localhost:6060/debug/pprof/block
*http://localhost:6060/debug/pprof/mutex

