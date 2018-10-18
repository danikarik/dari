package parser

import "time"

const (
	// MainTableIndex refers to "Основные данные"
	MainTableIndex = 0
	// ManufacturesTableIndex refers to "Производители"
	ManufacturesTableIndex = 2
	// PartsTableIndex refers to "Комплектность"
	PartsTableIndex = 3
	// TableHeaderIndex refers to table header row
	TableHeaderIndex = 0
)

const (
	// RegistryName refers to "Торговое название"
	RegistryName = 0
	// RegistryNumber refers to "Регистрационные номер"
	RegistryNumber = 1
	// IssueDate refers to "Дата регистрации"
	IssueDate = 2
	// RegistryDuration refers to "Срок регистрации"
	RegistryDuration = 3
	// ExpireDate refers to "Дата истечения"
	ExpireDate = 4
)

const (
	// ManufacturerName refers to "Наименование"
	ManufacturerName = 0
	// Country refers to "Страна"
	Country = 1
	// ManufacturerType refers to "Тип производителя"
	ManufacturerType = 2
	// ManufacturerRequired refers to "Производитель"
	ManufacturerRequired = "Производитель"
)

const (
	// Order refers to "№"
	Order = 0
	// PartName refers to "Наименование изделия"
	PartName = 1
	// Model refers to "Модель"
	Model = 2
)

// Status specifies result of action.
type Status int

const (
	scraped Status = iota
	failed
	added
	updated
	deleted
)

// Registry refers to "reestr" from dari.kz.
type Registry struct {
	Number        string
	IssueDate     time.Time
	ExpireDate    time.Time
	Duration      int
	Title         string
	Link          string
	Manufacturers []Manufacturer
	Builds        []Build
	Status        Status
}

// Manufacturer contains information about manufacturer.
type Manufacturer struct {
	Title   string
	Country string
	Type    string
}

// Build contains information about equipment part.
type Build struct {
	Order int
	Title string
	Model string
}
