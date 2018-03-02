package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var t *template.Template

type templatedata struct {
	Character           *RioChar
	CurrentTierBosses   int
	CurrentTierProgress int
	Role                string
	Realm               string
	Ilvl                int
	Twitch              string
}

// Member struct holds member
type Member struct {
	HTML   string
	Realm  string
	Role   string
	Twitch string
}

// Members holds member
type Members struct {
	Members []Member
}

func getMythicColor(score int) string {
	switch {
	case score > 0 && score < 600:
		return "#ffffff"

	case score > 599 && score < 700:
		return "#d2ffcc"
	case score > 699 && score < 800:
		return "#a5ff99"
	case score > 799 && score < 900:
		return "#78ff66"
	case score > 899 && score < 1000:
		return "#4bff33"
	case score > 999 && score < 1200:
		return "#1eff00"
	case score > 1199 && score < 1400:
		return "#18e22c"
	case score > 1399 && score < 1600:
		return "#12c558"
	case score > 1599 && score < 1800:
		return "#0ca984"
	case score > 1799 && score < 2000:
		return "#068cb0"
	case score > 1999 && score < 2200:
		return "#0070dd"
	case score > 2199 && score < 2400:
		return "#2064e0"
	case score > 2399 && score < 2600:
		return "#4158e3"
	case score > 2599 && score < 2800:
		return "#614ce7"
	case score > 2799 && score < 3000:
		return "#8240ea"
	case score > 2999 && score < 3200:
		return "#a335ee"
	case score > 3199 && score < 3400:
		return "#b544be"
	case score > 3399 && score < 3600:
		return "#c7538e"
	case score > 3599 && score < 3800:
		return "#da625f"
	case score > 3799 && score < 4000:
		return "#ec712f"
	case score > 3999 && score < 4500:
		return "#ff8800"
	case score > 4499 && score < 5000:
		return "#f89320"
	case score > 4999 && score < 5500:
		return "#f2a640"
	case score > 5499 && score < 6000:
		return "#ecb960"
	case score > 5999:
		return "#e6cc80"
	default:
		return "#ffffff"
	}
}

func getClassColor(class string) string {
	switch strings.ToLower(class) {
	case "warrior":
		return "#C79C6E"
	case "paladin":
		return "#F58CBA"
	case "hunter":
		return "#ABD473"
	case "rogue":
		return "#FFF569"
	case "priest":
		return "#FFFFFF"
	case "death knight":
		return "#C41F3B"
	case "shaman":
		return "#0070DE"
	case "mage":
		return "#69ccf0"
	case "warlock":
		return "#9482C9"
	case "monk":
		return "#00FF96"
	case "druid":
		return "#FF7D0A"
	case "demon hunter":
		return "#A330C9"
	default:
		return "#FFFFFF"
	}
}

//AddItem to the box
func (box *Members) AddItem(member Member) []Member {
	box.Members = append(box.Members, member)
	return box.Members
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func realmFormat(in string) string {
	return strings.Title(strings.Replace(in, "-", " ", -1))
}

func getChar(name string, realm string, role string, twitch string) string {
	character := getRioChar(name, realm)

	data := templatedata{
		Character:           character,
		Role:                role,
		CurrentTierBosses:   character.RaidProgression.AntorusTheBurningThrone.TotalBosses,
		CurrentTierProgress: character.RaidProgression.AntorusTheBurningThrone.MythicBossesKilled,
		Realm:               realm,
		Ilvl:                character.Gear.ItemLevelTotal,
		Twitch:              twitch,
	}
	var res bytes.Buffer

	err := t.Execute(&res, data)

	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "%s / %s / %s\n", character.Name, character.Realm, role)
	return res.String()
}

func mythicFilter(in int) string {
	if in > 3000 {
		return fmt.Sprintf("%sK", strconv.FormatFloat((float64(in)/1000), 'f', 1, 64))
	}
	return ""
}

func init() {
	fmap := template.FuncMap{
		"getClassColor":  getClassColor,
		"toUpper":        strings.ToUpper,
		"toLower":        strings.ToLower,
		"realmFormat":    realmFormat,
		"mythicFilter":   mythicFilter,
		"getMythicColor": getMythicColor,
	}
	t = template.Must(template.New("char.tmpl").Funcs(fmap).ParseFiles("char.tmpl"))

}

func main() {
	startT := time.Now()
	defer func() {
		fmt.Fprintf(os.Stderr, "Took: %s", time.Since(startT))
	}()
	members := Members{}

	file, err := os.Open("team.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "/")
		name, realm, role, twitch := s[0], strings.ToLower(s[1]), strings.ToLower(s[2]), strings.ToLower(s[3])
		member := Member{
			HTML:   getChar(name, realm, role, twitch),
			Role:   role,
			Realm:  realm,
			Twitch: twitch,
		}
		members.AddItem(member)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	box1 := Members{}
	box2 := Members{}
	box3 := Members{}
	box4 := Members{}

	pos := 0

	for _, item := range members.Members {
		pos++
		switch pos {
		case 1:

			box1.AddItem(item)
		case 2:
			box2.AddItem(item)
		case 3:
			box3.AddItem(item)
		case 4:
			box4.AddItem(item)
			pos = 0
		}

	}

	fmt.Println("[lgc_column grid=\"25\" tablet_grid=\"25\" mobile_grid=\"25\"]")
	for _, item := range box1.Members {
		fmt.Println(item.HTML)
	}
	fmt.Println("[/lgc_column]")

	fmt.Println("[lgc_column grid=\"25\" tablet_grid=\"25\" mobile_grid=\"25\"]")
	for _, item := range box2.Members {
		fmt.Println(item.HTML)
	}
	fmt.Println("[/lgc_column]")
	fmt.Println("[lgc_column grid=\"25\" tablet_grid=\"25\" mobile_grid=\"25\"]")
	for _, item := range box3.Members {
		fmt.Println(item.HTML)
	}
	fmt.Println("[/lgc_column]")
	fmt.Println("[lgc_column grid=\"25\" tablet_grid=\"25\" mobile_grid=\"25\" last=\"true\"]")
	for _, item := range box4.Members {
		fmt.Println(item.HTML)
	}
	fmt.Println("[/lgc_column]")
}
