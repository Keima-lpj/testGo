package project1

/**
学生的结构体
*/
type Student struct {
	Name  string  //姓名
	Score float64 //分数
}

/**
老师的结构体
*/
type teacher struct {
	Name  string  //姓名
	Score float64 //分数
}

func GetNewTeacher() *teacher {
	return &teacher{}
}

/**
校长的结构体
*/
type Master struct {
	Name  string  //姓名
	score float64 //分数
}

//这里申明两个校长的方法
func (m *Master) SetScore(score float64) {
	m.score = score
}
func (m *Master) GetScore() float64 {
	return m.score
}

/**
指导员的结构体
*/
type counselor struct {
	Name  string  //姓名
	score float64 //分数
}

func GetNewCounselor() *counselor {
	return &counselor{}
}

//这里申明一个指导员的方法，用于设置指导员的分数
func (c *counselor) SetScore(score float64) {
	c.score = score
}
