package configuration

const (
	EnvPrefix      = `wt`
	FullEnvPrefix  = EnvPrefix + "_"
	LayerSeparator = `_`

	// DevelopmentMode отвечает за запуск приложения в development режиме.
	DevelopmentMode        = `dev`
	DefaultDevelopmentMode = false

	// LogLayer Раздел настроек логирования
	LogLayer             = `log`
	LogLevel             = LogLayer + LayerSeparator + `level`
	DefaultLogLevel      = `info`
	LogMaxAge            = LogLayer + LayerSeparator + "maxAge"
	DefaultLogMaxAge     = 0
	LogMaxBackups        = LogLayer + LayerSeparator + "maxBackups"
	DefaultLogMaxBackups = 10
	LogMaxSize           = LogLayer + LayerSeparator + "maxSize"
	DefaultLogMaxSize    = 30

	// DirLayer Раздел настроек директорий
	DirLayer      = `dir`
	DirApp        = DirLayer + LayerSeparator + `app`
	DirBin        = DirLayer + LayerSeparator + `bin`
	DefaultDirBin = `bin`
	DirVar        = DirLayer + LayerSeparator + `var`
	DefaultDirVar = `var`
	DirLog        = DirLayer + LayerSeparator + `log`

	// FileLayer Раздел настроек файлов
	FileLayer          = `file`
	FileLogName        = FileLayer + LayerSeparator + `logName`
	DefaultFileLogName = `project.log`

	// HttpLayer Раздел настроек веб-сервера
	HttpLayer       = `http`
	HttpPort        = HttpLayer + LayerSeparator + `port`
	DefaultHttpPort = `80`
	HttpHost        = HttpLayer + LayerSeparator + `host`
	DefaultHttpHost = `localhost`

	// DatabaseLayer Раздел настроек базы данных
	DatabaseLayer = `database`
	DatabaseConn  = DatabaseLayer + LayerSeparator + `connection`

	// OpenWeatherLayer Раздел настроек API для сервиса OpenWeather
	OpenWeatherLayer      = `open_weather`
	OpenWeatherKey        = OpenWeatherLayer + LayerSeparator + `key`
	DefaultOpenWeatherKey = `981429fbfadfb6694e1ce1de10f128e7`
	OpenWeatherUrl        = OpenWeatherLayer + LayerSeparator + `url`
	DefaultOpenWeatherUrl = `http://api.openweathermap.org/`

	// SyncLayer Раздел настроек синхронизации
	SyncLayer                 = `sync`
	SyncGetCityWeatherTime    = SyncLayer + LayerSeparator + `getCityWeatherTime`
	DefaultGetCityWeatherTime = 1
	SyncСntDayArchiveTime     = SyncLayer + LayerSeparator + `cntDayArchiveTime`
	DefaultСntDayArchiveTime  = 1

	// WeatherLayer Раздел настроек для работы с данными по погоде
	WeatherLayer         = ``
	СntDayArchive        = `cntDayArchive`
	DefaultCntDayArchive = 1
)
