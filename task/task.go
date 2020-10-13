package task

import (
	"context"
	"errors"
	"log"
	"sync"
)

type Task struct {
	mu   sync.RWMutex
	jobs map[string]*Job
}

// 新建一个作业任务池
func NewTask() *Task {
	return &Task{jobs: make(map[string]*Job)}
}
func (t *Task) Get(jobId string) *Job {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if job, ok := t.jobs[jobId]; ok {
		return job
	}
	return nil
}

// Add 添加一个作业任务
func (t *Task) Add(job *Job) (jobId string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	jobId = job.id
	if j, ok := t.jobs[jobId]; ok {
		// 如果存在，则返回已存在的jobId
		return j.id
	}
	t.jobs[jobId] = job
	return jobId
}

// Run 运行一个作业任务
func (t *Task) Run(jobId string) error {

}
func (t *Task) runCtx(jobId string, ctx context.Context, cancelFunc context.CancelFunc) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if job, ok := t.jobs[jobId]; ok {
		if !job.IsRun() {
			return job.runCtx(ctx, cancelFunc)
		}
	} else {
		return errors.New("not found")

	}
	return nil
}

// Stop 停止作业任务
func (t *Task) Stop(jobId string) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if job, ok := t.jobs[jobId]; ok {
		if job.IsRun() {
			job.Stop()
		}
	}
}

// Delete 删除作业任务
func (t *Task) Delete(id string) {
	// 停止作业任务
	t.Stop(id)
	// 删除作业任务
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.jobs, id)
	log.Println("delete jobId:", id)
}
