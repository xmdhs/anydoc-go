// 基准测试：对比 wasm2go 生成代码在开启/未开启 -unsafe 时的转换性能。
//
// 用法（同一机器、同一 wasm、同一 wasm2go 版本下分别跑）：
//
//	go test -bench . -benchtime 10x -count 3 -run 'Benchmark' | tee bench-x.txt
//
// 注意：-run 必须匹配子基准名（BenchmarkConvert/xxx），`-run '^$'` 会把
// 全部子基准过滤掉，只剩父级空循环。
//
// 注意：core 包（wasm2go 生成物）编译参数会影响结果，两个对比版本必须用
// 完全相同的编译参数（见 build.sh 的 -gcflags=all="-N -l"）。
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	anydoc "anydoc-go/lib"
)

// BenchmarkConvert 用仓库样本逐文件转换，单次转换（每轮）从创建 converter
// 开始，模拟真实 CLI 的一次进程内转换成本。
func BenchmarkConvert(b *testing.B) {
	fixtures, err := filepath.Glob("testdata/fixtures/*/*.*")
	if err != nil {
		b.Fatal(err)
	}
	data := map[string][]byte{}
	for _, f := range fixtures {
		if strings.Contains(f, "malformed") || strings.HasSuffix(f, ".pdf") {
			continue // 错误路径/未编译 pdf 不参与基准
		}
		bb, err := os.ReadFile(f)
		if err != nil {
			b.Fatal(err)
		}
		data[f] = bb
		format := ""
		if strings.HasSuffix(f, ".csv") {
			format = "csv"
		}
		b.Run(filepath.Base(f), func(b *testing.B) {
			conv := anydoc.NewConverter()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if _, err := conv.ConvertBytes(data[f], format); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkConvertLargeCSV 用大输入（~2MB CSV）测稳定吞吐。
// 输入小文件时解释器启动开销占比大，大文件更能反映内存访问路径的差异。
func BenchmarkConvertLargeCSV(b *testing.B) {
	var sb strings.Builder
	for i := 0; i < 40000; i++ {
		fmt.Fprintf(&sb, "row%d,colA,colB,%0.3f\n", i, float64(i)/7)
	}
	input := []byte(sb.String())
	conv := anydoc.NewConverter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := conv.ConvertBytes(input, "csv"); err != nil {
			b.Fatal(err)
		}
	}
}