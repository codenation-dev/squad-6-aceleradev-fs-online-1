package db

import "models"

// We'll create a list of jokes
var jokes = []models.Joke{
	models.Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	models.Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	models.Joke{3, 0, "How many apples grow on a tree? All of them."},
	models.Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	models.Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	models.Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	models.Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

//GetJokes retorna lista de Piadas
func GetJokes() []models.Joke {
	return jokes
}
