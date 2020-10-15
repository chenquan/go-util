package task

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
	"sync"
	"time"
)

// IJobe 实现此接口用于执行指定作业任务
type IJob interface {
	IRunJob
	IStopJob
}
type IRunJob interface {
	RunJob() error
}
type IStopJob interface {
	StopJob()
}

// JobFunc 函数式编程实现 IJobe 接口
type JobFunc func() error

// RunJob 实现 IJobe 接口，执行 JobFunc 函数
func (j JobFunc) RunJob() error {
	return j()
}

// StopJob 实现 IJobe 接口，执行 JobFunc 函数
func (j JobFunc) StopJob() {}

// Job 作业任务
type Job struct {
	id     string             // 作业ID
	run    bool               //是否正在运行
	ctx    context.Context    // 上下文
	cancel context.CancelFunc // 用于停止任务通道
	job    IJob               // 任务
	mu     sync.RWMutex       // 任务信息锁
	once   sync.Once
}

// NewJob 新建一个作业
// 列如:
// var JobFunc1 JobFunc = func() error {
//		fmt.Println("第一个任务，我运行了")
//		return nil
//	}
// job1 := NewJob(jobFun1)
func NewJob(job IJob) *Job {
	u1 := uuid.Must(uuid.NewV4(), nil)
	uuidString := strings.ReplaceAll(u1.String(), "-", "")
	return &Job{id: uuidString, job: job}
}

// Run 执行作业,至到调用 Stop 关闭
func (j *Job) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	return j.runCtx(ctx, cancel)
}

// RunWithTimeout 执行作业任务,超时自动停止.可调用 Stop 手动关闭.
func (j *Job) RunWithTimeout(duration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	return j.runCtx(ctx, cancel)
}

// RunWithDeadline 执行作业任务,超过某个时间点自动停止.
func (j *Job) RunWithDeadline(d time.Time) error {
	ctx, cancel := context.WithDeadline(context.Background(), d)
	return j.runCtx(ctx, cancel)
}

// runCtx 执行作业任务，并指定上下文
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
				// 关闭
				j.job.StopJob()
				log.Println("stop jobId:", j.id)
				return
			default:
			}
		}
	}(j)
	return err
}

// Stop 停止作业任务
func (j *Job) Stop() {
	if j.IsRun() {
		j.mu.Lock()
		j.run = false
		j.cancel()
		j.mu.Unlock()
	} else {
		log.Println("not run jobId:", j.id)
	}

}

// IsRun 作业任务状态
func (j *Job) IsRun() bool {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.run
}

// JobId 作业任务ID
func (j *Job) JobId() string {
	return j.id
}
