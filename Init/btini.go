package Init

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type section map[string]string

type File struct {
	sections map[string]section
}

func InitReader(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return filestream(f), nil
}
func filestream(f io.Reader) *File {
	var linews, line string
	var err error
	single := make(map[string]section)
	r := bufio.NewReader(f)
	title := ""
	for err == nil {
		line, err = r.ReadString('\n')
		linews = strings.TrimSpace(line)
		if linews == "" || linews[0] == ';' || linews[0] == '#' {
			continue
		}
		if linews[0] == '[' && linews[len(linews)-1] == ']' {
			title = linews[1: len(linews)-1]
			_, ok := single[title]
			if !ok {
				single[title] = make(section)
			}
			continue
		}
		if title == "" {
			continue
		}
		if linews[0] == '\'' || linews[0] == '"' {
			key, value := keywithpoint(line, string(linews[0]))
			single[title][key] = value
			continue
		}
		key, value := keywithoutpoint(line)
		single[title][key] = value
	}
	return &File{single}
}

func keywithpoint(line, tou string) (string, string) {
	var key string
	var value string
	first := strings.IndexAny(line, tou)
	last := strings.IndexAny(line[first+1:], tou)
	if last == -1 {
		linews := strings.TrimSpace(line)
		location := strings.IndexAny(linews, "=")
		key = linews[:location]
	} else {
		key = line[first:last+2]
	}
	location := strings.IndexAny(line, "=")
	if location+1 != len(line) {
		value_source := line[location+1:]
		var value_f int
		var value_l int
		for i := 0; i < len(value_source); i++ {
			if value_source[i] != ' ' {
				value_f = i
				break
			}
		}
		for i := len(value_source) - 1; i > 0; i-- {
			if value_source[i] != ' ' {
				value_l = i
				break
			}
		}
		if value_f == value_l {
			value = string(value_source[value_f])
			value = value[:len(value)-1]
			if value[len(value)-1:] == "\r" {
				value = value[:len(value)-1]
			}
		} else {
			value = value_source[value_f:value_l]
			value = value[:len(value)-1]
			if value[len(value)-1:] == "\r" {
				value = value[:len(value)-1]
			}
		}
	}
	return key, value
}

func keywithoutpoint(line string) (string, string) {
	eqpoint := strings.IndexAny(line, "=")
	key_source := line[:eqpoint]
	value_source := line[eqpoint+1:]
	key := getstringwithspace(key_source)
	value := getstringwithspace(value_source)
	value = value[:len(value)-1]
	if value[len(value)-1:] == "\r" {
		value = value[:len(value)-1]
	}
	return key, value
}

func getstringwithspace(str string) string {
	var result string
	var first int
	var last int
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			first = i
			break
		}
	}
	for i := len(str) - 1; i > 0; i-- {
		if str[i] != ' ' {
			last = i
			break
		}
	}
	if first == last {
		result = string(str[first])
	} else {
		result = str[first:last+1]
	}
	return result
}

func (f *File) Value(title, key string) string {
	single, ok := f.sections[title]
	if ok == false {
		return ""
	}
	result, ok := single[key]
	if ok == false {
		return ""
	}
	return result
}
