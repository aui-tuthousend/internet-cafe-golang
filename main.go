package main

import (
	"fmt"
	"strconv"
)

type node struct {
	pass int
	// paket int
	nama string
	prev *node
	next *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (list *doublyLinkedList) pushBack(namaK string, pw int) {
	newNode := &node{pass: pw, nama: namaK}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.size++
}
func (list *doublyLinkedList) login(idc string) int {
	node := list.head
	var tempidd int
	var tempaidi string
	var tempnama int
	for node != nil {

		if node.nama == idc {
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&tempidd)
			if tempidd == node.pass {
				var pp int
				fmt.Print(
					"\nTindakan: ",
					"\n[1]. Ubah Username",
					"\n[2]. Ubah Password",
					"\n[3]. Delete member",
					"\n[4]. Menu utama",
					"\nMasukkan Pilihan: ",
				)
				fmt.Scanln(&pp)

				switch pp {
				case 1:
					fmt.Println("\nUpdate member: ")
					fmt.Print("\nMasukkan Nama baru: ")
					fmt.Scanln(&tempaidi)

					node.nama = tempaidi
					fmt.Print("\nUsername berhasil diubah! ")
					// continue

				case 2:
					var ew int
					fmt.Println("\nUpdate member: ")
					fmt.Print("Masukkan Password baru: ")
					fmt.Scanln(&tempnama)
					fmt.Print("Ulangi password baru: ")
					fmt.Scanln(&ew)

					if ew == tempnama {
						node.pass = tempnama
						fmt.Print("\nPassword berhasil diubah! ")
						// continue
					} else {
						fmt.Println("\nSalah bosku")
					}

				case 3:
					var afh string
					fmt.Print("\nYakin dihapus kids? [y/t]: ")
					fmt.Scanln(&afh)

					if afh == "y" {

						if node == list.head {
							list.head = node.next
						} else {
							node.prev.next = node.next
						}

						if node == list.tail {
							list.tail = node.prev
						} else {
							node.next.prev = node.prev
						}

						fmt.Println("\nMember Berhasil dihapus")
						node = nil
						list.size--
					} else {
						fmt.Println("\nOoo kocag")
						break
					}

				case 4:
					// break
					// continue
					return 1
				}

				return 0
			} else {
				fmt.Println("\nPassword Salah")
				return 0
			}
		}
		node = node.next
	}
	fmt.Println("bang udah banh salah bang")
	return 0
}

func (list *doublyLinkedList) PrintForward() {
	nom := 1
	current := list.head
	for current != nil {
		fmt.Print(
			nom,
			"  ",
			current.nama,
			"     \n",)

		current = current.next
		nom++
	}
}

func (list *doublyLinkedList) viewByID(id string, mem string) bool {
	node := list.head

	for node != nil {
		if node.nama == id {
			fmt.Print("\nNama Member: ", node.nama)
			fmt.Println("\nMember: ", mem)
			return true
		}

		node = node.next
	}

	fmt.Println("ID tidak ditemukan")
	return false
}

func (list *doublyLinkedList) getSize() int {
	return list.size
}

func menu() {
	fmt.Println("Pilihan Menu:")
	fmt.Println("[1]. Register Member")
	fmt.Println("[2]. Login")
	fmt.Println("[3]. View")
	fmt.Println("[4]. View By ID")
	fmt.Println("[5]. Menu Utama")
}

func (list *doublyLinkedList) clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func main() {
	list := &doublyLinkedList{}
	list2 := &doublyLinkedList{}
	var lanjut string
	var pilih, pilih1, pilih2, opd int

	fmt.Print(
		"\n\nSoftware Admin Warnet PMA",
		"\nMasukkan Pass Admin: ",
	)
	fmt.Scanln(&opd)
	if opd == 5656 {

		for {
			var tempPass int
			var lojin string

			fmt.Print("\033[H\033[2J")

			fmt.Println("\n------------------------------------------------")
			fmt.Println("          Selamat Datang di Warnet PMA")
			fmt.Println("------------------------------------------------")
			fmt.Println("Data yang tersedia: ")
			fmt.Println("[1]. Member Reguler")
			fmt.Println("[2]. Member VIP")
			fmt.Print("Masukkan data yang akan diolah: ")
			fmt.Scanln(&pilih)

			if pilih == 1 {
				fmt.Println("\nMenu Member Reguler")
				menu()
				fmt.Print("Pilih Menu: ")
				fmt.Scanln(&pilih1)
				switch pilih1 {

				case 1:
					var teem string
					fmt.Println("\nUsername minimal 5 huruf")
					fmt.Print("\nMasukkan Nama: ")
					fmt.Scanln(&teem)
					delapan := len(teem)
					if delapan < 4 {
						fmt.Println("Username minimal 5 huruf ngab hehe")
						break
					} else {
						fmt.Println("Password minimal 4 nomor")
						fmt.Print("Masukkan Password: ")
						fmt.Scanln(&tempPass)
						pasw := strconv.Itoa(tempPass)
						ayaya := len(pasw)

						if ayaya < 3 {
							fmt.Println("Password minimal 4 nomor atuh banh")
							break

						} else {
							list.pushBack(teem, tempPass)
							fmt.Println("\nMember berhasil dibuat")
						}
					}

				case 2:
					if list.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Print("\n   [Login] ")
						fmt.Print("\nMasukkan Nama Member: ")
						fmt.Scanln(&lojin)
						// list.login(lojin)
						con := list.login(lojin)
						if con == 1 {
							continue
						}
					}

				case 3:
					if list.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Println("\n----------------------------------------------------")
						fmt.Println("            Data Member Reguler PMA net")
						fmt.Println("----------------------------------------------------")
						fmt.Println("No Nama     Sisa Billing\n")
						list.PrintForward()
					}

				case 4:
					var id string
					if list.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Print("Masukkan Nama Member: ")
						fmt.Scanln(&id)
						list.viewByID(id, "Reguler")
					}

				case 5:
					lanjut = "y"
					continue

				default:
					fmt.Println("Maaf Menyu tida ada :))")
				}
			} else if pilih == 2 {
				fmt.Println("\nMenu Member VIP  ")
				menu()
				fmt.Print("Pilih Menu: ")
				fmt.Scanln(&pilih2)
				switch pilih2 {
				case 1:
					var teem string
					fmt.Println("\nUsername minimal 5 huruf")
					fmt.Print("\nMasukkan Nama: ")
					fmt.Scanln(&teem)
					delapan := len(teem)
					if delapan < 4 {
						fmt.Println("Username minimal 5 huruf ngab hehe")
						break
					} else {
						fmt.Println("Password minimal 4 nomor")
						fmt.Print("Masukkan Password: ")
						fmt.Scanln(&tempPass)
						pasw := strconv.Itoa(tempPass)
						ayaya := len(pasw)

						if ayaya < 3 {
							fmt.Println("Password minimal 4 nomor atuh banh")
							break

						} else {
							list2.pushBack(teem, tempPass)
							fmt.Println("\nMember berhasil dibuat")
						}
					}

				case 2:
					if list2.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Print("\n   [Login] ")
						fmt.Print("\nMasukkan Nama Member: ")
						fmt.Scanln(&lojin)
						// list.login(lojin)
						con := list2.login(lojin)
						if con == 1 {
							continue
						}
					}

				case 3:
					if list2.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Println("\n----------------------------------------------------")
						fmt.Println("            Data Member Reguler PMA net")
						fmt.Println("----------------------------------------------------")
						fmt.Println("No Nama     Sisa Billing\n")
						list2.PrintForward()
					}

				case 4:
					var id string
					if list2.getSize() == 0 {
						fmt.Println("\nBelum ada data")
						break
					} else {
						fmt.Print("Masukkan Nama Member: ")
						fmt.Scanln(&id)
						list2.viewByID(id, "Reguler")
					}

				case 5:
					lanjut = "y"
					continue

				default:
					fmt.Println("Maaf Menyu tida ada :))")
				}
			}

			fmt.Print("\napakah ingin lanjut? ")
			fmt.Scanln(&lanjut)
			if lanjut == "t" {
				break
			}
		}

		list.clear()
		list2.clear()

	} else {
		fmt.Print("\nPW salah wlee")
	}
}
