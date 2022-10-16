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

type Stats struct {
	health    int
	energy    int
	maxHealth int
	maxEnergy int
}

type Player struct {
	name  string
	stats Stats
}

func (player *Player) RestoreHealth(health int) {
	player.stats.health += health
	if player.stats.health > player.stats.maxHealth {
		player.stats.health = player.stats.maxHealth
	}
}

func (player *Player) RestoreEnergy(energy int) {
	player.stats.energy += energy
	if player.stats.energy > player.stats.maxEnergy {
		player.stats.energy = player.stats.maxEnergy
	}
}

func (player *Player) ApplyDamage(health int, energy int) {
	player.stats.health -= health
	player.stats.energy -= energy

	if player.stats.energy < 0 {
		player.stats.energy = 0
	}

	if player.stats.health < 0 {
		player.stats.health = 0
	}
}

func main() {
	player := Player{"shank", Stats{100, 100, 500, 100}}
	fmt.Println(player)
	player.RestoreHealth(50)
	player.RestoreEnergy(100)
	fmt.Println(player)
	player.ApplyDamage(70, 110)
	fmt.Println(player)
}
