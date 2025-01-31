# teacher_schedule
DATABASE create info:
1) psql -U postgres
2) \l
3) CREATE DATABASE teacher_schedule;
4) \l
5) psql -U postgres -d teacher_schedule

Endpoints:
1) / 
Get - return "АПИ Расписание Преподавателя"
2) /users
Get / - get all users

Get /teahchers - get users with role teacher

Post /register - add new user. 
Correct request example:
{
  "firstName": "Препод2222",
  "lastName": "Фамилия",
  "patronymic": "Отчество",
  "birthDate": "2020-12-12",
  "roleid": 2
}

Succesfull create: 201. 
{"ID":7,"FirstName":"Препод2222","LastName":"Фамилия","Patronymic":"Отчество","BirthDate":"2020-12-12","RoleID":2}

Incorrect request example:
{
  "firstName": "Препод2222",
  "lastName": "Фамилия",
  "patronymic": "Отчество",
  "birthDate": "2020-12-12",
  "roleid": "2"
}
400 Bad Request.
Неверный ввод

3) /schedule
GET / - get info about all schedule.

POST / - create info about schedule.
Correct request example:
{
  "teacherId": 1,
  "dayId": 1,
  "slotId": 1,
  "groupId": 1,
  "subjectId": 1,
  "classroomId": 1,
  "isOccupied": true
}
201 Created. {"ID":3,"TeacherID":1,"DayID":1,"SlotID":1,"GroupID":1,"SubjectID":1,"IsOccupied":true,"ClassroomID":1}
Incorrect request example:
{
  "teacherId": 1,
  "dayId": 1,
  "slotId": 1,
  "groupId": 1,
  "subjectId": 1,
  "classroomId": "1",
  "isOccupied": true
}
400 Bad Request
Неверный запрос

GET /{id} - get info about selected schedule
/2 - 200 OK {"ID":2,"TeacherID":2,"DayID":1,"SlotID":1,"GroupID":10,"SubjectID":10,"IsOccupied":false,"ClassroomID":10}
/222 - 404 Not Found Расписание не найдено
/sss - 500 Internal Server Error ОШИБКА: столбец "sss" не существует (SQLSTATE 42703)

PUT /{id} - edit selected schedule info
Correct request example:
{
  "groupId": 10,
  "subjectId": 10,
  "classroomId": 10,
  "isOccupied": false
}
200 OK 
{"ID":2,"TeacherID":2,"DayID":1,"SlotID":1,"GroupID":10,"SubjectID":10,"IsOccupied":false,"ClassroomID":10}
Incorrect request example:
{
  "groupId": "10",
  "subjectId": 10,
  "classroomId": 10,
  "isOccupied": false
}
400 Bad Request Неверное тело запроса

GET /teacher/{teacherId} - get schedule selected teacher
200 OK - [
    {
        "ID": 2,
        "TeacherID": 2,
        "DayID": 1,
        "SlotID": 1,
        "GroupID": 10,
        "SubjectID": 10,
        "IsOccupied": false,
        "ClassroomID": 10
    }
]
404 Not Found - Для данного преподавателя нет расписания
400 Bad Request - Нверный ID преподавтеля

GET /filter - get filteres schedule info
FIlter params:
dayId - schedule info by selected day
groupId - schedule info by selected group
subjectId - schedule info with selected subject
isOccupied - info about slot schedule
teacherId - info about schedule with selected teacher
Exapmple: 
/schedule/filter?teacherId=2&groupId=10
200 OK [{"ID":2,"TeacherID":2,"DayID":1,"SlotID":1,"GroupID":10,"SubjectID":10,"IsOccupied":false,"ClassroomID":10}]

POST /bulk-add - add bulk schedules
Correct request example:
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
200 OK [{"ID":4,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":1,"SubjectID":101,"IsOccupied":true,"ClassroomID":0},{"ID":5,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":1,"SubjectID":102,"IsOccupied":false,"ClassroomID":0},{"ID":6,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":2,"SubjectID":101,"IsOccupied":true,"ClassroomID":0},{"ID":7,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":3,"SubjectID":103,"IsOccupied":true,"ClassroomID":0},{"ID":8,"TeacherID":0,"DayID":0,"SlotID":0,"GroupID":3,"SubjectID":104,"IsOccupied":false,"ClassroomID":0}]

DELETE /{id} - delete selected schedule
204 No Content

4) /subjects
GET / - get info about subjects

POST / - create info about subject
Request example: 
{
    "subjectName":"math"
}
201 Created {"SubjectID":1,"SubjectName":"math"}

5) /group 
GET / - get info about group

POST / - create info about group
Request example:
{
  "groupNumber": "group_5",
  "courseNumber": 5
}
201 Created {"GroupID":5,"GroupNumber":"group_5","CourseNumber":5}

6) /attendance
GET / - get all attendances

POST / - add info about attendance
Request example:
{
  "studentId": 1,
  "scheduleId": 1,
  "isPresent": true
}
201 Created {"AttendanceID":1,"StudentID":1,"ScheduleID":1,"IsPresent":true}

POST /bulk-add - bulk add info about attendances
Request example:
[
    {
        "StudentID": 1,
        "ScheduleID": 10,
        "IsPresent": true
    },
    {
        "StudentID": 2,
        "ScheduleID": 10,
        "IsPresent": false
    },
    {
        "StudentID": 3,
        "ScheduleID": 11,
        "IsPresent": true
    },
    {
        "StudentID": 4,
        "ScheduleID": 12,
        "IsPresent": false
    },
    {
        "StudentID": 5,
        "ScheduleID": 12,
        "IsPresent": true
    }
]
200 OK [{"AttendanceID":8,"StudentID":1,"ScheduleID":10,"IsPresent":true},{"AttendanceID":9,"StudentID":2,"ScheduleID":10,"IsPresent":false},{"AttendanceID":10,"StudentID":3,"ScheduleID":11,"IsPresent":true},{"AttendanceID":11,"StudentID":4,"ScheduleID":12,"IsPresent":false},{"AttendanceID":12,"StudentID":5,"ScheduleID":12,"IsPresent":true}]

GET /student/{studentID} - get attendances selected student
Response example:
200 OK 
[
    {
        "AttendanceID": 1,
        "StudentID": 1,
        "ScheduleID": 1,
        "IsPresent": true
    },
    {
        "AttendanceID": 2,
        "StudentID": 1,
        "ScheduleID": 2,
        "IsPresent": false
    },
    {
        "AttendanceID": 3,
        "StudentID": 1,
        "ScheduleID": 10,
        "IsPresent": true
    },
    {
        "AttendanceID": 8,
        "StudentID": 1,
        "ScheduleID": 10,
        "IsPresent": true
    }
]

GET /schedule/{scheduleId} - get attendances by selected schedule
Response example: 
200 OK [
    {
        "AttendanceID": 3,
        "StudentID": 1,
        "ScheduleID": 10,
        "IsPresent": true
    },
    {
        "AttendanceID": 4,
        "StudentID": 2,
        "ScheduleID": 10,
        "IsPresent": false
    },
    {
        "AttendanceID": 8,
        "StudentID": 1,
        "ScheduleID": 10,
        "IsPresent": true
    },
    {
        "AttendanceID": 9,
        "StudentID": 2,
        "ScheduleID": 10,
        "IsPresent": false
    }
]

PUT /{id} - edit selected attendance
Request example:
{
    "AttendanceID": 3,
    "StudentID": 10,
    "ScheduleID": 100,
    "IsPresent": false
}

DELETE /{id} - delete seleceted attendance
DELETE /bulk-delete - bulk delete attendances
Request example:
[5,6,7,8,9,10]
200 OK{
    "count": 6,
    "message": "Посещаемости удалены успешно"
}