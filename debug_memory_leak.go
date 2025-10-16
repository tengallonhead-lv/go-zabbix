package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

// 内存泄漏排查工具函数
func startMemoryMonitoring() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				printDetailedMemStats()
			}
		}
	}()
}

func printDetailedMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	log.Printf(`=== Memory Stats ===
Heap Objects: %d
Heap In Use: %d KB
Heap Released: %d KB  
Total Alloc: %d KB
Sys Memory: %d KB
GC Cycles: %d
Goroutines: %d
===================`,
		m.HeapObjects,
		m.HeapInuse/1024,
		m.HeapReleased/1024,
		m.TotalAlloc/1024,
		m.Sys/1024,
		m.NumGC,
		runtime.NumGoroutine(),
	)
}

// 检查goroutine泄漏的辅助函数
func checkGoroutineGrowth() {
	initialCount := runtime.NumGoroutine()
	log.Printf("Initial goroutine count: %d", initialCount)
	
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				current := runtime.NumGoroutine()
				if current > initialCount*2 {
					log.Printf("WARNING: Goroutine count doubled! Current: %d, Initial: %d", 
						current, initialCount)
				}
			}
		}
	}()
}

func main() {
	// 启动pprof服务器
	go func() {
		log.Println("pprof server starting on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	// 启动内存监控
	startMemoryMonitoring()
	checkGoroutineGrowth()
	
	// 你的主程序逻辑
	select {}
}