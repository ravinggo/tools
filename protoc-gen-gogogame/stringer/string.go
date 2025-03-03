package stringer

import (
	"fmt"

	"github.com/gogo/protobuf/gogoproto"
	descriptorpb "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type stringer struct {
	*generator.Generator
	generator.PluginImports
	atleastOne bool
	localName  string
	jsonAny    generator.Single
	jwriter    generator.Single
	buffer     generator.Single
	unsafePkg  generator.Single
}

func NewStringer() *stringer {
	return &stringer{}
}

func (p *stringer) Name() string {
	return "newStringer"
}

func (p *stringer) Init(g *generator.Generator) {
	p.Generator = g

}

var keyStr = `w.RawByte('"')
w.RawString("%s")
w.RawByte('"')
w.RawByte(':')
`
var tempComma = ""

var writeF64Str = func() string {
	return `if math.Abs(float64(m.%s)) > 0.000001 {
` + tempComma + keyStr + `w.Float64(float64(m.%s))
needWriteComma = true
}
`
}
var writeF64StrValue = `w.Float64(float64(v))
`

var writeBoolStr = func() string {
	return `if m.%s {
` + tempComma + keyStr + `w.Bool(m.%s)
needWriteComma = true
}
`
}
var writeBoolStrValue = `w.Bool(v)
`

var writeI64Str = func() string {
	return `if m.%s != 0 {
` + tempComma + keyStr + `w.Int64(int64(m.%s))
needWriteComma = true
}
`
}
var writeI64StrValue = `w.Int64(int64(v))
`
var writeI64StrKey = `w.RawByte('"')
w.Int64(int64(k))
w.RawByte('"')
w.RawByte(':')
`

var writeUI64Str = func() string {
	return `if m.%s != 0 {
` + tempComma + keyStr + `w.Uint64(uint64(m.%s))
needWriteComma = true
}
`
}
var writeUI64StrValue = `w.Uint64(uint64(v))
`
var writeUI64StrKey = `w.RawByte('"')
w.Uint64(uint64(k))
w.RawByte('"')
w.RawByte(':')
`

var writeStringStr = func() string {
	return `if m.%s != "" {
` + tempComma + keyStr + `w.String(m.%s)
needWriteComma = true
}
`
}
var writeStringStrValue = `w.String(v)
`

var writeStringStrKey = `w.RawByte('"')
w.RawString(k)
w.RawByte('"')
w.RawByte(':')
`

var writeBytesStr = func() string {
	return `if len(m.%s) > 0 {
` + tempComma + keyStr + `w.Base64Bytes(m.%s)
needWriteComma = true
}
`
}
var writeBytesStrValue = `w.Base64Bytes(v)
`

var arrayStart = `w.RawByte('[')
`
var arrayEnd = `w.RawByte(']')
`

var objectStart = `w.RawByte('{')
`
var objectEnd = `w.RawByte('}')
`

var comma = `
w.RawByte(',')
`

func (p *stringer) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	proto3 := gogoproto.IsProto3(file.FileDescriptorProto)
	if !proto3 {
		return
	}

	p.jsonAny = p.NewImport("github.com/ravinggo/tools/jsonany")
	p.jwriter = p.NewImport("github.com/mailru/easyjson/jwriter")
	p.buffer = p.NewImport("github.com/mailru/easyjson/buffer")
	p.unsafePkg = p.NewImport("unsafe")

	p.atleastOne = false

	p.localName = generator.FileName(file)

	// fmtPkg := p.NewImport("fmt")
	// stringsPkg := p.NewImport("strings")
	// reflectPkg := p.NewImport("reflect")
	// sortKeysPkg := p.NewImport("github.com/gogo/protobuf/sortkeys")
	// protoPkg := p.NewImport("github.com/gogo/protobuf/proto")

	p.P(`var _ = `, p.jsonAny.Use(), ".Any{}")
	for _, message := range file.Messages() {
		if message.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.atleastOne = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (m *`, ccTypeName, `) JsonBytes(w*`, p.jwriter.Use(), `.Writer) {`)
		p.In()
		p.P(`if m == nil {`)
		p.In()
		p.P(`w.RawString("null") ;return `)
		p.Out()
		p.P(
			`}
`,
		)
		// if strings.Contains(message.GetName(), "MovePush") {
		//	fmt.Fprintln(os.Stderr, "1")
		// }
		jsonStr := objectStart
		jsonStr += p.doFields(message.Field)
		jsonStr += objectEnd

		p.P(jsonStr)
		p.P(
			`}
		
`,
		)
	}
	for _, message := range file.Messages() {
		if message.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(fmt.Sprintf(jsonStr, ccTypeName, p.jwriter.Use(), p.buffer.Use(), ccTypeName, p.unsafePkg.Use(), ccTypeName))
	}
}

var jsonStr = `
func (m* %s)MarshalJSON() ([]byte, error) {
	w:=%s.Writer{Buffer:%s.Buffer{Buf:make([]byte,0,2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func(m* %s) String()string  {
	d,_:=m.MarshalJSON()
	return *(*string)(%s.Pointer(&d))
}
func(m* %s) GoString()string  {
	return m.String()
}
`

func (p *stringer) doFields(fields []*descriptorpb.FieldDescriptorProto) string {
	jsonStr := ""
	if len(fields) > 0 {
		jsonStr += `needWriteComma := false
`
	}
	for i, field := range fields {
		jsonStr += p.doOneField(field, i == 0)
	}
	if len(fields) > 0 {
		jsonStr += `_ = needWriteComma
`
	}
	return jsonStr
}

func (p *stringer) doOneKeyField(field *descriptorpb.FieldDescriptorProto) string {
	jsonStr := ""
	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_INT64, descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32, descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		jsonStr += writeI64StrKey
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32, descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32, descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		jsonStr += writeUI64StrKey
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		jsonStr += writeStringStrKey
	default:
		p.Fail("unknown map key field type", field.GetType().String())
	}
	return jsonStr
}

func (p *stringer) doOneValueField(field *descriptorpb.FieldDescriptorProto) string {
	jsonStr := ""
	switch field.GetType() {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		jsonStr += writeF64StrValue

	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		jsonStr += writeBoolStrValue

	case descriptorpb.FieldDescriptorProto_TYPE_INT64, descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32, descriptorpb.FieldDescriptorProto_TYPE_SINT64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		jsonStr += writeI64StrValue

	case descriptorpb.FieldDescriptorProto_TYPE_UINT32, descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32, descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		jsonStr += writeUI64StrValue

	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		jsonStr += writeStringStrValue

	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		jsonStr += writeBytesStrValue

	default:
		if (field.IsMessage() && !gogoproto.IsCustomType(field)) || p.IsGroup(field) {
			tn := field.GetTypeName()
			if tn == ".google.protobuf.Any" {
				jsonStr += fmt.Sprintf(
					`(*%s.Any)(v).JsonBytes(w)
`, p.jsonAny.Use(),
				)
			} else {
				jsonStr += `v.JsonBytes(w)
`
			}
		}
	}
	return jsonStr
}

func (p *stringer) doOneField(field *descriptorpb.FieldDescriptorProto, isFirst bool) string {
	if !isFirst {
		tempComma = `if needWriteComma {
			w.RawByte(',')
		}
`
	} else {
		tempComma = ""
	}

	jsonStr := ""
	fieldName := generator.CamelCase(field.GetName())
	if field.IsRepeated() {
		if p.IsMap(field) {
			gmd := p.GoMapType(nil, field)
			jsonStr += tempComma
			jsonStr += fmt.Sprintf(keyStr, field.GetName())

			jsonStr += fmt.Sprintf(`if m.%s==nil {`, fieldName)
			jsonStr += fmt.Sprintf(
				`w.RawString("null")
}else if len(m.%s)==0{
	w.RawString("{}")				
} else {
	w.RawByte('{')			
`, fieldName,
			)
			temp := "ml" + generator.CamelCase(field.GetName())
			jsonStr += fmt.Sprintf("%s := len(m.%s)\n", temp, generator.CamelCase(field.GetName()))
			jsonStr += fmt.Sprintf(
				`for k,v:=range m.%s {
`, generator.CamelCase(field.GetName()),
			)
			jsonStr += p.doOneKeyField(gmd.KeyField)
			jsonStr += p.doOneValueField(gmd.ValueField)
			jsonStr += temp + "--\n"
			jsonStr += fmt.Sprintf(`if %s!=0{`, temp)
			jsonStr += comma
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += objectEnd
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += `needWriteComma = true
`
		} else {
			jsonStr += tempComma
			jsonStr += fmt.Sprintf(keyStr, field.GetName())
			jsonStr += fmt.Sprintf(`if m.%s==nil {`, fieldName)
			jsonStr += fmt.Sprintf(
				`w.RawString("null")
}else if len(m.%s)==0{
	w.RawString("[]")
} else {
`, fieldName,
			)
			jsonStr += arrayStart
			jsonStr += fmt.Sprintf(
				`for i,v:=range m.%s {
`, generator.CamelCase(field.GetName()),
			)
			jsonStr += p.doOneValueField(field)
			jsonStr += fmt.Sprintf(`if i != len(m.%s)-1 {`, generator.CamelCase(field.GetName()))
			jsonStr += comma
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += arrayEnd
			jsonStr += fmt.Sprintf(
				`}
`,
			)
			jsonStr += `needWriteComma = true
`
		}
	} else {
		// jsonStr += tempComma
		switch field.GetType() {
		case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE,
			descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
			jsonStr += fmt.Sprintf(writeF64Str(), fieldName, field.GetName(), fieldName)

		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
			jsonStr += fmt.Sprintf(writeBoolStr(), fieldName, field.GetName(), fieldName)

		case descriptorpb.FieldDescriptorProto_TYPE_INT64, descriptorpb.FieldDescriptorProto_TYPE_INT32,
			descriptorpb.FieldDescriptorProto_TYPE_SINT32, descriptorpb.FieldDescriptorProto_TYPE_SINT64,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
			descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			jsonStr += fmt.Sprintf(writeI64Str(), fieldName, field.GetName(), fieldName)

		case descriptorpb.FieldDescriptorProto_TYPE_UINT32, descriptorpb.FieldDescriptorProto_TYPE_UINT64,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED32, descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
			jsonStr += fmt.Sprintf(writeUI64Str(), fieldName, field.GetName(), fieldName)

		case descriptorpb.FieldDescriptorProto_TYPE_STRING:
			jsonStr += fmt.Sprintf(writeStringStr(), fieldName, field.GetName(), fieldName)

		case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			jsonStr += fmt.Sprintf(writeBytesStr(), fieldName, field.GetName(), fieldName)

		default:
			if (field.IsMessage() && !gogoproto.IsCustomType(field)) || p.IsGroup(field) {

				jsonStr += fmt.Sprintf(tempComma+keyStr, field.GetName())
				tn := field.GetTypeName()
				if tn == ".google.protobuf.Any" {
					jsonStr += fmt.Sprintf(
						`(*%s.Any)(m.%s).JsonBytes(w)
`, p.jsonAny.Use(), generator.CamelCase(field.GetName()),
					)
				} else {
					jsonStr += fmt.Sprintf(
						`m.%s.JsonBytes(w)
`, generator.CamelCase(field.GetName()),
					)
				}
				jsonStr += `needWriteComma = true
`
			}
		}
	}

	return jsonStr
}
