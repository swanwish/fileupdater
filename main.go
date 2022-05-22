package main

import (
	"flag"
	"runtime"

	"github.com/swanwish/fileupdaterdemo/rules"
	"github.com/swanwish/fileupdaterdemo/settings"
	"github.com/swanwish/fileupdaterdemo/updater"
	"github.com/swanwish/go-common/logs"
)

var (
	baseDir string
)

func parseCmdLineArgs() {
	flag.StringVar(&baseDir, "d", "", "The base dir to search files")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parseCmdLineArgs()

	settings.LoadAppSetting()

	fileUpdater := updater.NewFileUpdater(baseDir)

	fileUpdater.RegisterUpdateRule(rules.I18nJavaUpdateRule)
	fileUpdater.RegisterUpdateRule(rules.I18nKotlinUpdateRule)
	fileUpdater.RegisterUpdateRule(rules.I18nObjcUpdateRule)
	fileUpdater.RegisterUpdateRule(rules.I18nSwiftUpdateReule)

	if fileUpdater == nil {
		logs.Errorf("Failed to create i18n updater")
		return
	}

	// Update files
	fileUpdater.UpdateFiles()
}
