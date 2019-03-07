# Profiling a Go application
## Runtime Profiling
Profiling a program can tell you a lot about the memory and CPU usage of different functions in the program. This can greatly help you find the bottlenecks and optimize the program by hitting the bulls eye. 
### How I did it
1. Import the package github.com/pkg/profile
2. Install graphviz [graphviz](https://www.graphviz.org/)
3. Now, I've done profiling for CPU and memory one by one. First I profiled a sample Golang code (prime number calculation) for CPU usage. For that first I added this line of code at the top of the main:
`defer profile.Start().Stop()`
4. Compiled the program using `go install`.
5. Ran the following command to generate the profile in PDF format:-
`go tool pprof --pdf ~/go/bin/yourbinary /var/path/to/cpu.pprof > cpuProfile.pdf`
A PDF file with the name cpuProfile.pdf will be generated which will contain the CPU usage of different parts of the program.
6. Similary for memory profiling follow `3 --> 4 --> 5` but in step 3 add `defer profile.Start(profile.MemProfile).Stop()` instead of `defer profile.Start().Stop()`
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

