package entities

type Menu struct {
	ID, Description, Parent_menu_id, Path, Type, Label, Icon, In_use string
	Display_order                                                    int
	Children                                                         []Menu
}
