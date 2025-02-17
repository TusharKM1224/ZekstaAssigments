package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Loadconfig() {
	files := []string{"/home/yuka/ZekstaAssigments/Assignemt5/UserAuthWithArch/config/DBconfig.env", "/home/yuka/ZekstaAssigments/Assignemt5/UserAuthWithArch/config/OrgConfig.env"}
	for _, file := range files {
		err := godotenv.Load(file)
		fmt.Println(err)

	}
}
func GetEnv(key string) string {
	return os.Getenv(key)
}
