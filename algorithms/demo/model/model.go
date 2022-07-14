package model

type Q struct {
	ID         int
	Title      string
	As         []A
	ConditionS []Condition
}
type A struct {
	ID        int
	Title     string
	Link_Q_ID int
}
type Condition struct {
	Origin_Q_ID int
	Origin_A_ID int
}
