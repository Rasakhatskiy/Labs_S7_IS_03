package main

var subjectsMap map[DNA]Subject
var subjectsList []Subject

var teachersMap map[DNA]Teacher
var teachersList []Teacher

var lessonTypesMap map[DNA]LessonType
var lessonTypesList []LessonType

var classroomsMap map[DNA]Classroom
var classroomsList []Classroom

var groupsMap map[DNA]Group
var groupsList []Group

var timeSlotsMap map[DNA]Timeslot
var timeSlotsList []Timeslot

var weekDaysMap map[DNA]Weekday
var weekDaysList []Weekday

func initValues() {
	subjectsMap = map[DNA]Subject{
		/*00*/ "0000": {Name: "Iнформаційні технології в менеджментi"},
		/*01*/ "0001": {Name: "Телекомунікаційні технології"},
		/*02*/ "0010": {Name: "Вибрані розділи трудового права і основ підприємницької діяльності"},
		/*03*/ "0011": {Name: "Основи Data Mining"},
		/*04*/ "0100": {Name: "Коректність програм та логіки програмування"},
		/*05*/ "0101": {Name: "Композиційна семантика SQL-подібних мов"},
		/*06*/ "0110": {Name: "Методи специфікації програм"},
		/*07*/ "0111": {Name: "Інтелектуальні системи"},
		/*08*/ "1000": {Name: "Розробка бізнес-аналітичних систем"},
		/*09*/ "1001": {Name: "Вступ до університетських студій"},
		/*10*/ "1010": {Name: "Коректність програм та логіки програмування"},
	}
	teachersMap = map[DNA]Teacher{
		/*00*/ "000": {Name: "Вергунова Ірина Миколаївна", Subjects: []Subject{subjectsMap["0000"]}},
		/*01*/ "001": {Name: "Колєнов Сергій Олександрович", Subjects: []Subject{subjectsMap["0001"], subjectsMap["0010"], subjectsMap["0100"]}},
		/*02*/ "010": {Name: "Нікітченко Микола Степанович", Subjects: []Subject{subjectsMap["0011"], subjectsMap["1010"]}},
		/*03*/ "011": {Name: "Панченко Тарас Володимирович", Subjects: []Subject{subjectsMap["0101"]}},
		/*04*/ "100": {Name: "Глибовець Микола Миколайович", Subjects: []Subject{subjectsMap["0110"], subjectsMap["1000"], subjectsMap["1001"]}},
		/*05*/ "101": {Name: "Федорус О.М.", Subjects: []Subject{subjectsMap["0111"]}},
		/*06*/ "110": {Name: "Яковлев В.О.", Subjects: []Subject{subjectsMap["0111"]}},
	}
	lessonTypesMap = map[DNA]LessonType{
		/*00*/ "0": {Type: Lecture},
		/*01*/ "1": {Type: Practice},
	}
	classroomsMap = map[DNA]Classroom{
		/*00*/ "000": {Name: "101", Seats: 30},
		/*01*/ "001": {Name: "102", Seats: 30},
		/*02*/ "010": {Name: "201", Seats: 30},
		/*03*/ "011": {Name: "202", Seats: 30},
		/*04*/ "100": {Name: "01", Seats: 90},
		/*05*/ "101": {Name: "02", Seats: 90},
		/*06*/ "110": {Name: "403", Seats: 90},
	}
	groupsMap = map[DNA]Group{
		/*00*/ "000": {Name: "ТК-41"},
		/*01*/ "001": {Name: "ТК-42"},
		/*02*/ "010": {Name: "ТТП-41"},
		/*03*/ "011": {Name: "ТТП-42"},
		/*04*/ "100": {Name: "МІ"},
		/*05*/ "101": {Name: "ІПС-41"},
		/*06*/ "110": {Name: "ІПС-42"},
	}
	timeSlotsMap = map[DNA]Timeslot{
		"00": {Position: 1},
		"01": {Position: 2},
		"10": {Position: 3},
		"11": {Position: 4},
	}
	weekDaysMap = map[DNA]Weekday{
		"000": {Day: 1},
		"001": {Day: 2},
		"010": {Day: 3},
		"011": {Day: 4},
		"100": {Day: 5},
	}

	for _, subject := range subjectsMap {
		subjectsList = append(subjectsList, subject)
	}
	for _, subject := range teachersMap {
		teachersList = append(teachersList, subject)
	}
	for _, subject := range lessonTypesMap {
		lessonTypesList = append(lessonTypesList, subject)
	}
	for _, subject := range classroomsMap {
		classroomsList = append(classroomsList, subject)
	}
	for _, subject := range groupsMap {
		groupsList = append(groupsList, subject)
	}
	for _, subject := range timeSlotsMap {
		timeSlotsList = append(timeSlotsList, subject)
	}
	for _, subject := range weekDaysMap {
		weekDaysList = append(weekDaysList, subject)
	}

}
