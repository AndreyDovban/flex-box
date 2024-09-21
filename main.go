package main

import (
	"flexbox/flex"
	"flexbox/mywidgets"
	"image/color"
	"log"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func buidHashList(objects [][]string) map[string][]string {
	hash := map[string][]string{}

	for _, arr := range objects {
		for i := 0; i < len(arr)-1; i++ {
			_, ok := hash[arr[i]]
			if ok {
				children := hash[arr[i]]
				if !slices.Contains(children, arr[i+1]) {
					children = append(children, arr[i+1])
					hash[arr[i]] = children
				}
			} else {
				if i != len(arr)-1 {
					hash[arr[i]] = []string{}
				}
			}
		}
	}

	hash[""] = []string{objects[0][0]}

	return hash
}

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Learn")

	/** Flex block */

	data := binding.NewString()

	direction := "column"
	align := "center"
	justify := "center"

	vid := mywidgets.NewMyWidget("click")
	vid.Resize(fyne.NewSize(200, 40))

	// vid := widget.NewLabel("click")
	// vid.Resize(fyne.NewSize(200, 40))

	num := mywidgets.NewNumericalEntry()
	num.Resize(fyne.NewSize(200, 40))

	text1 := widget.NewLabel("First Label")
	text1.Wrapping = fyne.TextWrapBreak
	text2 := widget.NewLabel("Middle Label")
	text3 := widget.NewLabel("bottomright")

	paragraf := widget.NewRichText(
		&widget.TextSegment{
			Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."})
	paragraf.Resize(fyne.NewSize(160, 100))
	paragraf.Wrapping = fyne.TextWrapBreak
	paragraf.Truncation = fyne.TextTruncateEllipsis

	flexLayout := flex.NewFlexBox(direction, align, justify, 10, 10)
	flexBlock := container.New(
		flexLayout,
		num,
		vid,
		text3,
		paragraf,
		text2,
	)

	buttons := container.NewVBox(
		widget.NewButton("Direction row", func() { data.Set("row"); direction = "row" }),
		widget.NewButton("Direction column", func() { data.Set("column"); direction = "column" }),
		widget.NewButton("Align start", func() { data.Set("start"); align = "start" }),
		widget.NewButton("Align center", func() { data.Set("center"); align = "center" }),
		widget.NewButton("Align end", func() { data.Set("end"); align = "end" }),
		widget.NewButton("Justify start", func() { data.Set("start"); justify = "start" }),
		widget.NewButton("Justify center", func() { data.Set("center"); justify = "center" }),
		widget.NewButton("Justify emd", func() { data.Set("end"); justify = "end" }),
		widget.NewButton("Justify around", func() { data.Set("around"); justify = "around" }),
		widget.NewButton("Justify between", func() { data.Set("between"); justify = "between" }),
	)

	flexTabContent := container.NewBorder(
		nil, nil, buttons, nil, flexBlock)

	data.AddListener(binding.NewDataListener(func() {

		log.Println(flexBlock.Layout)

		flexTabContent.Remove(flexBlock)
		flexLayout.Dir = direction
		flexLayout.Align = align
		flexLayout.Justify = justify

		flexTabContent.Add(flexBlock)
		flexTabContent.Refresh()
		log.Println(direction, align, justify)
	}))

	/** Shortcat example */

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}

	mainWindow.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})

	/** Tree Block */

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

	treeList := buidHashList(arr)

	log.Println(treeList)

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

	treeContent := container.New(
		flex.NewFlexBox("column", "center", "start", 10, 10), tree,
	)

	/** List With Data block */

	hr := canvas.NewLine(color.White)
	hr.Resize(fyne.NewSize(200, 0))

	names := binding.NewStringList()
	names.Set([]string{"Jon", "Bob", "Kate", "Frank", "Andrey"})

	addElem := widget.NewButton("Add elem", func() {
		names.Append("New text")
	})

	// list := widget.NewList(
	// 	func() int { return len(names) },
	// 	func() fyne.CanvasObject { return widget.NewLabel("") },
	// 	func(id widget.ListItemID, co fyne.CanvasObject) {
	// 		co.(*widget.Label).SetText(names[id])
	// 	},
	// )

	list := widget.NewListWithData(
		names,
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(di binding.DataItem, co fyne.CanvasObject) {
			txt, _ := di.(binding.String).Get()
			co.(*widget.Label).SetText(txt)
		},
	)
	list.Resize(fyne.NewSize(200, 400))
	// list.OnSelected = func(id widget.ListItemID) { log.Println(names[id]) }
	list.OnSelected = func(id widget.ListItemID) {
		s, _ := names.GetItem(id)
		text, _ := s.(binding.String).Get()
		log.Println(text)

	}

	listContent := container.New(
		flex.NewFlexBox("column", "center", "start", 10, 10),
		addElem,
		hr,
		list,
	)

	/** Custom wibget block*/

	customWidgetContent := container.New(
		flex.NewFlexBox("column", "center", "center", 10, 10),
		widget.NewButton("click", nil),
	)

	/** App Tabs block */

	tab := container.NewAppTabs(
		container.NewTabItem("Other", treeContent),
		container.NewTabItem("List", listContent),
		container.NewTabItem("CustomWidget", customWidgetContent),
		container.NewTabItem("Flex", flexTabContent),
	)

	mainWindow.SetContent(tab)

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(1000, 600))
	mainWindow.Show()

	app.Run()
}
