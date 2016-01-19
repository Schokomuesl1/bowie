package bowie

import (
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

type DiceType interface {
    Roll() int32
}

type NormalDice struct {
    Sides int32
}

func (d NormalDice) Roll() int32 {
    return rand.Int31n(d.Sides) + 1
}
