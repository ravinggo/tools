//go:generate go run github.com/ravinggo/tools/deepcopy-gen --use-pool=0 --output-file zz_generated.deepcopy.go --go-header-file=../boilerplate.go.txt github.com/ravinggo/tools/keepalive-gen/output_tests_no_ka/output

//go:generate go run github.com/ravinggo/tools/keepalive-gen --use-pool=0 --output-file zz_generated.keepalive.go --go-header-file=../boilerplate.go.txt github.com/ravinggo/tools/keepalive-gen/output_tests_no_ka/output

package output_tests_no_ka
