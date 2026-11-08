# Port Scanner - Micro-Tool 2

This is Micro-Tool 2. I built it to understand the core mechanics of network ports and how to scan them for open or closed states using Go's native libraries.

## Usage

`
./portscanner <IP>
`
## What I Learned
While building this Micro Tool, I understood several concurrent programming and networking mechanics:

- Implementing a worker pool architecture to limit simultaneous connections and prevent OS resource exhaustion.

- Using unbuffered channels to safely synchronize data between the main program and multiple background workers.

- Decoupling the job-feeding loop into a background GoRoutine to prevent circular deadlocks.

- Utilizing connection timeouts to handle stealth firewalls and silently dropped packets without freezing the scanner.
