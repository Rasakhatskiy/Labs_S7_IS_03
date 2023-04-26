package main

import "time"

const (
	Lecture  string = "lecture"
	Practice        = "practice"
)

type DNA string

type Subject struct {
	Name string
}

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

type TimeTable []Lesson
