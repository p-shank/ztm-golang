//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestRestoreHealth(t *testing.T) {
	testPlayer := Player{"test", Stats{0, 0, 100, 200}}
	testPlayer.RestoreHealth(200)
	if testPlayer.stats.health > testPlayer.stats.maxHealth {
		t.Errorf("Player health can not be greater than: %v", testPlayer.stats.maxHealth)
	}
}

func TestRestoreEnergy(t *testing.T) {
	testPlayer := Player{"test", Stats{0, 0, 100, 200}}
	testPlayer.RestoreEnergy(300)
	if testPlayer.stats.energy > testPlayer.stats.maxEnergy {
		t.Errorf("Player health can not be greater than: %v", testPlayer.stats.maxEnergy)
	}
}

func TestApplyDamage(t *testing.T) {
	testPlayer := Player{"test", Stats{200, 100, 400, 400}}
	testPlayer.ApplyDamage(300, 110)

	if testPlayer.stats.health < 0 {
		t.Errorf("Player health can not be less than: 0")
	}

	if testPlayer.stats.energy < 0 {
		t.Errorf("Player energy can not be less than: 0")
	}
}
