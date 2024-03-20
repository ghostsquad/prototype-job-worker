<h1 align="center">
  <br>
  <a href="http://github.com/go-fn/fn"><img src="./docs/assets/logo.png" alt="github.com/ghostsquad/prototype-job-worker" width="200px" /></a>
  <br>
  Prototype Job Worker
  <br>
</h1>

<p align="center">
  <a href="#introduction">Introduction</a> •
  <a href="#getting-started">Getting Started</a> •
  <a href="#design">Design</a>
</p>

## Introduction

This project is an implementation of this [Teleport Challenge](https://github.com/gravitational/careers/blob/main/challenges/systems/challenge-1.md)

## Getting Started

### CLI Usage

starting a job returns the job id, which is needed for other commands

```shell
$ pjw start 'echo "hello world"'
1
```

you can list jobs, which returns just the job id's (for now). Some simple bash could be used to get more information for each.

```shell
$ pjw list
1
2
3
```

get a specific job (and it's status).

_Note: `end` and `exitCode` may be `null` if the job is still running_

```shell
$ pjw get 1
{
  "id": "1",
  "command": "echo \"hello world\"",
  "start": "1990-12-31T23:59:60Z",
  "end": "1990-12-31T23:59:60Z",
  "exitCode": 0,
}
```

stop a job that's still running will return once the process has terminated

```shell
$ pjw stop 1
{
  "id": "1",
  "command": "echo \"hello world\"",
  "start": "1990-12-31T23:59:60Z",
  "end": "1990-12-31T23:59:60Z",
  "exitCode": 0,
}
```

you can stream the job output. This always returns full results, starting from the beginning of the output, regardless of when you run the command

```shell
$ pjw logs 1
hello world
```

### Mac M1 Host

If you are running from a Mac M1/M2 that is using an ARM processor, you'll need virtualization that supports ARM such as [multipass](https://multipass.run/docs/mac-tutorial).

```shell
brew install multipass
multipass launch 22.04
```

** _TODO More Steps Here_

## Design

### Scope

#### In-Scope

##### Library

- Worker library with methods to start/stop/query status and get the output of a job.
- Library should be able to stream the output of a running job.
  - Output should be from start of process execution.
  - Multiple concurrent clients should be supported.

##### API

- GRPC API to start/stop/get status/stream output of a running process.
- Use mTLS authentication and verify client certificate. Set up strong set of cipher suites for TLS and good crypto setup for certificates. Do not use any other authentication protocols on top of mTLS.
- Use a simple authorization scheme.
- Add resource control for CPU, Memory and Disk IO per job using cgroups.
- Add resource isolation for using PID, mount, and networking namespaces.

##### CLI

- CLI should be able to connect to worker service and start, stop, get status, and output of a job.

#### Out-Of-Scope

- HTTPS API to start/stop/get status of a running process.
- Server/Client on separate machines
- User-provided cgroup resources / isolation parameters
- Persistence - Process logs are stored in memory, not persisted or read from disk
- Scalability - Not designed to handle any sort of scale, just a simple demo
- Availability - Not highly available
- Cross-Platform - Only works on Linux AMD64
- CI/CD pipeline

### Server

All job information is stored in memory. There's no persistence. If the server/application restarts, all job information and history is lost.

A Job looks like this:

```go
type Job struct {
  // for simplicity, Id is an int64, however, a GUID, UUID or other type might be better for a production-ready application
  Id int64

  // This is the `clientName` part of the client certificate CN
  // Used to distinguish who created this job
  ClientName string
  Command string
  StartTime date
  // EndTime is a date pointer to allow it to be nil
  EndTime *date
  // ExitCode is a int32 pointer to allow it to be nil
  ExitCode *int32

  Logs []string
  mu sync.Mutex
}
```

Each job needs a mutex so that the goroutine that can is running the job can periodically update it with stored logs. When a client connects (via the `logs` command), then can get the full output of this slice, followed by either periodic updates (when the slice length changes), or subscribe to the job runner that's actively running the job.

Job starts need to be handled by an object that is thread safe, for consistency with creating jobIds. Lowering coordination requirements here may be provide a performance improvement, but that is currently not in scope.

```go

type JobRunner struct {
  jobs map[int64]*Job
  jobsByClient map[string]*Job
  mu sync.Mutex
}

func (j *JobRunner) Start(command, client string) *JobRunner {
  j.mu.Lock()
  defer j.mu.Unlock()

  // do stuff
}
```

### CLI

```text
pjw list          Provides a list of jobs and their status
pjw get ID        Provides a detailed status of a specific job
pjw start CMD     Create and start a job; returns the job id
pjw stop ID       Stop a job. Returns success if the job is already stopped/done
pjw logs ID       Stream the logs of a job, always starting at the beginning
```

### Security

#### Authentication

Authentication is based on mTLS and is also completely manual to setup. The server needs a Certificate Authority (CA) and public and private server certificates. It does not generate any of this on it's own. Securely storing and presenting the CA and private server certs is the responsibility of the application operator.

Examples for creating certs:

```shell
DAYS=20
SUBJECT="/C=US/ST=Washington/L=Seattle/O=GhostSquad/OU=Prototype Job Worker/CN=localhost"

# CA
openssl req -newkey rsa:4096 \
  -new -nodes -x509 \
  -days "${DAYS}" \
  -out ca.crt \
  -keyout ca.key \
  -subj "${SUBJECT}"

# Server
openssl genrsa -out server.key 4096
openssl req -new -key server.key -days "${DAYS}" -out server.csr \
    -subj "${SUBJECT}"

openssl x509  -req -in server.csr \
    -extfile <(printf "subjectAltName=DNS:localhost") \
    -CA ca.crt -CAkey ca.key  \
    -days "${DAYS}" -sha256 -CAcreateserial \
    -out server.crt

# Client
CLIENT=john

openssl genrsa -out "${CLIENT}.key" 4096

openssl req -new -key "${CLIENT}.key" -days "${DAYS}" -out "${CLIENT}.csr" \
    -subj "${SUBJECT}/clientName=${CLIENT}"

openssl x509  -req -in "${CLIENT}.csr" \
    -extfile <(printf "subjectAltName=DNS:localhost") \
    -CA ca.crt -CAkey ca.key -out "${CLIENT}.crt" -days "${DAYS}" -sha256 -CAcreateserial
```

Out of scope: Improvements to this process including easy to use script/binary for generating client certificates. Generating client certificates from application itself, and having the user download it. Using external certificate management software such as Vault.

All actions from the CLI require providing the client cert as an option (e.g. `pjw --client-crt /path/to/crt --client-key /path/to/key`)

Out of scope: Improvement to authorizing every request would be a "login" command, which generates a short-lived token, and that token is what is validated.

#### Authorization Schema

##### Actions

- list
- get
- start
- stop
- logs

All actions are scoped to the current user. As an example, that means `list` will only ever return jobs that were started by that user.

Out of scope: Getting jobs across all users (admin-like)

All authorization configuration is statically hard-coded, no configuration file or "ease of use" patterns have been implemented.
