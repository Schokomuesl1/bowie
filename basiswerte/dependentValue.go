package basiswerte

import (
    "fmt"
)

type DependentValue interface {
    Register(DependingValue)
    NotifyValueChanged()
    Value() int
}

type DependingValue interface {
    Update()
}

type ValueType interface {
    Value() int
}

type dependenceStorage struct {
    dependsOnMe []DependingValue
}

func (s dependenceStorage) Register(d DependingValue) {
    // first check if d is already registered
    for _, v := range s.dependsOnMe {
        if d == v {
            return
        }
    }
    n := len(s.dependsOnMe)
    // possibly extend slice/buffer
    if n == cap(s.dependsOnMe) {
        newDependsOnMe := make([]DependingValue, len(s.dependsOnMe), 2*len(s.dependsOnMe)+1)
        copy(newDependsOnMe, s.dependsOnMe)
        s.dependsOnMe = newDependsOnMe
    }
    s.dependsOnMe = s.dependsOnMe[0 : n+1]
    s.dependsOnMe[n] = d
}

func (s dependenceStorage) NotifyValueChanged() {
    for _, v := range s.dependsOnMe {
        v.Update()
    }
}

type CalculatedDependentValue struct {
    Name          string
    BaseValues    []ValueType
    Mult          float32
    dependingOnMe dependenceStorage
}

func MakeCalculatedDependentValue(name string, mult float32, base []DependentValue) *CalculatedDependentValue {
    var cal CalculatedDependentValue
    cal.Name = name
    cal.Mult = mult
    if cal.Mult == 0 {
        cal.Mult = 1
    }
    cal.BaseValues = make([]ValueType, len(base))
    for i, v := range base {
        cal.BaseValues[i] = v
        v.Register(cal)
    }
    return &cal
}

func (c *CalculatedDependentValue) Value() int {
    var sum int
    for _, v := range c.BaseValues {
        sum += v.Value()
    }
    return int((float32(sum) * c.Mult) + 0.5)
}

func (c CalculatedDependentValue) Update() {
    c.dependingOnMe.NotifyValueChanged() // no re-calculation, this happens in Value(). Only notify we might have to be recalculated
}

func (c CalculatedDependentValue) String() string {
    return fmt.Sprintf("%s: %d", c.Name, c.Value())
}
