package defs

//reqeusts
type UserCredential struct {
	userName string `json:"user_name"`
	pwd      string `json:"pwd"`
}

//Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
