package commonNames

import (
	"errors"
	"fmt"
	"math/rand"
)

// names lista de nombre
var names = map[int]string{
	1:  "Filemón",
	2:  "Gertrudis",
	3:  "Epifanio",
	4:  "Herculano",
	5:  "Pascualina",
	6:  "Eleuterio",
	7:  "Desiderio",
	8:  "Avelina",
	9:  "Eulogio",
	10: "Gaudencio",
}

// lastNames lista de apellidos
var lastNames = map[int]string{
	1:  "Delgadillo",
	2:  "Pájaro",
	3:  "Pimiento",
	4:  "Pitufo",
	5:  "Raposo",
	6:  "Fartachón",
	7:  "Tocino",
	8:  "Troncoso",
	9:  "Chicharrón",
	10: "Cachirulo",
}

/*NewName Agrega un nuevo nombre*/
func NewName(name string) error {
	if isValid(name, names) {
		names[len(names)] = name
		fmt.Printf("Name added correctly: %s\n", names[len(lastNames)])
		return nil
	}
	return errors.New("invalid name")
}

/*NewLastName Agrega un nuevo apellido*/
func NewLastName(lastName string) error {
	if isValid(lastName, lastNames) {
		lastNames[len(lastNames)] = lastName
		fmt.Printf("Last name added correctly: %s\n", lastNames[len(lastNames)])
		return nil
	}
	return errors.New("invalid last name")
}

/*RandomName retorna un nombre al azar*/
func RandomName() string {
	return names[rand.Intn(len(names))]
}

/*RandomLastName retorna un apellido al azar*/
func RandomLastName() string {
	return lastNames[rand.Intn(len(lastNames))]
}

/*RandomNameAndLastName retorna un nombre y apellido al azar*/
func RandomNameAndLastName() string {
	return fmt.Sprintf("%s %s", RandomName(), RandomLastName())
}

/*RandomNamesAndLastNames retorna dos nombres y dos apellidos al azar*/
func RandomNamesAndLastNames() string {
	return fmt.Sprintf("%s %s %s %s", RandomName(), RandomName(), RandomLastName(), RandomLastName())
}

func isValid(valueName string, mapa map[int]string) bool {
	for _, value := range mapa {
		if value == valueName {
			return false
		}
	}
	return true
}
