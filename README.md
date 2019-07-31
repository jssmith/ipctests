# Unix IPC Testing

Testing the performance of Unix IPC.
Adapted from [blog post by Eli Bendersky](https://eli.thegreenplace.net/2019/unix-domain-sockets-in-go/).

## Running

First build
```
go build ./bin/server
go build ./bin/client
```

Run the client and server programs separately
```
./server --unixdomain --msgsize 10000000 &
./client --unixdomain --msgsize 10000000
```

## Some test results

### IPC

Unix Domain Sockets on an AWS c5n.large instance.

**128 byte messages**
```
50000 pingpongs took 790975853 ns
avg. rt latency 15819 ns
avg. thoughput 0.007536 GB/sec
```

**10,000,000 byte messages**
```
50000 pingpongs took 37768188364 ns
avg. rt latency 755363 ns
avg. thoughput 12.329458 GB/sec
```

### Network
Running over the network with a pair of c5n.large instances (placement group locality).

**128 byte messages**
```
50000 pingpongs took 3120816262 ns
avg. rt latency 62416 ns
avg. thoughput 0.001910 GB/sec
```

**10,000,000 byte messages**
```
1000 pingpongs took 8487414031 ns
avg. rt latency 8487414 ns
avg. thoughput 1.097298 GB/sec
```