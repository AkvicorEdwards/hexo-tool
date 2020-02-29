package main

import (
	"bufio"
	"fmt"
	"hexo-tool/utilities"
	"os"
	"strings"
	"time"
)

const dir string = "/home/akvicor/Dropbox/Hexo/blog"
const dirRoot string = "/home/akvicor/Dropbox/Hexo"
const terminal string = "xfce4-terminal -x "
const fileManager string = "ranger "

const menu = `
***********Menu*****************
* 0. Parse and upload          *
* 1. Generate new blog         *
* 2. Clean up old files        *
* 3. Run locally               *
* 4. Compress backup to tgz    *
* 5. Open blog file directory  *
* 6. Open resource directory   *
* 7. Open the root directory   *
* 8. Open backup folder        *
********************************`

const menuBlog = `
***********Menu*****************
* 0. Page                      *
* 1. OI                        *
* 2. Normal                    *
********************************`


func handleChoice(choice string) {
	if choice == "0" {
		utilities.Execute("hexo g -d", dir)
	}else if choice == "1" {
		blogType := ""
		blogTitle := ""
		fmt.Println(menuBlog)
		fmt.Println("Enter your choice: ")
		_, err := fmt.Scanln(&blogType)
		if err != nil {
			fmt.Println("Illegal input")
		}
		fmt.Println("Enter title: ")
		blogTitle, err = bufio.NewReader(os.Stdin).ReadString('\n')
		blogTitle = strings.TrimSpace(blogTitle)
		if err != nil {
			fmt.Println("Illegal input")
		}
		if blogType == "0" {
			utilities.Execute(`hexo new page "`+blogTitle+`"`, dir)
		}else if blogType == "1" {
			utilities.Execute(`hexo new oi "`+blogTitle+`"`, dir)
		}else if blogType == "2" {
			utilities.Execute(`hexo new "`+blogTitle+`"`, dir)
		}
		utilities.Execute(terminal + fileManager + `"` + dir + `/source/_posts"`, dir)
	}else if choice == "2" {
		utilities.Execute("hexo clean", dir)
	}else if choice == "3" {
		utilities.Execute("hexo s", dir)
	}else if choice == "4" {
		utilities.Execute("hexo clean", dir)
		utilities.Execute(fmt.Sprintf("tar zcvf '%s.tar.gz' blog", time.Now().Format("2006-01-02 15-04-05")), dirRoot)
	}else if choice == "5" {
		utilities.Execute(terminal + fileManager + `"` + dir + `/source/_posts" &`, dir)
	}else if choice == "6" {
		utilities.Execute(terminal + fileManager + `"` + dir + `/source" &`, dir)
	}else if choice == "7" {
		utilities.Execute(terminal + fileManager + `"` + dir + `"  &`, dir)
	}else if choice == "8" {
		utilities.Execute(terminal + fileManager + `. &`, dirRoot)
	}

}

func main() {
	var choice string

	for {
		fmt.Println(menu)
		fmt.Println("Enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Illegal Input")
		}
		handleChoice(choice)
	}
}
