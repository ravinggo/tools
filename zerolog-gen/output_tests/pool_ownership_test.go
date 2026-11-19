package output_tests

import (
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/mailru/easyjson/jwriter"
	"github.com/ravinggo/objectpool"
	"github.com/ravinggo/zerolog"

	"github.com/ravinggo/tools/zerolog-gen/output_tests/output"
)

// easyjson 的 chunk 依次为 128/256/512/1024/2048，写满前四块后取到的就是 2048 的池化 chunk。
const bytesToReachEasyjson2048Chunk = 128 + 256 + 512 + 1024 + 1

// JSON 超过 2048 字节时，objectpool 的缓冲不能再被 easyjson BuildBytes 放进它自己的池：
// 否则同一块内存被两个池同时持有，easyjson 写入的 JSON 会覆盖 objectpool 记在头 4 字节的容量，
// 下次 GetBytes 返回伪造容量的切片，append 越界写坏堆。
func TestMarshalZerologObjectLargeJSONDoesNotSharePooledBuffer(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	domain := &output.Domain{A: 1, C: []string{strings.Repeat("x", 4096)}}
	domain.MarshalZerologObject(zerolog.Dict())

	var w jwriter.Writer
	for i := 0; i < bytesToReachEasyjson2048Chunk; i++ {
		w.RawByte('x')
	}
	easyjsonChunk := w.Buffer.Buf

	pooled := objectpool.GetBytes(2048)
	defer objectpool.PutBytes(pooled)

	if cap(pooled) != 2048 {
		t.Fatalf("objectpool.GetBytes(2048) cap = %d (0x%x)，容量头被覆盖", cap(pooled), cap(pooled))
	}
	if unsafe.SliceData(pooled) == unsafe.SliceData(easyjsonChunk) {
		t.Fatal("同一块内存同时被 objectpool 与 easyjson 池持有")
	}
}

func TestMarshalZerologObjectLargeJSONKeepsFullPayload(t *testing.T) {
	payload := strings.Repeat("y", 4096)
	domain := &output.Domain{A: 7, C: []string{payload}}

	var sb strings.Builder
	logger := zerolog.New(&sb)
	logger.Info().Object("domain", domain).Send()

	if !strings.Contains(sb.String(), `"C":["`+payload+`"]`) {
		t.Fatalf("日志缺少完整 payload: len=%d tail=%q", sb.Len(), sb.String()[sb.Len()-80:])
	}
}
