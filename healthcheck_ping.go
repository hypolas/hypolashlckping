package hypolashlckping

import (
	"github.com/go-ping/ping"
	helpers "github.com/hypolas/hypolashlckhelpers"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Call is the entry module point
func Call() helpers.Result {
	var w sync.WaitGroup
	timeInt64, err := strconv.ParseInt(testTimeoutInSeconds, 10, 32)
	if err != nil {
		log.Err.Fatalln(err)
	}
	d := time.Duration(timeInt64) * time.Second

	if timeout(&w, d) {
		result.Output = "TIMEOUT"
		log.Err.Println("TIMEOUT")
	}

	w.Wait()

	return result
}

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	c := make(chan int)
	w.Add(1)
	go func() {
		defer close(c)
		pinger, err := ping.NewPinger(host)
		pinger.SetPrivileged(true)
		if err != nil {
			log.Err.Fatalln(err)
		}

		pinger.Count, err = strconv.Atoi(pingCount)
		if err != nil {
			log.Err.Fatalln(err)
		}

		err = pinger.Run()
		if err != nil {
			log.Err.Println(err)
		}

		stats := pinger.Statistics()
		if stats.PacketLoss <= float64(maxPercentLost) {
			result.IsUP = true
		}

		result.Output = strconv.Itoa(int(stats.PacketLoss))
		if err != nil {
			log.Err.Fatalln(err)
		}
		w.Done()
	}()
	select {
	case <-c:
		return false
	case <-time.After(t):
		return true
	}
}
