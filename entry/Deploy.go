package entry
type Deploy struct {
	// 日报地址
	FilePath string `json:"file_path"`
	// 模板地址
	ModePath string	`json:"mode_path"`
	// 输出路径
	OutPath string `json:"out_path"`
	// 作者
	Author string `json:"author"`
	// 当前周数
	NowWeekNum int `json:"now_week_num"`
	// 邮件主送人
	Sender []string `json:"sender"`
	// 邮件抄送人
	EmailCc []string `json:"email_cc"`
	// 邮件发送地址
	EmailAddress string `json:"email_address"`
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

