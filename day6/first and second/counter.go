package main

func Process(maxDays int, f Fish) int64 {
	var num int64 = 1

	curDay := f.ObserveDay

	for f.ObserveDay+f.Timer < maxDays {
		// each timer interval we spawn a new fish, until we reach max days
		curDay += f.Timer + 1

		// reset the fish's timer to the global reset time
		f.Timer = TimerReset - 1

		// spawn a new fish with the current day as observation day
		nf := Fish{
			Timer:      8,
			ObserveDay: curDay,
		}

		num += Process(maxDays, nf)

		// the fish's observe day is moved forward
		f.ObserveDay = curDay
	}

	return num
}
