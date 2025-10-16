# 内存增长排查指南

## 1. 监控多个指标对比

```bash
# 同时监控这些指标，观察增长趋势
curl -s http://localhost:6060/debug/pprof/heap?debug=1 | head -20
```

观察以下指标的变化：
- `process_resident_memory_bytes` (RSS)
- Go heap size (从pprof获取)
- goroutine数量
- 文件描述符数量

## 2. 内存快照对比分析

```bash
# 获取两个时间点的内存快照
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap

# 或者保存快照文件进行对比
curl -o heap1.pprof http://localhost:6060/debug/pprof/heap
# 等待一段时间后
curl -o heap2.pprof http://localhost:6060/debug/pprof/heap

# 对比两个快照
go tool pprof -base heap1.pprof heap2.pprof
```

## 3. 检查常见泄漏点

### goroutine泄漏检查
```bash
curl http://localhost:6060/debug/pprof/goroutine?debug=1
```

### 文件描述符检查
```bash
lsof -p [PID] | wc -l
# 或
ls /proc/[PID]/fd | wc -l
```

## 4. 运行时内存统计

在代码中添加定期打印：
```go
func printMemStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    log.Printf("Heap: %d KB, Sys: %d KB, NumGC: %d", 
        m.HeapInuse/1024, m.Sys/1024, m.NumGC)
}
```