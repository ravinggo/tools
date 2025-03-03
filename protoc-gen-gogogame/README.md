# protoc-gen-gogogame
Compatible with gogo protobuf of google.golang.org/protobuf/proto.Message by generating ProtoReflect() function</br>
通过生成ProtoReflect() 函数 来兼容google.golang.org/protobuf/proto.Message的gogo protobuf

```go
vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
vanity.ForEachFile(files, vanity.TurnOnSizerAll)
vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

vanity.ForEachFieldInFilesExcludingExtensions(vanity.OnlyProto2(files), vanity.TurnOffNullableForNativeTypesWithoutDefaultsOnly)
vanity.ForEachFile(files, vanity.TurnOffGoUnrecognizedAll)
vanity.ForEachFile(files, vanity.TurnOffGoUnkeyedAll)
vanity.ForEachFile(files, vanity.TurnOffGoSizecacheAll)

vanity.ForEachFile(files, vanity.TurnOffGoEnumStringerAll)
vanity.ForEachFile(files, vanity.TurnOnEnumStringerAll)

vanity.ForEachFile(files, vanity.TurnOnEqualAll)
vanity.ForEachFile(files, vanity.TurnOffGoStringerAll)
vanity.ForEachFile(files, vanity.TurnOnMessageNameAll)
 // generate 'func ProtoReflect() protoreflect.Message'
generator.RegisterPlugin(google_proto.NewGoogleProto())
 // generate :
// 'func String() string'
// 'func MarshalJSON() ([]byte, error)'
generator.RegisterPlugin(stringer.NewStringer())
```

## usage
```shell
protoc -I. -I./proto  --gogogame_out=plugins=googleProto+newStringer:.. ./proto/*.proto
```

## sample

[output_test](./output_test)