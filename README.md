# weatherAPI

### Введение:
**weatherAPI** - сервис, написанный на Golang, представляющий из себя небольшой набор функционала, который позволяет получать информацию о погоде, для разных городов. 
Информация о погоде берется с сайта:
https://openweathermap.org/api  
Для реализации cli оболочки используется библиотека https://github.com/kardianos/service  
Для реализации веб интерфейсов используется фреймворк https://echo.labstack.com/


### Архитектура проекта
При разработке архитектуры был использован следующий шаблон:  
https://github.com/golang-standards/project-layout


### Реализованный функционал:

#### API:

**[GET]**

*/api/v1/weather/city* - получить статистику по списку подписки (сколько и какие города мониторятся)

**[GET]**

*/api/v1/weather/city/{name}* - получить среднюю температуру за последние дни

**[POST]**

*/api/v1/weather/city/{name}* - добавить город в подписки

**[DELETE]**

*/api/v1/weather/city/{name}* - удалить город из подписок

#### Фоновые задачи:

*Записать в кеш текущую температуру по каждому городу*
- Фоновая служба запускается с определенной периодичностью и кеширует в бд текущую температуру каждого города из подписки.  
Для установки частоты запуска фоновой службы используется переменная окружения "getCityWeatherTime". Частота запуска измеряется в минутах.  

*Записать в кеш среднюю температуру по каждому городу*
- Фоновая служба запускается с определенной периодичностью и кеширует в бд среднюю температуру каждого города из подписки.  
Средняя температура высчитывается за последние дни.  
Кол-во дней, за которые считается средняя температура, можно указать в переменной окружения "cntDayArchive".
Для установки частоты запуска фоновой службы используется переменная окружения "cntDayArchiveTime". Частота запуска измеряется в днях.  

#### Фитчи:

- Добавлен функционал, который выводит в консоль переданные сообщения, выполняется в фоновом режиме.  
В данном проекте, функционал используется для протоколирования текущей погоды каждого города из подписки.

- Так же в проекте используется логирование. В качестве библиотеки используется zap.Logger. Логи лежат в папке var/log.

#### Конфигурация проекта.
Для чтения конфигурации проекта используется https://github.com/spf13/viper.  
Конфиги читаются из env, если конфигов нет, то устанавливаются дефолтные, посмотреть которые можно в пакете **/internal/project/domain/configuration/**

#### Описание пакетов internal
##### interfaces  
Здесь хранится описание интерфейсов, позволяющих взаимодействовать с программой.
cli оболочка https://github.com/kardianos/service
web https://echo.labstack.com/
##### infrastructure
Здесь хранится реализация зависимостей (сервисов), которые использует приложение.
##### domain
Здесь лежат структуры, константы и т.д. Используемые по всему проекту
##### core
Прокладка между пакетом вызывающим метод, какого либо сервиса и этим сервисом соответственно. 

#### Документация REST
Для спецификации API используется Swagger.  
Документация доступна по следующему пути:  
http://localhost/docs/index.html

#### Описание остальных корневых пакетов проекта

##### docs
- Здесь хранятся сгенерированные файлы документации  
##### pkg
- Здесь хранятся дополнительные, не объемные функции, используемые в проекте
##### schema
- Здесь лежат файлы миграции
##### var
- Логи проекта
