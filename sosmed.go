package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
	Friends  [10]int
	Status   [10]string
}

type Comment struct {
	StatusID int
	UserID   int
	Text     string
}

const MaxComments = 100

var comments [MaxComments]Comment
var commentCount int

const MaxUsers = 100

var users [MaxUsers]User
var userCount int

func main() {
	mainMenu()
}

func mainMenu() {
	var choice int
	for {
		fmt.Println("Menu Utama")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		if choice == 1 {
			registerUser()
		} else if choice == 2 {
			loginUser()
		} else if choice == 3 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func registerUser() {
	if userCount >= MaxUsers {
		fmt.Println("Pendaftaran tidak bisa dilakukan, pengguna sudah penuh!")
		return
	}

	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	users[userCount] = User{
		ID:       userCount + 1,
		Username: username,
		Password: password,
	}
	userCount++

	fmt.Println("Registrasi berhasil!")
}

func loginUser() {
	var username, password string
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			fmt.Println("Login berhasil!")
			userHome(users[i].ID)
			return
		}
	}
	fmt.Println("Username atau password salah!")
}

func userHome(userID int) {
	var choice int
	for {
		fmt.Println("Menu Home")
		fmt.Println("1. Lihat Status")
		fmt.Println("2. Tambah Teman")
		fmt.Println("3. Hapus Teman")
		fmt.Println("4. Edit Profil")
		fmt.Println("5. Lihat Teman Terurut")
		fmt.Println("6. Cari Pengguna")
		fmt.Println("7. Tambah Komentar")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		if choice == 1 {
			viewStatus(userID)
		} else if choice == 2 {
			addFriend(userID)
		} else if choice == 3 {
			removeFriend(userID)
		} else if choice == 4 {
			editProfile(userID)
		} else if choice == 5 {
			viewSortedFriends(userID)
		} else if choice == 6 {
			searchUser()
		} else if choice == 7 {
			addComment(userID)
		} else if choice == 8 {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func viewStatus(userID int) {
	fmt.Println("Status Pengguna:")
	for i := 0; i < userCount; i++ {
		for j := 0; j < len(users[i].Status); j++ {
			if users[i].Status[j] != "" {
				fmt.Printf("%s: %s\n", users[i].Username, users[i].Status[j])
				// Menampilkan komentar untuk status ini
				for k := 0; k < commentCount; k++ {
					if comments[k].StatusID == j+1 && comments[k].UserID == users[i].ID {
						commenter := getUserByID(comments[k].UserID)
						if commenter != nil {
							fmt.Printf("  - %s: %s\n", commenter.Username, comments[k].Text)
						}
					}
				}
			}
		}
	}
}

func getUserByID(userID int) *User {
	for i := 0; i < userCount; i++ {
		if users[i].ID == userID {
			return &users[i]
		}
	}
	return nil
}

func addFriend(userID int) {
	var friendUsername string
	fmt.Print("Masukkan username teman: ")
	fmt.Scan(&friendUsername)

	for i := 0; i < userCount; i++ {
		if users[i].Username == friendUsername {
			for j := 0; j < len(users[userID-1].Friends); j++ {
				if users[userID-1].Friends[j] == 0 {
					users[userID-1].Friends[j] = users[i].ID
					fmt.Println("Teman berhasil ditambahkan!")
					return
				}
			}
			fmt.Println("Daftar teman penuh!")
			return
		}
	}
	fmt.Println("Pengguna tidak ditemukan!")
}

func removeFriend(userID int) {
	var friendUsername string
	fmt.Print("Masukkan username teman yang ingin dihapus: ")
	fmt.Scan(&friendUsername)

	for i := 0; i < userCount; i++ {
		if users[i].Username == friendUsername {
			for j := 0; j < len(users[userID-1].Friends); j++ {
				if users[userID-1].Friends[j] == users[i].ID {
					users[userID-1].Friends[j] = 0
					fmt.Println("Teman berhasil dihapus!")
					return
				}
			}
			fmt.Println("Teman tidak ditemukan di daftar teman!")
			return
		}
	}
	fmt.Println("Pengguna tidak ditemukan!")
}

func editProfile(userID int) {
	var newUsername, newPassword string
	fmt.Print("Masukkan username baru: ")
	fmt.Scan(&newUsername)
	fmt.Print("Masukkan password baru: ")
	fmt.Scan(&newPassword)

	users[userID-1].Username = newUsername
	users[userID-1].Password = newPassword
	fmt.Println("Profil berhasil diubah!")
}

func viewSortedFriends(userID int) {
	var choice int
	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Username (Ascending)")
	fmt.Println("2. Username (Descending)")
	fmt.Print("Pilihan: ")
	fmt.Scan(&choice)

	var friends []User
	for _, friendID := range users[userID-1].Friends {
		if friendID != 0 {
			for _, user := range users {
				if user.ID == friendID {
					friends = append(friends, user)
					break
				}
			}
		}
	}

	if choice == 1 {
		selectionSort(friends, true)
	} else if choice == 2 {
		selectionSort(friends, false)
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("Daftar Teman Terurut:")
	for _, friend := range friends {
		fmt.Println(friend.Username)
		for j := 0; j < len(friend.Status); j++ {
			if friend.Status[j] != "" {
				fmt.Printf("%s: %s\n", friend.Username, friend.Status[j])
				for k := 0; k < commentCount; k++ {
					if comments[k].StatusID == j+1 && comments[k].UserID == friend.ID {
						commenter := getUserByID(comments[k].UserID)
						if commenter != nil {
							fmt.Printf("  - %s: %s\n", commenter.Username, comments[k].Text)
						}
					}
				}
			}
		}
	}
}

func searchUser() {
	var username string
	fmt.Print("Masukkan username yang dicari: ")
	fmt.Scan(&username)

	index := binarySearch(users[:userCount], username)
	if index != -1 {
		fmt.Printf("Pengguna ditemukan: %s\n", users[index].Username)
	} else {
		fmt.Println("Pengguna tidak ditemukan!")
	}
}

func addComment(userID int) {
	var statusOwner, commentText string
	var statusID int

	fmt.Print("Masukkan username pemilik status: ")
	fmt.Scan(&statusOwner)

	// Mencari pengguna berdasarkan username
	ownerIndex := -1
	for i := 0; i < userCount; i++ {
		if users[i].Username == statusOwner {
			ownerIndex = i
			break
		}
	}

	if ownerIndex == -1 {
		fmt.Println("Pengguna tidak ditemukan!")
		return
	}

	// Menampilkan status milik pengguna tersebut
	fmt.Println("Status:")
	for i := 0; i < len(users[ownerIndex].Status); i++ {
		if users[ownerIndex].Status[i] != "" {
			fmt.Printf("%d: %s\n", i+1, users[ownerIndex].Status[i])
		}
	}

	fmt.Print("Masukkan ID status yang ingin dikomentari: ")
	fmt.Scan(&statusID)

	if statusID < 1 || statusID > len(users[ownerIndex].Status) || users[ownerIndex].Status[statusID-1] == "" {
		fmt.Println("ID status tidak valid!")
		return
	}

	fmt.Print("Masukkan komentar: ")
	fmt.Scan(&commentText)

	comments[commentCount] = Comment{
		StatusID: statusID,
		UserID:   userID,
		Text:     commentText,
	}
	commentCount++

	fmt.Println("Komentar berhasil ditambahkan!")
}

func selectionSort(arr []User, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j].Username < arr[extremeIdx].Username) ||
				(!ascending && arr[j].Username > arr[extremeIdx].Username) {
				extremeIdx = j
			}
		}
		arr[i], arr[extremeIdx] = arr[extremeIdx], arr[i]
	}
}

func binarySearch(arr []User, username string) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid].Username == username {
			return mid
		} else if arr[mid].Username < username {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
