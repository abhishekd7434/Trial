package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	// If the file doesn't exist, create it or append to the file
	t := time.Now()
	cursec := (t.Hour()*60+t.Minute())*60 + t.Second()
	atsec := (12*60 + 44) * 60
	r := atsec - cursec
	name := "initial"
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(r)
		if r > 0 {

			ioutil.WriteFile("./"+name, []byte("Contents"), 0)
		} else {
			name = GetFilenameDate()
			ioutil.WriteFile("./"+name, []byte("Contents"), 0)
			fmt.Println("New Day New File")
			atsec = atsec + 5*60
		}

	}

}

func GetFilenameDate() string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return "file-" + t.Format(layout) + ".txt"
}

// func LogWithTime(bs []byte) (n int, err error) {
// 	t := time.Now()
// 	cursec := (t.Hour()*60+t.Minute())*60 + t.Second()
// 	atsec := 24 * 60 * 60
// 	r := atsec - cursec
// 	name := "initial"
// 	if r > 0 {
// 		file, err := os.OpenFile("./logs"+name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		log.SetOutput(file)
// 	} else {
// 		name = GetFilenameDate()
// 		file, err := os.OpenFile("./logs"+name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.SetOutput(file)
// 		fmt.Println("New Day New File")
// 	}

// 	return len(bs), nil
// }
