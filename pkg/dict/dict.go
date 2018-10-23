package dict

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
	_ = iota
	// Processing means "В обработке"
	Processing Status = iota
	// Failed means "Ошибка при записи"
	Failed
	// Inserted means "Новая запись"
	Inserted
	// Updated means "Запись обновлена"
	Updated
	// Deleted means "Запись удалена из реестра"
	Deleted
	// ParsingFailed means "Ошибка при чтении веб-страницы"
	ParsingFailed
)

const (
	// RegistryNameLabel refers to "Торговое название"
	RegistryNameLabel = "Торговое название"
	// RegistryNumberLabel refers to "Регистрационный номер"
	RegistryNumberLabel = "Регистрационный номер"
	// RegistryIssueDateLabel refers to "Дата регистрации"
	RegistryIssueDateLabel = "Дата регистрации"
	// RegistryDurationLabel refers to "Срок регистрации"
	RegistryDurationLabel = "Срок регистрации"
	// RegistryExpireDateLabel refers to "Дата истечения"
	RegistryExpireDateLabel = "Дата истечения"
	// RegistryParsingLabel refers to "Обновление данных"
	RegistryParsingLabel = "Обновление данных"
)

const (
	// ProcessRunning refers to "Запущен"
	ProcessRunning = 1
	// ProcessFailed refers to "Ошибка"
	ProcessFailed = 2
	// ProcessFinished refers to "Завершен"
	ProcessFinished = 3
)
