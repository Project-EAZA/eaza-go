package course

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

type Grades struct {
	A     int `bson:"A"`
	Ab    int `bson:"AB"`
	Nw    int `bson:"NW"`
	I     int `bson:"I"`
	Other int `bson:"OTHER"`
	D     int `bson:"D"`
	S     int `bson:"S"`
	Cr    int `bson:"CR"`
	Bc    int `bson:"BC"`
	N     int `bson:"N"`
	Nr    int `bson:"NR"`
	C     int `bson:"C"`
	U     int `bson:"U"`
	B     int `bson:"B"`
	F     int `bson:"F"`
	P     int `bson:"P"`
}
