package staff

import "log"

var OverpaidSalary = 70000
var UnderpaidSalary = 60000

// Starts with a capital letter to be exported
type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	logFun()
	var overpaid []Employee
	for _, emp := range e.AllStaff {
		if emp.Salary >= OverpaidSalary {
			overpaid = append(overpaid, emp)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	logFun()
	var underpaid []Employee
	for _, emp := range e.AllStaff {
		if emp.Salary <= UnderpaidSalary {
			underpaid = append(underpaid, emp)
		}
	}
	return underpaid
}

func (e *Office) notVisible() {
	log.Println("This function is not exported and cannot be accessed outside this package.")
}

func logFun(){
	log.Println("This is a log function")
	// This function is not exported and cannot be accessed outside this package.
	// It can be used internally within the package for logging purposes.
}