package model

type Course struct {
	Title         string     `bson:"title"`
	CatalogNumber string     `bson:"catalogNumber"`
	Subject       Subject    `bson:"subject"`
	GeneralEd     GeneralEd  `bson:"generalEd"`
	Level         []Level    `bson:"level"`
	Breadths      []Breadth  `bson:"breadths"`
	Teachings     []Teaching `bson:"teachings"`
	Repeatable    string     `bson:"repeatable"`
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
	Description string `bson:"description"`
	Code        string `bson:"code"`
}

type Level struct {
	Description string `bson:"description"`
	Code        string `bson:"code"`
}

type GeneralEd struct {
	Description string `bson:"description"`
	Code        string `bson:"code"`
}

type EthnicStudies struct {
	Description string `bson:"description"`
	Code        string `bson:"code"`
}

type Instructor struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}

type Teaching struct {
	TermCode int `bson:"termCode"`
	//Subjects []Subject `bson:"subjects"`
	Sections []Section `bson:"sections"`
}

type Section struct {
	TermCode      int          `bson:"termCode"`
	CourseNumber  int          `bson:"courseNumber"`
	SectionNumber int          `bson:"sectionNumber"`
	Instructors   []Instructor `bson:"instructors"`
	Grades        Grades
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

type Subject struct {
	SubjectCode       string `bson:"subjectCode"`
	Description       string `bson:"description"`
	ShortDescription  string `bson:"shortDescription"`
	FormalDescription string `bson:"formalDescription"`
}

type SchoolCollege struct {
	AcademicOrgCode   string      `bson:"academicOrgCode"`
	AcademicGroupCode string      `bson:"academicGroupCode"`
	ShortDescription  string      `bson:"shortDescription"`
	FormalDescription string      `bson:"formalDescription"`
	UddsCode          interface{} `bson:"uddsCode"`
	SchoolCollegeURI  string      `bson:"schoolCollegeURI"`
}
