package main

import (
	"fmt"
	"time"
)

const (
	Lecture  string = "lecture"
	Practice        = "practice"
)

type DNA string

type Subject struct {
	Name string
	Dna  DNA
}

type Teacher struct {
	Subjects []Subject
	Name     string
	Dna      DNA
}

type LessonType struct {
	Type string
	Dna  DNA
}

type Group struct {
	Name string
	Dna  DNA
}

type Weekday struct {
	Day time.Weekday
	Dna DNA
}

type Classroom struct {
	Name  string
	Seats int
	Dna   DNA
}

type Timeslot struct {
	Position int
	Dna      DNA
}

type LessonGene struct {
	Subject    Subject
	Teacher    Teacher
	LessonType LessonType
	Group      Group
	Classroom  Classroom
	Timeslot   Timeslot
	Weekday    Weekday
}

type TimeTable []LessonGene

func (t *TimeTable) String() string {
	res := ""
	for _, gene := range *t {
		res += fmt.Sprintf("\t%v\n", gene)
	}
	return res
}
