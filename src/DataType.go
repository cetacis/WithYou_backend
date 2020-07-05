package src

type TogetherTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	Comment string `json:"comment"`
	FriendEmail string `json:"friend_email"`
	IsFinished bool `json:"is_finished"`
}

type PrivateTask struct {
	Name string `json:"name"`
	Number int `json:"number"`
	IsFinished bool `json:"is_finished"`
}

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Age string `json:"age"`
	Bio string `json:"bio"`
	Constellation string `json:"constellation"`
	Blood string `json:"blood"`
	Birthday string `json:"birthday"`
	Sex string `json:"sex"`
	ImgPath string `json:"img_path"`
	TogetherTasks []TogetherTask `json:"together_tasks"`
	PrivateTasks []PrivateTask `json:"private_tasks"`
}
