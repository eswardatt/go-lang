package services

import (
	"encoding/json"
	"fmt"

	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

// type Menu struct {
// 	ID, Description, Parent_menu_id, Path, Type, Label, Icon, In_use string
// 	Display_order                                                    int
// 	Children                                                         []Menu `json:"Children,omitempty"`
// }

// func (this Menu) Size() int {
// 	var size int = len(this.Children)
// 	for _, c := range this.Children {
// 		size += c.Size()
// 	}
// 	return size
// }
// func (this Menu) Add(nodes ...Menu) bool {
// 	fmt.Println(this)
// 	fmt.Println("entered into add ..")
// 	var size = this.Size()
// 	for _, n := range nodes {
// 		fmt.Println(n)
// 		fmt.Println(this.ID)
// 		if n.Parent_menu_id == this.ID {
// 			fmt.Println("entered into if n.parent=this.id..")
// 			this.Children = append(this.Children, n)
// 			fmt.Println(this.Children)
// 		} else {
// 			for _, c := range this.Children {
// 				fmt.Println("entered in into else")
// 				if c.Add(n) {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return this.Size() == size+len(nodes)
// }

func (serv ServiceImpl) AddMenu(menu entities.Menu) {
	serv.Repo.CreateMenu(menu)
}

func (serv ServiceImpl) GetMenu() []entities.Menu {
	parentList := []entities.Menu{}
	childrenList := []entities.Menu{}

	for _, row := range serv.Repo.GetMenu() {
		if row.Parent_menu_id == "" {
			parentList = append(parentList, row)
		} else {
			childrenList = append(childrenList, row)
		}
	}
	for _, parent := range parentList {
		for _, children := range childrenList {
			if children.Parent_menu_id == parent.ID {
				fmt.Println("parent.ID=children.Parent_menu_id")
				parent.Children = append(parent.Children, children)
				json.Marshal(parent)
				fmt.Println(parent)

			}
		}
	}
	fmt.Println(parentList)
	fmt.Println(childrenList)
	//	fmt.Println(menuvo)

	// fmt.Println(objmenu.Add(menuvo...), objmenu.Size())
	// bytes, _ := json.MarshalIndent(objmenu, "", "\t") //formated output
	// //bytes, _ := json.Marshal(root)
	// fmt.Println(string(bytes))
	return serv.Repo.GetMenu()
}
