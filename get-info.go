package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	// "sync"
)

type GosumemoryResponse struct {
	Settings struct {
		ShowInterface bool `json:"showInterface"`
		Folders       struct {
			Game  string `json:"game"`
			Skin  string `json:"skin"`
			Songs string `json:"songs"`
		} `json:"folders"`
	} `json:"settings"`
	Menu struct {
		MainMenu struct {
			BassDensity float64 `json:"bassDensity"`
		} `json:"mainMenu"`
		State         float64 `json:"state"`
		GameMode      float64 `json:"gameMode"`
		IsChatEnabled float64 `json:"isChatEnabled"`
		Bm            struct {
			Time struct {
				FirstObj float64 `json:"firstObj"`
				Current  float64 `json:"current"`
				Full     float64 `json:"full"`
				Mp3      float64 `json:"mp3"`
			} `json:"time"`
			ID           float64 `json:"id"`
			Set          float64 `json:"set"`
			Md5          string  `json:"md5"`
			RankedStatus float64 `json:"rankedStatus"`
			Metadata     struct {
				Artist         string `json:"artist"`
				ArtistOriginal string `json:"artistOriginal"`
				Title          string `json:"title"`
				TitleOriginal  string `json:"titleOriginal"`
				Mapper         string `json:"mapper"`
				Difficulty     string `json:"difficulty"`
			} `json:"metadata"`
			Stats struct {
				Ar  float64 `json:"AR"`
				Cs  float64 `json:"CS"`
				Od  float64 `json:"OD"`
				Hp  float64 `json:"HP"`
				Sr  float64 `json:"SR"`
				Bpm struct {
					Min float64 `json:"min"`
					Max float64 `json:"max"`
				} `json:"BPM"`
				MaxCombo float64 `json:"maxCombo"`
				FullSR   float64 `json:"fullSR"`
				MemoryAR float64 `json:"memoryAR"`
				MemoryCS float64 `json:"memoryCS"`
				MemoryOD float64 `json:"memoryOD"`
				MemoryHP float64 `json:"memoryHP"`
			} `json:"stats"`
			Path struct {
				Full   string `json:"full"`
				Folder string `json:"folder"`
				File   string `json:"file"`
				Bg     string `json:"bg"`
				Audio  string `json:"audio"`
			} `json:"path"`
		} `json:"bm"`
		Mods struct {
			Num float64 `json:"num"`
			Str string  `json:"str"`
		} `json:"mods"`
		Pp struct {
			Num95  float64 `json:"95"`
			Num96  float64 `json:"96"`
			Num97  float64 `json:"97"`
			Num98  float64 `json:"98"`
			Num99  float64 `json:"99"`
			Num100 float64 `json:"100"`
			// Strains []float64 `json:"strains"`
		} `json:"pp"`
	} `json:"menu"`
	Gameplay struct {
		GameMode float64 `json:"gameMode"`
		Name     string  `json:"name"`
		Score    float64 `json:"score"`
		Accuracy float64 `json:"accuracy"`
		Combo    struct {
			Current float64 `json:"current"`
			Max     float64 `json:"max"`
		} `json:"combo"`
		Hp struct {
			Normal float64 `json:"normal"`
			Smooth float64 `json:"smooth"`
		} `json:"hp"`
		Hits struct {
			Num0         float64 `json:"0"`
			Num50        float64 `json:"50"`
			Num100       float64 `json:"100"`
			Num300       float64 `json:"300"`
			Geki         float64 `json:"geki"`
			Katu         float64 `json:"katu"`
			SliderBreaks float64 `json:"sliderBreaks"`
			Grade        struct {
				Current     string `json:"current"`
				MaxThisPlay string `json:"maxThisPlay"`
			} `json:"grade"`
			UnstableRate float64 `json:"unstableRate"`
			// HitErrorArray any `json:"hitErrorArray"`
		} `json:"hits"`
		Pp struct {
			Current     float64 `json:"current"`
			Fc          float64 `json:"fc"`
			MaxThisPlay float64 `json:"maxThisPlay"`
		} `json:"pp"`
		KeyOverlay struct {
			K1 struct {
				IsPressed bool    `json:"isPressed"`
				Count     float64 `json:"count"`
			} `json:"k1"`
			K2 struct {
				IsPressed bool    `json:"isPressed"`
				Count     float64 `json:"count"`
			} `json:"k2"`
			M1 struct {
				IsPressed bool    `json:"isPressed"`
				Count     float64 `json:"count"`
			} `json:"m1"`
			M2 struct {
				IsPressed bool    `json:"isPressed"`
				Count     float64 `json:"count"`
			} `json:"m2"`
		} `json:"keyOverlay"`
		Leaderboard struct {
			HasLeaderboard bool `json:"hasLeaderboard"`
			IsVisible      bool `json:"isVisible"`
			Ourplayer      struct {
				Name      string  `json:"name"`
				Score     float64 `json:"score"`
				Combo     float64 `json:"combo"`
				MaxCombo  float64 `json:"maxCombo"`
				Mods      string  `json:"mods"`
				H300      float64 `json:"h300"`
				H100      float64 `json:"h100"`
				H50       float64 `json:"h50"`
				H0        float64 `json:"h0"`
				Team      float64 `json:"team"`
				Position  float64 `json:"position"`
				IsPassing float64 `json:"isPassing"`
			} `json:"ourplayer"`
			// Slots any `json:"slots"`
		} `json:"leaderboard"`
	} `json:"gameplay"`
	ResultsScreen struct {
		Num0     float64 `json:"0"`
		Num50    float64 `json:"50"`
		Num100   float64 `json:"100"`
		Num300   float64 `json:"300"`
		Name     string  `json:"name"`
		Score    float64 `json:"score"`
		MaxCombo float64 `json:"maxCombo"`
		Mods     struct {
			Num float64 `json:"num"`
			Str string  `json:"str"`
		} `json:"mods"`
		Geki float64 `json:"geki"`
		Katu float64 `json:"katu"`
	} `json:"resultsScreen"`
	Tourney struct {
		Manager struct {
			IpcState float64 `json:"ipcState"`
			BestOF   float64 `json:"bestOF"`
			TeamName struct {
				Left  string `json:"left"`
				Right string `json:"right"`
			} `json:"teamName"`
			Stars struct {
				Left  float64 `json:"left"`
				Right float64 `json:"right"`
			} `json:"stars"`
			Bools struct {
				ScoreVisible bool `json:"scoreVisible"`
				StarsVisible bool `json:"starsVisible"`
			} `json:"bools"`
			Chat     any `json:"chat"`
			Gameplay struct {
				Score struct {
					Left  float64 `json:"left"`
					Right float64 `json:"right"`
				} `json:"score"`
			} `json:"gameplay"`
		} `json:"manager"`
		IpcClients any `json:"ipcClients"`
	} `json:"tourney"`

	// Added data
	MapData string `json:"mapData"`
}

type StreamCompanionResponse struct {
	Acc                          float64   `json:"acc"`
	AccPpIfMapEndsNow            float64   `json:"accPpIfMapEndsNow"`
	AimPpIfMapEndsNow            float64   `json:"aimPpIfMapEndsNow"`
	Ar                           float64   `json:"ar"`
	ArtistRoman                  string    `json:"artistRoman"`
	ArtistUnicode                string    `json:"artistUnicode"`
	BackgroundImageFileName      string    `json:"backgroundImageFileName"`
	BackgroundImageLocation      string    `json:"backgroundImageLocation"`
	BanchoCountry                string    `json:"banchoCountry"`
	BanchoID                     float64   `json:"banchoId"`
	BanchoIsConnected            float64   `json:"banchoIsConnected"`
	BanchoStatus                 float64   `json:"banchoStatus"`
	BanchoUsername               string    `json:"banchoUsername"`
	Bpm                          string    `json:"bpm"`
	C100                         float64   `json:"c100"`
	C300                         float64   `json:"c300"`
	C50                          float64   `json:"c50"`
	ChatIsEnabled                float64   `json:"chatIsEnabled"`
	Circles                      float64   `json:"circles"`
	Combo                        float64   `json:"combo"`
	ComboLeft                    float64   `json:"comboLeft"`
	ConvertedUnstableRate        float64   `json:"convertedUnstableRate"`
	Creator                      string    `json:"creator"`
	Cs                           float64   `json:"cs"`
	CurrentBpm                   float64   `json:"currentBpm"`
	CurrentMaxCombo              float64   `json:"currentMaxCombo"`
	DiffName                     string    `json:"diffName"`
	Dir                          string    `json:"dir"`
	Dl                           string    `json:"dl"`
	Drainingtime                 float64   `json:"drainingtime"`
	FirstHitObjectTime           float64   `json:"firstHitObjectTime"`
	GameMode                     string    `json:"gameMode"`
	Geki                         float64   `json:"geki"`
	Grade                        float64   `json:"grade"`
	HitErrors                    []float64 `json:"hitErrors"`
	Hp                           float64   `json:"hp"`
	IngameInterfaceIsEnabled     float64   `json:"ingameInterfaceIsEnabled"`
	IsBreakTime                  float64   `json:"isBreakTime"`
	Katsu                        float64   `json:"katsu"`
	KeyOverlay                   string    `json:"keyOverlay"`
	Lb                           string    `json:"lb"`
	LeaderBoardMainPlayer        string    `json:"leaderBoardMainPlayer"`
	LeaderBoardPlayers           string    `json:"leaderBoardPlayers"`
	LiveStarRating               float64   `json:"liveStarRating"`
	LocalTime                    string    `json:"localTime"`
	LocalTimeISO                 string    `json:"localTimeISO"`
	MAR                          float64   `json:"mAR"`
	MBpm                         string    `json:"mBpm"`
	MCS                          float64   `json:"mCS"`
	MHP                          float64   `json:"mHP"`
	MMainBpm                     float64   `json:"mMainBpm"`
	MMaxBpm                      float64   `json:"mMaxBpm"`
	MMinBpm                      float64   `json:"mMinBpm"`
	MOD                          float64   `json:"mOD"`
	MStars                       float64   `json:"mStars"`
	MainBpm                      float64   `json:"mainBpm"`
	Mania1000000PP               float64   `json:"mania_1_000_000PP"`
	ManiaM1000000PP              float64   `json:"mania_m1_000_000PP"`
	MapArtistTitle               string    `json:"mapArtistTitle"`
	MapArtistTitleUnicode        string    `json:"mapArtistTitleUnicode"`
	MapDiff                      string    `json:"mapDiff"`
	MapPosition                  string    `json:"mapPosition"`
	Mapid                        float64   `json:"mapid"`
	Mapsetid                     float64   `json:"mapsetid"`
	MaxBpm                       float64   `json:"maxBpm"`
	MaxCombo                     float64   `json:"maxCombo"`
	MaxGrade                     float64   `json:"maxGrade"`
	Md5                          string    `json:"md5"`
	MinBpm                       float64   `json:"minBpm"`
	Miss                         float64   `json:"miss"`
	Mode                         string    `json:"mode"`
	Mods                         string    `json:"mods"`
	ModsEnum                     float64   `json:"modsEnum"`
	Mp3Name                      string    `json:"mp3Name"`
	NoChokePp                    float64   `json:"noChokePp"`
	Od                           float64   `json:"od"`
	OsuFileLocation              string    `json:"osuFileLocation"`
	OsuFileName                  string    `json:"osuFileName"`
	OsuIsRunning                 float64   `json:"osuIsRunning"`
	Osu90PP                      float64   `json:"osu_90PP"`
	Osu95PP                      float64   `json:"osu_95PP"`
	Osu96PP                      float64   `json:"osu_96PP"`
	Osu97PP                      float64   `json:"osu_97PP"`
	Osu98PP                      float64   `json:"osu_98PP"`
	Osu99PP                      float64   `json:"osu_99PP"`
	Osu999PP                     float64   `json:"osu_99_9PP"`
	OsuSSPP                      float64   `json:"osu_SSPP"`
	OsuM90PP                     float64   `json:"osu_m90PP"`
	OsuM95PP                     float64   `json:"osu_m95PP"`
	OsuM96PP                     float64   `json:"osu_m96PP"`
	OsuM97PP                     float64   `json:"osu_m97PP"`
	OsuM98PP                     float64   `json:"osu_m98PP"`
	OsuM99PP                     float64   `json:"osu_m99PP"`
	OsuM999PP                    float64   `json:"osu_m99_9PP"`
	OsuMSSPP                     float64   `json:"osu_mSSPP"`
	PlayerHp                     float64   `json:"playerHp"`
	PlayerHpSmooth               float64   `json:"playerHpSmooth"`
	Plays                        float64   `json:"plays"`
	PpIfMapEndsNow               float64   `json:"ppIfMapEndsNow"`
	PpIfRestFced                 float64   `json:"ppIfRestFced"`
	Previewtime                  float64   `json:"previewtime"`
	RankedStatus                 float64   `json:"rankedStatus"`
	RawStatus                    float64   `json:"rawStatus"`
	Retries                      float64   `json:"retries"`
	Score                        float64   `json:"score"`
	SimulatedPp                  float64   `json:"simulatedPp"`
	Skin                         string    `json:"skin"`
	SkinPath                     string    `json:"skinPath"`
	Sl                           float64   `json:"sl"`
	SliderBreaks                 float64   `json:"sliderBreaks"`
	Sliders                      float64   `json:"sliders"`
	SongSelectionMainPlayerScore float64   `json:"songSelectionMainPlayerScore"`
	SongSelectionRankingType     float64   `json:"songSelectionRankingType"`
	// SongSelectionScores          float64   `json:"songSelectionScores"`
	// SongSelectionTotalScores     float64   `json:"songSelectionTotalScores"`
	Source               string  `json:"source"`
	SpeedPpIfMapEndsNow  float64 `json:"speedPpIfMapEndsNow"`
	Spinners             float64 `json:"spinners"`
	StarsNomod           float64 `json:"starsNomod"`
	Status               float64 `json:"status"`
	StrainPpIfMapEndsNow float64 `json:"strainPpIfMapEndsNow"`
	Sv                   float64 `json:"sv"`
	Tags                 string  `json:"tags"`
	Test                 string  `json:"test"`
	Threadid             float64 `json:"threadid"`
	Time                 float64 `json:"time"`
	TimeLeft             string  `json:"timeLeft"`
	TitleRoman           string  `json:"titleRoman"`
	TitleUnicode         string  `json:"titleUnicode"`
	TotalAudioTime       float64 `json:"totalAudioTime"`
	Totaltime            float64 `json:"totaltime"`
	UnstableRate         float64 `json:"unstableRate"`
	Username             string  `json:"username"`

	// Added data
	MapData string `json:"mapData"`
}

// func getDataFromUrl(url string, wg *sync.WaitGroup, out chan<- bytes[]) ([]byte, error) {
func getDataFromUrl(url string) ([]byte, error) {
	// this is kinda bad because for gosumemory delay will be at the very least 2s
	client := http.Client{
		Timeout: time.Millisecond * 2000, // this will almost never timeout
	}
	response, err := client.Get(url)
	// response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Print("Successfully read data from ", url)
	return data, nil
}

func addMapPathSC(jsonBytes []byte) ([]byte, error) {
	var osuData StreamCompanionResponse
	if err := json.Unmarshal(jsonBytes, &osuData); err != nil {
		return jsonBytes, errors.New("Error decoding JSON: " + err.Error())
	}

	mapPath := osuData.OsuFileLocation
	mapData, err := os.ReadFile(mapPath)
	if err != nil {
		return jsonBytes, errors.New("Error reading map data: " + err.Error())
	}

	// b64MapData := b64.StdEncoding.EncodeToString(mapData)
	// osuData.MapData = b64MapData
	osuData.MapData = string(mapData)
	returnJson, err := json.Marshal(osuData)
	if err != nil {
		return jsonBytes, errors.New("Error reencoding JSON: " + err.Error())
	}
	return returnJson, nil
}

func addMapPathGM(jsonBytes []byte) ([]byte, error) {
	var osuData GosumemoryResponse
	if err := json.Unmarshal(jsonBytes, &osuData); err != nil {
		return jsonBytes, errors.New("Error decoding JSON: " + err.Error())
	}

	mapPath := filepath.Join(osuData.Settings.Folders.Songs, osuData.Menu.Bm.Path.Folder, osuData.Menu.Bm.Path.File)
	mapData, err := os.ReadFile(mapPath)
	if err != nil {
		return jsonBytes, errors.New("Error reading map data: " + err.Error())
	}

	// b64MapData := b64.StdEncoding.EncodeToString(mapData)
	// osuData.MapData = b64MapData
	osuData.MapData = string(mapData)
	returnJson, err := json.Marshal(osuData)
	if err != nil {
		return jsonBytes, errors.New("Error reencoding JSON: " + err.Error())
	}
	return returnJson, nil
}

func OsuDataWithMapData(osuData []byte) []byte {
	// For some reason StreamCompanion adds UTF-8 BOM
	// which is not defined in JSON standard
	osuData = bytes.TrimPrefix(osuData, []byte("\xef\xbb\xbf"))

	if fullData, err := addMapPathSC(osuData); err == nil {
		return fullData
	} else {
		log.Print(err)
	}
	if fullData, err := addMapPathGM(osuData); err == nil {
		return fullData
	} else {
		log.Print(err)
	}
	return osuData
}

// TODO: Rewrite to use goroutines
func getOsuData(urls []string) []byte {
	var errors []string
	for _, url := range urls {
		data, err := getDataFromUrl(url)
		if err == nil {
			return OsuDataWithMapData(data)
		} else {
			log.Print(err)
			errors = append(errors, err.Error())
		}
	}
	errors_map := map[string]interface{}{}
	errors_map["error"] = errors
	errors_json, _ := json.Marshal(errors_map)

	log.Print(string(errors_json))
	return errors_json
}

// var wg sync.WaitGroup
// wg.add(1)
// bytes_chan := make(chan bytes[], len(urls))
// for i, url := range urls {
//     go getDataFromUrl(url, wg, bytes_chan)
// }
// out := <-bytes_chan
// return out,

// func getData()([]byte, error) {
//     jsonData, err := getDataFromGosumemory()
//     if err != nil {
//         log.Fatal("Error getting and unmarshaling JSON data:", err)
//     }

//     // log.Printf("Received JSON data: %+v", jsonData)
//     // log.Printf(jsonData["message"].(string))
// 	// log.Printf(jsonData["status"].(string))
//     return jsonData
// }
