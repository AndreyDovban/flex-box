package block

import (
	"flexbox/flex"
	"flexbox/tree"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/** Tree Block */
func TreeBlock() *fyne.Container {

	arr := [][]string{
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___zlobanov205"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___zlobanov205"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___zlobanov205_exxff7hzhu74irsh"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___lkudryavcev209"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___lkudryavcev209"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___lkudryavcev2_3cabm0o6yrkhh5a"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___vpopov266"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___vpopov266"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___vpopov266_2sendaaixue7ib5q48"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___lbelyaev986"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___lbelyaev986"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___lbelyaev986_qlz6zu0wlwjfxbd2"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___agorshkov203"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___agorshkov203"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___agorshkov203_5gh5pko0gcbb9ko"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___visaev948"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___visaev948"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___visaev948_5vdz1xy39lseo0hv45"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users", "uid=___vdementev422"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=groups", "cn=___vdementev422"},
		{"dc=granulex,dc=test", "cn=accounts", "cn=users_history", "cn=___vdementev422_3xzsryytfzs7dm8"},
	}

	treeList := tree.BuidHashList(arr)

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return treeList[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := treeList[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			_, ok := treeList[uid]
			if !ok {
				// fyne.LogError("Missing tutorial panel: "+uid, nil)
				obj.(*widget.Label).SetText(uid)
			}
			obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			_, ok := treeList[uid]
			if ok {
				log.Println(uid)
			} else {
				log.Println(uid)
			}
		},
	}
	tree.Resize(fyne.NewSize(900, 600))
	tree.HideSeparators = true

	treeContent := container.New(
		flex.NewFlexBox("column", "center", "start", 10, 10), tree,
	)

	return treeContent
}
