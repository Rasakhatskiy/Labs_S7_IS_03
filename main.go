package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var randDevice *rand.Rand

const lessonsInTable = 16
const starterSet = 100

func main() {
	initValues()

	randDevice = rand.New(rand.NewSource(time.Now().UnixNano()))

	gen0 := make([]string, starterSet)

	for i := 0; i < starterSet; i++ {
		for {
			genChromo := CreateRandomChromosome(lessonsInTable)
			//tt := timetableFromChromosome(genChromo, lessonsInTable)
			//if tt.ValidateTimeSlots() {
			gen0[i] = genChromo
			fmt.Printf(".")
			break
			//}
		}
	}
	fmt.Println()

	gen0 = removeDuplicate(gen0)
	//fmt.Printf("Got %d unique chromosomes", len(gen0))

	for _, gen := range gen0 {
		if fitness(gen) != 1 {
			break
		}
	}

}

func fitness(chromosome string) int {
	tt := timetableFromChromosome(chromosome, lessonsInTable)
	fit := tt.ValidateTimeSlots()
	if fit != 1 {
		fmt.Println(tt.String())
		return fit
	}
	return 1
}

func CreateRandomChromosome(number int) string {
	result := string("")

	for i := 0; i < number; i++ {
		subjectKey := subjectsList[randDevice.Intn(len(subjectsList))].Dna
		teacherKey := teachersList[randDevice.Intn(len(teachersList))].Dna
		lessonTypeKey := lessonTypesList[randDevice.Intn(len(lessonTypesList))].Dna
		classroomKey := classroomsList[randDevice.Intn(len(classroomsList))].Dna
		groupKey := groupsList[randDevice.Intn(len(groupsList))].Dna
		timeSlotKey := timeSlotsList[randDevice.Intn(len(timeSlotsList))].Dna
		weekDayKey := weekDaysList[randDevice.Intn(len(weekDaysList))].Dna

		result += subjectKey + teacherKey + lessonTypeKey + classroomKey + groupKey + timeSlotKey + weekDayKey
	}
	return string(result)
}

func timetableFromChromosome(chromosome string, n int) TimeTable {
	chromosomes := divideString(chromosome, n)
	timetable := TimeTable{}
	for _, s := range chromosomes {
		lesson, err := GetLesson(s)
		if err != nil {
			fmt.Println(err)
		} else {
			timetable = append(timetable, *lesson)
		}
	}
	return timetable
}

func divideString(mystr string, size int) []string {
	var parts []string
	partSize := len(mystr) / size
	for i := 0; i < size; i++ {
		start := i * partSize
		end := start + partSize
		if i == size-1 {
			end = len(mystr)
		}
		parts = append(parts, mystr[start:end])
	}
	return parts
}

func GetLesson(chromosome string) (*LessonGene, error) {
	subject, ok := subjectsMap[chromosome[0:4]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	teacher, ok := teachersMap[chromosome[4:7]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	lessontype, ok := lessonTypesMap[chromosome[7:8]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	classroom, ok := classroomsMap[chromosome[8:11]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	group, ok := groupsMap[chromosome[11:14]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	timeSlot, ok := timeSlotsMap[chromosome[14:16]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	weekDay, ok := weekDaysMap[chromosome[16:19]]
	if !ok {
		return nil, errors.New("wrong dna")
	}

	return &LessonGene{
		Subject:    subject,
		Teacher:    teacher,
		LessonType: lessontype,
		Classroom:  classroom,
		Group:      group,
		Timeslot:   timeSlot,
		Weekday:    weekDay,
	}, nil
}

type ConflictingTypes struct {
	Classroom Classroom
	Subject   Subject
	Type      LessonType
}

func (t *TimeTable) ValidateTimeSlots() int {
	groupAtDayAtTime := make(map[string]ConflictingTypes) // Map to keep track of which students are in each class

	for _, gene := range *t {
		stamp := gene.Group.Dna + gene.Weekday.Dna + gene.Timeslot.Dna
		classroom, subject, lessonType := gene.Classroom, gene.Subject, gene.LessonType

		//for _, groupKey := range groupsKeys {
		if val, ok := groupAtDayAtTime[stamp]; ok && val.Timeslot == timeslot && val.Weekday == weekday {
			fmt.Println(val)
			fmt.Println(groupAtDayAtTime)
			return 0 // Return a fitness of 0 if the constraint is not satisfied
		}
		groupAtDayAtTime[group.Name] = ConflictingTypes{Classroom: classroom, Timeslot: timeslot, Weekday: weekday}
		//}
	}

	return 1
}
