package entry

type WorkInfo struct {
	// 工作日
	WorkingDay []string `json:"working_day"`
	// 工作地点
	WorkPlace []string `json:"work_place"`
	// 项目名称
	ProjectName []string `json:"project_name"`
	// 本周工作计划
	NowWeekWorkContext []string `json:"now_week_work_context"`
	// 学习计划
	StudyPlan string `json:"study_plan"`
	// 下周工作计划
	NextWeekPlan string `json:"next_week_plan"`
	// 费用
	Cost []float64 `json:"cost"`
	// 备注
	Remark []string `json:"remark"`
}
