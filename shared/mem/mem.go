package mem
import (
	"runtime"
	"fmt"
	"github.com/pivotal-golang/bytefmt"
)


func Report() {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	fmt.Println("Alloc: ", bytefmt.ByteSize(m.Alloc * bytefmt.BYTE))
	fmt.Println("Total Alloc: ", bytefmt.ByteSize(m.TotalAlloc * bytefmt.BYTE))
	fmt.Println("Sys: ", bytefmt.ByteSize(m.Sys * bytefmt.BYTE))
	fmt.Println("Lookups:", bytefmt.ByteSize(m.Lookups * bytefmt.BYTE))
}
