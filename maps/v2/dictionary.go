package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}

func Search(dictionary map[string]string, word string) string {

	return dictionary[word]

}

func main() {}
