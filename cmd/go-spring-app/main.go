package main

import (
	_ "gitlab.com/BPeters58/go-spring-app/pkg/controller"
	_ "gitlab.com/BPeters58/go-spring-app/pkg/repository"
	_ "gitlab.com/BPeters58/go-spring-app/pkg/service"
	"gitlab.com/BPeters58/spring-go"
)

func main() {
	spring.Run()
}
