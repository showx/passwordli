// 1password 密码读取
// Author:shengsheng

package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"encoding/json"
	"strings"
)

type PasswordInfo struct {
	Uuid string
	UpdatedAt     int `json:"updatedAt"`
	LocationKey	string
	SecurityLevel string
	ContentsHash string
	Title string
	Location string
	OpenContents OpenContents
	CreatedAt int
	TypeName string
	SecureContents SecureContents
	Fields map[string]Fields
}

type Fields struct {
	Id string
	Name string
	Value string
	T string
	Designation string
}

type Tags struct {
	
}

type Sections struct {
	Title string
	Name string
}

type SecureContents struct {
	NotesPlain string
	Sections []Sections
	Fields []Fields
}

type OpenContents struct {
//	Tags map[string]string
	ContentsHash string
	Title string
	SecureContents SecureContents
}


func main() {
	
	
	if len(os.Args) < 2 {
		fmt.Println("请输入要获取的URL")
		os.Exit(1)
	}
	keyword := os.Args[1]
	
	fi, err := os.Open("data.1pif")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		
		jsonStr := string(a)
		if jsonStr[0:3] == "***" {
			continue;
		}
		
		var onepassword PasswordInfo
		err := json.Unmarshal([]byte(jsonStr), &onepassword)
		if err != nil {
			fmt.Println(err.Error())
		}
		
		// 不包含的字符串跳过
		if strings.Contains(onepassword.Location, keyword) == false {
			continue;
		}
		
		fmt.Println("-----------------------------");
//		fmt.Println(jsonStr);
		fmt.Println("Uuid:",onepassword.Uuid);
		fmt.Println("Location:",onepassword.Location)
		
		for _,v :=range onepassword.SecureContents.Fields {
			fmt.Println("----------------------------- ^ ^");
			fmt.Println("Designation:", v.Designation)
			fmt.Println("Name:", v.Name)
			fmt.Println("Value:", v.Value)
		}
	}
}