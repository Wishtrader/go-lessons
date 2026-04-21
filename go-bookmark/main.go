package main

import "fmt"

type StringMap = map[string]string

func main() {
	fmt.Println("------ Bookmark App ------")
	bookmarks := StringMap{}

Menu:
	for {
		choice := getMenu()
		switch choice {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var choice int
	fmt.Println("Choose an option:")
	fmt.Println("1. Show all bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Delete a bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return choice
}

func printBookmarks(bookmarks StringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("The bookmarks collection is empty.")
		return
	}
	for key, value := range bookmarks {
		fmt.Println(key, "-", value)
	}
	fmt.Println("---------------")
}

func addBookmark(bookmarks StringMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter a bookmark name: ")
	fmt.Scanln(&newBookmarkKey)
	fmt.Print("Enter a bookmark link: ")
	fmt.Scanln(&newBookmarkValue)
	if newBookmarkKey != "" {
		bookmarks[newBookmarkKey] = newBookmarkValue
	}
}

func deleteBookmark(bookmarks StringMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Enter bookmark name to delete: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
