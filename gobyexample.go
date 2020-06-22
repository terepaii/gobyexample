package main

import (
	"fmt"
	"errors"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
	//"strconv"
)

// 11. Functions
func mult(a int, b int) (int, string) {
	res := a * b
	return res, "Successfully returned a result!"
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func passByVal(i int) {
	i++;
}
func passByRef(i *int) {
	*i++
}

type animal interface {
	makeSound() string
	tellMeYourSize() string
}

type dog struct {
	name string
	size string
}

type cat struct {
	name string
	size string
}

func newDog(name, size string) *dog {
	d := dog{name: name, size: size}
	return &d
}

func (d *dog) makeSound() string {
	return "Bark"
}

func (c *cat) makeSound() string {
	return "Meow"
}

func (d *dog) tellMeYourSize() string {
	s := fmt.Sprintf("I am %s!\n", d.size)
	return s
}

func (c *cat) tellMeYourSize() string {
	s := fmt.Sprintf("I am %s!\n", c.size)
	return s
}

func animalRunner(a animal) {
	fmt.Println(a.makeSound())
	fmt.Println(a.tellMeYourSize())
}

func returningANumberWithAnError() (int, error) {
	return 1, errors.New("error")
}

/*type intArgError stuct {
	int arg
	prob string
}*/


/*func (e *intArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}*/

func incrementBy(n *int, incBy int, syncType string) {
	*n += incBy
	fmt.Println(syncType, *n)
}

/*func worker(val int, done chan bool) {
	fmt.Printf("I am worker %d..\n", val)
	time.Sleep(time.Second)
	fmt.Printf("I am worker %d and I am done...\n", val)
	
	done <- true
}*/

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {

	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker [%d] started job [%d]\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker [%d] finished job [%d]\n", id, j)
		results <- j * 2
	}
}

func ping(pings chan<- string) {
	pings <- "ping"
}

func pong(pings <-chan string, pongs chan<- string) {
	ping := <-pings
	pongs <- ping
}

func incrementAtomicCounter(id int, wg *sync.WaitGroup, counter *uint64) {
	for i := 0; i < 1000; i++ {
		atomic.AddUint64(counter, 1)
		if *counter >= 50000 {
			break
		}
		fmt.Printf("Worker [%d] has incremented the atomic counter to [%d]\n", id, *counter)
	}

	// Decrement the wait group
	wg.Done()
}

func main() {
	// 1. Hello World
	// fmt.Println("Hello World!")

	// 2. Variables
	/*var b bool
	fmt.Println(b)

	something := "1" + "2"
	fmt.Println(something)*/

	// 3. Constants
	/* const b int = 0
	b = 2
	fmt.Println(b)*/

	// 4. Loops
	/*i := 0
	fmt.Println("First Loop:")
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nSecond Loop:")
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	fmt.Println("\nThird Loop:")
	limit := 10
	for i := 0; i < limit; i++ {
		if i > limit/2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("\nFourth Loop:")
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\nFinished Looping...")*/

	// 5. If statements

	/*var a int = 5
	if num := a; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num is in single digits")
	} else {
		fmt.Println("num is in double digits")
	}*/

	// 6. Switch statements

	/*whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool!")
		case int:
			fmt.Println("I'm an int!")
		case string:
			fmt.Println("I'm a string!")
		default:
			fmt.Println("Who am I?!", t)
		}
	}

	whatAmI(true)
	whatAmI("String")
	whatAmI(99)
	whatAmI(3.4)*/

	// 7. Initialize Array

	/*arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}*/

	// 8. Slices
	/*fmt.Println("Making slice s")
	s := make([]string, 100)
	for i := 0; i < len(s); i++ {
		s[i] = strconv.Itoa(i)
	}
	fmt.Println(s)

	fmt.Println("\ns[1:10]")
	fmt.Println(s[1:10])

	fmt.Println("\nt")
	t := append(s, "101")
	fmt.Println(t)

	fmt.Println("\nPrinting s again..")
	fmt.Println("", s)

	fmt.Println("\nAppending two slices together")
	a := []int{1, 2, 3, 4}
	fmt.Println("a = ", a)
	b := []int{5, 6, 7, 8}
	fmt.Println("b = ", b)
	fmt.Println("Appending b to a...")
	a = append(a, b...)
	fmt.Println("a = ", a)*/

	// 9. Maps
	/*m := make(map[int]string)
	m[1] = "One"
	m[2] = "Two"
	m[3] = "Three"

	fmt.Println("This is a number to string map")
	fmt.Println(m)

	fmt.Println("The string for 1 is: ", m[1])
	fmt.Println("The string for 2 is: ", m[2])
	fmt.Println("The string for 3 is: ", m[3])

	fmt.Println("Deleting m[3]...")
	delete(m, 3)

	_, present := m[3]
	fmt.Println("Is the value for m[3] present?")
	if present {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No :(")
	}*/

	// 10. Range
	/*s := []int{1, 2, 3, 4}

	for i, v := range s {
		fmt.Printf("%d + %d = %d\n", i, v, i + v)
	}

	for i, c := range "Stephen Terepaii Tangi O'Sullivan" {
		fmt.Printf("%d, %s == %x\n", i, string(c), c)
	}*/

	// 11. Functions
	/*
	fmt.Println("Multiplying 1 and 2")
	fmt.Println(mult(1, 2))

	// Variadic
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("Adding up these numbers: %d\n", s)
	fmt.Println(sum(s...))*/

	// 12. Closures
	/*nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4 
	fmt.Println(nextInt()) // 5

	anotherNextInt := intSeq()
	fmt.Println(anotherNextInt()) // 1, new state separate from nextInt()*/

	// 13. Recursion
	//fmt.Println(fact(1))

	// 14. Pointers
	/*i := 0
	fmt.Printf("i is now %d\n", i)
	fmt.Println("Passing by value, incrementing by 1...")
	passByVal(i)
	fmt.Printf("i is now %d\n", i)
	fmt.Println("Passing by reference, incrementing by 1...")
	passByRef(&i)
	fmt.Printf("i is now %d\n", i)*/

	// 15. Structs
	//f := newDog("Fido", "medium")
	/*fmt.Println(a)*/
	
	// 16. Methods
	//fmt.Printf("%s makes a %s sound\n", f.name, f.makeSound())

	// 17. Interfaces
	//c := cat{"Mittens", "small"}
	//animalRunner(f)
	//animalRunner(&c)

	// 18. Errors
	//fmt.Println(returningANumberWithAnError())

	// 19. Goroutines
	/*i :=  0
	incrementBy(&i, 10, "sync")
	go incrementBy(&i, 5, "async")
	for j := 0; j < 10; j++ {
		go incrementBy(&i, j, "async")
		go func(msg string) {
			fmt.Println(msg)
		}("going")
	}
	time.Sleep(time.Second)
	fmt.Println("done")*/

	// 20. Channels

	/*messages := make(chan int)

	go func() {messages <- 200}()

	fmt.Println(<-messages)
	//fmt.Println(<-messages)

	done := make(chan bool, 3)
	go worker(1, done)
	go worker(2, done)
	go worker(3, done)

	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings)
	pong(pings, pongs)

	fmt.Println(<-pongs)*/

	// 21. Select
	/*c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "c3"
	}()

	for i:= 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <- c3:
			fmt.Println("received", msg3)
		}	
	}*/

	// 22. Timeouts
	/*c1 := make(chan string, 1)
	go func () {
		time.Sleep(2 * time.Second)
		c1 <- "I got the result"
	}()

	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timed out...")
	}

	c2 := make(chan string, 1)

	go func () {
		time.Sleep(time.Second)
		c2 <- "I got the result the second time around!"
	}()

	select {
	case res:= <-c2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out again...")
	}*/

	// 23. Non-blocking select statements
	/*messages := make(chan string)
	
	go func () {
		//time.Sleep(time.Second)
		messages <- "Hello!"
	}()

	loop:
	for {
		select {
		case res:= <-messages:
			fmt.Println(res)
			break loop
		case <-time.After(2 * time.Second):
			fmt.Println("Timed out...")
			break loop
		default:
			fmt.Println("No messages received")
		}
	}*/

	// 24. Closing Channels 
	/*jobs := make(chan int, 5)
	done := make(chan bool, 1)



	for j := 0; j < 5; j++ {
		jobs <- j
		fmt.Printf("Sending job %d\n", j)
	}

	go func () {
		for {
			job, nextJob := <- jobs
			if nextJob {
				fmt.Printf("Received job %d\n", job)
			} else {
				fmt.Println("Received all jobs!")
				done <- true
				return
			}
		}
	} ()
	close(jobs)
	fmt.Println("Sent all jobs!")

	// Block until done channel is populated
	<-done*/

	// 25. Range over channels
	/*queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	} */

	// 26. Timers
	/*duration := 2 * time.Second
	t1 := time.NewTimer(duration)

	<-t1.C

	fmt.Printf("t1 Fired after [%.f] seconds", duration.Seconds())*/

	// 27. Tickers
	/*ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool, 1)


	go func () {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at: ", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true

	fmt.Println("Timer stopped"	)*/

	// 28. Worker Pools
	/*numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 1; i <= numJobs; i++ {
		go worker(i, jobs, results)
	}

	for j:=1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}*/

	// 29. Wait Groups
	/*var wg sync.WaitGroup

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < numJobs; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()*/

	// 30. Rate Limiting
	/*numRequests := 5
	requests := make(chan int, numRequests)
	for i := 0; i < numRequests; i++ {
		requests <- i
	}

	limiter := time.Tick(1000 * time.Millisecond)

	for j := 0; j < numRequests; j++ {
		<-limiter
		fmt.Printf("Processing request [%d]\n", <-requests)
	}

	burstLimit := 3
	burstReqQueue := make(chan time.Time, burstLimit)

	for k := 0; k < burstLimit; k++ {
		burstReqQueue <- time.Now()
	}*/

	// 31. Atomic Counters
	/*var atomicCounter uint64
	var wg sync.WaitGroup
	// Start 50 goroutines
	for workerNum := 0; workerNum < 16; workerNum++ {
		// Add this process to the wait group
		wg.Add(1)
		go incrementAtomicCounter(workerNum, &wg, &atomicCounter)
	}

	wg.Wait()*/

	// 32. Mutexes

	/*var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var wg sync.WaitGroup

	var readOps uint64
	var writeOps uint64

	for w := 0; w < 100; w++ {
		go func () {
			wg.Add(1)
			num := rand.Intn(50)
			key := rand.Intn(50)
			mutex.Lock()
			state[key] = num
			mutex.Unlock()
			atomic.AddUint64(&writeOps, 1)

			// Wait between writes
			time.Sleep(time.Millisecond)
		} ()
	}
	
	go func () {
		for r := 0; r < 100; r++ {
			mutex.Lock()
			fmt.Println(state[rand.Intn(5)])
			mutex.Unlock()
			atomic.AddUint64(&readOps, 1)

			// Wait between reads
			time.Sleep(time.Millisecond)
		}
		
	} ()
	time.Sleep(2 * time.Second)

	mutex.Lock()
	fmt.Println("State map: ", state)
	mutex.Unlock()*/
}

