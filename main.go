package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	const op = "main()"

	var pathFile = "./names.csv"
	csvFile, err := getCSVFile(pathFile)
	if err != nil {
		log.Printf("%s: open csv file error > %s", op, err)
	}

	defer csvFile.Close()

	props, err := getCsvProps(csvFile)
	if err != nil {
		log.Printf("%s: get csv props error > %s", op, err)
	}

	pathResult := "./RESULT.DAT"

	err = writePropsToResult(props, pathResult)
	if err != nil {
		log.Printf("%s: write props to result error > %s", op, err)
	}
}

// getCSVFile return file from path
func getCSVFile(path string) (*os.File, error) {
	const op = "openCSV"

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("%s: open csv file error > %s", op, err)
		return nil, err
	}

	return f, nil
}

// getCsvProps return all records from csv file
func getCsvProps(f *os.File) ([][]string, error) {
	const op = "getCsvProps()"

	// create new reader
	reader := csv.NewReader(f)
	reader.Comma = rune(';')
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("%s: get csv props error > %s", op, err)
		return nil, err
	}

	return records, nil
}

// writePropsToResult write slice op props to result
func writePropsToResult(props [][]string, result string) error {
	const op = "writePropsToResult()"

	f, err := os.OpenFile(result, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("%s: open result file error > %s", op, err)
		return err
	}
	defer f.Close()

	head := "CDA7359AB6408E7F0088CAB68470D5FE" +  "\n" +  "\n"
	_, err = f.WriteString(head)
		if err != nil {
			log.Printf("%s: write head to file error > %s", op, err)
			return err
		}
	var index string

	l := len(props)
	for i := 0; i < l; i++ {
		switch {
		case i < 9:
			index = "[00" + strconv.Itoa(i + 1) + "]" + "\n"
		case i < 99 && i >= 9 :
			index = "[0" + strconv.Itoa(i + 1) + "]" + "\n"
		case i < 999 && i >= 99:
			index = "[" + strconv.Itoa(i + 1) + "]" + "\n"
		case i > 99:
			index = "[" + strconv.Itoa(i + 1) + "]" + "\n"
		}

		name := props[i][0] + ",,34567;" + "\n"

		str := index + name

		_, err := f.WriteString(str)
		if err != nil {
			log.Printf("%s: write string to file error > %s", op, err)
			return err
		}
	}

	return nil
}