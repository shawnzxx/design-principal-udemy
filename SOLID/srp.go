package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

/* This file includes of demonstration about Single Responsibility Principal
 */

// Journal consider a SRP struct which store entries and manipulate entries, so all realted functions within same package
var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(test string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, test)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	fmt.Printf("removing index %d\n", index)
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Below operation function all break SRP, it is separation of concerns which related to persistence
// if we put below persistence logic under same Journal object, how about other struct also needs persistence capability?

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// we shall put all persistence concern into separate type/package let say persistence package
// this way we can have cutting cross settings for all kinds of type, like you want persistence Books
// we create common separator which used by other types also
var lineSeparator = "\n"

// SaveToFile put into different package
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// or put into different type
type Persistence struct {
	linSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.linSeparator)), 0644)
}

func main() {
	// Journal related operations
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	j.AddEntry("I coding.")
	fmt.Println(strings.Join(j.entries, "\n"))
	j.RemoveEntry(2)
	fmt.Println(strings.Join(j.entries, "\n"))

	// 1st way to keep SRP: use a separate function
	SaveToFile(&j, "journal1.txt")

	// 2nd way to keep SRP: use a new struct
	p := Persistence{"\n"}
	p.saveToFile(&j, "journal2.txt")
}
