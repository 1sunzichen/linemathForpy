package main

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}
	if bits[len(bits)-2] == 0 {
		return true
	}
	return false
}
