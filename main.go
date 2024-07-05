package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Root struct {
	Notify   bool           `json:"notify"`
	Schedule []ScheduleItem `json:"schedule"`
	Groups   []string       `json:"groups"`
}

type ScheduleItem struct {
	Name   string   `json:"name"`
	Couple []Couple `json:"couples"`
}

type Couple struct {
	I      string `json:"i"`
	Name   string `json:"name"`
	Office string `json:"office"`
}

func main() {
	httpPostUrl := "https://nggtk.ru/api/v2/GetScheduleGroup/?vk_access_token_settings=&vk_app_id=7688110&vk_are_notifications_enabled=0&vk_is_app_user=1&vk_is_favorite=1&vk_language=ru&vk_platform=desktop_web&vk_ref=catalog_recent&vk_testing_group_id=3&vk_ts=1715447998&vk_user_id=491552018&sign=jHrws9PeP528Ijpeo8IEv5mVtqOB7j-kbnmyO_64bAo"
	//fmt.Println("URL:", httpPostUrl)

	res, err := http.PostForm(httpPostUrl, url.Values{"group": {"17"}})

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(body))

	var root Root

	json.Unmarshal(body, &root)

	for _, item := range root.Schedule {
		fmt.Println(item.Name)
		for _, couple := range item.Couple {
			fmt.Println(couple.Name, " ", couple.Office)
		}
		fmt.Println("")
	}

	fmt.Scanln()
}
