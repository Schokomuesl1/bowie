package basiswerte

func Kosten(SK string, stufe int) int {
	_, exists := kosten[SK]
	if !exists {
		return -1
	}
	if stufe > 25 || stufe < 0 {
		return -1
	}
	return kosten[SK][stufe]
}
