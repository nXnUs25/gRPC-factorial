
# gRPC-factorial
GoLang Test Task

# Usage

To run the application after building binaries open separate window in terminal and execute the command `./calculater_server`

```shell
❯ ./calculater_server
2022/04/26 14:04:07 gRPCServerFactorial: Starting... GRPC client on port 50001
2022/04/26 14:04:07 gRPCServerFactorial: Server started at port [[::]:50001]
```

this will execute the server side application on port **5100** by default unless env variable is specified/exported

```shell
❯ export GRPC_PORT=50001
❯ echo $GRPC_PORT
50001
```

then open second window and run the client application by executing the following command `./factorial_client` 
it will calculate the factorial number as example

```shell
❯ ./factorial_client -p 50001
2022/04/26 14:11:18 FactorialReciver: Starting... GRPC Client to talk with GRPC Factorial Server on port 50001
The factorial of [5] is [120]
The factorial of [20] is [2432902008176640000]
The factorial of [1000] is [4.023872600770937735365106377186724593273707471886710737496506544e+2567]
The factorial of [10000000] is [e^151180956.50958320419886149466037750244140625]
The factorial of [999999999999999] is [e^33538776394910650.537109375]
```

To find out more about the operations use the `-h|-help` option.

```shell
❯ ./factorial_client -h
  -h
        Prints usage information.
  -help
        Prints usage information.
  -i value
        Numbers to calculate 1,2,3,4,... . (default: 5,20,1000,10000000,999999999999999)
  -integers value
        Numbers to calculate 1,2,3,4,... . (default: 5,20,1000,10000000,999999999999999)
  -p int
        Port to connect to Factorial Server. (default 5100)
  -port int
        Port to connect to Factorial Server. (default 5100)
```

## Docker image

https://hub.docker.com/repository/docker/nonus25/grpc-factorial-server

`nonus25/grpc-factorial-server`

## Build

To build the project please specify in the terminal `make` in the root of the project

```shell
❯ make
go build factorial/*.go
go build -o calculater_server rpc_factorial/server/*.go 
go build cli/*.go
go build -o factorial_client rpc_factorial/client/*.go
./calculater_server
2022/04/26 13:45:39 gRPCServerFactorial: Starting... GRPC client on port 50001
2022/04/26 13:45:39 gRPCServerFactorial: Server started at port [[::]:50001]
```

### Build Client

Run `make client` to build a client binary.

```shell
❯ make client
go build cli/*.go
go build -o factorial_client rpc_factorial/client/*.go
```

### Build Server

```shell
❯ make server
go build factorial/*.go
go build -o calculater_server rpc_factorial/server/*.go 
```

## Test

To run the unit tests execute following command/commands:

#### via `make` 

```shell
❯ make test
go test ./cli/ ./factorial/
ok      github.com/nXnUs25/gRPC-factorial/cli   (cached)
ok      github.com/nXnUs25/gRPC-factorial/factorial     (cached)
```

#### via `go test`

```shell
❯ go test ./cli/ ./factorial/
ok      github.com/nXnUs25/gRPC-factorial/cli   (cached)
ok      github.com/nXnUs25/gRPC-factorial/factorial     (cached)
```

#### with verbose output

```shell
❯ go test ./cli/ ./factorial/ -v
=== RUN   Test_integers_String
=== RUN   Test_integers_String/ToString_values:_1_2_3
    cli_test.go:32: [PASS] integers.String() = 1 2 3, want 1 2 3
--- PASS: Test_integers_String (0.00s)
    --- PASS: Test_integers_String/ToString_values:_1_2_3 (0.00s)
=== RUN   Test_integers_Set
=== RUN   Test_integers_Set/SetSrings_Used_Twice
    cli_test.go:100: [PASS] ntegers.Set() error = Cannot use [-i|--integers] flag more than once, wantErr true
=== RUN   Test_integers_Set/SetSrings()_Used_once
    cli_test.go:100: [PASS] ntegers.Set() error = <nil>, wantErr false
=== RUN   Test_integers_Set/SetSrings()_Used_once,_with_negative_int
    cli_test.go:100: [PASS] ntegers.Set() error = Accepting only positive integer numbers, wantErr true
=== RUN   Test_integers_Set/SetSrings()_Used_once,_with_invalid_int
    cli_test.go:100: [PASS] ntegers.Set() error = strconv.Atoi: parsing "adnf": invalid syntax, wantErr true
--- PASS: Test_integers_Set (0.00s)
    --- PASS: Test_integers_Set/SetSrings_Used_Twice (0.00s)
    --- PASS: Test_integers_Set/SetSrings()_Used_once (0.00s)
    --- PASS: Test_integers_Set/SetSrings()_Used_once,_with_negative_int (0.00s)
    --- PASS: Test_integers_Set/SetSrings()_Used_once,_with_invalid_int (0.00s)
=== RUN   Test_isNotValidInt
    cli_test.go:171: [PASS] Input OK, Input is [0]: got [false], want [false] AND got err [false], want err [false]
    cli_test.go:171: [PASS] Input OK, Input is [a]: got [true], want [true] AND got err [true], want err [true]
    cli_test.go:171: [PASS] Input OK, Input is [-342]: got [true], want [true] AND got err [true], want err [true]
    cli_test.go:171: [PASS] Input OK, Input is [3.5]: got [true], want [true] AND got err [true], want err [true]
    cli_test.go:171: [PASS] Input OK, Input is [-0.1]: got [true], want [true] AND got err [true], want err [true]
    cli_test.go:171: [PASS] Input OK, Input is [3455245]: got [false], want [false] AND got err [false], want err [false]
    cli_test.go:171: [PASS] Input OK, Input is [0.00001]: got [true], want [true] AND got err [true], want err [true]
--- PASS: Test_isNotValidInt (0.00s)
PASS
ok      github.com/nXnUs25/gRPC-factorial/cli   (cached)
=== RUN   TestMakeCalculate
    calculate_test.go:37: [PASS] Initialized Types are as expected got: [*factorial.Number] and want it [*factorial.Number]
    calculate_test.go:37: [PASS] Initialized Types are as expected got: [*factorial.BigNumber] and want it [*factorial.BigNumber]
    calculate_test.go:37: [PASS] Initialized Types are as expected got: [*factorial.HugeNumber] and want it [*factorial.HugeNumber]
--- PASS: TestMakeCalculate (0.00s)
=== RUN   TestCalculate
    calculate_test.go:114: [PASS] Values calculated [120] and want it [120] are same.
    calculate_test.go:114: [PASS] Values calculated [2.84625967962859768163906785941e+35659] and want it [2.84625967962859768163906785941e+35659] are same.
    calculate_test.go:114: [PASS] Values calculated [e^106663518720] and want it [e^106663518720] are same.
--- PASS: TestCalculate (0.01s)
=== RUN   TestFactorialApproximationNlnN
    factApproxNlnN_test.go:54: [PASS] Passed with tolerance [+1e-30] for values [got: 106663518720 - want: 1.066635187e+11]
    factApproxNlnN_test.go:54: [PASS] Passed with tolerance [+1e-30] for values [got: 82103.4036865234375 - want: 82103.40369]
    factApproxNlnN_test.go:54: [PASS] Passed with tolerance [+1e-30] for values [got: 106663518720 - want: 1.066635187e+11]
    factApproxNlnN_test.go:54: [PASS] Passed with tolerance [+1e-30] for values [got: 289336061853696 - want: 2.893360619e+14]
--- PASS: TestFactorialApproximationNlnN (0.00s)
=== RUN   TestFactorialBigMul
    factBigMul_test.go:53: [PASS] Passed with tolerance [+1e-30] for values [got: 120 - want: 120]
    factBigMul_test.go:53: [PASS] Passed with tolerance [+1e-30] for values [got: 4.02387259951360593775022331467e+2567 - want: 4.02387260077093773543702433923e+2567]
    factBigMul_test.go:53: [PASS] Passed with tolerance [+1e-30] for values [got: 2.84625967962859768163906785941e+35659 - want: 2.846259680917054519e+35659]
    factBigMul_test.go:53: [PASS] Passed with tolerance [+1e-30] for values [got: 8.26393169070469505888138658331e+5565708 - want: 8.26393168833124006237664610317e+5565708]
--- PASS: TestFactorialBigMul (48.19s)
=== RUN   TestFactorialIteraion
    factIteration_test.go:76: [PASS] Passed with tolerance [+1e-30] for values [got: 2.65252859812191068217601719009e+32 - want: 2.65252859812191068217601719009e+32]
    factIteration_test.go:76: [PASS] Passed with tolerance [+1e-30] for values [got: 5.71338395644585538318709988603e+262 - want: 5.71338395644580014701987431176e+262]
    factIteration_test.go:76: [PASS] Passed with tolerance [+1e-30] for values [got: 7.25741561530799778760021609319e+306 - want: 7.2574156153078992429849328164e+306]
    factIteration_test.go:76: [PASS] Passed with tolerance [+1e-30] for values [got: 7.15694570462638505039051537884e+118 - want: 7.15694570462638594628841224996e+118]
    factIteration_test.go:76: [PASS] Passed with tolerance [+1e-30] for values [got: 120 - want: 120]
--- PASS: TestFactorialIteraion (0.00s)
PASS
ok      github.com/nXnUs25/gRPC-factorial/factorial     (cached)
```


## Create Docker image

```shell
❯ docker build -t nonus25/grpc-factorial-server .
[+] Building 16.9s (18/18) FINISHED                                                                                                    
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => => transferring dockerfile: 37B                                                                                               0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => => transferring context: 2B                                                                                                   0.0s
 => resolve image config for docker.io/docker/dockerfile:1                                                                        1.7s
 => [auth] docker/dockerfile:pull token for registry-1.docker.io                                                                  0.0s
 => CACHED docker-image://docker.io/docker/dockerfile:1@sha256:91f386bc3ae6cd5585fbd02f811e295b4a7020c23c7691d686830bf6233e91ad   0.0s
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => [internal] load metadata for docker.io/library/golang:1.16-alpine                                                             1.0s
 => [auth] library/golang:pull token for registry-1.docker.io                                                                     0.0s
 => [internal] load build context                                                                                                 1.0s
 => => transferring context: 23.83MB                                                                                              1.0s
 => [1/7] FROM docker.io/library/golang:1.16-alpine@sha256:5616dca835fa90ef13a843824ba58394dad356b7d56198fb7c93cbe76d7d67fe       0.0s
 => CACHED [2/7] WORKDIR /factorial                                                                                               0.0s
 => CACHED [3/7] COPY go.mod ./                                                                                                   0.0s
 => CACHED [4/7] COPY go.sum ./                                                                                                   0.0s
 => CACHED [5/7] RUN go mod download                                                                                              0.0s
 => [6/7] COPY . ./                                                                                                               0.2s
 => [7/7] RUN go build -o calculater_server rpc_factorial/server/*.go                                                            12.3s
 => exporting to image                                                                                                            0.4s
 => => exporting layers                                                                                                           0.4s
 => => writing image sha256:ab6ef5ab1b5c2022e086c65eb732cac04a9e9ecf493d382d843ad686b6a6e8b4                                      0.0s
 => => naming to docker.io/nonus25/grpc-factorial-server                                                                          0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
```


