/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/20
 */

package lol

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func ExampleDoThenRepeat() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := DoThenRepeat(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, false)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	fmt.Println(i)
	// Output:
	// 7
}

func ExampleRepeatTask() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := RepeatTask(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, false)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	fmt.Println(i)
	// Output:
	// 10
}
