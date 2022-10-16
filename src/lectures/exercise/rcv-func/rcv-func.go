//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health int
type Energy int
type Stats struct {
	maxHealth Health
	maxEnergy Energy
}

type Player struct {
	name  string
	stats Stats
}

func (player *Player) updateHealth(health Health) {
	player.stats.maxHealth = health
}

func (player *Player) updateEnergy(energy Energy) {
	player.stats.maxEnergy = energy
}

func main() {
	player := Player{"shank", Stats{100, 200}}
	fmt.Println(player)
	player.updateHealth(50)
	player.updateEnergy(100)
	fmt.Println(player)
}
