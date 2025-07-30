package parser

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/emicklei/proto"
	"github.com/xuri/excelize/v2"

	"github.com/ravinggo/tools/utils"
)

func Walk(srcPath string) ([]string, error) {
	var protos []string
	err := filepath.Walk(
		srcPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info == nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".proto") {
				protos = append(protos, path)
				return nil
			}
			return err
		},
	)
	return protos, err
}

type Parser struct {
	readers     []*proto.Parser
	messages    map[string]*MessageBase
	enums       map[string]*ErrorInfo
	ext         map[string]*proto.Message
	enumsModels map[string][]*ErrorInfo
	Files       []*File
	dirs        []string
	sheetName   string
}

func NewParser(dir string) (*Parser, error) {
	dirs := strings.Split(dir, ",")
	var files []string
	for _, v := range dirs {
		fs, err := Walk(v)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}

	readers := make([]*proto.Parser, 0, len(files))
	for _, v := range files {
		f, err := os.Open(v)
		if err != nil {
			return nil, err
		}
		p := proto.NewParser(f)
		_, fn := filepath.Split(v)
		p.Filename(fn)
		readers = append(readers, p)
	}
	return &Parser{
		readers:     readers,
		messages:    map[string]*MessageBase{},
		enums:       map[string]*ErrorInfo{},
		ext:         map[string]*proto.Message{},
		enumsModels: map[string][]*ErrorInfo{},
		dirs:        dirs,
		sheetName:   "languages",
	}, nil
}

func (this_ *Parser) Parse() error {
	for _, v := range this_.readers {
		definition, err := v.Parse()
		if err != nil {
			return err
		}
		fmt.Println(definition.Filename)
		file := &File{FileName: definition.Filename}
		definition.Accept(file)
		this_.Files = append(this_.Files, file)
	}
	return nil
}

func (this_ *Parser) Check() {
	commonErrorMap := map[string]bool{}
	checkDump := map[string]*ErrorInfo{}
	for _, file := range this_.Files {
		for _, mb := range file.Mbs {
			if oldMb, ok := this_.messages[mb.Msg.Name]; ok {
				Errorf("message '%s'[%s:%d] had defined [%s:%d]", mb.Msg.Name, file.FileName, mb.Msg.Position.Line, oldMb.FileName, oldMb.Msg.Position.Line)
			}
			this_.messages[mb.Msg.Name] = mb
		}
		for _, ei := range file.Eis {
			if oldEi, ok := this_.enums[ei.Enum.Name]; ok {
				Errorf(
					"enum value '%s'[%s:%d] had defined [%s:%d]", ei.Enum.Name, file.FileName, ei.Enum.Position.Line, oldEi.FileName, oldEi.Enum.Position.Line,
				)
			}
			this_.enums[ei.Enum.Name] = ei
			if oldEi, ok := checkDump[ei.TipName]; ok {
				Errorf("enum value '%s'[%s:%d] had defined [%s:%d]", ei.TipName, file.FileName, ei.Enum.Position.Line, oldEi.FileName, oldEi.Enum.Position.Line)
			}
			checkDump[ei.TipName] = ei
			this_.enumsModels[ei.Module] = append(this_.enumsModels[ei.Module], ei)
			if ei.Enum.Parent.(*proto.Enum).Name == "CommonErrorCode" {
				commonErrorMap[ei.Enum.Name] = true
			}
		}
		for k, v := range file.Ext {
			this_.ext[k] = v
		}
	}

	for _, v := range this_.Files {
		v.Check(commonErrorMap)
	}

	for k := range this_.enumsModels {
		em := this_.enumsModels[k]
		sort.Slice(
			em, func(i, j int) bool {
				return em[i].TipName < em[j].TipName
			},
		)
	}
}

func Errorf(format string, a ...interface{}) {
	fmt.Fprintln(os.Stderr, "Parse Failed:", fmt.Sprintf(format, a...))
	os.Exit(-1)
}

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
func (this_ *Parser) OutputErrorCodeTxt(outPath string) {
	b := &bytes.Buffer{}
	keys := make([]string, 0, len(this_.enumsModels))
	for k := range this_.enumsModels {
		keys = append(keys, k)
	}
	sort.Slice(
		keys, func(i, j int) bool {
			return keys[i] < keys[j]
		},
	)
	for i := 0; i < len(keys); i++ {
		vs := this_.enumsModels[keys[i]]
		for _, v := range vs {
			b.WriteString(fmt.Sprintf("%s=%s\n", v.TipName, v.TipDesc))
		}
	}
	err := os.WriteFile(outPath, b.Bytes(), 0666)
	if err != nil {
		Errorf("保存[%s]失败：%v", outPath, err)
	}
}

func (this_ *Parser) OutputErrorCodeExcel(outPath string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			Errorf("close excel file error:%v", err)
		}
	}()

	index, err := f.NewSheet(this_.sheetName)
	if err != nil {
		Errorf("new sheet error:%v", err)
	}
	f.SetActiveSheet(index)
	this_.setValue(f, "A1", "##var")
	this_.setValue(f, "B1", "id")
	this_.setValue(f, "C1", "zhCN")
	this_.setValue(f, "A2", "##type")
	this_.setValue(f, "B2", "string")
	this_.setValue(f, "C2", "string")
	this_.setValue(f, "A3", "##")
	this_.setValue(f, "C3", "ChineseSimplified")
	this_.setValue(f, "A4", "##comment")
	this_.setValue(f, "B4", "key")
	this_.setValue(f, "C4", "(简体中文)40")

	keys := make([]string, 0, len(this_.enumsModels))
	for k := range this_.enumsModels {
		keys = append(keys, k)
	}
	sort.Slice(
		keys, func(i, j int) bool {
			return keys[i] < keys[j]
		},
	)
	indexCell := 4
	for i := 0; i < len(keys); i++ {
		vs := this_.enumsModels[keys[i]]
		for _, v := range vs {
			indexCell++
			indexStr := strconv.Itoa(indexCell)
			this_.setValue(f, "B"+indexStr, v.TipName)
			this_.setValue(f, "A"+indexStr, v.TipDesc)
		}
	}
	err = f.DeleteSheet("Sheet1")
	if err != nil {
		Errorf("delete excel Sheet1 error:%v", err)
	}
	if err := f.SaveAs(outPath); err != nil {
		Errorf("save excel file error:%v", err)
	}
}

func (this_ *Parser) setValue(file *excelize.File, cell string, value any) {
	err := file.SetCellValue(this_.sheetName, cell, value)
	if err != nil {
		Errorf("set value error:%v", err)
	}
}

func (this_ *Parser) OutputErrorCodeGoCode(pkgName, outPath string) {
	m := map[string][]*ErrorInfo{}
	for _, v := range this_.enums {
		n := v.Enum.Parent.(*proto.Enum).Name
		m[n] = append(m[n], v)
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
		ei := m[k]
		sort.Slice(
			ei, func(i, j int) bool {
				return ei[i].Enum.Name < ei[j].Enum.Name
			},
		)
	}
	sort.Strings(keys)
	var stackError string
	for i := 0; i < len(keys); i++ {
		for _, v := range m[keys[i]] {
			stackError += fmt.Sprintf(
				`
func New%s(extra ...string) *berror.ErrMsg {
	e := &berror.ErrMsg{}
	e.ErrCode = %s
	e.ErrMsg = "%s"
	if berror.IsOpenStack() {
		e.WithStackTrace()
	}
	if len(extra) > 0 {
		e.ErrExtraInfo = extra[0]
	}
	return e
}
`, v.Enum.Name, v.ErrorCode, v.TipName,
			)
		}
	}

	str := fmt.Sprintf(
		`// Code generated by gen-proto-error. DO NOT EDIT.
// source: %s
package %s

import (
	berror "github.com/ravinggo/game/common/berror"
)

%s
`, strings.Join(this_.dirs, ","), pkgName, stackError,
	)
	ext := filepath.Ext(outPath)
	if ext != ".go" {
		if !utils.IsDirExists(outPath) {
			err := os.MkdirAll(outPath, os.ModePerm)
			if err != nil {
				Errorf("create dir[%s] failded:%v", outPath, err)
			}
		}
		outPath = filepath.Join(outPath, "error_code.go")
	} else {
		dir, _ := filepath.Split(outPath)
		if !utils.IsDirExists(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				Errorf("create dir[%s] failded:%v", dir, err)
			}
		}
	}

	err := os.WriteFile(outPath, []byte(str), os.ModePerm)
	if err != nil {
		Errorf("save [%s] failed:%v", outPath, err)
	}
}

func HeadToLower(s string) string {
	if s[0] >= 'A' && s[0] <= 'Z' {
		x := &strings.Builder{}
		x.WriteByte(s[0] + 32)
		x.WriteString(s[1:])
		return x.String()
	}
	return s
}
