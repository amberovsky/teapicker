package employee

import (
	"amberovsky/teapicker/pkg/employee"
	"github.com/google/uuid"
)

/*
In-memory storage implementation for employees
*/

type inMemoryRepository struct {
	employees map[string]*employee.Employee
}

func NewInMemoryRepository() *inMemoryRepository {
	employees := make(map[string]*employee.Employee)

	employees["5167fa62-72b3-4397-b69a-2c0bef9391d4"] = employee.NewEmployee(uuid.MustParse("5167fa62-72b3-4397-b69a-2c0bef9391d4"), "Tom Richardson")
	employees["eb3973c6-ed75-4a62-b9e8-4ba7226b5469"] = employee.NewEmployee(uuid.MustParse("eb3973c6-ed75-4a62-b9e8-4ba7226b5469"), "Adrian Shedden")
	employees["659a3d79-3ec3-42cc-8540-310a2599ad63"] = employee.NewEmployee(uuid.MustParse("659a3d79-3ec3-42cc-8540-310a2599ad63"), "Charlie Richardson")

	employees["77345581-08be-4b85-8805-57e42bd01f20"] = employee.NewEmployee(uuid.MustParse("77345581-08be-4b85-8805-57e42bd01f20"), "Anton")

	employees["0a2206ef-4f84-48a8-a7fc-18bae866a488"] = employee.NewEmployee(uuid.MustParse("0a2206ef-4f84-48a8-a7fc-18bae866a488"), "Titus Flavus")
	employees["b6625257-19eb-4f98-a05e-c0674db6d870"] = employee.NewEmployee(uuid.MustParse("b6625257-19eb-4f98-a05e-c0674db6d870"), "Aulus Regillensis")
	employees["63c5646c-e212-4d14-98e4-1b0979a1314f"] = employee.NewEmployee(uuid.MustParse("63c5646c-e212-4d14-98e4-1b0979a1314f"), "Manius Maximus")
	employees["872d921f-ad5a-4354-8564-4657178dda01"] = employee.NewEmployee(uuid.MustParse("872d921f-ad5a-4354-8564-4657178dda01"), "Gaius Mamercus")
	employees["b7977c1d-7528-4a12-8e47-044840ae5988"] = employee.NewEmployee(uuid.MustParse("b7977c1d-7528-4a12-8e47-044840ae5988"), "Lucius Cincinnatus")
	employees["9d0f24ac-f7c0-40b8-97e4-b3bdcb1bd49e"] = employee.NewEmployee(uuid.MustParse("9d0f24ac-f7c0-40b8-97e4-b3bdcb1bd49e"), "Mamercus Mamercinus")
	employees["a96dd9f6-8e45-4b69-ac32-a153f590c62c"] = employee.NewEmployee(uuid.MustParse("a96dd9f6-8e45-4b69-ac32-a153f590c62c"), "Quintus Fidenas")
	employees["8a86fa37-64c7-461a-b860-852041a1b9a5"] = employee.NewEmployee(uuid.MustParse("8a86fa37-64c7-461a-b860-852041a1b9a5"), "Aulus Tubertus")
	employees["216551a6-11d1-4d30-89fd-983fc8daf345"] = employee.NewEmployee(uuid.MustParse("216551a6-11d1-4d30-89fd-983fc8daf345"), "Publius Cossus")

	return &inMemoryRepository{employees: employees}
}

func (r *inMemoryRepository) GetAll() ([]*employee.Employee, error) {
	employees := []*employee.Employee{}
	for _, value := range r.employees {
		employees = append(employees, value)
	}

	return employees, nil
}

func (r *inMemoryRepository) Add(e *employee.Employee) error {
	r.employees[e.UUID.String()] = e
	return nil
}

func (r *inMemoryRepository) Delete(UUID uuid.UUID) error {
	delete(r.employees, UUID.String())
	return nil
}
