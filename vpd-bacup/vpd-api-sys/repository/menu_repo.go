package repository

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func (repo RepositoryImpl) GetMenu() []entities.Menu {
	menuT, err := repo.Db.Query("SELECT menu_id, menu_description, parent_menu_id, menu_path, menu_type, display_order,display_order,menu_label, in_use FROM sys.menu")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	menu := entities.Menu{}
	var menuList []entities.Menu
	for menuT.Next() {
		errr := menuT.Scan(&menu.ID, &menu.Description, &menu.Parent_menu_id, &menu.Path, &menu.Type, &menu.Display_order, &menu.Label, &menu.Icon, &menu.In_use)

		if errr != nil {
			panic("Error while Reading  : " + errr.Error())
		}
		menuList = append(menuList, menu)
	}

	//defer repo.Db.Close()

	return menuList
}

func (repo RepositoryImpl) CreateMenu(menu entities.Menu) {
	_, err := repo.Db.Query("INSERT INTO sys.menu(menu_description, parent_menu_id, menu_path, menu_type, display_order,menu_label,icon, in_use) VALUES ($1,$2, $3,$4, $5, $6,$7,$8)", menu.Description, menu.Parent_menu_id, menu.Path, menu.Type, menu.Display_order, menu.Label, menu.Icon, menu.In_use)
	//defer repo.Db.Close()
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}

}
