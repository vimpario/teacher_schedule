# Teacher Schedule API

## Описание
Teacher Schedule API предназначен для управления расписанием преподавателей, пользователей, посещаемостью, группами и предметами.

---

## Установка базы данных
```bash
psql -U postgres
\l
CREATE DATABASE teacher_schedule;
\l
psql -U postgres -d teacher_schedule
```

---

## Установка
1. Установите Go
2. Клонируйте репозиторий
3. В корне проекта создайте .env файл со следующими данными:
- `DB_HOST=` хост базы данных, по умолчанию localhost
- `DB_PORT=` порт базы данных, по умолчанию 5432
- `DB_USER=`имя пользователя базы данных
- `DB_PASSWORD=`пароль базы данных
- `DB_NAME=`название базы данных


## Запуск проекта
Находясь в корне проекта, выполните команду:
```sh
go run cmd/main.go
```

## Эндпоинты

### Главная страница
`GET /`  
Ответ: `"АПИ Расписание Преподавателя"`

---

### Пользователи

#### Получение всех пользователей
`GET /users`

#### Получение всех преподавателей
`GET /users/teachers`

#### Регистрация нового пользователя
`POST /users/register`
##### Пример запроса:
```json
{
  "firstName": "Препод2222",
  "lastName": "Фамилия",
  "patronymic": "Отчество",
  "birthDate": "2020-12-12",
  "roleid": 2
}
```
##### Ответ:
201 Created
```json
{"ID":7,"FirstName":"Препод2222","LastName":"Фамилия","Patronymic":"Отчество","BirthDate":"2020-12-12","RoleID":2}
```

---

### Расписание

#### Получение всех расписаний
`GET /schedule`

#### Создание нового расписания
`POST /schedule`
##### Пример запроса:
```json
{
  "teacherId": 1,
  "dayId": 1,
  "slotId": 1,
  "groupId": 1,
  "subjectId": 1,
  "classroomId": 1,
  "isOccupied": true
}
```
##### Ответ:
201 Created
```json
{"ID":3,"TeacherID":1,"DayID":1,"SlotID":1,"GroupID":1,"SubjectID":1,"IsOccupied":true,"ClassroomID":1}
```

#### Массовое создание новых расписаний
`POST /schedule/bulk-add`
##### Пример запроса:
```json
[
  {
    "groupId": 1,
    "subjectId": 101,
    "isOccupied": true
  },
  {
    "groupId": 1,
    "subjectId": 102,
    "isOccupied": false
  },
  {
    "groupId": 2,
    "subjectId": 101,
    "isOccupied": true
  },
  {
    "groupId": 3,
    "subjectId": 103,
    "isOccupied": true
  },
  {
    "groupId": 3,
    "subjectId": 104,
    "isOccupied": false
  }
]
```
##### Ответ:
200 OK 
```json
[{"ID":4,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":1,"SubjectID":101,"IsOccupied":true,"ClassroomID":0},{"ID":5,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":1,"SubjectID":102,"IsOccupied":false,"ClassroomID":0},{"ID":6,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":2,"SubjectID":101,"IsOccupied":true,"ClassroomID":0},{"ID":7,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":3,"SubjectID":103,"IsOccupied":true,"ClassroomID":0},{"ID":8,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":3,"SubjectID":104,"IsOccupied":false,"ClassroomID":0}]
```

#### Получение расписания по ID
`GET /schedule/{id}`

#### Изменение расписания
`PUT /schedule/{id}`
##### Пример запроса:
```json
{
  "groupId": 10,
  "subjectId": 10,
  "classroomId": 10,
  "isOccupied": false
}
```
##### Ответ:
200 OK
```json
{"ID":2,"TeacherID":2,"DayID":1,"SlotID":1,"GroupID":10,"SubjectID":10,"IsOccupied":false,"ClassroomID":10}
```

#### Получение расписания преподавателя
`GET /schedule/teacher/{teacherId}`

#### Фильтрация расписания
`GET /schedule/filter`
##### Параметры фильтрации:
- `dayId` - расписание по выбранному дню
- `groupId` - расписание для группы
- `subjectId` - расписание по предмету
- `isOccupied` - занятость слота
- `teacherId` - расписание конкретного преподавателя

##### Пример запроса:
`GET /schedule/filter?teacherId=2&groupId=10`

---

### Предметы

#### Получение всех предметов
`GET /subjects`

#### Добавление нового предмета
`POST /subjects`
##### Пример запроса:
```json
{
    "subjectName":"math"
}
```
##### Ответ:
201 Created
```json
{"SubjectID":1,"SubjectName":"math"}
```

---

### Группы

#### Получение всех групп
`GET /group`

#### Создание группы
`POST /group`
##### Пример запроса:
```json
{
  "groupNumber": "group_5",
  "courseNumber": 5
}
```
##### Ответ:
201 Created
```json
{"GroupID":5,"GroupNumber":"group_5","CourseNumber":5}
```

---

### Посещаемость

#### Получение всех посещаемостей
`GET /attendance`

#### Добавление посещаемости
`POST /attendance`
##### Пример запроса:
```json
{
  "studentId": 1,
  "scheduleId": 1,
  "isPresent": true
}
```
##### Ответ:
201 Created
```json
{"AttendanceID":1,"StudentID":1,"ScheduleID":1,"IsPresent":true}
```

#### Групповое добавление посещаемости
`POST /attendance/bulk-add`
##### Пример запроса:
```json
[
    { "StudentID": 1, "ScheduleID": 10, "IsPresent": true },
    { "StudentID": 2, "ScheduleID": 10, "IsPresent": false }
]
```
##### Ответ:
200 OK
```json
[{"AttendanceID":8,"StudentID":1,"ScheduleID":10,"IsPresent":true},
{"AttendanceID":9,"StudentID":2,"ScheduleID":10,"IsPresent":false}]
```

#### Получение посещаемости студента
`GET /attendance/student/{studentID}`

#### Получение посещаемости по расписанию
`GET /attendance/schedule/{scheduleId}`

#### Редактирование посещаемости
`PUT /attendance/{id}`
##### Пример запроса:
```json
{
    "AttendanceID": 3,
    "StudentID": 10,
    "ScheduleID": 100,
    "IsPresent": false
}
```

#### Удаление расписания
`DELETE /schedule/{id}`

---

## Ошибки
- **400 Bad Request** – Некорректный запрос
- **404 Not Found** – Данные не найдены
- **500 Internal Server Error** – Ошибка сервера

---

