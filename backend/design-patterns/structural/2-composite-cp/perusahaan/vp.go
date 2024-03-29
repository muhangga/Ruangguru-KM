package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	total := vp.GetSalary()
	for _, employee := range vp.Subordinate {
		total += employee.TotalDivisonSalary()
	}
	return total
}
