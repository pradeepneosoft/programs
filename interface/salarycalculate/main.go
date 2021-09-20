package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}
type Permanent struct {
	Empid  int
	Salary int
	Pf     int
}
type Contract struct {
	Empid  int
	Salary int
}
type Freelance struct {
	PerHour int
	Hours   int
}

func (p Permanent) CalculateSalary() int {
	return p.Salary + p.Pf
}
func (c Contract) CalculateSalary() int {
	return c.Salary
}
func (f Freelance) CalculateSalary() int {
	return f.PerHour * f.Hours

}
func expenses(e []SalaryCalculator) {
	var total int
	for _, v := range e {
		total = v.CalculateSalary() + total
	}
	fmt.Println(total)
}

func main() {
	p := Permanent{Empid: 1, Salary: 25000, Pf: 300}
	c := Contract{Salary: 30000, Empid: 2}
	f := Freelance{PerHour: 25, Hours: 200}
	Sl := []SalaryCalculator{p, c, f}
	expenses(Sl)

}
