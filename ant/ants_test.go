package main

import (
	ants "github.com/panjf2000/ants/v2"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"testing"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

const (
	taskNum  = 5000000
	pollSize = ants.DefaultAntsPoolSize
)

var curMem uint64

//测试使用常规方案T
func TestNoPool(t *testing.T) {
	f, _ := os.Create("cpu_no.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	wg := sync.WaitGroup{}
	wg.Add(taskNum)
	for i := 1; i <= taskNum; i++ {
		go func() {
			demoFunc()
			wg.Done()
		}()
	}
	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)

	ff, _ := os.Create("mem_no.prof")
	pprof.WriteHeapProfile(ff)
}

//测试使用ants
func TestAntsPool(t *testing.T) {

	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	p, _ := ants.NewPool(pollSize)
	defer p.Release()
	wg := sync.WaitGroup{}
	wg.Add(taskNum)
	for i := 1; i <= taskNum; i++ {
		_ = p.Submit(func() {
			demoFunc()
			wg.Done()
		})
	}
	wg.Wait()
	t.Logf("pool, capacity:%d", p.Cap())
	t.Logf("pool, running workers number:%d", p.Running())
	t.Logf("pool, free workers number:%d", p.Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)

	ff, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(ff)
}
