package server

func Init() {
	r := CreateRouter()
	r.Run(":80")
}
