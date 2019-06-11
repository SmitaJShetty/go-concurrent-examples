package main

import ("fmt"
"sync"
"time"
)

func main(){
	usingChannel()

	usingWaitGroup()

	usingCond()
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