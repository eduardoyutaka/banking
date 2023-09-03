package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Eduardo", "Curitiba", "110011", "2000-01-01", "1"},
		{"1001", "Rob", "Curitiba", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
