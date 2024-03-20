package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	bonus := 1 * float64(j.BaseSalary) * float64(j.WorkingMonth) / 12
	return bonus
}
func (s Senior) GetBonus() float64 {
	bonus := 2*float64(s.BaseSalary)*float64(s.WorkingMonth)/12 + (s.PerformanceRate * float64(s.BaseSalary))
	return bonus
}
func (m Manager) GetBonus() float64 {
	bonus := 2*float64(m.BaseSalary)*float64(m.WorkingMonth)/12 + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
	return bonus
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0

	for _, employee := range employees {
		sum += employee.GetBonus()
	}
	return sum
}

func main() {
	rangga := Senior{
		Name:            "Rangga",
		BaseSalary:      1000000,
		WorkingMonth:    36,
		PerformanceRate: 0.5,
	}

	fmt.Println(int(EmployeeBonus(rangga)))

}
