package canonical

type User struct {
	Id        string
	Document  string
	Name      string
	Email     string
	Password  string
	CreatedAt string
}
type Product struct {
	ID        string
	Name      string
	Desc      string
	Price     string
	Category  string
	Status    string
	ImagePath string
}

type Order struct {
	ID         string
	CustomerID string
	PaymentID  string
	Status     string
	CreatedAt  string
	UpdatedAt  string
	Total      string
	OrderItems []OrderItem
}

type OrderItem struct {
	Product  Product
	Quantity int
}
