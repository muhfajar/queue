package queue

import "sync"

type Worker struct {
	Thread int
	Alloc  int
	Set    Callback

	tasks chan func() interface{}
}

type Callback struct {
	TaskDone  func(interface{})
	QueueDone func()
}

const (
	DefaultMaxThread = 25
	DefaultMaxAlloc  = 1000
)

func NewQueue(w *Worker) *Worker {
	if w.Thread == 0 {
		w.Thread = DefaultMaxThread
	}

	if w.Alloc == 0 {
		w.Alloc = DefaultMaxAlloc
	}

	return &Worker{
		tasks:  make(chan func() interface{}, w.Alloc),
		Thread: w.Thread,
		Alloc:  w.Alloc,
		Set: Callback{
			TaskDone:  w.Set.TaskDone,
			QueueDone: w.Set.QueueDone,
		},
	}
}

func (w *Worker) Start() {
	var wg sync.WaitGroup
	var mx = &sync.Mutex{}

	defer func() {
		close(w.tasks)
		if w.Set.QueueDone != nil {
			w.Set.QueueDone()
		}
	}()

	wg.Add(len(w.tasks))
	for i := 0; i < w.Thread; i++ {
		go func() {
			for {
				select {
				case task, ok := <-w.tasks:
					if !ok {
						break
					}

					rs := task()

					if w.Set.TaskDone != nil {
						mx.Lock()
						w.Set.TaskDone(rs)
						mx.Unlock()
					}

					wg.Done()
				default:
					break
				}
			}
		}()
	}

	wg.Wait()
}

func (w *Worker) Append(task func() interface{}) {
	w.tasks <- task
}
