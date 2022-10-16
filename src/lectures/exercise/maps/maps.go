//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func displayServer(servers []string, serversRegisterd map[string]int) {
	var numberOfServers int = len(servers)
	serverStatus := make(map[int]int)

	serverStatusToName := map[int]string{
		0: "Online",
		1: "Offline",
		2: "Maintenance",
		3: "Retired",
	}

	for _, value := range serversRegisterd {
		_, found := serverStatus[value]
		if !found {
			serverStatus[value] = 1
		} else {
			serverStatus[value] = serverStatus[value] + 1
		}
	}

	fmt.Println("Number of servers", numberOfServers)

	for key, value := range serverStatus {
		fmt.Println(serverStatusToName[key], "servers", value)
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverRegisterd := map[string]int{
		"darkstar": Online,
		"aiur":     Online,
		"omicron":  Online,
		"w359":     Online,
		"baseline": Online,
	}

	displayServer(servers, serverRegisterd)

	serverRegisterd["darkstar"] = Retired
	serverRegisterd["aiur"] = Offline

	displayServer(servers, serverRegisterd)

	for key, _ := range serverRegisterd {
		serverRegisterd[key] = Maintenance
	}

	displayServer(servers, serverRegisterd)
}
