package mapreduce
import "container/list"
import "fmt"

type WorkerInfo struct {
  address string
  // You can add definitions here.
}

type JobsCtl struct {
  jobId int
  jobStatus bool
}

// Clean up all workers by sending a Shutdown RPC to each one of them Collect
// the number of jobs each work has performed.
func (mr *MapReduce) KillWorkers() *list.List {
  l := list.New()
  for _, w := range mr.Workers {
    DPrintf("DoWork: shutdown %s\n", w.address)
    args := &ShutdownArgs{}
    var reply ShutdownReply;
    ok := call(w.address, "Worker.Shutdown", args, &reply)
    if ok == false {
      fmt.Printf("DoWork: RPC %s shutdown error\n", w.address)
    } else {
      l.PushBack(reply.Njobs)
    }
  }
  return l
}

func (mr *MapReduce) RunMaster() *list.List {
  // Your code here

  MapCtl := make(chan JobsCtl)

  ReduceCtl := make(chan JobsCtl)

  doMap := func(jobId int) {
    addr := <-mr.registerChannel
	var args DoJobArgs
	var reply DoJobReply
	args.File = mr.file
	args.Operation = Map
	args.JobNumber = jobId
	args.NumOtherPhase = mr.nReduce
	res := call(addr, "Worker.DoJob", args, &reply)
	MapCtl <- JobsCtl{jobId, res}
	mr.registerChannel <- addr
  }

  doReduce := func(jobId int) {
    addr := <-mr.registerChannel
	var args DoJobArgs
	var reply DoJobReply
	args.File = mr.file
	args.Operation = Reduce
	args.JobNumber = jobId
	args.NumOtherPhase = mr.nMap
	res := call(addr, "Worker.DoJob", args, &reply)
	ReduceCtl <- JobsCtl{jobId, res}
	mr.registerChannel <- addr
  }

  for i := 0; i < mr.nMap; i ++ {
    go doMap(i)
  }

  for i := 0; i < mr.nMap; i ++ {
    <-MapCtl
  }

  fmt.Println("Map Finish")

  for i := 0; i < mr.nReduce; i ++ {
    go doReduce(i) 
  }

  for i := 0; i < mr.nReduce; i ++ {
    <-ReduceCtl
  }

  return mr.KillWorkers()
}
