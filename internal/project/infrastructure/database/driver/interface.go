package driver

type (
	// Driver описывает обязательный функционал, которым должен обладать драйвер общения с б/д.
	Driver interface {
		Getter
		Selector
		Querier
		Close()       // Close закрывает соединение с б/д
		Name() string // Name возвращает строковый идентификатор Driver
	}

	// Getter получает строку вывода результата запроса и сканирует ее в dst.
	// Подходит для вызова запросов с одной строкой (как правило сущностью) в выводе.
	// Если результат вывода будет содержать больше одной строки, возвращает ошибку.
	Getter interface {
		Get(dst interface{}, query string, args ...interface{}) error
	}

	// Selector проходится по всем строкам результата запроса и сканирует их в dst.
	// Подходит для вызова запросов с множеством строк в выводе.
	Selector interface {
		Select(dst interface{}, query string, args ...interface{}) error
	}

	Querier interface {
		Query(query string, args ...interface{}) error
	}
)
