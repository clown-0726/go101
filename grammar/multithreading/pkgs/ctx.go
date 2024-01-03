package pkgs

import (
	"context"
	"fmt"
	"time"
)

func CtxTest() {
	baseCtx := context.Background()

	// 传递请求相关值的结构体：key/value 的 context
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)

	// 设置截止日期：声明一个 timeout 的 context
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 + time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("Child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	// 设置同步信号：
	select {
	case <-timeoutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("Main process exit!")
	}
}
