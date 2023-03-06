package hypolashlckping

import (
	"runtime"
	"strconv"

	"github.com/go-ping/ping"
	helpers "github.com/hypolas/hypolashlckhelpers"
)

func Call() helpers.Result {
	pinger, err := ping.NewPinger(host)
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}
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

	return result
}
