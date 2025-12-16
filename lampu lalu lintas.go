package main

import (
	"fmt"
	"time"
)

func main() {
	currentLight := "RED" // Mulai dengan merah

	for {
		switch currentLight {
		case "RED":
			fmt.Println("Lampu Merah: Berhenti")
			time.Sleep(20 * time.Second)
			currentLight = "GREEN"
		case "YELLOW":
			fmt.Println("Lampu Kuning: Siap-siap")
			time.Sleep(20 * time.Second)
			currentLight = "RED"
		case "GREEN":
			fmt.Println("Lampu Hijau: Jalan")
			time.Sleep(5 * time.Second)
			currentLight = "YELLOW"
		}
	}
}