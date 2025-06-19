//go:generate go run github.com/mailru/easyjson/easyjson -all ./output

//go:generate go run github.com/ravinggo/tools/zerolog-gen -v 3 --output-file zz_generated.zerolog.go --go-header-file=../boilerplate.go.txt github.com/ravinggo/tools/zerolog-gen/output_tests/output
package output_tests
