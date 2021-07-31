package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type AlcoholType string

const (
	Wine      AlcoholType = "Wine"
	Beer      AlcoholType = "Beer"
	Vodka     AlcoholType = "Vodka"
	Water     AlcoholType = "Water"
	Moonshine AlcoholType = "Moonshine"
)

func GenerateAnswer(uid uint, username string) string {
	foundAnswer, result := CheckIfResultAlreadyGeneratedForToday(uid)
	if foundAnswer {
		return result
	}

	intro := generateIntro()
	alcoholType := generateType()
	amount := generateAmount(alcoholType)
	result = parseAmount(alcoholType, amount)
	result = intro + " " + result + " for today " + fmt.Sprintf("(%0.1f l.)", amount)
	InsertResult(username, result, uid)
	return result
}

func generateTruncatedNormal(mean, sd, a, b float64) float64 {
	rand.Seed(time.Now().UnixNano())
	result := rand.NormFloat64()*sd + mean
	for result < b || result > a {
		result = rand.NormFloat64()*sd + mean
	}

	return result
}

func generateIntro() string {
	rand.Seed(time.Now().UnixNano())
	variants := []string{
		"Hello!",
		"Today is your lucky day!",
		"Hi!",
		"There is no other option.",
		"Still alive after yesterday?",
	}
	return variants[rand.Intn(len(variants))]
}

func generateType() AlcoholType {
	rand.Seed(time.Now().UnixNano())
	variants := []AlcoholType{
		Wine,
		Beer,
		Vodka,
		Water,
		Moonshine,
	}
	return variants[rand.Intn(len(variants))]
}

func generateAmount(alcoholType AlcoholType) (result float64) {
	if alcoholType == Wine {
		result = generateTruncatedNormal(1.7, 1.4, 3.5, 0.2)
	} else if alcoholType == Beer {
		result = generateTruncatedNormal(3, 3, 6, 0.5)
	} else if alcoholType == Vodka {
		result = generateTruncatedNormal(1.5, 0.7, 3, 0.1)
	} else if alcoholType == Water {
		result = generateTruncatedNormal(2.5, 0.5, 3, 2)
	} else if alcoholType == Moonshine {
		result = generateTruncatedNormal(0.7, 0.5, 1.1, 0.4)
	}
	return result
}

func parseAmount(alcoholType AlcoholType, amount float64) string {
	fullBottles, frac := math.Modf(amount)

	glasses := 0

	if alcoholType == Wine || alcoholType == Water {
		for frac > 0 {
			glasses++
			frac = frac - 0.2
		}
	} else if alcoholType == Beer {
		for frac > 0 {
			glasses++
			frac = frac - 0.5
		}
	} else if alcoholType == Vodka || alcoholType == Moonshine {
		for frac > 0 {
			glasses++
			frac = frac - 0.1
		}
	}

	fullBottlesInt := int(fullBottles)

	if alcoholType == Wine || alcoholType == Water {
		if glasses == 5 {
			fullBottlesInt++
			glasses = 0
		}
	} else if alcoholType == Beer {
		if glasses == 2 {
			fullBottlesInt++
			glasses = 0
		}
	} else if alcoholType == Vodka || alcoholType == Moonshine {
		if glasses == 10 {
			fullBottlesInt++
			glasses = 0
		}
	}

	if alcoholType == Wine {
		return createAnswerString("🍾", "🍷", "wine", fullBottlesInt, glasses)
	} else if alcoholType == Beer {
		return createAnswerString("🍾", "🍺", "beer", fullBottlesInt, glasses)
	} else if alcoholType == Vodka {
		return createAnswerString("🍾", "🍸", "beer", fullBottlesInt, glasses)
	} else if alcoholType == Water {
		return createAnswerString("🍾", "🥃", "beer", fullBottlesInt, glasses)
	} else if alcoholType == Moonshine {
		return createAnswerString("🍾", "🥛", "beer", fullBottlesInt, glasses)
	}

	return ""
}

func createAnswerString(bigEmoji, smallEmoji, name string, fullBottles, glasses int) string {
	if fullBottles == 0 {
		if glasses == 1 {
			return fmt.Sprintf("Only %v of %v", smallEmoji, name)
		}
		return fmt.Sprintf("%v of %v", strings.Repeat(smallEmoji, glasses), name)
	} else if fullBottles == 1 {
		if glasses == 0 {
			return fmt.Sprintf("Whole %v of %v", bigEmoji, name)
		}
		return fmt.Sprintf("Whole %v and %v of %v", bigEmoji, strings.Repeat(smallEmoji, glasses), name)
	}
	if glasses == 0 {
		return fmt.Sprintf("%v of %v", strings.Repeat(bigEmoji, fullBottles), name)
	}
	return fmt.Sprintf("%v and %v of %v", strings.Repeat(bigEmoji, fullBottles), strings.Repeat(smallEmoji, glasses), name)
}
