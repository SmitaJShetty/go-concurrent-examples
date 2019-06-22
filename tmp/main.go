package main

import ("fmt"
"sync"
"time"
)

func main(){
	usingChannel()

	usingWaitGroup()

	usingCond()

	usingOnce()

	usingPool()
}

func usingChannel(){
	fmt.Println("main started....")
	fileName := "testfile.csv"	
	fileReader := NewFileReader()
	ch:= make(chan int,1)
	go fileReader.DoWorkChannel(fileName, ch)		
	waitFunc := func(ch chan int){
		fmt.Println("in waitfunc with channel")
		<- ch
	}				
	fmt.Println("in main")
	//time.Sleep(500000)
	waitFunc(ch)	
}

func usingWaitGroup(){
	var wg sync.WaitGroup
	wg.Add(2)
	waitFunc:= func(wg *sync.WaitGroup){
		fmt.Println("in waitfunc with waitgroup")
		wg.Done()
	}
	fileReader := NewFileReader()
	go fileReader.DoWorkWaitGroup(&wg)
	go waitFunc(&wg)
	wg.Wait()
	fmt.Println("in main")
}

//using condition, one go routine adds rocketboosters, one counts. 
//main thread waits for the counter to finish
func usingCond(){
	const optimumRockerBoosterCount int = 2
	con := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(1)
	var rocketBoosters []int
	addRockerBoosters := func(){
		time.Sleep(10000)
		con.L.Lock()
		rocketBoosters=append(rocketBoosters, 1)
		fmt.Println("Added one rocket booster")
		con.L.Unlock()
		con.Signal()
	}

	startEngines:= func(){
		con.L.Lock()
		for len(rocketBoosters)<optimumRockerBoosterCount{
			con.Wait()
			fmt.Println("waiting ....")
		}
		con.L.Unlock()
		wg.Done()
	}
	go startEngines()
	for i:=0;i<3;i++ {
		go addRockerBoosters()
	}
	
	wg.Wait()
	fmt.Println("end of main")
}

func usingOnce(){
	/*sync.Once only counts number of times a Do is called, not how many times unique functions are passed to Do */
	counter := 0
	print:= func (){fmt.Println("counter value was: ", counter) }
	increment := func() { 
		counter++
		print()
	}

	decrement := func() { 
		counter--
		print()
	}

	var testSync sync.Once
	testSync.Do(increment)
	testSync.Do(decrement)
	
}

func usingPool(){
	/*
		Get is primary interface - checks for any existing instances. 
		If none are available, calls the New method to create a new instance.
		Put puts instances back into pool
	*/

	testPool := &sync.Pool{
		New: func() interface{}{
			fmt.Println("usingPool,new instance created")
			return struct{}{}
		},
	}
	testInstance:= testPool.Get()
	testPool.Put(testInstance)
	testPool.Get()
}