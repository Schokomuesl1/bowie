package basiswerte

func Kosten(SK string, stufe int) int {
	_, exists := Kostentable[SK]
	if !exists {
		return -1
	}
	if stufe > 25 || stufe < 0 {
		return -1
	}
	return Kostentable[SK][stufe]
}
