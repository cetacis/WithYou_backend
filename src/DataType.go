package src

type TogetherTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	Comment string `json:"comment"`
	FriendEmail string `json:"friend_email"`
	IsFinished bool `json:"is_finished"`
	TaskType int `json:"task_type"`
}

type PrivateTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	IsFinished bool `json:"is_finished"`
}

type Message struct {
	Msg string `json:"msg"`
	IsUser bool `json:"is_user"`
	IsRead bool `json:"is_read"`
}

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Age string `json:"age"`
	Bio string `json:"bio"`
	Constellation string `json:"constellation"`
	Birthday string `json:"birthday"`
	Sex string `json:"sex"`
	ImgPath string `json:"img_path"`
	TogetherTasks []TogetherTask `json:"together_tasks"`
	PrivateTasks []PrivateTask `json:"private_tasks"`
	Friends []string `json:"friends"`
	Messages []Message `json:"messages"`
}

type TaskQueue struct {
	TaskId string `json:"task_id"`
	Email string `json:"email"`
}

type MatchInfo struct {
	First string `json:"first"`
	Second string `json:"second"`
}

type RtMsg struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
}
