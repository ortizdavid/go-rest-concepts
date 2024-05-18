package models

type ReportModel struct {
}

type TableReport struct {
	Title  	string
	Rows	interface{}
	Count   int
}

func (ReportModel) GetAllUsers() TableReport {
	rows, _ := UserModel{}.FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Users",
		Rows: rows,
		Count: count,
	}
}


func (ReportModel) GetAllTasks() TableReport {
	rows, _ := TaskModel{}.FindAllData()
	count := len(rows)
	return TableReport {
		Title: "Tasks",
		Rows: rows,
		Count: count,
	}
}
