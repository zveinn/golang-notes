package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"
)

// Signal ..
// This type is used to signal to the main program loop
// that some goroutine or process needs to be executed.
type Signal struct {
	ctx context.Context
	ID  int
	Tag string
}

// goroutineMonitor is used to deliver signals
// to the main program loop.
var goroutineMonitor = make(chan *Signal, 100)

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

var (
	workerCount = 1
	jobChannel  = make(chan int, 100)
)

func main() {
	flag.IntVar(&workerCount, "workerCount", 1, "define how many workers you want to start")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	// This will trigger the goroutine in "case 1" inside the main program loop
	for i := 0; i < workerCount; i++ {
		goroutineMonitor <- &Signal{ID: 1, ctx: ctx, Tag: "Worker " + strconv.Itoa(i)}
	}

	goroutineMonitor <- &Signal{ID: 2, ctx: ctx, Tag: "Work Generator "}

	// This channel is response for catching operating system signal
	// like ctrl+c and then exit the program cleanely using cancel()
	// and os.Exit(1)
	operatingSystemSignal := make(chan os.Signal, 10)
	signal.Notify(
		operatingSystemSignal,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGILL,
	)

	for {
		select {

		case osSignal := <-operatingSystemSignal:
			logger.Debug("Application is exiting with signal:", osSignal)
			cancel()
			os.Exit(1)

		case signal := <-goroutineMonitor:
			logger.Debug("Starting goroutine", "ID", strconv.Itoa(signal.ID), "Tag", signal.Tag)

			switch signal.ID {
			case 1:
				go worker(signal)
			case 2:
				go sendWork(signal)
			default:
				logger.Warn("Unknown signal caught", "ID", strconv.Itoa(signal.ID), "Tag", signal.Tag)
			}

		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func worker(signal *Signal) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, string(debug.Stack()))
		}

		// If the goroutine exits or panics we drop the signal back
		// into the goroutine monitor to restart it.
		goroutineMonitor <- signal
	}()

	for job := range jobChannel {
		logger.Info("New job", "Worker", signal.Tag, "JobID", strconv.Itoa(job))
	}
}

func sendWork(signal *Signal) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, string(debug.Stack()))
		}

		// If the goroutine exits or panics we drop the signal back
		// into the goroutine monitor to restart it.
		goroutineMonitor <- signal
	}()

	counter := 0
	for {
		counter++
		jobChannel <- counter
		time.Sleep(1 * time.Second)
	}
}

func startAnAPI(signal *Signal) {
	defer func() {
		r := recover()
		if r != nil {
			log.Println(r, string(debug.Stack()))
		}

		// If the goroutine exits or panics we drop the signal back
		// into the goroutine monitor to restart it.
		goroutineMonitor <- signal
	}()

	logger.Info("Start you webserver here and it will not exit the program if it fails")
	time.Sleep(30 * time.Second)
}
