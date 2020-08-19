package aptitude

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
)

const basePCAptitudePoints = 6

type Aptitude int

const (
	Vigor Aptitude = iota
	Focus
	Agility
	Perception
	Empathy
	Cunning
)

func (d Aptitude) String() string {
	return [...]string{"Vigor", "Focus", "Agility", "Perception", "Empathy", "Cunning"}[d]
}

func AptitudeValue(s string) (Aptitude, error) {
	switch strings.ToLower(s) {
	case "vigor":
		return Vigor, nil
	case "focus":
		return Focus, nil
	case "agility":
		return Agility, nil
	case "perception":
		return Perception, nil
	case "empathy":
		return Empathy, nil
	case "cunning":
		return Cunning, nil
	}
	return 0, fmt.Errorf("not found")
}

type Aptitudes map[Aptitude]int

func New() Aptitudes {
	a := make(map[Aptitude]int, 6)
	for i := Aptitude(0); i < 6; i++ {
		a[i] = 1
	}
	return a
}

func (a Aptitudes) DefaultAptitudeList() {
	a = make(map[Aptitude]int, 6)
	for i := Aptitude(0); i < 6; i++ {
		a[i] = 1
	}
}

func (a Aptitudes) RandomAptitudes() int {
	if a.TotalAptitudePoints() == basePCAptitudePoints {
		for i := 0; i < 2; i++ {
			aptIndex := rand.Intn(6)
			if a[Aptitude(aptIndex)] == 1 {
				a[Aptitude(aptIndex)] = 2
			} else {
				i--
			}
		}
		for i := 0; i < 1; i++ {
			aptIndex := rand.Intn(6)
			if a[Aptitude(aptIndex)] == 1 {
				a[Aptitude(aptIndex)] = 3
			} else {
				i--
			}
		}
	} else {
		fmt.Println("not starting from base aptitudes")
	}

	return a.TotalAptitudePoints() - basePCAptitudePoints
}

func (a Aptitudes) TotalAptitudePoints() int {
	var total int
	for _, score := range a {
		for i := score; i > 0; i-- {
			total += CalculateAptitudePoints(i)
		}
	}
	return total
}

func CalculateAptitudePoints(i int) int {
	return int(math.Pow(float64(i), 2))
}

func (a Aptitudes) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(a)
	count := 0

	keys := make([]int, 0, len(a))
	for key := range a {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		buffer.WriteString(fmt.Sprintf("\"%s\":%d", strings.ToLower(Aptitude(key).String()), a[Aptitude(key)]))
		count++
		if count < length {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (a Aptitudes) UnmarshalJSON(b []byte) error {
	var aptitudes map[string]int
	err := json.Unmarshal(b, &aptitudes)
	if err != nil {
		return err
	}
	for aptitude, score := range aptitudes {
		aptitudeID, err := AptitudeValue(aptitude)
		if err != nil {
			return err
		}
		a[aptitudeID] = score
	}
	return nil
}
