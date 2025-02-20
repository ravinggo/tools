# deepcopy-gen
## install
```shell
go install github.com/ravinggo/tools/deepcopy-gen@latest
```

### generate package
```go
// +k8s:deepcopy-gen=package
```
### Ignore a struct
```go
// +k8s:deepcopy-gen=false
```
### More information in the [output_tests](./output_tests) directory

## command help
```go
 deepcopy-gen --help
Usage of B:\opensource\tools\deepcopy-gen\deepcopy-gen:
      --add_dir_header                   If true, adds the file directory to the header of the log messages
      --alsologtostderr                  log to standard error as well as files (no effect when -logtostderr=true)
      --bounding-dirs strings            Comma-separated list of import paths which bound the types for which deep-copies will be generated.
      --go-header-file string            the path to a file containing boilerplate header text; the string "YEAR" will be replaced with the current 4-digit year
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory (no effect when -logtostderr=true)     
      --log_file string                  If non-empty, use this log file (no effect when -logtostderr=true)
      --log_file_max_size uint           Defines the maximum size a log file can grow to (no effect when -logtostderr=true). Unit is megabytes. If the value is 0, the maximum file size is unlimited. (default 1800)
      --logtostderr                      log to standard error instead of files (default true)
      --one_output                       If true, only write logs to their native severity level (vs also writing to each lower severity level; no effect when -logtostderr=true)
      --output-file string               the name of the file to be generated (default "generated.deepcopy.go")
      --skip_headers                     If true, avoid header prefixes in the log messages
      --skip_log_headers                 If true, avoid headers when opening log files (no effect when -logtostderr=true)       
      --stderrthreshold severity         logs at or above this threshold go to stderr when writing to files and stderr (no effect when -logtostderr=true or -alsologtostderr=true) (default 2)
      --use-pool                         use sync.Pool mode (default true)
  -v, --v Level                          number for the log level verbosity
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
pflag: help requested

```

### example
```go
deepcopy-gen --bounding-dirs=./output_tests/test --output-file zz_generated.deepcopy.go ./output_tests/test
```

## NODE: 
### Struct1 has a reference type(slice,map,pointer) field. However, as the value type of Struct2 field Map1, it will cause additional allocation, which cannot be solved yet.<br>
### Please use Struct3 to use pointers
```
type Struct1 struct {
    Field1 []string //  
    Field2 int
} 

type Struct2 struct {
    Map1 map[string]Struct1
    Field2 int
}

type Struct3 struct {
    Map1 map[string]*Struct1
    Field2 int
}

```
