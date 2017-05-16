package cache

import (
	"math/rand"
	"strconv"
	"testing"
)

// func Test_GetItem(t *testing.T) {
// 	v := Instance().GetItem("1")
// 	if v != "" {
// 		t.Error("居然能获取到值")
// 	}
// }

// func Test_SetItem(t *testing.T) {
// 	f := Instance().SetItem("1", "yidane")
// 	if !f {
// 		t.Error("保存数据失败")
// 	}
// }

// func Test_BuilkSetItem(t *testing.T) {
// 	for i := 0; i < 2000000; i++ {
// 		f := Instance().SetItem(strconv.Itoa(i), time.Now().String())
// 		if !f {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func Test_SingleSton(t *testing.T) {

// }

// func Benchmark_T(b *testing.B) {
// 	lock := &sync.Mutex{}
// 	for i := 0; i < b.N; i++ {
// 		lock.Lock()
// 		//fmt.Println(time.Now())
// 		lock.Unlock()
// 	}
// }
func Benchmark_GetItem(b *testing.B) {
	v := Instance().GetItem(strconv.Itoa(rand.Intn(b.N)))
	if v != "" {
	}
}

// func Benchmark_TOne(b *testing.B) {
// 	once := &sync.Once{}
// 	count := 0
// 	pArr := make(map[int]bool)
// 	runtime.GOMAXPROCS(1)
// 	for i := 0; i < b.N; i++ {
// 		once.Do(func() {
// 			count++
// 			pid := os.Getppid()
// 			//fmt.Println("NumCPU:", runtime.NumCPU())
// 			if p := pArr[pid]; p == false {
// 				pArr[pid] = true
// 			}
// 		})
// 	}
// 	fmt.Println(pArr)
// }
