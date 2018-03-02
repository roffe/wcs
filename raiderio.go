package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url *url.URL, target interface{}) error {
	r, err := myClient.Get(url.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//RioChar result from raider.io
type RioChar struct {
	AchievementPoints int    `json:"achievement_points"`
	ActiveSpecName    string `json:"active_spec_name"`
	ActiveSpecRole    string `json:"active_spec_role"`
	Class             string `json:"class"`
	Faction           string `json:"faction"`
	Gear              struct {
		ArtifactTraits    int `json:"artifact_traits"`
		ItemLevelEquipped int `json:"item_level_equipped"`
		ItemLevelTotal    int `json:"item_level_total"`
	} `json:"gear"`
	Gender           string `json:"gender"`
	HonorableKills   int    `json:"honorable_kills"`
	MythicPlusScores struct {
		All    int `json:"all"`
		Dps    int `json:"dps"`
		Healer int `json:"healer"`
		Tank   int `json:"tank"`
	} `json:"mythic_plus_scores"`
	Name            string `json:"name"`
	ProfileURL      string `json:"profile_url"`
	Race            string `json:"race"`
	RaidProgression struct {
		AntorusTheBurningThrone struct {
			HeroicBossesKilled int    `json:"heroic_bosses_killed"`
			MythicBossesKilled int    `json:"mythic_bosses_killed"`
			NormalBossesKilled int    `json:"normal_bosses_killed"`
			Summary            string `json:"summary"`
			TotalBosses        int    `json:"total_bosses"`
		} `json:"antorus-the-burning-throne"`
		TheEmeraldNightmare struct {
			HeroicBossesKilled int    `json:"heroic_bosses_killed"`
			MythicBossesKilled int    `json:"mythic_bosses_killed"`
			NormalBossesKilled int    `json:"normal_bosses_killed"`
			Summary            string `json:"summary"`
			TotalBosses        int    `json:"total_bosses"`
		} `json:"the-emerald-nightmare"`
		TheNighthold struct {
			HeroicBossesKilled int    `json:"heroic_bosses_killed"`
			MythicBossesKilled int    `json:"mythic_bosses_killed"`
			NormalBossesKilled int    `json:"normal_bosses_killed"`
			Summary            string `json:"summary"`
			TotalBosses        int    `json:"total_bosses"`
		} `json:"the-nighthold"`
		TombOfSargeras struct {
			HeroicBossesKilled int    `json:"heroic_bosses_killed"`
			MythicBossesKilled int    `json:"mythic_bosses_killed"`
			NormalBossesKilled int    `json:"normal_bosses_killed"`
			Summary            string `json:"summary"`
			TotalBosses        int    `json:"total_bosses"`
		} `json:"tomb-of-sargeras"`
		TrialOfValor struct {
			HeroicBossesKilled int    `json:"heroic_bosses_killed"`
			MythicBossesKilled int    `json:"mythic_bosses_killed"`
			NormalBossesKilled int    `json:"normal_bosses_killed"`
			Summary            string `json:"summary"`
			TotalBosses        int    `json:"total_bosses"`
		} `json:"trial-of-valor"`
	} `json:"raid_progression"`
	Realm        string `json:"realm"`
	Region       string `json:"region"`
	ThumbnailURL string `json:"thumbnail_url"`
}

func getRioChar(name string, realm string) *RioChar {
	char := new(RioChar)
	// var URL *url.URL
	URL, err := url.Parse("https://raider.io")
	check(err)
	URL.Path += "/api/v1/characters/profile"
	parameters := url.Values{}
	parameters.Add("region", "eu")
	parameters.Add("realm", realm)
	parameters.Add("name", name)
	parameters.Add("fields", "gear,raid_progression,mythic_plus_scores")
	URL.RawQuery = parameters.Encode()
	err = getJSON(URL, &char)
	check(err)
	return char
}
