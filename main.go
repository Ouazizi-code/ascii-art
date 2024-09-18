package main

import (
	"fmt"
	//"os"
)

type Books struct {
	Title string
	Author string
	ISBN string
	Available bool
}
type Member struct {
	Name string
	MemberID int
	BorrowedBooks []Books
}

func addBook(title, author, isbn string) Books {
    return Books{Title: title, Author: author, ISBN: isbn, Available: true}
}

func registerMember(name string, id int) Member {
    return Member{Name: name, MemberID: id}
}

func borrowBook(member *Member, book *Books) {
    if book.Available {
        book.Available = false
        member.BorrowedBooks = append(member.BorrowedBooks, *book)
        fmt.Printf("%s borrowed %s\n", member.Name, book.Title)
    } else {
        fmt.Printf("%s is not available\n", book.Title)
    }
}

func returnBook(member *Member, book *Books) {
    for i, b := range member.BorrowedBooks {
        if b.ISBN == book.ISBN {
            book.Available = true
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            fmt.Printf("%s returned %s\n", member.Name, book.Title)
            return
        }
    }
    fmt.Printf("%s did not borrow %s\n", member.Name, book.Title)
}

func main() {
	book1 := addBook("The Go Programming Language", "Alan A. A. Donovan", "978-0134190440")
    book2 := addBook("Learning Go", "Jon Bodner", "978-1492077213")

    member1 := registerMember("Alice", 1)
    member2 := registerMember("Bob", 2)

    borrowBook(&member1, &book1)
    borrowBook(&member2, &book2)
    returnBook(&member1, &book1)
    borrowBook(&member2, &book1)

    fmt.Println("Books:", book1, book2)
    fmt.Println("Members:", member1, member2)
}


