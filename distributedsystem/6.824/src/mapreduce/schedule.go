package mapreduce

import (
	"fmt"
	"sync"
)

//
// schedule() starts and waits for all tasks in the given phase (mapPhase
// or reducePhase). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers. Once all tasks
	// have completed successfully, schedule() should return.
	//
	// Your code here (Part III, Part IV).
	//

	wg := &sync.WaitGroup{}

	taskCh := make(chan DoTaskArgs, ntasks)
	for i := 0; i < ntasks; i++ {
		wg.Add(1)
		taskCh <- DoTaskArgs{
			JobName:       jobName,
			File:          mapFiles[i],
			Phase:         phase,
			TaskNumber:    i,
			NumOtherPhase: n_other,
		}
	}
	go func() {
		for task := range taskCh {
			go doTask(task, <-registerChan, taskCh, registerChan, wg)
		}
	}()

	wg.Wait()
	close(taskCh)
	fmt.Printf("Schedule: %v done\n", phase)
}

func doTask(request DoTaskArgs, server string, taskChan chan DoTaskArgs, registerChan chan string, wg *sync.WaitGroup) {
	if success := call(server, "Worker.DoTask", request, nil); success {
		wg.Done()
		registerChan <- server
		return
	}

	taskChan <- request
	registerChan <- server
}
