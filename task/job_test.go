package task

import (
	"log"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	var jobFunc JobFunc = func() error {
		time.Sleep(time.Second * 5)
		return nil
	}
	job := NewJob(jobFunc)
	log.Println("1.运行")
	_ = job.Run()
	time.Sleep(time.Second)
	log.Println("2.停止")
	job.Stop()
	time.Sleep(time.Second)
	log.Println("2.运行")
	_ = job.RunWithDeadline(time.Now().Add(time.Millisecond * 2))
	_ = job.RunWithTimeout(time.Millisecond * 2)

}
