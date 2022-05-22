package updater

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/swanwish/fileupdaterdemo/common"
	"github.com/swanwish/fileupdaterdemo/rules"
	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
)

type fileUpdater struct {
	baseDir      string
	fileExts     []string
	updateReules []rules.UpdateRule
	extRuleMap   map[string][]rules.UpdateRule
}

func NewFileUpdater(baseDir string) *fileUpdater {
	if !utils.FileExists(baseDir) {
		logs.Errorf("The base dir %s does not exists", baseDir)
		return nil
	}

	return &fileUpdater{
		baseDir:      baseDir,
		fileExts:     make([]string, 0),
		updateReules: make([]rules.UpdateRule, 0),
		extRuleMap:   make(map[string][]rules.UpdateRule),
	}
}

func (updater *fileUpdater) RegisterUpdateRule(updateRule rules.UpdateRule) {
	for _, ext := range updateRule.Exts {
		updateRules, exists := updater.extRuleMap[ext]
		if !exists {
			updateRules = make([]rules.UpdateRule, 0)
		}
		updateRules = append(updateRules, updateRule)
		updater.extRuleMap[ext] = updateRules
		updater.addFileExt(ext)
	}
}

func (updater fileUpdater) UpdateFiles() error {
	if !utils.FileExists(updater.baseDir) {
		logs.Errorf("base dir %s does not exists", updater.baseDir)
		return common.ErrNotExists
	}

	files := updater.listFiles(updater.baseDir)
	for _, file := range files {
		ext := path.Ext(file)
		if updateRules, exists := updater.extRuleMap[ext]; exists {
			for _, updateRule := range updateRules {
				if err := updater.updateFile(file, updateRule); err != nil {
					logs.Errorf("Failed to update file %s, the error is %#v", file, err)
				}
			}
		}
	}
	return nil
}

func (updater *fileUpdater) addFileExt(ext string) {
	for _, fileExt := range updater.fileExts {
		if fileExt == ext {
			return
		}
	}
	updater.fileExts = append(updater.fileExts, ext)
}

func (updater fileUpdater) listFiles(searchDir string) []string {
	items, err := ioutil.ReadDir(searchDir)
	if err != nil {
		logs.Errorf("Failed to get items from dir %s, the error is %#v", searchDir, err)
		return nil
	}

	files := make([]string, 0)
	for _, item := range items {
		itemName := item.Name()
		subDir := filepath.Join(searchDir, itemName)
		if item.IsDir() {
			subFiles := updater.listFiles(subDir)
			if err != nil {
				logs.Errorf("Failed to get sub file, the error is %#v", err)
				return nil
			}
			files = append(files, subFiles...)
		} else {
			fileExt := path.Ext(itemName)
			for _, ext := range updater.fileExts {
				if fileExt == ext {
					files = append(files, filepath.Join(searchDir, itemName))
				}
			}
		}
	}
	return files
}

func (updater fileUpdater) updateFile(filePath string, updateRule rules.UpdateRule) error {
	if !utils.FileExists(filePath) {
		logs.Errorf("The file %s does not exists", filePath)
		return common.ErrNotExists
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		logs.Errorf("Failed to get file stat, the error is %#v", err)
		return err
	}

	byteContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		logs.Errorf("Failed to read file %s, the error is %#v", filePath, err)
		return err
	}

	content := string(byteContent)

	ruleRegexp, err := regexp.Compile(updateRule.Pattern)
	if err != nil {
		logs.Errorf("Failed to compile pattern %s, the error is %#v", updateRule.Pattern, err)
		return err
	}
	result := ruleRegexp.FindAllStringSubmatch(content, -1)
	if len(result) == 0 {
		logs.Debugf("There are no match on file %s", filePath)
		return nil
	}

	updated := false
	for _, match := range result {
		if len(match) > 1 {
			replacer := updateRule.GetMatchReplacer(match)
			if replacer == nil {
				continue
			}
			content = strings.ReplaceAll(content, match[0], *replacer)
			updated = true
		} else {
			logs.Errorf("The length of the match is 0")
		}
	}

	if updated {
		if err = ioutil.WriteFile(filePath, []byte(content), fileInfo.Mode().Perm()); err != nil {
			logs.Errorf("Failed to update file %s", filePath)
		}
	}

	return nil
}
