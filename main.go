package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

var randDevice *rand.Rand

type Species struct {
	Gene    string
	Fitness int
}

func printSpecies(species []Species) {
	for i2, s := range species {
		if s.Fitness == 0 {
			fmt.Printf(" and %d zeros\n", len(species)-i2)
			break
		}
		fmt.Printf("%d ", s.Fitness)
	}

	fmt.Println()
}

//func (s Species) String() string {
//	return fmt.Sprintf("%d ", s.Fitness /*, s.Gene[0:16]*/)
//}

func main() {
	initValues()
	randDevice = rand.New(rand.NewSource(time.Now().UnixNano()))

	gen0 := make([]string, starterSet)
	var species []Species
	created := 0
	for created < starterSet {
		gene := CreateRandomGene(lessonsInTable)
		fit := fitness(gene)
		//if fit != 0 {
		gen0[created] = gene
		created++
		species = append(species, Species{
			Gene:    gene,
			Fitness: fit,
		})
		//fmt.Printf(".")
		//}
	}

	fmt.Println()

	fmt.Println("fallG ", fallG)
	fmt.Println("fallT ", fallT)
	fmt.Println("fallC ", fallC)
	fmt.Println("fallS ", fallS)
	fmt.Println("fallW ", fallW)
	fmt.Println("invalidG ", invalidG)

	species = removeDuplicateSpecies(species)
	//fmt.Printf("Got %d unique chromosomes", len(gen0))

	fmt.Println("OK: ", len(species))

	sort.Slice(species, func(i, j int) bool {
		return (species)[i].Fitness > (species)[j].Fitness
	})

	fmt.Println("GEN 0")
	printSpecies(species)
	//species = removeInvalid(species)

	const MaxGen = 10

	for i := 0; i < MaxGen; i++ {
		species = crossover(species)
		species = mutate(species)
		fmt.Println("GEN ", i+1)
		printSpecies(species)
	}

	a := divideString(species[0].Gene, lessonsInTable)

	for _, s := range a {
		fmt.Println(lessonFromGene(s))
	}

	showTeachers()
}

func CreateRandomGene(number int) string {
	result := ""

	for i := 0; i < number; i++ {
		result += validSubTeach[randDevice.Intn(len(validSubTeach))]

		//fmt.Println(lt.Type, ct.Seats)

		for j := 0; j < (SingleGenomeLength - (TeachersLen + SubjectLen)); j++ {
			result += strconv.Itoa(randDevice.Intn(2))
		}
	}
	//for i := 0; i < number; i++ {
	//	subjectKey := subjectsList[randDevice.Intn(len(subjectsList))].Dna
	//	teacherKey := teachersList[randDevice.Intn(len(teachersList))].Dna
	//	lessonTypeKey := lessonTypesList[randDevice.Intn(len(lessonTypesList))].Dna
	//	classroomKey := classroomsList[randDevice.Intn(len(classroomsList))].Dna
	//	groupKey := groupsList[randDevice.Intn(len(groupsList))].Dna
	//	timeSlotKey := timeSlotsList[randDevice.Intn(len(timeSlotsList))].Dna
	//	weekDayKey := weekDaysList[randDevice.Intn(len(weekDaysList))].Dna
	//
	//	result += subjectKey + teacherKey + lessonTypeKey + classroomKey + groupKey + timeSlotKey + weekDayKey
	//}
	//fmt.Println(result)
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
	offset := 0

	subject, ok := subjectsMap[gene[offset:offset+SubjectLen]]
	offset += SubjectLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	teacher, ok := teachersMap[gene[offset:offset+TeachersLen]]
	offset += TeachersLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	lessontype, ok := lessonTypesMap[gene[offset:offset+LessonTypeLen]]
	offset += LessonTypeLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	classroom, ok := classroomsMap[gene[offset:offset+ClassroomsLen]]
	offset += ClassroomsLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	group, ok := groupsMap[gene[offset:offset+GroupsLen]]
	offset += GroupsLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	timeSlot, ok := timeSlotsMap[gene[offset:offset+TimeSlotsLen]]
	offset += TimeSlotsLen
	if !ok {
		return nil, errors.New("wrong dna")
	}
	weekDay, ok := weekDaysMap[gene[offset:offset+WeekDaysLen]]
	offset += WeekDaysLen
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

var fallG, fallT, fallC, fallS, fallW, invalidG int = 0, 0, 0, 0, 0, 0

func fitness(chromosome string) int {
	groupAtDayAtTime := make(map[string]LessonGene) // Map to keep track of which students are in each class
	teacehrAtDayAtTime := make(map[string]LessonGene)

	//targetLessons := lessonsInTable / len(groupsList)
	groupsNumberOfLessons := make(map[string]int)
	for _, group := range groupsList {
		groupsNumberOfLessons[group.Dna] = 0
	}

	fit := 0

	for _, gene := range divideString(chromosome, lessonsInTable) {
		lesson, err := lessonFromGene(gene)
		if err != nil {
			invalidG++
			return 0
		}

		// group in two lessons at same time
		groupDayTimeStamp := lesson.Group.Dna + lesson.Weekday.Dna + lesson.Timeslot.Dna
		if _, ok := groupAtDayAtTime[groupDayTimeStamp]; ok {
			fallG++
			return 0
		}
		groupAtDayAtTime[groupDayTimeStamp] = *lesson

		// teacher in two lessons at same time
		teacherDayTimeStamp := lesson.Teacher.Dna + lesson.Weekday.Dna + lesson.Timeslot.Dna
		if _, ok := teacehrAtDayAtTime[teacherDayTimeStamp]; ok {
			fallT++
			return 0
		}
		teacehrAtDayAtTime[teacherDayTimeStamp] = *lesson

		//if lesson.Classroom.Seats == 30 && lesson.LessonType.Type == Lecture {
		//	fallC++
		//	return 0
		//}

		hasLesson := false
		teacher := teachersMap[lesson.Teacher.Dna]
		for _, subject := range teacher.Subjects {

			if lesson.Subject.Dna == subject.Dna {
				hasLesson = true
				break
			}
		}
		if !hasLesson {
			fallS++
			return 0
		}

		// earlier - better
		//fit += (8 - lesson.Timeslot.Position) / 2

		// saturdays less preferable
		if lesson.Weekday.Day < 6 {
			fit += 1
		}

		if lesson.Weekday.Day == 0 {
			fallW++
			return 0
		}

		// each group should have more lessons
		groupsNumberOfLessons[lesson.Group.Dna]++
	}

	for _, i := range groupsNumberOfLessons {
		fit += i
	}

	if fit < 0 {
		fit = 0
	}

	return fit
}

func breed(p1, p2 string) (string, string) {
	//splitPoint := SplitPositions[randDevice.Intn(len(SplitPositions))]
	splitPoint := randDevice.Intn(TotalGenomeLength)
	return p1[:splitPoint] + p2[splitPoint:], p2[:splitPoint] + p1[splitPoint:]
}

func crossover(species []Species) []Species {
	startLen := len(species)
	for i := 0; i < startLen; i += 2 {
		//if species[i].Fitness == 0 {
		//	continue
		//}

		c1, c2 := breed(species[i].Gene, species[i+1].Gene)
		f1 := fitness(c1)
		f2 := fitness(c2)
		species = append(species, Species{Gene: c1, Fitness: f1})
		species = append(species, Species{Gene: c2, Fitness: f2})
	}

	species = removeDuplicateSpecies(species)

	sort.Slice(species, func(i, j int) bool {
		return (species)[i].Fitness > (species)[j].Fitness
	})

	return species[:startLen]
}

func mutate(species []Species) []Species {
	for i := range species {
		n := TotalGenomeLength / 10
		gene := species[i].Gene
		for j := 0; j < n; j++ {
			mutatePoint := randDevice.Intn(TotalGenomeLength)
			c := string(gene[mutatePoint])
			if c == "0" {
				gene = gene[:mutatePoint] + "1" + gene[mutatePoint+1:]
			} else {
				gene = gene[:mutatePoint] + "0" + gene[mutatePoint+1:]
			}
		}
		if validateFullGene(gene) == nil {
			species[i].Gene = gene
		} else {
			fmt.Println("mutation failed")
		}
	}
	return species
}

func removeInvalid(species []Species) []Species {
	p := 0
	for i2, s := range species {
		if s.Fitness == 0 {
			p = i2
			break
		}
	}

	return species[:p]
}

func validateFullGene(gene string) error {
	split := divideString(gene, lessonsInTable)
	for _, s := range split {
		_, err := lessonFromGene(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func showTeachers() {
	fmt.Println("======")
	for _, teacher := range teachersList {
		fmt.Printf("%s ", teacher.Name)
		for _, subject := range teacher.Subjects {
			fmt.Printf("%s; ", subject.Name)
		}
		fmt.Println()
	}
	fmt.Println()

}

func genValidTeacherSubjectPairs() []string {
	var res []string
	for _, teacher := range teachersList {
		for _, subject := range teacher.Subjects {
			res = append(res, subject.Dna+teacher.Dna)
		}
	}
	fmt.Println(res)
	return res
}
