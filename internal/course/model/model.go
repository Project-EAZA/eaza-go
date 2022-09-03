package model

type Course struct {
	CourseId      string        `bson:"courseId"`
	Title         string        `bson:"title"`
	CatalogNumber string        `bson:"catalogNumber"`
	Subject       Subject       `bson:"subject"`
	GeneralEd     GeneralEd     `bson:"generalEd"`
	Level         []Level       `bson:"level"`
	Breadths      []Breadth     `bson:"breadths"`
	Repeatable    string        `bson:"repeatable"`
	EthnicStudies EthnicStudies `json:"ethnicStudies"`
	Teachings     []Teaching    `bson:"teachings"`
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
func (c *Course) bind(other *Course) {
	if other != nil {
		c.CourseId = other.CourseId
		c.Title = other.Title
		c.CatalogNumber = other.CatalogNumber
		c.Subject = other.Subject
		c.GeneralEd = other.GeneralEd
		c.Level = other.Level
		c.Breadths = other.Breadths
		c.Repeatable = other.Repeatable
		c.Teachings = other.Teachings
	}
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
	TermCode int       `bson:"termCode"`
	Sections []Section `bson:"sections"`
}

type Section struct {
	TermCode      int          `bson:"termCode"`
	CourseNumber  int          `bson:"courseNumber"`
	SectionNumber int          `bson:"sectionNumber"`
	Instructors   []Instructor `bson:"instructors"`
	Grades        Grades       `bson:"grades"`
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
	SubjectCode      string `bson:"subjectCode"`
	Description      string `bson:"description"`
	ShortDescription string `bson:"shortDescription"`
}

type SchoolCollege struct {
	AcademicOrgCode   string      `bson:"academicOrgCode"`
	AcademicGroupCode string      `bson:"academicGroupCode"`
	ShortDescription  string      `bson:"shortDescription"`
	FormalDescription string      `bson:"formalDescription"`
	UddsCode          interface{} `bson:"uddsCode"`
	SchoolCollegeURI  string      `bson:"schoolCollegeURI"`
}
