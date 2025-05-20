package main

type Guess int

func (g Guess) IsValid() bool {
	if 0 < g && g < 101 {
		return true
	}
	return false
}
