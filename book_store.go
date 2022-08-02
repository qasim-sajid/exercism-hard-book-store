package bookstore

import (
	"sort"
)

const (
	Book_PRICE = 800
)

var bookDiscounts []float64 = []float64{1, 0.95, 0.9, 0.8, 0.75}

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	if len(books) == 1 {
		return Book_PRICE
	}

	sortedBooks := getArrangedSortedBooksSlice(books)

	minLevel := 2
	maxLevel := getMaxLevel(sortedBooks)

	cheapestCost := len(books) * Book_PRICE

	for i := minLevel; i <= maxLevel; i++ {
		booksCombination := getBooksCombination(sortedBooks, i)
		booksCombinationCost := getBooksCombinationCost(booksCombination)

		if booksCombinationCost < cheapestCost {
			cheapestCost = booksCombinationCost
		}
	}

	return cheapestCost
}

//Get a books combination
func getBooksCombination(books []int, maxLevel int) []int {
	booksCombination := make([]int, books[0])

	currentIndex := 0
	for i := 0; i < len(books); i++ {
		if books[i] <= 0 {
			return booksCombination
		}

		for j := 0; j < books[i]; j++ {
			booksCombination[currentIndex]++

			currentIndex++
			if currentIndex >= len(booksCombination) {
				currentIndex = 0
			}
		}

		if booksCombination[0] < maxLevel {
			currentIndex = 0
		}
	}

	return booksCombination
}

//Get maximum book combinations level
func getMaxLevel(books []int) int {
	maxLevel := len(books)

	for i := len(books) - 1; i >= 0; i-- {
		if books[i] > 0 {
			break
		}

		maxLevel--
	}

	return maxLevel
}

//Get the total cost of a specific books combination
func getBooksCombinationCost(booksCombination []int) int {
	booksCombinationCost := 0
	for i := 0; i < len(booksCombination); i++ {
		booksCombinationCost += int(float64(booksCombination[i]*Book_PRICE) * bookDiscounts[booksCombination[i]-1])
	}
	return booksCombinationCost
}

//Arrange and sort the books slice
func getArrangedSortedBooksSlice(books []int) []int {
	b := make([]int, 5)

	for _, v := range books {
		b[v-1]++
	}

	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	return b
}
