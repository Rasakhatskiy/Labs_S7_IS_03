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
	fmt.Printf("Got %d unique chromosomes", len(gen0))
}

func fitness(chromosome string) int {
	tt := timetableFromChromosome(chromosome, lessonsInTable)

}

func CreateRandomChromosome(number int) string {
	result := ""

	for i := 0; i < number; i++ {
		subjectKey := subjectKeys[randDevice.Intn(len(subjectKeys))]
		teacherKey := teachersKeys[randDevice.Intn(len(teachersKeys))]
		lessontypeKey := lessonTypesKeys[randDevice.Intn(len(lessonTypesKeys))]
		classroomKey := classroomsKeys[randDevice.Intn(len(classroomsKeys))]
		groupKey := groupsKeys[randDevice.Intn(len(groupsKeys))]
		timeSlotKey := timeSlotsKeys[randDevice.Intn(len(timeSlotsKeys))]
		weekDayKey := weekDaysKeys[randDevice.Intn(len(weekDaysKeys))]

		result += subjectKey + teacherKey + lessontypeKey + classroomKey + groupKey + timeSlotKey + weekDayKey
	}
	return result
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

func GetLesson(chromosome string) (*Lesson, error) {
	subject, ok := subjectsMap[DNA(chromosome[0:4])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	teacher, ok := teachersMap[DNA(chromosome[4:7])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	lessontype, ok := lessonTypes[DNA(chromosome[7:8])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	classroom, ok := classroomsMap[DNA(chromosome[8:11])]
	if !ok {
		return nil, errors.New("wrong dna")
	}
	group, ok := groupsMap[DNA(chromosome[11:14])]
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

type VerTimeTable struct {
	Days []VerDay
}

type VerDay struct {
	Slots []VerSlot
}

type VerSlot struct {
	//Subject Subject
	//Teacher Teacher
	//Classroom Classroom
	//LessonType LessonType
	//Group Group
	Taken bool
}

func (t *TimeTable) ValidateTimeSlots() bool {
	v := VerTimeTable{}
	v.Days = make([]VerDay, 7)
	for i := range v.Days {
		v.Days[i].Slots = make([]VerSlot, 4)
	}

	for _, lesson := range *t {
		if v.Days[lesson.Weekday.Day].Slots[lesson.TimeSlot.Position-1].Taken {
			return false
		}
		v.Days[lesson.Weekday.Day].Slots[lesson.TimeSlot.Position-1].Taken = true
	}
	return true
}
