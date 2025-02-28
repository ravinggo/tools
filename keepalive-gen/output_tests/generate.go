//go:generate go run github.com/ravinggo/tools/deepcopy-gen --output-file zz_generated.deepcopy.go --go-header-file=../boilerplate.go.txt github.com/ravinggo/tools/keepalive-gen/output_tests/output

//go:generate go run github.com/ravinggo/tools/keepalive-gen --output-file zz_generated.keepalive.go --go-header-file=../boilerplate.go.txt github.com/ravinggo/tools/keepalive-gen/output_tests/output

package output_tests
