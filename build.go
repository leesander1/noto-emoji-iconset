package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var notoDirectory = "noto-emoji"

// Emoji contains the data of an emoji
type Emoji struct {
	Unicode string
	SVG     string
}

func updateNotoEmoji(update bool) {
	_, err := os.Stat(notoDirectory)
	exists := !os.IsNotExist(err)
	if exists && update {
		os.RemoveAll(notoDirectory)
	}
	if !exists || update {
		fmt.Println("Downloading noto-emoji...")
		exec.Command("git", "clone", "https://github.com/googlei18n/noto-emoji", notoDirectory).Run()
		fmt.Println("noto-emoji finished downloading")
	}
}

func unicodeToEmoji(unicode string, separator string) (string, error) {
	unicodeStr := strings.Split(unicode, separator)
	for index, unicodeChar := range unicodeStr {
		hex, err := strconv.ParseInt(unicodeChar, 16, 64)
		if err != nil {
			return "", err
		}
		unicodeStr[index] = string(hex)
	}
	return strings.Join(unicodeStr, ""), nil
}

func fileToEmoji(filePath string, fileName string) (*Emoji, error) {
	emoji := &Emoji{}

	file, err := os.Open(filePath + fileName)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	regexpNewLine := regexp.MustCompile(`\r?\n`)
	regexpTabs := regexp.MustCompile(`\t`)

	regexpHeader := regexp.MustCompile(".*<svg.*?>")
	regexpFooter := regexp.MustCompile("</svg>")

	code := regexpNewLine.ReplaceAllString(string(bytes), "")
	code = regexpTabs.ReplaceAllString(code, "")
	code = regexpHeader.ReplaceAllString(code, "")
	code = regexpFooter.ReplaceAllString(code, "")

	unicode := fileName[len("emoji_u") : len(fileName)-len(".svg")]
	unicodedEmoji, err := unicodeToEmoji(unicode, "_")
	if err != nil {
		return nil, err
	}
	emoji.Unicode = unicodedEmoji
	emoji.SVG = code

	return emoji, nil
}

func readEmojis() []Emoji {
	emojis := []Emoji{}

	path := notoDirectory + "/svg/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsing ", len(files), " elements...")
	for index, file := range files {
		emoji, err := fileToEmoji(path, file.Name())
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error parsing element ", index+1, " of ", len(files))
			continue
		}
		emojis = append(emojis, *emoji)
		fmt.Println("Parsed element ", index+1, " of ", len(files))
	}

	return emojis
}

func writeIconset(data []Emoji) {
	content := `<link rel="import" href="../iron-icon/iron-icon.html">
<link rel="import" href="../iron-iconset-svg/iron-iconset-svg.html">

<iron-iconset-svg name="emoji" size="128">

<svg><defs>
{{ range $element := . -}}
<g id="{{ .Unicode }}">{{ .SVG }}</g>
{{ end -}}
</defs></svg>
</iron-iconset-svg>`

	t := template.New("t")
	t, err := t.Parse(content)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("noto-emoji-iconset.html")
	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}

// Emojione represents an emoji in Emojione data
type Emojione struct {
	Unicode   string `json:"unicode"`
	Shortname string `json:"shortname"`
}

func writeDictionary() {
	fmt.Println("Writing Emoji Dictionary...")
	resp, err := http.Get("https://raw.githubusercontent.com/Ranks/emojione/master/emoji.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var emojione map[string]Emojione
	err = json.Unmarshal(data, &emojione)
	if err != nil {
		panic(err)
	}

	emojis := make(map[string]string)

	fmt.Println("Parsing ", len(emojione), " emojis...")

	index := 0
	for _, emoji := range emojione {
		unicodedEmoji, err := unicodeToEmoji(emoji.Unicode, "-")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error parsing element ", index+1, " of ", len(emojione))
			continue
		}
		emojis[emoji.Shortname] = unicodedEmoji
		fmt.Println("Parsed element ", index+1, " of ", len(emojione))
		index++
	}

	fmt.Println("Parsing emojis as JSON...")
	jsonData, err := json.Marshal(emojis)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("emoji.json", jsonData, 0644)
	fmt.Println("emoji.json file created")
}

func main() {
	updateNoto := flag.Bool("update", false, "update noto emoji")

	flag.Parse()

	updateNotoEmoji(*updateNoto)
	emojis := readEmojis()
	writeIconset(emojis)
	writeDictionary()
}
