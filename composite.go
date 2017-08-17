package main

import "composite"

func main() {
	solderA := composite.Soldier{"soldierA"}
	solderB := composite.Soldier{"soldierB"}
	solderC := composite.Soldier{"soldierC"}
	solderD := composite.Soldier{"soldierD"}

	soldier_queue_A := []composite.Soldier{solderA, solderB}
	captainA := composite.Captain{"captainA", soldier_queue_A}
	soldier_queue_B := []composite.Soldier{solderC, solderD}
	captainB := composite.Captain{"captainB", soldier_queue_B}

	captain_queue := []composite.Captain{captainA, captainB}
	general := composite.General{"GENERAL", captain_queue}
	general.Fight()
	/*
		GENERAL is fighting!
		captainA is fighting
		soldierA is fighting!
		soldierB is fighting!
		captainB is fighting
		soldierC is fighting!
		soldierD is fighting!
	*/
}
