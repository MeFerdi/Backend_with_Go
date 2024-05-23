package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)

	if length >= 3 {
		podium[0], podium[2] = podium[2], podium[0]
	}
	return podium
}
