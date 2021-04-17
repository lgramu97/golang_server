package _go

//package main
// Para usar directamente, sin necesidad de ser un modulo, cambiar todos los package por main.

func Main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Login()))

	server.Listen()
}
