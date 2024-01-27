package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cast"

	"github.com/kdancybot/np-manager/config"

	"github.com/kdancybot/np-manager/gui"
	"github.com/kdancybot/np-manager/mem"
	"github.com/kdancybot/np-manager/memory"
	"github.com/kdancybot/np-manager/np"
	"github.com/kdancybot/np-manager/updater"
)

func ChangeLogDestinationToFile() {
	f, err := os.OpenFile("npc.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	// Not explicitly closing file is bad,
	// but it shouldn't become a problem with only one config file opened
}

func main() {
	ChangeLogDestinationToFile()
	config.Init()
	updateTimeFlag := flag.Int("update", cast.ToInt(config.Config["update"]), "How fast should we update the values? (in milliseconds)")
	shouldWeUpdate := flag.Bool("autoupdate", true, "Should we auto update the application?")
	isRunningInWINE := flag.Bool("wine", cast.ToBool(config.Config["wine"]), "Running under WINE?")
	songsFolderFlag := flag.String("path", config.Config["path"], `Path to osu! Songs directory ex: /mnt/ps3drive/osu\!/Songs`)
	memDebugFlag := flag.Bool("memdebug", cast.ToBool(config.Config["memdebug"]), `Enable verbose memory debugging?`)
	memCycleTestFlag := flag.Bool("memcycletest", cast.ToBool(config.Config["memcycletest"]), `Enable memory cycle time measure?`)
	flag.Parse()
	mem.Debug = *memDebugFlag
	memory.MemCycle = *memCycleTestFlag
	memory.UpdateTime = *updateTimeFlag
	memory.SongsFolderPath = *songsFolderFlag
	memory.UnderWine = *isRunningInWINE
	if runtime.GOOS != "windows" && memory.SongsFolderPath == "auto" {
		log.Fatalln("Please specify path to osu!Songs (see --help)")
	}
	if memory.SongsFolderPath != "auto" {
		if _, err := os.Stat(memory.SongsFolderPath); os.IsNotExist(err) {
			log.Fatalln(`Specified Songs directory does not exist on the system! (try setting to "auto" if you are on Windows or make sure that the path is correct)`)
		}
	}
	if *shouldWeUpdate {
		updater.DoSelfUpdate()
	}

	go memory.Init()
	// err := db.InitDB()
	// if err != nil {
	// 	log.Println(err)
	// 	time.Sleep(5 * time.Second)
	// 	os.Exit(1)
	// }
	go np.SetupStructure()
	go np.WsConnectionHandler()

	gui.Start()
}
