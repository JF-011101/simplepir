package db

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Empty struct{}

func GenData() {
	start := time.Now() // 获取当前时间
	wg := &sync.WaitGroup{}
	limiter := make([]chan Empty, 101)
	for o := 0; o < 101; o++ {
		limiter[o] = make(chan Empty, 1)
	}
	limiter[0] <- Empty{}
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func(first, last, k int, wg *sync.WaitGroup, limiter []chan Empty) {
			defer wg.Done()
			var insertRecords []Pi
			for i := first; i < last; i++ {
				gid := GetGID()
				a := time.Now().UnixNano() + gid + int64(i)
				b := a * (gid / 10)
				n1 := fmt.Sprintf("1%05v%05v", rand.New(rand.NewSource(a)).Int31n(100000), rand.New(rand.NewSource(b)).Int31n(100000))
				n2, _ := strconv.ParseUint(n1, 10, 64)
				insertRecords = append(insertRecords,
					Pi{
						PhoneNumber: n2,
					},
				)

			}

			<-limiter[k]
			err := DB.Create(insertRecords).Error
			if err != nil {
				fmt.Print(err)
				panic(err)
			}
			limiter[k+1] <- Empty{}
		}(j*1000, j*1000+1000, j, wg, limiter)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("创建10万数据完成耗时：", elapsed)

}

func GetGID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseInt(string(b), 10, 64)
	return n
}
