package storage

//通用记录内容
type GenerationRecordContent struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}

//人物记录内容集
type PersonageRecordContentContentCollection struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}

//经济记录内容集
type EconomicsRecordContentCollection struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}

//体育记录内容集
type SportRecordContentCollection struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}

//军事记录内容集
type MilitaryRecordContentCollection struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}


//娱乐记录内容集
type EntertainmentRecordContentCollection struct {
	Uuid string `json:"uuid"`
	Content string `json:"content"`
}
