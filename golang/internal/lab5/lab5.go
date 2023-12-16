package lab5

import (
	"fmt"
)

type dog struct {
	age   int
	name  string
	breed string
}

func NewDog(ageVar int, nameVar, breedVar string) (dog, error) {
	var c dog = dog{
		name:  nameVar,
		breed: breedVar,
	}
	var err = c.SetAge(ageVar)
	return c, err
}

func (c dog) GetAge() int {
	return c.age
}

func (c *dog) SetAge(age int) error {
	if age >= 0 && age <= 25 {
		c.age = age
		return nil
	}
	return fmt.Errorf("cat named \"%s\" has invalid age", c.GetName())
}

func (c dog) GetName() string {
	return c.name
}

func (c *dog) SetName(name string) {
	c.name = name
}

func (c dog) GetBreed() string {
	return c.breed
}

func (c dog) PetThedog() {
	fmt.Printf("Dog %s says WOOF!\n", c.GetName())
}

//АААаааАаААаААаА
