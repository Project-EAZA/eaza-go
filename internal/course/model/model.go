package model

// Course is the model of a course
type Course struct {
	UUID         string     `bson:"uuid"`
	CourseNumber int        `bson:"courseNumber"`
	Name         string     `bson:"name"`
	Breadths     []Breadth  `bson:"breadths"`
	Ethnic       Ethnic     `bson:"ethnic"`
	Level        Level      `bson:"level"`
	GE           GE         `bson:"GE"`
	Description  string     `bson:"description"`
	Requirement  string     `bson:"requirement"`
	Teachings    []Teaching `bson:"teachings"`
}

func (c *Course) AvgGPA() float32 {
	total := DefaultGrades()
	for _, t := range c.Teachings {
		for _, s := range t.Sections {
			total.Add(&s.Grades)
		}
	}
	return total.GPA()
}

type Breadth struct {
	Name string `bson:"name"`
	Code string `bson:"code"`
}

type Level struct {
	Name string `bson:"name"`
	Code string `bson:"code"`
}

type Ethnic struct {
	Name string `bson:"name"`
	Code string `bson:"code"`
}

type GE struct {
	Name string `bson:"name"`
	Code string `bson:"code"`
}

type Instructor struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}

type Teaching struct {
	TermCode int       `bson:"termCode"`
	Subjects []Subject `bson:"subjects"`
	Sections []Section `bson:"sections"`
}

type Section struct {
	TermCode      int          `bson:"termCode"`
	CourseNumber  int          `bson:"courseNumber"`
	SectionType   string       `bson:"sectionType"`
	SectionNumber int          `bson:"sectionNumber"`
	Schedule      Schedule     `bson:"schedule"`
	Room          Room         `bson:"room"`
	Instructors   []Instructor `bson:"instructors"`
	Grades        Grades
}

type Subject struct {
	Name         string `bson:"name"`
	Abbreviation string `bson:"abbreviation"`
	Code         string `bson:"code"`
}

type Times struct {
	StartTime int `bson:"startTime"`
	EndTime   int `bson:"endTime"`
}
type Room struct {
	FacilityCode string `bson:"facilityCode"`
	RoomCode     string `bson:"roomCode"`
	UUID         string `bson:"uuid"`
}

type Days struct {
	Days []string `bson:"days"`
}

type Schedule struct {
	Times Times  `bson:"times"`
	Days  Days   `bson:"days"`
	UUID  string `bson:"uuid"`
}
