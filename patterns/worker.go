package patterns

import (
	"net"
	"sync"
)

// A "Worker" is anything that runs a go routine that loops over an open chanel
// and sends incoming values off to some kind of end-point, whether it be a
// remote system, disk file, email, or something else.
//
// Generally, if you have a worker manager with values the worker needs that may
// change, use an atomic pointer to load those values. For example a worker
// manager may have a dynamically assembled list of end points the worker needs
// to know to forward data to. Then the worker could use an atomic pointer to
// load those values in the middle of the goroutine loop.
//
// Just be sure to replace the w.in to the correct type, as long as it remains
// a pointer.
//
// Also remember to replace w.endPoints with the correct type.
//
// Lastly, w.loop() will need to be updated with the actual instructions.
type Worker struct {
	// Replace with type to receive. Must be pointer, but to any type.
	in chan *string

	// Replace with whatever the worker is sending data to.
	endPoints []net.Conn

	wg     sync.WaitGroup
	IsIdle bool // Is the loop paused?
}

func NewWorker() *Worker {
	return &Worker{
		IsIdle: true,
	}
}

// Start starts the worker loop on a go routine.
func (w *Worker) Start() { w.IsIdle = false; w.wg.Add(1); go w.loop() }

// Stop closes the in channel and tells the WaitGroup to wait so that any
// outstanding goroutines will finish and then the worker is idled.
func (w *Worker) Stop() { close(w.in); w.wg.Wait(); w.IsIdle = true }

// loop is the primary way the worker "works". It is started as a goroutine in
// w.Start() and will continue for as long as the channel is open.
//
// The WaitGroup waiting isn't strictly necessary but added for boilerplate.
//
// The channel value nil check is required to make sure the loop doesn't send
// junk data to an endPoint or cause an error, hence the in type needing to be
// chan *[T].
func (w *Worker) loop() {
	defer w.wg.Done()

	for s := range w.in {
		if s == nil {
			continue
		}

		for _, i := range w.endPoints {
			i.Write([]byte(*s))
		}

		w.wg.Wait()
	}
}
