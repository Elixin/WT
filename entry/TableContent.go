package entry

type TableContent struct {
	//标题
	Title string
	// 当前星期
	NowWeek string
	// 工作日
	WorkingDay []string
	// 工作地点
	WorkPlace []string
	// 项目名称
	ProjectName []string
	// 本周工作计划
	NowWeekWorkContext []string
	// 完成内容
	OverContext [5]string
	// 工作总结
	WorkSummary string
	// 学习计划
	StudyPlan string
	// 下周工作计划
	NextWeekPlan string
	// 费用
	Cost []float64
	// 备注
	Remark []string
}

func (TableContent) SetValueInEntry(deploy *Deploy, week []string, parsing []string, allContext string) TableContent {
	content := TableContent{}
	if len(deploy.WorkPlace)>0 {
		content.WorkPlace = append(content.WorkPlace, deploy.WorkPlace...)
	}else {
		content.WorkPlace = make([]string,5)
	}

	if len(week)>0 {
		content.WorkingDay = append(content.WorkingDay, week...)
	}else {
		content.WorkingDay = make([]string,5)
	}

	if len(deploy.ProjectName)>0 {
		content.ProjectName = append(content.ProjectName, deploy.ProjectName...)
	}else {
		content.ProjectName = make([]string,5)
	}

	if len(deploy.NowWeekWorkContext)>0 {
		content.NowWeekWorkContext = append(content.NowWeekWorkContext, deploy.NowWeekWorkContext...)
	}else {
		content.NowWeekWorkContext = make([]string,5)
	}


	if len(parsing) == 1 {
		for i := 0; i < 5; i++ {
			content.OverContext[i] = parsing[0]
		}
	}else {
		for i := 0; i < 5; i++ {
			content.OverContext[i] = parsing[i]
		}
	}

	if len(deploy.Cost)>0 {
		content.Cost = append(content.Cost, deploy.Cost...)
	}else {
		content.Cost = make([]float64,5)
	}

	if len(deploy.Remark)>0 {
		content.Remark = append(content.Remark, deploy.Remark...)
	}else {
		content.Remark = make([]string,5)
	}


	switch deploy.NowWeekNum {
	case 1:
		content.NowWeek = "第一周"
	case 2:
		content.NowWeek = "第二周"
	case 3:
		content.NowWeek = "第三周"
	case 4:
		content.NowWeek = "第四周"
	case 5:
		content.NowWeek = "第五周"
	case 6:
		content.NowWeek = "第六周"
	}

	content.WorkSummary = allContext
	content.StudyPlan = deploy.StudyPlan
	content.NextWeekPlan = deploy.NextWeekPlan
	//20210517-20210521工作周报及计划
	content.Title =week[0]+"-"+week[len(week)-1]+"工作周报及计划"
	return content
}
