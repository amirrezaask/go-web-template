package pool

type Task interface{
	Run() error
}

//Worker is the definition of every worker in the system, it's similar to unix/linux processes, where every process
//has stdin, stdout and stderror.
type Worker interface{
	Start(Task)
	StdIn() chan interface{}
	StdOut() chan interface{}
	StdErr() chan error
}

type WorkerPool struct {
	workerPool chan Worker
	taskPool chan Task
}

func (w *WorkerPool) Start() {
	for t := range w.taskPool {
		worker := <-w.workerPool
		go func(t Task) {
			worker.Start(t)
		}(t)
	}
}