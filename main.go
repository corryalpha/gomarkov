package gomarkov

import (
	"math/rand"
	"strings"
	"time"
)

var model []string

// Clear clears the model completely.
func Clear() {
	model = []string{}
}

// Feed is used to add ONE string to the model.
func Feed(data string) {
	model = append(model, strings.ToLower(data))
}

// FeedMany is used to add more than one string to the model.
func FeedMany(data []string) {

	for i := 0; i <= len(data)-1; i++ {

		data[i] = strings.ToLower(data[i])

	}

	model = append(model, data...)

}

// Sentence is used to generate a sentence using the model.
func Sentence(wordcount int) string {
	rand.Seed(time.Now().UnixNano())

	var sentence []string
	for i := 0; i <= wordcount; i++ { // Repeat this as many times as our wordcount

		randNum := rand.Intn(len(model)) // Generates a random number to use as index of our model

		sentence = append(sentence, model[randNum]) // Append our random word in model to our new sentence

	}

	// Lets format our sentence now

	var finalSentence string

	for i := 0; i <= len(sentence)-1; i++ { // Loop thru each word in our random words
		finalSentence = finalSentence + " " + sentence[i] // Append it
	}

	return finalSentence

}
