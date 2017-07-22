package auth

import (
	"conf"
	"time"
)

func ValidationKeyGenerate(id int) ValidationKeyPairTime {
	pair := ValidationKeyPairTime{
		ValidationKeyPair:ValidationKeyPair{
			ID:id,
			Key:conf.RandString(20),
		},
		GeneratedTime:time.Now(),
	}
	return pair
}
func ValidationKeyUpdate(pairs []ValidationKeyPairTime,outDateSeconds float64) {
	for {
		validationKeyClear(pairs,outDateSeconds)
		time.Sleep(300 * time.Second)
	}
}
func validationKeyClear(pairs []ValidationKeyPairTime,outDateSeconds float64) {
	j := 0
	i := 0
	for k := j; k < len(pairs); k++ {
		if isValidationKeyAvailable(pairs[k],outDateSeconds) {
			// swap [swapper] and [k]
			temp := pairs[i]
			pairs[i] = pairs[k]
			pairs[k] = temp
			// i指针自增
			i++
		}
	}
	pairs = pairs[i:]
}

func isValidationKeyAvailable(pair ValidationKeyPairTime,outDateSeconds float64) bool {
	return time.Since(pair.GeneratedTime).Seconds() > outDateSeconds
}

func FindValidationKey(pairs []ValidationKeyPairTime,target int) int {
	for i := 0; i < len(pairs); i++ {
		if pairs[i].ValidationKeyPair.ID == target {
			return i
		}
	}
	return -1
}
