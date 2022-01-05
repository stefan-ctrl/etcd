package raft

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	prefix   = "[leaderelection]"
	timepath = "/tmp/time.log"
)

type Event string

const (
	SERVICE_START     Event = "t0" // + done outside of this
	ELECTION_START    Event = "t1" // +
	LEADER_ELECTED    Event = "t2" // +
	READY             Event = "t3" // + done outside of this. Will be done in Ready handler
	STABLE_LEADER     Event = "t4" // +
	SERVICE_STOP      Event = "t5" // +
	T2_START          Event = "t6" // + done outside of this
	T2_ELECTION_START Event = "t7" // +
	T2_LEADER_ELECTED Event = "t8" // +
	T2_READY          Event = "t9" // + as just in case measurement
	T2_STABLE         Event = "tA" // +
	T2_STOPPED        Event = "tB" // +
	T3_START          Event = "tC" // + done outside of this
	T3_ELECTION_START Event = "tD" // +
	T3_LEADER_ELECTED Event = "tE" // +
	T3_READY          Event = "tF" // + done outside of this. Will be done in Ready handler
	T3_STABLE         Event = "tG" // +
	T3_STOPPED        Event = "tH" // +

	UNSTABLE       Event = "UNSTABLE"       // +
	LEADER_STOPPED Event = "LEADER_STOPPED" // +
)

//Print logs with [leaderelection] prefix
func PrintTiming(event Event) {
	time := time.Now().UnixMilli()
	var str string
	switch event {
	case LEADER_ELECTED:
		str = fmt.Sprintf("%s:%d\n%s:%d\n%s:%d", LEADER_ELECTED, time,
			T2_LEADER_ELECTED, time,
			T3_LEADER_ELECTED, time)
	case ELECTION_START:
		str = fmt.Sprintf("%s:%d\n%s:%d\n%s:%d", ELECTION_START, time,
			T2_ELECTION_START, time,
			T3_ELECTION_START, time)
	case STABLE_LEADER:
		str = fmt.Sprintf("%s:%d\n%s:%d\n%s:%d", STABLE_LEADER, time,
			T2_STABLE, time,
			T3_STABLE, time)
	case SERVICE_STOP:
		str = fmt.Sprintf("%s:%d\n%s:%d\n%s:%d", SERVICE_STOP, time,
			T2_STOPPED, time,
			T3_STOPPED, time)
	default:
		str = fmt.Sprintf("%s:%d", event, time)
	}
	log.Print(str)
	writeTimingToFile(str)
}

func PrintDebug(msg string) {
	writeTimingToFile(msg)
}

func writeTimingToFile(msg string) {
	msg = msg + "\n"
	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(timepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	b := []byte(msg)
	if _, err := file.Write(b); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
