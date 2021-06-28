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

	content.WorkPlace = append(content.WorkPlace, deploy.WorkPlace...)
	content.WorkingDay = append(content.WorkingDay, week...)
	content.ProjectName = append(content.ProjectName, deploy.ProjectName...)
	content.NowWeekWorkContext = append(content.NowWeekWorkContext, deploy.NowWeekWorkContext...)

	if len(parsing) == 1 {
		for i := 0; i < 5; i++ {
			content.OverContext[i] = parsing[0]
		}
	}else {
		for i := 0; i < 5; i++ {
			content.OverContext[i] = parsing[i]
		}
	}
	content.Cost = append(content.Cost, deploy.Cost...)
	content.Remark = append(content.Remark, deploy.Remark...)


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
