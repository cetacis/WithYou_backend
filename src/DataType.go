package src

type TogetherTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	Comment string `json:"comment"`
	FriendEmail string `json:"friendEmail"`
	IsFinished bool `json:"IsFinished"`
}

type PrivateTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	IsFinished bool `json:"IsFinished"`
}

type Message struct {
	Msg string `json:"msg"`
	IsUser bool `json:"IsUser"`
	IsRead bool `json:"IsRead"`
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
	TogetherTasks []TogetherTask `json:"TogetherTasks"`
	PrivateTasks []PrivateTask `json:"PrivateTasks"`
	Friends []string `json:"Friends"`
	Messages []Message `json:"Messages"`
	Partner string `json:"partner"`
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
