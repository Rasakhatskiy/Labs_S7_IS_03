package main

import (
	"fmt"
	"time"
)

const (
	Lecture  string = "lecture "
	Practice        = "practice"
)

type Subject struct {
	Name string
	Dna  string
}

type Teacher struct {
	Subjects []Subject
	Name     string
	Dna      string
}

type LessonType struct {
	Type string
	Dna  string
}

type Group struct {
	Name string
	Dna  string
}

type Weekday struct {
	Day time.Weekday
	Dna string
}

type Classroom struct {
	Name  string
	Seats int
	Dna   string
}

type Timeslot struct {
	Position int
	Dna      string
}

type LessonGene struct {
	Subject    *Subject
	Teacher    *Teacher
	LessonType *LessonType
	Group      *Group
	Classroom  *Classroom
	Timeslot   *Timeslot
	Weekday    *Weekday
	Gene       string
}

func formattedWeekDay(d int) string {
	if d == 0 {
		return "Неділя   "
	}
	if d == 1 {
		return "Понеділок"
	}
	if d == 2 {
		return "Віторок  "
	}
	if d == 3 {
		return "Середа   "
	}
	if d == 4 {
		return "Четвер   "
	}
	if d == 5 {
		return "П'ятниця "
	}
	if d == 6 {
		return "Субота   "
	}
	return ""
}

func (l LessonGene) String() string {
	return fmt.Sprintf("%s %s %d %s %s %s %s %s",
		l.Gene,
		formattedWeekDay(int(l.Weekday.Day)),
		l.Timeslot.Position,
		l.Group.Name,
		l.LessonType.Type,
		l.Classroom.Name,
		l.Teacher.Name,
		l.Subject.Name)
}

type TimeTable []LessonGene

func (t *TimeTable) String() string {
	res := ""
	for _, gene := range *t {
		res += fmt.Sprintf("\t%v\n", gene)
	}
	return res
}
