package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var randDevice *rand.Rand

const lessonsInTable = 16
const starterSet = 1000000

func main() {
	initValues()
	randDevice = rand.New(rand.NewSource(time.Now().UnixNano()))

	gen0 := make([]string, starterSet)
	for i := 0; i < starterSet; i++ {
		gen0[i] = CreateRandomGene(lessonsInTable)
	}
	fmt.Println()

	gen0 = removeDuplicate(gen0)
	//fmt.Printf("Got %d unique chromosomes", len(gen0))

	count := 0
	for _, gen := range gen0 {
		if fitness(gen) == 1 {
			count++
			fmt.Println()
		}
	}
	fmt.Println("SUCC: ", count)
}

func CreateRandomGene(number int) string {
	result := ""

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
	return result
}

func timetableFromChromosome(chromosome string, n int) TimeTable {
	chromosomes := divideString(chromosome, n)
	timetable := TimeTable{}
	for _, s := range chromosomes {
		lesson, err := lessonFromGene(s)
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

func lessonFromGene(gene string) (*LessonGene, error) {
	subject, ok := subjectsMap[gene[0:4]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	teacher, ok := teachersMap[gene[4:7]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	lessontype, ok := lessonTypesMap[gene[7:8]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	classroom, ok := classroomsMap[gene[8:11]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	group, ok := groupsMap[gene[11:14]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	timeSlot, ok := timeSlotsMap[gene[14:16]]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	weekDay, ok := weekDaysMap[gene[16:19]]
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
		Gene:       gene,
	}, nil
}

type ConflictingTypes struct {
	Classroom Classroom
	Subject   Subject
	Type      LessonType
}

type tmp struct {
}

func fitness(chromosome string) int {
	groupAtDayAtTime := make(map[string]LessonGene) // Map to keep track of which students are in each class
	teacehrAtDayAtTime := make(map[string]LessonGene)

	for _, gene := range divideString(chromosome, lessonsInTable) {
		lesson, err := lessonFromGene(gene)
		if err != nil {
			return 0
		}

		// group in two lessons at same time
		groupDayTimeStamp := lesson.Group.Dna + lesson.Weekday.Dna + lesson.Timeslot.Dna
		if _, ok := groupAtDayAtTime[groupDayTimeStamp]; ok {
			return 0
		}
		groupAtDayAtTime[groupDayTimeStamp] = *lesson

		// teacher in two lessons at same time
		teacherDayTimeStamp := lesson.Teacher.Dna + lesson.Weekday.Dna + lesson.Timeslot.Dna
		if _, ok := teacehrAtDayAtTime[teacherDayTimeStamp]; ok {
			return 0
		}
		teacehrAtDayAtTime[teacherDayTimeStamp] = *lesson

		if lesson.Classroom.Seats < 90 && lesson.LessonType.Type == Lecture {
			return 0
		}

	}

	return 1
}
