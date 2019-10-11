package go_pool

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"time"
)

const (
	DefaultGPoolMaxWorker = 50
	DefaultIdelTime       = 1 * time.Second
	DefaultDispatchPeriod = 200 * time.Millisecond
	bucketSize            = 5
)

//pool config
type Options struct {
	MinWorker int
	MaxWorker int
	//queue size
	JobQueueSize   int
	IdleTimeOut    time.Duration
	DispatchPeriod time.Duration
}

//待运行任务
type Runnable func()

type taskModel struct {
	call Runnable
	res  chan bool
}

type GoPool struct {
	sync.Mutex
	ctx            context.Context
	cancel         context.CancelFunc
	isClose        bool
	maxWorker      int
	minWorker      int
	idleTimeOut    time.Duration
	dispatchPeriod time.Duration
	JobQueueSize   int
	jobChan        chan taskModel
	killChan       chan bool
	call           Runnable
	currWorker     int
}

func (p GoPool) dispatch(maxWorker int) {
	if p.currWorker == p.maxWorker {
		return
	}
	p.Lock()
	defer p.Unlock()

	for i := 0; i < maxWorker; i++ {
		if p.currWorker >= p.maxWorker {
			return
		}
		go p.handle()
		p.currWorker++
	}
}

func (p *GoPool) handle() {
	timer := time.NewTimer(p.idleTimeOut)
	for {
		select {
		case job := <-p.jobChan:
			job.call()
			if job.res != nil {
				job.res <- true
			}
			timer.Reset(p.idleTimeOut)
		case <-timer.C:
			p.killWorker(timer)
			return
		case <-p.killChan:
			p.killWorker(timer)
			return
		}
	}
}

func (p GoPool) killWorker(timer *time.Timer) {
	p.Lock()
	defer p.Unlock()
	if p.currWorker <= p.minWorker {
		timer.Reset(p.idleTimeOut)
		return
	}
	p.currWorker--
	timer.Stop()
}

func NewPool(opt *Options) (*GoPool, error) {
	//config
	pool := GoPool{}
	pool.minWorker = opt.MinWorker
	pool.maxWorker = opt.MaxWorker
	pool.JobQueueSize = opt.JobQueueSize
	pool.idleTimeOut = opt.IdleTimeOut
	pool.dispatchPeriod = opt.DispatchPeriod

	if pool.maxWorker <= 0 {
		pool.maxWorker = DefaultGPoolMaxWorker
	}

	if pool.dispatchPeriod < time.Millisecond || pool.dispatchPeriod > time.Second {
		pool.dispatchPeriod = DefaultDispatchPeriod
	}

	if pool.minWorker < 1 {
		pool.minWorker = pool.maxWorker / 5
	}

	if pool.JobQueueSize <= 0 {
		pool.JobQueueSize = pool.maxWorker
	}

	if pool.maxWorker < pool.minWorker {
		return nil, errors.New("max size less than min size!")
	}

	//init
	pool.jobChan = make(chan taskModel, pool.JobQueueSize)
	pool.killChan = make(chan bool, 0)

	pool.isClose = false
	ctx, cancel := context.WithCancel(context.Background())
	pool.ctx = ctx
	pool.cancel = cancel

	pool.dispatch(pool.maxWorker)

	return &pool, nil

}

func (p GoPool) SupplyAsync(fn Runnable) error {
	task := taskModel{
		call: fn,
		res:  nil,
	}

	p.trySpawnWorker()

	go func() {
		p.jobChan <- task
	}()
	return nil
}

func (p GoPool) trySpawnWorker() {
	if len(p.jobChan) > 0 && p.currWorker < p.maxWorker {
		p.dispatch(p.maxWorker / bucketSize)
	}
}

func (p GoPool) Close() error {
	p.Lock()
	if p.isClose {
		p.Unlock()
		return nil
	}
	p.isClose = true
	p.cancel()
	p.Unlock()
	return nil
}
