package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

}

const (
	Lecture  string = "lecture"
	Practice        = "practice"
)

type DNA string

// 4
type Subject struct {
	Name string
}

// 3
type Teacher struct {
	Subjects []Subject
	Name     string
}

type LessonType struct {
	Type string
}

type Group struct {
	Name string
}

type Weekday struct {
	Day time.Weekday
}

type Classroom struct {
	Name  string
	Seats int
}

type TimeSlot struct {
	Position int
}

type Lesson struct {
	Subject   Subject
	Teacher   Teacher
	Type      LessonType
	Group     Group
	Classroom Classroom
	TimeSlot  TimeSlot
	Weekday   Weekday
}

func CreateRandomChromosome(number int) []string {
	subList
}

//func (l *Lesson) Chromosome() string {
//	return fmt.Sprintf("%s%s%s%s%s%s",
//		l.Subject.Dna,
//		l.Teacher.Dna,
//		l.Type.Dna,
//		l.Classroom.Dna,
//		l.TimeSlot.Dna,
//		l.Weekday.Dna,
//	)
//}

func GetLesson(chromosome string) (*Lesson, error) {
	subject, ok := subjects[DNA(chromosome[0:4])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	teacher, ok := teachers[DNA(chromosome[4:7])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	lessontype, ok := lessonTypes[DNA(chromosome[7:8])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	classroom, ok := classrooms[DNA(chromosome[8:11])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	group, ok := groups[DNA(chromosome[11:14])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	timeSlot, ok := timeSlots[DNA(chromosome[14:16])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	weekDay, ok := weekDays[DNA(chromosome[16:19])]
	if !ok {
		return nil, errors.New("wrong dna")
	}

	return &Lesson{
		Subject:   subject,
		Teacher:   teacher,
		Type:      lessontype,
		Classroom: classroom,
		Group:     group,
		TimeSlot:  timeSlot,
		Weekday:   weekDay,
	}, nil
}

var subjects = map[DNA]Subject{
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

var teachers = map[DNA]Teacher{
	/*00*/ "000": {Name: "Вергунова Ірина Миколаївна", Subjects: []Subject{subjects["0000"]}},
	/*01*/ "001": {Name: "Колєнов Сергій Олександрович", Subjects: []Subject{subjects["0001"], subjects["0010"], subjects["0100"]}},
	/*02*/ "010": {Name: "Нікітченко Микола Степанович", Subjects: []Subject{subjects["0011"], subjects["1010"]}},
	/*03*/ "011": {Name: "Панченко Тарас Володимирович", Subjects: []Subject{subjects["0101"]}},
	/*04*/ "100": {Name: "Глибовець Микола Миколайович", Subjects: []Subject{subjects["0110"], subjects["1000"], subjects["1001"]}},
	/*05*/ "101": {Name: "Федорус О.М.", Subjects: []Subject{subjects["0111"]}},
	/*06*/ "110": {Name: "Яковлев В.О.", Subjects: []Subject{subjects["0111"]}},
}

var lessonTypes = map[DNA]LessonType{
	/*00*/ "0": {Type: Lecture},
	/*01*/ "1": {Type: Practice},
}

var classrooms = map[DNA]Classroom{
	/*00*/ "000": {Name: "101", Seats: 30},
	/*01*/ "001": {Name: "102", Seats: 30},
	/*02*/ "010": {Name: "201", Seats: 30},
	/*03*/ "011": {Name: "202", Seats: 30},
	/*04*/ "100": {Name: "01", Seats: 90},
	/*05*/ "101": {Name: "02", Seats: 90},
	/*06*/ "110": {Name: "403", Seats: 90},
}

var groups = map[DNA]Group{
	/*00*/ "000": {Name: "ТК-41"},
	/*01*/ "001": {Name: "ТК-42"},
	/*02*/ "010": {Name: "ТТП-41"},
	/*03*/ "011": {Name: "ТТП-42"},
	/*04*/ "100": {Name: "МІ"},
	/*05*/ "101": {Name: "ІПС-41"},
	/*06*/ "110": {Name: "ІПС-42"},
}

var timeSlots = map[DNA]TimeSlot{
	"00": {Position: 1},
	"01": {Position: 2},
	"10": {Position: 3},
	"11": {Position: 4},
}

var weekDays = map[DNA]Weekday{
	"000": {Day: 1},
	"001": {Day: 2},
	"010": {Day: 3},
	"011": {Day: 4},
	"100": {Day: 5},
}
