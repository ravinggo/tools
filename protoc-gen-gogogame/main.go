package main

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"

	"github.com/ravinggo/tools/protoc-gen-gogogame/google-proto"
	"github.com/ravinggo/tools/protoc-gen-gogogame/stringer"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)
	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

	vanity.ForEachFieldInFilesExcludingExtensions(vanity.OnlyProto2(files), vanity.TurnOffNullableForNativeTypesWithoutDefaultsOnly)
	vanity.ForEachFile(files, vanity.TurnOffGoUnrecognizedAll)
	vanity.ForEachFile(files, vanity.TurnOffGoUnkeyedAll)
	vanity.ForEachFile(files, vanity.TurnOffGoSizecacheAll)

	vanity.ForEachFile(files, vanity.TurnOffGoEnumStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnEnumStringerAll)

	vanity.ForEachFile(files, vanity.TurnOffGoStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnMessageNameAll)
	generator.RegisterPlugin(google_proto.NewGoogleProto())
	generator.RegisterPlugin(stringer.NewStringer())

	resp := command.Generate(req)
	command.Write(resp)
}
