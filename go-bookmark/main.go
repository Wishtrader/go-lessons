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
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		case 4:
			break Menu
		}	
	}
}

func getMenu() int8 {
	var choice int8
	fmt.Println("Chose an option:")
	fmt.Println("1. Show all bookmarks")
	fmt.Println("2. Add a bookmark")
	fmt.Println("3. Delete a bookmark")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return choice
}

func printBookmarks(bookmarks StringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("The bookmarks is empty.")
	}
	for key, value := range bookmarks {
		fmt.Println(key, "-", value)
	}
	fmt.Println("---------------")
}

func addBookmark(bookmarks StringMap) StringMap{
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Enter a bookmark name: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Enter a bookmark link: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	return bookmarks
}

func deleteBookmark(bookmarks StringMap) StringMap{
	var bookmarkKeyToDelete string
	fmt.Print("Enter bookmark name to delete: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
