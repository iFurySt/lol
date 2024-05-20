/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/20
 */

package lol

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestDoThenRepeatParallel(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := DoThenRepeat(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, false)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	assert.Equal(t, 7, i)
}

func TestDoThenRepeatSequential(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := DoThenRepeat(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, true)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	assert.Equal(t, 3, i)
}

func TestRepeatTaskParallel(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := RepeatTask(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, false)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	assert.Equal(t, 10, i)
}

func TestRepeatTaskSequential(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	i := 0
	wait := RepeatTask(ctx, time.Millisecond*100, func() {
		i += 1
		time.Sleep(300 * time.Millisecond)
	}, true)

	time.Sleep(1 * time.Second)
	cancel()
	wait()

	assert.Equal(t, 3, i)
}
