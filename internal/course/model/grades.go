package model

type Grades struct {
	A     int `bson:"A"`
	AB    int `bson:"AB"`
	B     int `bson:"B"`
	BC    int `bson:"BC"`
	C     int `bson:"C"`
	Cr    int `bson:"CR"`
	D     int `bson:"D"`
	Nw    int `bson:"NW"`
	I     int `bson:"I"`
	Other int `bson:"OTHER"`
	S     int `bson:"S"`
	N     int `bson:"N"`
	Nr    int `bson:"NR"`
	U     int `bson:"U"`
	F     int `bson:"F"`
	P     int `bson:"P"`
}

func DefaultGrades() Grades {
	return Grades{
		A:     0,
		AB:    0,
		Nw:    0,
		I:     0,
		Other: 0,
		D:     0,
		S:     0,
		Cr:    0,
		BC:    0,
		N:     0,
		Nr:    0,
		C:     0,
		U:     0,
		B:     0,
		F:     0,
		P:     0,
	}
}

func (g *Grades) Add(other *Grades) {
	g.A += other.A
	g.AB += other.AB
	g.Nw += other.Nw
	g.I += other.I
	g.Other += other.Other
	g.D += other.D
	g.S += other.S
	g.Cr += other.Cr
	g.BC += other.BC
	g.N += other.N
	g.Nr += other.Nr
	g.C += other.C
	g.U += other.U
	g.B += other.B
	g.F += other.F
	g.P += other.P
}

func (g *Grades) GPA() float32 {
	// in case total is zero, otherwise answer will be NaN
	if g.Total() == 0 {
		return 0.0
	}
	return g.GradesSum() / float32(g.Total())
}

func (g *Grades) GradesSum() float32 {
	return float32(g.A)*4.0 + float32(g.AB)*3.5 + float32(g.B)*3.0 + float32(g.BC)*2.5 + float32(g.C)*2.0 + float32(g.D)*1.0
}

func (g *Grades) Total() int {
	return g.A + g.AB + g.B + g.BC + g.C + g.D + g.F
}
