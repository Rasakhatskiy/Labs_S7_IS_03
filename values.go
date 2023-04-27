package main

var subjectsMap map[string]*Subject
var subjectsList []Subject

var teachersMap map[string]*Teacher
var teachersList []Teacher

var lessonTypesMap map[string]*LessonType
var lessonTypesList []LessonType

var classroomsMap map[string]*Classroom
var classroomsList []Classroom

var groupsMap map[string]*Group
var groupsList []Group

var timeSlotsMap map[string]*Timeslot
var timeSlotsList []Timeslot

var weekDaysMap map[string]*Weekday
var weekDaysList []Weekday

var validSubTeach []string

const lessonsInTable = 20
const starterSet = 1000

const SubjectLen = 4
const TeachersLen = 3
const LessonTypeLen = 1
const ClassroomsLen = 4
const GroupsLen = 3
const TimeSlotsLen = 3
const WeekDaysLen = 3

const SingleGenomeLength = SubjectLen + TeachersLen + LessonTypeLen + ClassroomsLen + GroupsLen + TimeSlotsLen + WeekDaysLen
const TotalGenomeLength = SingleGenomeLength * lessonsInTable

var SplitPositions = []int{
	SubjectLen,
	SubjectLen + TeachersLen,
	SubjectLen + TeachersLen + LessonTypeLen,
	SubjectLen + TeachersLen + LessonTypeLen + ClassroomsLen,
	SubjectLen + TeachersLen + LessonTypeLen + ClassroomsLen + GroupsLen,
	SubjectLen + TeachersLen + LessonTypeLen + ClassroomsLen + GroupsLen + TimeSlotsLen,
}

func initValues() {
	subjectsMap = map[string]*Subject{
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

		/*11*/ "1011": {Name: "Розробка ПЗ під мобільні платформи"},
		/*12*/ "1100": {Name: "Теорія прийняття рішень"},
		/*13*/ "1101": {Name: "Нейронні мережі"},
		/*14*/ "1110": {Name: "Об'єктно-орієнтоване програмування"},
		/*15*/ "1111": {Name: "Розподілене та паралельне програмування"},
	}
	teachersMap = map[string]*Teacher{
		/*00*/ "000": {Name: "Вергунова І.М. ", Subjects: []Subject{{Name: (subjectsMap["0000"]).Name, Dna: "0000"}}},
		/*01*/ "001": {Name: "Колєнов С.О.   ", Subjects: []Subject{{Name: (subjectsMap["0001"]).Name, Dna: "0001"}, {Name: (subjectsMap["0010"]).Name, Dna: "0010"}, {Name: (subjectsMap["0100"]).Name, Dna: "0100"}}},
		/*02*/ "010": {Name: "Нікітченко М.С.", Subjects: []Subject{{Name: (subjectsMap["0011"]).Name, Dna: "0011"}, {Name: (subjectsMap["1100"]).Name, Dna: "1100"}}},
		/*03*/ "011": {Name: "Панченко Т.В.  ", Subjects: []Subject{{Name: (subjectsMap["0101"]).Name, Dna: "0101"}, {Name: (subjectsMap["1101"]).Name, Dna: "1101"}}},
		/*04*/ "100": {Name: "Глибовець М.М. ", Subjects: []Subject{{Name: (subjectsMap["0110"]).Name, Dna: "0110"}, {Name: (subjectsMap["1000"]).Name, Dna: "1000"}, {Name: (subjectsMap["1001"]).Name, Dna: "1001"}}},
		/*05*/ "101": {Name: "Федорус О.М.   ", Subjects: []Subject{{Name: (subjectsMap["0111"]).Name, Dna: "0111"}, {Name: (subjectsMap["1110"]).Name, Dna: "1110"}}},
		/*06*/ "110": {Name: "Яковлев В.О.   ", Subjects: []Subject{{Name: (subjectsMap["0111"]).Name, Dna: "0111"}, {Name: (subjectsMap["1111"]).Name, Dna: "1111"}}},
		/*07*/ "111": {Name: "Ткаченко О.М.  ", Subjects: []Subject{{Name: (subjectsMap["1011"]).Name, Dna: "1011"}, {Name: (subjectsMap["1010"]).Name, Dna: "1010"}}},
	}
	lessonTypesMap = map[string]*LessonType{
		/*00*/ "0": {Type: Lecture},
		/*01*/ "1": {Type: Practice},
	}
	classroomsMap = map[string]*Classroom{
		/*00*/ "0000": {Name: "101", Seats: 30},
		/*01*/ "0001": {Name: "102", Seats: 30},
		/*02*/ "0010": {Name: "201", Seats: 30},
		/*03*/ "0011": {Name: "202", Seats: 30},
		/*04*/ "0100": {Name: "203", Seats: 30},
		/*05*/ "0101": {Name: "204", Seats: 30},
		/*06*/ "0110": {Name: "205", Seats: 30},
		/*07*/ "0111": {Name: "206", Seats: 30},
		/*08*/ "1000": {Name: "207", Seats: 90},
		/*09*/ "1001": {Name: "01 ", Seats: 90},
		/*10*/ "1010": {Name: "02 ", Seats: 90},
		/*11*/ "1011": {Name: "403", Seats: 90},
		/*12*/ "1100": {Name: "03 ", Seats: 90},
		/*13*/ "1101": {Name: "04 ", Seats: 90},
		/*14*/ "1110": {Name: "05 ", Seats: 90},
		/*15*/ "1111": {Name: "06 ", Seats: 90},
	}
	groupsMap = map[string]*Group{
		/*00*/ "000": {Name: "ТК-41 "},
		/*01*/ "001": {Name: "ТК-42 "},
		/*02*/ "010": {Name: "ТТП-41"},
		/*03*/ "011": {Name: "ТТП-42"},
		/*04*/ "100": {Name: "МІ    "},
		/*05*/ "101": {Name: "ІПС-41"},
		/*06*/ "110": {Name: "ІПС-42"},
		/*07*/ "111": {Name: "ДО    "},
	}
	timeSlotsMap = map[string]*Timeslot{
		"000": {Position: 1},
		"001": {Position: 2},
		"010": {Position: 3},
		"011": {Position: 4},
		"100": {Position: 5},
		"101": {Position: 6},
		"110": {Position: 7},
		"111": {Position: 8},
	}
	weekDaysMap = map[string]*Weekday{
		"000": {Day: 1},
		"001": {Day: 2},
		"010": {Day: 3},
		"011": {Day: 4},
		"100": {Day: 5},
		"101": {Day: 6},
		"110": {Day: 0},
		"111": {Day: 0},
	}

	for dna, subject := range subjectsMap {
		subject.Dna = dna
		subjectsMap[dna].Dna = dna
		subjectsList = append(subjectsList, *subject)
	}
	for dna, subject := range teachersMap {
		subject.Dna = dna
		teachersMap[dna].Dna = dna
		teachersList = append(teachersList, *subject)
	}
	for dna, subject := range lessonTypesMap {
		subject.Dna = dna
		lessonTypesMap[dna].Dna = dna
		lessonTypesList = append(lessonTypesList, *subject)
	}
	for dna, subject := range classroomsMap {
		subject.Dna = dna
		classroomsMap[dna].Dna = dna
		classroomsList = append(classroomsList, Classroom{
			Name:  classroomsMap[dna].Name,
			Seats: classroomsMap[dna].Seats,
			Dna:   classroomsMap[dna].Dna,
		})
	}
	for dna, subject := range groupsMap {
		subject.Dna = dna
		groupsMap[dna].Dna = dna
		groupsList = append(groupsList, *subject)
	}
	for dna, subject := range timeSlotsMap {
		subject.Dna = dna
		timeSlotsMap[dna].Dna = dna
		timeSlotsList = append(timeSlotsList, *subject)
	}
	for dna, subject := range weekDaysMap {
		subject.Dna = dna
		weekDaysMap[dna].Dna = dna
		weekDaysList = append(weekDaysList, *subject)
	}

	validSubTeach = genValidTeacherSubjectPairs()
}
