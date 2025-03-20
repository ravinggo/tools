package google_proto

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type googleProto struct {
	*generator.Generator
	generator.PluginImports

	protoreflect generator.Single
}

func NewGoogleProto() *googleProto {
	return &googleProto{}
}

func (p *googleProto) Init(g *generator.Generator) {
	p.Generator = g
	p.PluginImports = generator.NewPluginImports(g)
	p.protoreflect = p.NewImport("google.golang.org/protobuf/reflect/protoreflect")
}

func (p *googleProto) Name() string { return "googleProto" }

func (p *googleProto) Generate(file *generator.FileDescriptor) {
	proto3 := gogoproto.IsProto3(file.FileDescriptorProto)
	if !proto3 {
		return
	}

	for _, message := range file.Messages() {
		if message.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (m *`, ccTypeName, `) ProtoReflect() `, p.protoreflect.Use(), `.Message {`)
		p.P(`   return nil`)
		p.P(`}`)
	}
	p.P(`var _ `, p.protoreflect.Use(), `.FullName = ""`)
}
