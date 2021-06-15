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
	// 邮件实体类
	Email `json:"email"`
	// 工作内容
	WorkInfo `json:"work_info"`
}

