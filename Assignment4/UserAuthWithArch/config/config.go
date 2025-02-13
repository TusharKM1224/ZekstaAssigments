package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Loadconfig(){
	files:=[]string{"E:\\PracticeHub\\Go tutorial\\ZekstaAssignments\\Assignment4\\UserAuthWithArch\\config\\DBconfig.env","E:\\PracticeHub\\Go tutorial\\ZekstaAssignments\\Assignment4\\UserAuthWithArch\\config\\OrgConfig.env"}
	for _,file:=range files{
		err:=godotenv.Load(file)
		fmt.Println(err)


	}
}
func GetEnv(key string) string {
	return os.Getenv(key)
}

