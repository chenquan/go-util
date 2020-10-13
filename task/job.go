package task

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"sync"
	"time"
)

// IJobe 实现此接口用于执行指定作业任务
type IJobe interface {
	RunJob() error
}

// JobFunc 函数式编程实现 IJobe 接口
type JobFunc func() error

// RunJob 实现 IJobe 接口，执行 JobFunc 函数
func (j JobFunc) RunJob() error {
	return j()
}

type Job struct {
	id     string // 作业ID
	run    bool   //是否正在运行
	ctx    context.Context
	cancel context.CancelFunc // 用于停止任务通道
	job    IJobe
	mu     sync.RWMutex
}

func NewJob(job IJobe) *Job {
	u1 := uuid.Must(uuid.NewV4(), nil)
	uuidString := strings.ReplaceAll(u1.String(), "-", "")
	return &Job{id: uuidString, job: job}
}

func (j *Job) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	return j.runCtx(ctx, cancel)
}
func (j *Job) RunWithTimeout(duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	return j.runCtx(ctx, cancel)
}
func (j *Job) RunWithDeadline(d time.Time) error {
	ctx, cancel := context.WithDeadline(context.Background(), d)
	return j.runCtx(ctx, cancel)
}

func (j *Job) runCtx(ctx context.Context, cancel context.CancelFunc) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	err := j.job.RunJob()
	log.Println("run jobId:", j.id)

	if err != nil {
		return err
	}
	j.ctx = ctx
	j.cancel = cancel
	j.run = true

	go func(job *Job) {
		for {
			select {
			case <-job.ctx.Done():
				fmt.Println("guan")
				return

			default:

			}
		}
	}(j)
	return err
}
func (j *Job) Stop() {

	if j.IsRun() {
		j.mu.Lock()
		j.run = false
		j.cancel()
		j.mu.Unlock()
		log.Println("stop jobId:", j.id)
	}
	log.Println("not run jobId:", j.id)

}
func (j *Job) IsRun() bool {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.run
}
func (j *Job) JobId() string {
	return j.id
}
