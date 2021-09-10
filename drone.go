package go_unit_test_bootcamp

// ----- Find drone func -----
func FindMissingDrone(droneIds []int) int {
	dict := make(map[int]int)

	for _, id := range droneIds {
		// val is the return value; ok is a bool that returns
		// 		if the key is in the map or not
		if val, ok := dict[id]; ok {
			dict[id] = val + 1
		} else {
			dict[id] = 1
		}
	}

	for key, value := range dict {
		// if not even then return that value since it's missing
		if value%2 != 0 {
			return key
		}
	}
	return 0
}
